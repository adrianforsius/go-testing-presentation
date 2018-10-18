package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Starter interface {
	Start() (string, error)
}

type Engine struct{}

func (e Engine) Start() (string, error) {
	// Simulating reading of text
	time.Sleep(500 * time.Millisecond)
	return "wrooom", nil
}

type Car struct {
	Engine Starter
}

func (c Car) Drive(s string) (string, error) {
	status, err := c.Engine.Start()
	if s == "like I know PHP" {
		return "", errors.New("not cool enough to drive")
	}
	return fmt.Sprintf("You want to drive %s? ..ok.. enjoy I just started the engine: %s", s, status), err
}

func main() {
	style := "like a boss"
	car := Car{
		Engine: Engine{},
	}
	_, err := car.Drive(style)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print("Success\n")
}
