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
	time.Sleep(3000 * time.Millisecond)
	return "cool entry", nil
}

type coolWriter struct {
	Reader MyReader
}

func (c coolWriter) Write(b []byte) (string, error) {
	readText, err := c.Reader.Read()
	return fmt.Sprintf("%b, coool I just read this %s", b, readText), err
}

func main() {
	text := []byte("hello world")
	writer := coolWriter{
		Reader: coolReader{},
	}
	_, err := writer.Write(text)
	if err != nil {
		fmt.Print("Ooops something went wrong, bye")
		os.Exit(1)
	}
	fmt.Print("Success!")
}
