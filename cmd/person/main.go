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
	var errpatherr *os.PathError

	//using error.Is() to make more meaning full error massage in cmd
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("you need to provide the name \"%s\" of a valid file in your directory next to the executable - %s\n", fileName, errpatherr.Path)
	} else if errors.As(err, &errpatherr) {
		//using error.As() to get more detailed erorr massage from error struct
		fmt.Printf("error in copyFile: %s - OPERATION: %s - %s\n", errpatherr.Path, errpatherr.Op, err)
	} else if err != nil {
		log.Fatal("in func main while calling handOnExercise3() error val :", err)
	}

	temperr := foo()
	fmt.Println(temperr) //all
	temperr = errors.Unwrap(temperr)
	fmt.Println(temperr) //barr moo cat
	temperr = errors.Unwrap(temperr)
	fmt.Println(temperr) // moo cat
	temperr = errors.Unwrap(temperr)
	fmt.Println(temperr) // cat
	temperr = errors.Unwrap(temperr)
	fmt.Println(temperr) //nil

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

func foo() error { return fmt.Errorf("eror from foo \t %w", bar()) }
func bar() error { return fmt.Errorf("eror from bar \t %w", moo()) }
func moo() error { return fmt.Errorf("eror from moo \t %w", cat()) }
func cat() error { return errors.New("eror from cat \t") }
