package main

import (
	"log"
	"os"
)

type DivideByZeroError struct {
	ErrorMessage string
}

var divideByZeroError = DivideByZeroError{ErrorMessage: "Divide by zero"}

func (c DivideByZeroError) Error() string {
	return c.ErrorMessage
}

func divide(a, b int) (int, error) {
	log.Printf("[INFO]: Diving %d by %d", a, b)
	if b == 0 {
		return 0.0, divideByZeroError
	}

	return a / b, nil
}

func logResult(val int, err error) {
	if err != nil {
		log.Printf("[ERROR]: %+v\n", err)
	} else {
		log.Printf("[INFO]: %+v\n", val)
	}
}

func main() {
	log.Println("Error handling started...")
	f, err := os.OpenFile("./errorlog/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatalf("[ERROR] cannot open log file: %v", err)
	}
	defer func() {
		_ = f.Close()
	}()

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	logResult(divide(10, 0))
	logResult(divide(10, 2))
}
