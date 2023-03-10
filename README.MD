Tag v1.7.xx
===========

Hand On Exercise (Handling Error in golang)
--------------------------------------------
Hands-on exercise #1
-------------------
● create a type
● have it implement the error interface
● create a value of that type
● pass that into a function, or return it from a function, that takes type error

Hands-on exercise #2
--------------------
Answer these questions
● Do you ever want to use
    ○ panic(error)?
        ■ programmer error happened
        ■ prints out stack trace
    ○ log.Panic(error)?
        ■ programmer error happened
        ■ prints out stack trace
    ○ log.Fatal(error)?
        ■ turns off the program
        ■ does not print out stack trace
        ■ appropriate for, “Something has gone so catastrophically wrong, the
        shouldn’t try to recover.”
● Do you ever want to return an error without adding any information to the error?
    ○ always add information about the error
● Do you ever want to log.Print or fmt.Print an error without adding any information to the
error?
    ○ always good to add info
● Do you always want to use fmt.Errorf()?
     whenever you’re returning an error, wrap it with more info
    ○ always doesn’t necessarily as you might have created your own error struct, or
something

Hands-on exercise #3
-------------------
Using fmt.Errorf(), handle all errors in a program that does the following:
● open file #1
● create file #2
● copy the contents of file #1 to file #2


Hands-on exercise #4
--------------------
Using the code in the previous example, use errors.Is() to give a special message if the file you
are trying to open does not exist. Notes:
● https://godoc.org/os#pkg-variables


Hands-on exercise #5
--------------------
Using the code in the previous example, use errors.As() to access *PathError and print out some
fields or run some methods attached to *PathError. Notes:
● https://godoc.org/os#Open
● https://godoc.org/os#Create
● https://godoc.org/os#PathError


Hands-on exercise #6
--------------------
Recreate the foo bar moo cat example. To do this, follow these general steps:
● create these functions, each of which returns a value of type error
    ■ main calls foo
        ○ foo
    ■ foo calls bar
        ○ bar
    ■ bar calls moo
        ○ moo
    ■ moo calls cat
        ○ cat
    ■ cat calls nothing
● include the error information from any previous calls
● use errors.Unwrap() to access all of the individual errors
    ○ name the variables that hold the unwrapped error
        ■ fooErr
        ■ barErr
        ■ mooErr
        ■ catErr
        ■ baseErr
● this err will be nil

