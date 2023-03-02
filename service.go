package architecture

import "fmt"

//type person serve as model
type Person struct {
	First string
}

//Accessor how to store and retive a person
//when Retrieve person, if person data don not exist, return empty people {} ,error
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
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

//manipulate person data
func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" { //add validattion ad create a new desire error massage
		return Person{}, fmt.Errorf("no person in index %d", n)
	}
	return p, nil
}
