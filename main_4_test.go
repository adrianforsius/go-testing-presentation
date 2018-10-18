package main

import (
	"testing"
	"time"
)

// Show case a concurrent testing for slow tests
func TestSlowParallel(t *testing.T) {
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
			// simulating slownes
			time.Sleep(100 * time.Millisecond)
			if result != tc.output {
				t.Fatalf("test failed expected '%s' got '%s'", tc.output, result)
			}
		})
	}
}
