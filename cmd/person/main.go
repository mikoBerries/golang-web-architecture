package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"os"
)

func main() {
	err1 := fs.ErrExist
	data := "some erorr information"
	if errors.Is(err1, fs.ErrExist) {
		fmt.Printf("Erros exist %v", data)
	}
	fmt.Println()
	//a wraapped error massage
	fsNotExist := fs.PathError{
		Op:   "opening file .txt",
		Path: "pages/files/",
		Err:  fs.ErrNotExist,
	}
	// var err string
	//error checking from wrapped error massage and costuming each error massage
	if errors.Is(fsNotExist.Unwrap(), fs.ErrClosed) { //are err closed ?
		fmt.Printf("file %v%v Are Closed", fsNotExist.Path, fsNotExist.Op)
	} else if errors.Is(fsNotExist.Unwrap(), fs.ErrInvalid) { //are this err invalid ?
		fmt.Printf("file %v%v Are invalid", fsNotExist.Path, fsNotExist.Op)
	} else if errors.Is(fsNotExist.Unwrap(), fs.ErrNotExist) { //are this err not exist ?
		fmt.Printf("file %v/%v Are not exist please do semothing", fsNotExist.Path, fsNotExist.Op)
	}
	var addErr *net.AddrError
	if errors.As(err1, &addErr) { // are err1 is net.AddrError?
		// do something
		if addErr.Timeout() {

		}
	}
	//some other eror

	f, err := os.Open("rumi.txt")
	var perr *os.PathError

	switch {
	case errors.Is(err, os.ErrPermission) && errors.As(err, &perr): // it is os err.permission erorr? && err is type os.PathError?
		err = fmt.Errorf("you do not have permission to open the file: %w and the path is %s", err, perr.Path)
		log.Println(err)
	case errors.Is(err, os.ErrNotExist) && errors.As(err, &perr): // it is os err.ErrNotExist erorr? && err is type os.PathError?
		err = fmt.Errorf("the file does not exist: %w and the path is %s", err, perr.Path)
		log.Println(err)
	case errors.As(err, &perr):
		err = fmt.Errorf("here is the orignial error %w and the path %s", err, perr.Path)
		log.Println(err)
	case err != nil:
		err = fmt.Errorf("file couldn't be opened: %w", err)
		log.Println(err)
	}
	defer f.Close()

	//unwrap help to getting depper detail error in diffrent class / method that produce eror massage
	unwrapErr := foo()
	fmt.Println(unwrapErr)
}

func cat() error {
	return errors.New("cat is an error")
}

func moo() error {
	return fmt.Errorf("moo is an error: %w", cat())
}

func bar() error {
	return fmt.Errorf("bar is an error: %w", moo())
}

func foo() error {
	return fmt.Errorf("foo is an error: %w", bar())
}
