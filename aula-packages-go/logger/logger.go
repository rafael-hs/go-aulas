package logger

import "fmt"

func init() {
	fmt.Println("Logger initialized")
	privateLog()
}

func privateLog() {
	fmt.Println("This is a private log function")
}

func Log() {
	fmt.Println("This is a public log function")
}

func OtherLog() {
	fmt.Println("This is another public log function")
}
