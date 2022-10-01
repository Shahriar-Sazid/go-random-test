package pack

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id            int64     `json:"id" gorm:"primarykey"`
	UserName      string    `json:"userName"`
	Phone         string    `json:"phone"`
	Password      string    `json:"-"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	FullName      string    `json:"fullName"`
	DisplayName   string    `json:"displayName"`
	Company       string    `json:"company"`
	Department    string    `json:"department"`
	Photo         string    `json:"photo"`
	Addresses     string    `json:"addresses"`
	PhoneNumber   string    `json:"phoneNumber"`
	Active        bool      `json:"active"`
	LastTimeLogin time.Time `json:"lastTimeLogin"`
}

type Response struct {
	Status  int             `json:"status"`
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Message string          `json:"message,omitempty"`
	Input   interface{}     `json:"input,omitempty"`
	Error   error           `json:"error,omitempty"`
	Paging  interface{}     `json:"paging,omitempty"`
}

type pingDataFormat struct {
	Id                       int64  `json:"id"`
	UploadStartTimeInSeconds int    `json:"uploadStartTimeInSeconds"`
	UploadEndTimeInSeconds   int    `json:"uploadEndTimeInSeconds"`
	CallbackURL              string `json:"callbackURL"`
}

func MarshalTest() {
	jsonData := `{"status":200,"success":true,"data":{"id":1606146401088049402}}`
	response := Response{}
	err := json.Unmarshal([]byte(jsonData), &response)

	if err != nil {
		panic(err)
	}

	user := User{}

	json.Unmarshal(response.Data, &user)

	fmt.Println(user.Id)
	fmt.Printf("%T", user.Id)
}
