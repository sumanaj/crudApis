package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(filename, text string) {
	fmt.Printf("Writing to a file in Go lang\n")

	// Creating the file using Create() method
	// with user inputted filename and err
	// variable catches any error thrown
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// closing the running file after the main
	// method has completed execution and
	// the writing to the file is complete
	defer file.Close()

	// writing data to the file using
	// WriteString() method and the
	// length of the string is stored
	// in len variable
	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile(filename string) {

	fmt.Printf("\n\nReading a file in Go lang\n")

	// file is read using ReadFile()
	// method of ioutil package
	data, err := ioutil.ReadFile(filename)

	// in case of an error the error
	// statement is printed and
	// program is stopped
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filename)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

func RenameFile(filename string, rename_filename string) {
	err := os.Rename(filename, rename_filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nNew File Name: %s\n", filename)
}

// main function
func main() {
	// user input for filename
	fmt.Println("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)
	if fileExists(filename) {
		fmt.Println("File already exists in this location, please use different fileName")
	} else {
		// user input for file content
		fmt.Println("Enter text: ")
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')

		// file is created and read
		CreateFile(filename, input)
		ReadFile(filename)
		fmt.Println("Do you want to rename a file, if yes please enter the fileName you wish to update with: ")
		var rename_filename string
		fmt.Scanln(&rename_filename)
		RenameFile(filename, rename_filename)
	}
}
