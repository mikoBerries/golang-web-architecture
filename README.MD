Golang web architecture with todd mc
====================================

testing in golang using gomock (mock genetaor)
-----------------------------
1. go get "github.com/golang/mock"
2. Best practice to populate tools:
    -make tools.go and empty import any tools you use
    -adding command "// + build tools" so tools will ignored while compiler build code 
3. Mockgen will generate "Mock" of interface from .go class
4. in service_test.go we can use a "Mock" of Accessor to testing func Put(a Accessor, n int, p Person) without manualy making type that implementing Accessor interface.
5. (1) Define an interface that you wish to mock.
      type MyInterface interface {
        SomeMethod(x int64, y string)
      }
    (2) Use mockgen to generate a mock from the interface.
    (3) Use the mock in a test:
        func TestMyThing(t *testing.T) {
            mockCtrl := gomock.NewController(t)
            defer mockCtrl.Finish()

            mockObj := something.NewMockMyInterface(mockCtrl)
            mockObj.EXPECT().SomeMethod(4, "blah")
            // pass mockObj to a real object and play with it.
        }

6.  Example of using Call.After to chain expected call order:
    firstCall := mockObj.EXPECT().SomeMethod(1, "first")
    secondCall := mockObj.EXPECT().SomeMethod(2, "second").After(firstCall)
    mockObj.EXPECT().SomeMethod(3, "third").After(secondCall)

7. Example of using InOrder to declare expected call order:
    gomock.InOrder(
    mockObj.EXPECT().SomeMethod(1, "first"),
    mockObj.EXPECT().SomeMethod(2, "second"),
    mockObj.EXPECT().SomeMethod(3, "third"),
    )
 8. Modifying Failure Messages (When a matcher reports a failure, it prints the received (Got) vs the expected (Want) value.)
 - The Want value comes from the matcher's String() method. If the matcher's default output doesn't meet your need.
 ```go
 gomock.WantFormatter(s fmt.Stringer, m Matcher) Matcher
 ```
 - The Got value comes from the object's String() method if it is available. In some cases the output of an object is difficult to read (e.g., []byte) and it   would be helpful for the test to print it differently.use
 ```go
gomock.GotFormatterAdapter(s GotFormatter, m Matcher) Matcher 
```
mockgen command
----------------
For an example of the use of mockgen, see the sample/ directory. In simple cases, you will need only the -source flag.
mockgen -source service.go

