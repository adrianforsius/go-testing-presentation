package main

import "testing"

func TestWithTables(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "like I'm the coolest kid on the block",
			output: "You want to drive like I'm the coolest kid on the block? ..ok.. enjoy I just started the engine: wrooom",
		},
		{
			input:  "like a boss",
			output: "You want to drive like a boss? ..ok.. enjoy I just started the engine: wrooom",
		},
	}

	car := Car{
		Engine: Engine{},
	}
	for _, tc := range testCases {
		result, _ := car.Drive(tc.input)
		if result != tc.output {
			t.Fatalf("test failed expected '%s' got '%s'", tc.output, result)
		}
	}
}
