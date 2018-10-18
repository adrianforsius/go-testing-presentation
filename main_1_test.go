package main

import "testing"

func TestDriveLikeMe(t *testing.T) {
	input := "like me"
	output := "You want to drive like me? ..ok.. enjoy I just started the engine: wrooom"
	car := Car{
		Engine: Engine{},
	}
	result, _ := car.Drive(input)
	if result != output {
		t.Fatalf("test failed expected '%s' got '%s'", output, result)
	}
}
