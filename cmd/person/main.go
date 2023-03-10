package main

import (
	"fmt"
	"io"
	"os"
)

type meow int

func (m meow) error() string {
	return fmt.Sprint("error from meow ", m)
}
func main() {
	var m meow = 1
	fmt.Println(m.error())
	fileName := "this.txt"
	createFile := "createthis.txt"
	err := handOnExercise3(fileName, createFile)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func handOnExercise3(fileName string, createFile string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error while opening file %s: %w", fileName, err)
	}
	defer f.Close()
	cf, err := os.Create(createFile)
	if err != nil {
		return fmt.Errorf("error while creating file %s: %w", createFile, err)
	}
	defer cf.Close()

	n, err := io.Copy(cf, f)
	if err != nil {
		return fmt.Errorf("error while copy file %s: %w", fileName, err)
	}
	fmt.Println("writen data :", n)
	return nil
}
