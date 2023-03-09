package fileWriter

import (
	"fmt"
	"io"
	"os"
)

//wrapper for file Error massage
type WriteFileError struct {
	Op    string
	Value interface{}
	err   error
}

func (w WriteFileError) Error() string {
	return w.err.Error()
}

func (w WriteFileError) Unwrap() error {
	return w.err
}

//wrapper class for writer
type WriteFile struct {
	f   *os.File
	err error
}

func NewWriteFile(filename string) *WriteFile {
	f, err := os.Create(filename)
	if err != nil {
		return &WriteFile{
			f: nil,
			err: WriteFileError{
				Op:  "NewWriteFile-Create",
				err: fmt.Errorf("error while creating a file: %w", err),
			},
		}
	}
	return &WriteFile{
		f:   f,
		err: nil,
	}
}

func (w *WriteFile) WriteString(text string) {
	if w.err != nil {
		return
	}

	_, err := io.WriteString(w.f, text)
	if err != nil {
		w.err = WriteFileError{
			Op:    "WriteString",
			Value: text,
			err:   fmt.Errorf("failed while writing a string: %w", err),
		}
	}
}

func (w *WriteFile) Close() {
	if w.err != nil {
		return
	}

	err := w.f.Close()
	if err != nil {
		w.err = WriteFileError{
			Op:    "WriteString",
			Value: nil,
			err:   fmt.Errorf("failed while Closing: %w", err),
		}
	}
}

// All errors returning from Err should be of type *WriteFileError
func (w *WriteFile) Err() error {
	return w.err
}
