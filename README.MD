Golang web architecture with todd mc
====================================

testing in golang
-----------------------------
1. <filename>_test.go is file to use to testing some func in file.
2. import "testing"
3. _test.go will ignored by builder.
4. Few write test function
    -Test<funcname>(t *testing.T){}
    -Example<funcname>(){}
go test command
----------------
go test
go test -v