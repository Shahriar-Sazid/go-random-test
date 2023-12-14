package iface

import "fmt"

type Writer interface {
	Write(message string)
	Print(message string)
	Test()
}

type FileLogger struct {
	ID int
}

func (f *FileLogger) Write(message string) {
	fmt.Println("Writing:", message)
	f.Test()
}

func (f *FileLogger) Print(message string) {
	fmt.Println("Logging:", message)
}

func (f *FileLogger) Test() {
	fmt.Println("Called Test:")
}

type SecondLogger struct {
	FileLogger
}

func (f *SecondLogger) Print(message string) {
	fmt.Println("SecondLogger:", message)
}

func (f *SecondLogger) Test() {
	fmt.Println("Updated Test")
}

type ThirdLogger struct {
	SecondLogger
}

func (f *ThirdLogger) Test() {
	fmt.Println("ThirdLogger Test")
}

func CreateLogger(kind interface{}) Writer {
	switch kind.(type) {
	case FileLogger:
		return &FileLogger{}
	case SecondLogger:
		return &SecondLogger{}
		// here v has type S
	case ThirdLogger:
		return &ThirdLogger{}
	default:
		return nil
		// no match; here v has the same type as i
	}
}

func TestIface() {
	var writer Writer

	writer = CreateLogger(FileLogger{})
	writer.Print("This is a log message.")

	writer = CreateLogger(SecondLogger{})
	writer.Print("2nd log message.")

	writer = CreateLogger(ThirdLogger{})
	writer.Print("3rd log message.")
	writer.(*ThirdLogger).Test()
}
