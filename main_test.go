package main

import (
	"fmt"
	"testing"
)

// A full show case testing success and failure in a concurrent manner
func TestSuccessMain(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "chill driving",
			input:  "like I'm the coolest kid on the block",
			output: "You want to drive like I'm the coolest kid on the block? ..ok.. enjoy I just started the engine: wrooom",
		},
		{
			name:   "boss driving",
			input:  "like a boss",
			output: "You want to drive like a boss? ..ok.. enjoy I just started the engine: wrooom",
		},
	}

	car := Car{
		Engine: Engine{},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := car.Drive(tc.input)
			if result != tc.output {
				t.Fatalf("test failed expected '%s' got '%s'", tc.output, result)
			}
		})
	}
}

func TestErrorMain(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "coolness fail",
			input:  "like I know PHP",
			output: "not cool enough to drive",
		},
	}

	car := Car{
		Engine: Engine{},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := car.Drive(tc.input)
			if fmt.Sprintf("%s", err) != tc.output {
				t.Fatalf("test failed expected '%s' got '%v'", tc.output, err)
			}
		})
	}
}
