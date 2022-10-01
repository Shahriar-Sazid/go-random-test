package pack

import (
	"encoding/json"
	"fmt"
	"time"
)

type Name1 struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}
type Name2 struct {
	SecondName string
	FirstName  string
}

func json_test() {

	for i := 0; i < 100; i++ {
		n1, _ := json.Marshal(Name1{
			FirstName:  "Shahriar",
			SecondName: "Ahmad",
		})

		n2, _ := json.Marshal(Name1{
			SecondName: "Ahmad",
			FirstName:  "Shahriar",
		})

		fmt.Println((string)(n1) == (string)(n2))
		js, err := json.Marshal(n1)
		if err != nil {
			fmt.Println("error occurred in marshalling json")
			return
		}
		var t1 *Name1
		err = json.Unmarshal(js, t1)
		if err != nil {
			fmt.Println("error occurred in unmarshalling json")
			fmt.Printf("%+v", err)
			return
		}
	}
}

type Arg struct {
	Name  string      `bson:"name"`
	Type  string      `bson:"type"`
	Value interface{} `bson:"value"`
}
type Headers map[string]interface{}

type Signature struct {
	UUID           string
	Name           string
	RoutingKey     string
	ETA            *time.Time
	GroupUUID      string
	GroupTaskCount int
	Args           []Arg
	Headers        Headers
	Priority       uint8
	Immutable      bool
	RetryCount     int
	RetryTimeout   int
	OnSuccess      []*Signature
	OnError        []*Signature
	ChordCallback  *Signature
	//MessageGroupId for Broker, e.g. SQS
	BrokerMessageGroupId string
	//ReceiptHandle of SQS Message
	SQSReceiptHandle string
	// StopTaskDeletionOnError used with sqs when we want to send failed messages to dlq,
	// and don't want machinery to delete from source queue
	StopTaskDeletionOnError bool
	// IgnoreWhenTaskNotRegistered auto removes the request when there is no handeler available
	// When this is true a task with no handler will be ignored and not placed back in the queue
	IgnoreWhenTaskNotRegistered bool
}

func json_test2() {
	t := Signature{
		UUID: "sjlf-ldjfsld-dslfjsdf",
		Name: "abc",
		Args: []Arg{
			{
				Name:  "asfd",
				Type:  "string",
				Value: "lsdjfsldfjlsdjf sldfjsldjf sdlfj",
			},
		},
		Priority:                    0,
		Immutable:                   false,
		RetryCount:                  0,
		RetryTimeout:                0,
		SQSReceiptHandle:            "",
		StopTaskDeletionOnError:     false,
		IgnoreWhenTaskNotRegistered: false,
	}

	js, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error occurred in marshalling json")
		return
	}
	var t1 *Signature
	err = json.Unmarshal(js, t1)
	if err != nil {
		fmt.Println("error occurred in unmarshalling json")
		fmt.Printf("%+v", err)
		return
	}
}
