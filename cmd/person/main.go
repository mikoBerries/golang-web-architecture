package main

import (
	"errors"
	"fmt"
	"io/fs"
	"net"
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
	var netErr *net.Error
	if errors.As(err1, &netErr) { // are err1 is net.error?
		//do something
		// if err1.Timeout() {

		// }
	}
	//some toher eror
}
