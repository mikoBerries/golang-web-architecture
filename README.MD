TAG v1.5.xx
===========
Hands-on exercise #1
--------------------
- Create a background context.
- Create a child context with that background as a parent
- store two values in the child context
    * key of type string
    * key of type int
- pull the values out

Hands-on exercise #2
--------------------
There are two versions of this video: a first try and a second try. You will see the
SECOND try first. After that, if you want to see more, you can watch the longer and
slower, and with more discussion, FIRST try.
- SECOND try is first
- FIRST try is second
Using the code from the previous exercise
- have each key be its own type
- abstract all of the context code into its own
    * functions as applicable
    * package, separate from main
- pull the values out

Hands-on exercise #3
-------------------
Answer these two questions:
- you should store LESS or MORE values in context? LESS
- what are the four common values stored in context that Daniel mentioned? userID sessionID IPAdress isAdmin


Hands-on exercise #4
--------------------
There are two versions of this video: a first try and a second try. You will see the
SECOND try first. After that, if you want to see more, you can watch the longer and
slower, and with more discussion, FIRST try.
- SECOND try is first
- FIRST try is second
Context is not magic. It doesn’t end stuff automatically. If something takes in a context, it
knows about the context, and then can end stuff if it is coded to do so.
- create a background context
- create a child context with a timeout
- show the code working two ways
    * not timing out
    * timing out
Things that will be useful to use in this example:
- time.Time is a specific point in time
- time.Duration is a period of time
- time.Millisecond
    * this is a duration
- time.Sleep
    * takes in a duration


Hands-on exercise #5
--------------------
- create a background context
- create a child context with cancel
    * assign this to a variable “ctx”   
- launch 100 goroutines
    * give a unique number to each goroutine
        ■ print “running” and the number
    * have an eternal loop with
        ■ select
- case <- ctx.Done():
    * return
- default:
    * print “still working” and the number
        * sleep for a 1/10 of a second
        ■ show the number of goroutines running
- after your eternal for loop
    * sleep for a bit so that all goroutines can be launched
    * call your cancel func
        ■ this will cancel your context, and all launched goroutines
    * sleep for a bit so that your program can close goroutines
- show the goroutines running
- exit main
Things that will be useful to use in this example:
- runtime.NumGoroutine()
- time.Sleep


func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
WithCancel returns a copy of parent with a new Done channel. The returned
context's Done channel is closed when the returned cancel function is called or
when the parent context's Done channel is closed, whichever happens first. Canceling
this context releases resources associated with it, so code should call cancel as
soon as the operations running in this Context complete. source
A receive expression used in an assignment or initialization of the special form
x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch
yields an additional untyped boolean result reporting whether the
communication succeeded. The value of ok is true if the value received was
delivered by a successful send operation to the channel, or false if it is a zero
value generated because the channel is closed and empty.
