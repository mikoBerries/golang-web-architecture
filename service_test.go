package architecture

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

type Db map[int]Person

func (m Db) Save(n int, p Person) {
	m[n] = p
}

func (m Db) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	putUsingGomock(t)
	p := Person{
		First: "James",
	}

	mdb := Db{}
	Put(mdb, 1, p)

	got := mdb.Retrieve(1)
	if got != p {
		t.Fatalf("Want %v, got %v", p, got)
	}
}

func ExamplePut() {
	mdb := Db{}
	p := Person{
		First: "James",
	}

	Put(mdb, 1, p)
	got := mdb.Retrieve(1)
	fmt.Println(got)
	// Output: {James}
}

func putUsingGomock(t *testing.T) {
	// modify Want error massage
	gomock.WantFormatter(
		gomock.StringerFunc(func() string { return "is equal to fifteen" }),
		gomock.Eq(15),
	)
	// Modify Got error massage
	gomock.GotFormatterAdapter(
		gomock.GotFormatterFunc(func(i interface{}) string {
			// Leading 0s
			return fmt.Sprintf("%02d", i)
		}),
		gomock.Eq(15),
	)
	ctrl := gomock.NewController(t)

	//If you are using a Go version of 1.14+, a mockgen version of 1.5.0+, and are passing a *testing.T into gomock.NewController(t) you no longer need to call ctrl.Finish()
	// defer ctrl.Finish()

	acs := NewMockAccessor(ctrl)

	p := Person{
		First: "James",
	}

	acs.
		EXPECT().    //expected
		Save(1, p).  //Func save() will called with (1,p) params
		MinTimes(1). //Min times  func save called in func put()
		MaxTimes(1)  //Max times func save called in func put()

	// Accessor func save eror example :
	// Unexpected call to *architecture.MockAccessor.Save([3 {James}])
	// doesn't match the argument at index 0.
	// Got: 3 (int)
	// Want: is equal to 1 (int)

	// acs.
	// 	EXPECT().
	// 	Retrieve(1).
	// 	Return(p)
	// 	//error  aborting test due to missing call(s)

	Put(acs, 1, p)
}
