package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "This is gonna go in file"

	file, err := os.Create("./myFile.txt")

	CheckNilError(err)

	length, err := io.WriteString(file, content)
	CheckNilError(err)

	fmt.Println("Length is ", length)
	defer file.Close()

	readFile("./myFile.txt")

}

func readFile(filename string) {
	// returns byte data
	dataByte, err := ioutil.ReadFile(filename)

	CheckNilError(err)

	fmt.Println("Text data inside the file is ", string(dataByte))

}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
