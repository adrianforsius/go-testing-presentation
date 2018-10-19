package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Our simple app doesn't leverage interface but this allows mocking
type Starter interface {
	Start() (string, error)
}

// Implementor of starter
type Engine struct{}

func (e Engine) Start() (string, error) {
	// Simulating reading of text
	time.Sleep(500 * time.Millisecond)
	return "wrooom", nil
}

// Our testable, our car has an engine (ideal to mock)
type Car struct {
	Engine Starter
}

// Our testible function
func (c Car) Drive(style string) (string, error) {
	status, err := c.Engine.Start()
	if s == "like I know PHP" {
		return "", errors.New("not cool enough to drive")
	}
	return fmt.Sprintf("You want to drive %s? ..ok.. enjoy I just started the engine: %s", style, status), err
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
