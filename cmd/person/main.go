package main

import (
	"errors"
	"fmt"
	"io"
	"log"
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

	//using error is to make more meaning full error massage in cmd
	if err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Println("Error: file not found")
	} else {
		log.Fatal("in func main while calling handOnExercise3() error val :", err)
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
