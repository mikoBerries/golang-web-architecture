//decoupling for database func using interface
package main

import "fmt"

//type person serve as model
type person struct {
	first string
}

//dummy database connection
type mongo map[int]person
type postg map[int]person

// `m mongo` is a receiver
// `n int, p person` is a key/value pair
func (m mongo) save(n int, p person) {
	m[n] = p
}

func (m mongo) retrieve(n int) person {
	return m[n]
}

func (pg postg) save(n int, p person) {
	pg[n] = p
}

func (pg postg) retrieve(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.retrieve(n)
}

//person service as business logic
type personService struct {
	a accessor
}

//manipulate person data
func (ps personService) get(n int) (person, error) {
	p := ps.a.retrieve(n)
	if p.first == "" { //add validattion ad create a new desire error massage
		return person{}, fmt.Errorf("no person in index %d", n)
	}
	return p, nil
}

func main() {
	dbm := mongo{}
	dbp := postg{}
	p1 := person{
		first: "Jenny",
	}

	p2 := person{
		first: "James",
	}
	ps := personService{
		//using mongo
		a: dbm,
		//or  using postgre
		//a: dbp,
	}

	//store person to mongo db
	put(dbm, 1, p1)
	put(dbm, 2, p2)

	fmt.Println(get(dbm, 1))
	fmt.Println(get(dbm, 3))

	fmt.Println(ps.get(1))
	fmt.Println(ps.get(3))

	//or store person to postgre db
	put(dbp, 1, p1)
	put(dbp, 2, p2)

	fmt.Println(get(dbp, 1))
	fmt.Println(get(dbp, 2))

}
