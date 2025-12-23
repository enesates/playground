package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
)

type ValidationError struct {
	msg string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("[Error]: %s", e.msg)
}

func InvalidError() error {
	return ValidationError{
		msg: "Invalid Error",
	}
}

var (
	Error400 = ValidationError{msg: "400"}
	Error500 = ValidationError{msg: "500"}
)

func DoIt() error {
	//return InvalidError()
	return Error400
}

func main() {
	log.Println("Hello")

	logfile, _ := os.Create("./sysglobal.log")
	log.SetOutput(logfile)
	log.Println("Hello")

	err := DoIt()
	if errors.Is(err, Error400) {
	} else if errors.Is(err, Error500) {
	}

	log.Println(err)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger = logger.With("userId", "12323")
	logger.Error("ERROR")
}
