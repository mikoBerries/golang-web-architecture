Golang web architecture with todd mc
====================================

Code Organization Conventions
-----------------------------
1. separating file to folder
    *database
    *controler
    *cmd(main.go file here)
2. using constructor on accessor:
    -By using a constructor function, we are able to customize how our instance is created so it's not have default value
3. separate database folder but not using constructor method /func because it's already specify which database connected in this ca se mongodb/ postgresql

go build command
----------------
go build -o person.exe -i main.go