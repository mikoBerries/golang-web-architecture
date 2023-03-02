//decoupling for database func using interface
package main

import (
	"fmt"

	architecture "github.com/golang-web-architecture"
	"github.com/golang-web-architecture/storage/mongo"
	"github.com/golang-web-architecture/storage/postgres"
)

func main() {
	dbm := mongo.DB{}
	dbp := postgres.DB{}
	p1 := architecture.Person{
		First: "Jenny",
	}

	p2 := architecture.Person{
		First: "James",
	}
	ps := architecture.PersonService{
		//using mongo
		A: dbm,
		//or  using postgre
		//a: dbp,
	}

	//store person to mongo db
	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)

	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbm, 3))

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	//or store person to postgre db
	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)

	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))

}
