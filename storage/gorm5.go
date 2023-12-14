package storage

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Promo struct {
	Model
	ID         uint       `gorm:"primaryKey;autoIncrement:false"`
	Code       string     `gorm:"primaryKey"`
	Countries  []Country  `gorm:"ForeignKey:PromoID,PromoCode;References:ID,Code"`
	Businesses []Business `gorm:"ForeignKey:PromoID,PromoCode;References:ID,Code"`
}

type Country struct {
	Model
	ID        uint `gorm:"primaryKey"`
	Code      int64
	PromoID   uint
	PromoCode string
}

type Business struct {
	Model
	ID        uint `gorm:"primaryKey"`
	Code      int64
	PromoID   uint
	PromoCode string
	Brands    []Brand
}

type ModelA struct {
	ID         uint
	ModelBList []ModelB
}

type ModelB struct {
	ID         uint
	ModelAID   uint
	ModelCList []ModelC
}

type ModelC struct {
	ID       uint
	ModelBID uint
}

func (m *ModelA) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("model_a_id = ?", m.ID).Delete(&ModelB{})
	return
}

func (m *ModelB) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("model_b_id = ?", m.ID).Delete(&ModelC{})
	return
}

// Updating data in same transaction
func (p *Promo) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("promo_id = ? and promo_code = ?", p.ID, p.Code).Delete(&Business{})
	tx.Clauses(clause.Returning{}).Where("promo_id = ? and promo_code = ?", p.ID, p.Code).Delete(&Country{})
	return
}

// Updating data in same transaction
func (c *Country) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("PromoID", "sfs")
	return
}

func (b *Business) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("business_id = ?", b.ID).Delete(&Brand{})
	return
}

func (b *Brand) AfterDelete(tx *gorm.DB) (err error) {
	tx.Where("brand_id = ?", b.ID).Delete(&Branch{})
	return
}

type Brand struct {
	gorm.Model `json:"-"`
	Code       int64    `json:"code" gorm:"index:,unique,composite:brand"`
	BusinessID uint     `json:"-" gorm:"index:,unique,composite:brand"`
	Branches   []Branch `json:"branches,omitempty"`
}

type Branch struct {
	gorm.Model `json:"-"`
	Code       int64 `json:"code" gorm:"index:,unique,composite:branch"`
	BrandID    uint  `json:"-" gorm:"index:,unique,composite:branch"`
}

type PostgresConfig struct {
	IsSQLite    bool
	DropEnabled bool
	Host        string
	Port        int
	Name        string
	User        string
	Password    string
	MaxIdleConn int
	MaxOpenConn int
	BatchSize   int
	LogDB       bool
	MaxConnTime time.Duration
}

type Adapter struct {
	db *gorm.DB
}

func newAdapterWithDialector(dialect gorm.Dialector) (*Adapter, error) {
	d, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now()
		},
	})
	if err != nil {
		return nil, err
	}

	adapter := &Adapter{
		db: d,
	}
	if err := adapter.migrate(); err != nil {
		return nil, err
	}

	return adapter, nil
}

func NewAdapterWithConfig(cfg PostgresConfig) (*Adapter, error) {
	var adapter *Adapter
	var err error

	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	dialect := postgres.Open(uri)
	adapter, err = newAdapterWithDialector(dialect)
	if err != nil {
		return nil, err
	}

	sqlDB, err := adapter.db.DB()
	if err != nil {
		return nil, err
	}
	adapter.db.CreateBatchSize = cfg.BatchSize
	sqlDB.SetConnMaxLifetime(cfg.MaxConnTime)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	return adapter, nil
}

func (a Adapter) migrate() error {
	db := a.db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Error),
	})

	err := db.AutoMigrate(&Promo{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Country{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Business{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Brand{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Branch{})
	if err != nil {
		return err
	}
	return err
}

func TestGorm5() {
	adapter, err := NewAdapterWithConfig(PostgresConfig{
		Host:        "localhost",
		Port:        5432,
		Name:        "gorm",
		User:        "postgres",
		Password:    "admin",
		MaxIdleConn: 2,
		MaxOpenConn: 2,
		BatchSize:   100,
		LogDB:       true,
		MaxConnTime: time.Second * 5,
	})

	if err != nil {
		fmt.Errorf("%+v", err.Error())
		return
	}

	promo := Promo{
		ID:   1,
		Code: "abc",
		Countries: []Country{
			{
				Code: 123,
			},
			{
				Code: 456,
			},
		},
		Businesses: []Business{
			{
				Code: 789,
				Brands: []Brand{
					{
						Code: 1,
						Branches: []Branch{
							{
								Code: 2,
							},
						},
					},
				},
			},
		},
	}
	if err := adapter.db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&promo).Error; err != nil {
		fmt.Errorf("%+v\n", err.Error())
	}
	adapter.db.Delete(&Promo{ID: 1, Code: "abc"})
}
