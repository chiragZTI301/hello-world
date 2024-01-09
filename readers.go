package main

import (
	"fmt"
	"io"
	"strings"
	"os"
)

func readData() {
	reader := strings.NewReader("Hello, Golang!")

	// Create a buffer to read into
	buffer := make([]byte, 8)

	// Read from the reader into the buffer
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
		return
	}

	// Print the number of bytes read and the content
	fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
}

func readDataFromFile() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 8)

	// Read from the file into the buffer
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error:", err)
			return
		}

		// If no more content to read, break the loop
		if n == 0 {
			break
		}

		// Print the content read from the file
		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
	}
}
