package main

import (
	"fmt"
	"testing"
)

// Integration test
func TestSuccessMain(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "normalized string",
			input:  "totally normal string",
			output: "success",
		},
	}

	writer := coolWriter{
		Reader: coolReader{},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := writer.Write(tc.input)
			if result != tc.output {
				t.Fatalf("test failed expected %s got %s", tc.output, result)
			}
		})
	}
}

// Integration test
func TestErrorMain(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "coolness fail",
			input:  "I like PHP",
			output: "error not cool enough",
		},
	}

	writer := coolWriter{
		Reader: coolReader{},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := writer.Write(tc.input)
			_ = err
			if fmt.Sprintf("%s", err) != tc.output {
				t.Fatalf("test failed expected %s got %v", tc.output, err)
			}
		})
	}
}
