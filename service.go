package architecture

import "fmt"

//type person serve as model
type Person struct {
	First string
}

type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}

//person service as business logic
type PersonService struct {
	A Accessor
}

//manipulate person data
func (ps PersonService) Get(n int) (Person, error) {
	p := ps.A.Retrieve(n)
	if p.First == "" { //add validattion ad create a new desire error massage
		return Person{}, fmt.Errorf("no person in index %d", n)
	}
	return p, nil
}
