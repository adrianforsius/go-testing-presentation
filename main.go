package main

import (
	"fmt"
	"os"
	"time"
)

type MyReader interface {
	Read() (string, error)
}

type coolReader struct{}

func (c coolReader) Read() (string, error) {
	// Simulating reading of text
	time.Sleep(500 * time.Millisecond)
	return "cool entry", nil
}

type coolWriter struct {
	Reader MyReader
}

func (c coolWriter) Write(s string) (string, error) {
	readText, err := c.Reader.Read()
	return fmt.Sprintf("%s, coool I just read this %s", s, readText), err
}

func main() {
	text := "hello world"
	writer := coolWriter{
		Reader: coolReader{},
	}
	_, err := writer.Write(text)
	if err != nil {
		fmt.Printf("error %s", err)
		os.Exit(1)
	}
	fmt.Print("Success\n")
}
