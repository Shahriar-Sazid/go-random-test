package storage

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataSourceType string

const (
	POSTGRES DataSourceType = "POSTGRES"
	MYSQL    DataSourceType = "MYSQL"
)

type PostgresConnectionInfo struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type MySQLConnectionInfo struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type ConnectionInfo struct {
	Postgres *PostgresConnectionInfo `gorm:"-" json:"postgres,omitempty"`
	Mysql    *MySQLConnectionInfo    `gorm:"-" json:"mysql,omitempty"`
}
type DataSource struct {
	gorm.Model
	Name           string
	Type           DataSourceType `sql:"type:ENUM('POSTGRES')" gorm:"column:data_source_type"`
	ConnectionInfo ConnectionInfo `gorm:"type:json" `
}

func (a *ConnectionInfo) Scan(src any) error {
	switch src := src.(type) {
	case nil:
		return nil
	case []byte:
		var res ConnectionInfo
		err := json.Unmarshal(src, &res)
		*a = res
		return err

	default:
		return fmt.Errorf("scan: unable to scan type %T into struct", src)
	}

}

func (a ConnectionInfo) Value() (driver.Value, error) {
	ba, err := json.Marshal(a)
	return ba, err
}

func GormTest2() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("could not open database")
	}
	err = db.AutoMigrate(&DataSource{})
	if err != nil {
		log.Fatal("could not migrate database")
	}
	createTestData1(db)
	fetchData1(db)
}

func createTestData1(db *gorm.DB) {
	ds := []DataSource{
		{
			Name: "Postgres",
			Type: POSTGRES,
			ConnectionInfo: ConnectionInfo{
				Postgres: &PostgresConnectionInfo{
					Host:     "localhost",
					Port:     333,
					Username: "sdlfj",
					Password: "sdfs",
					DBName:   "sdfsd",
				},
			},
		},
		{
			Name: "Mysql",
			Type: MYSQL,
			ConnectionInfo: ConnectionInfo{
				Mysql: &MySQLConnectionInfo{
					Host:     "localhost",
					Port:     333,
					Username: "sdlfj",
					Password: "sdfs",
					DBName:   "sdfsd",
				},
			},
		},
	}
	err := db.Create(&ds).Error
	if err != nil {
		log.Println("failed to create user data")
	}
}

func fetchData1(db *gorm.DB) {
	var dsList []DataSource
	if err := db.Find(&dsList).Error; err != nil {
		log.Println("failed to load post")
	}
	log.Println(dsList)
}
