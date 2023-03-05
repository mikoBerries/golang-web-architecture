tag v1.4.0
Hands-on Exercises - ninja level 2
Hands-on exercise #1
Create an interface and use it.
video: 56 Hands On 1


tag v1.4.1
Hands-on exercise #2
Recreate the database example using two maps to simulate the databases.
● Create an “accessor” interface that has two methods
    ○ save
    ○ retrieve
● have your databases implement the “accessor” interface
● create functions that take in values of type accessor
    ○ put
    ○ get
● use “put” and “get” to store data in either database
video: 57 Hands On 2


tag v1.4.2
Hands-on exercise #3
-------------------
Follow these steps
● fork
● clone
● checkout v1.4.1
● branch
● make changes:
    ○ Organize your code to follow the proper go conventions
● cmd folder
    ■ subfolder for what you want to call your executable
    ■ put in here
● package main
● func main
● create package architecture
● create database implementations in separate folders
    ■ mongo
        ● package mongo
     ● db.go
    ■ harddrive
    ● package harddrive
    ● db.go
video: 58 Hands On 3

tag v1.4.3
Hands-on exercise #4
Add some documentation, both package level and on your interface
● Tip: place the package documentation in a doc.go
video: 59 Hands On 4


tag v1.4.4
Hands-on exercise #5
-------------------
In this exercise, copy the contents from one file to a new file. Create a file “file-01.txt” from which
to copy. To complete this exercise, use
● io.Copy
● os.Create
● os.Open
video: 60 Hands On 5