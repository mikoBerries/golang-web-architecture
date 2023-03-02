package postgres

import architecture "github.com/golang-web-architecture"

//dummy database connection
type DB map[int]architecture.Person

//some save func to postgres db
func (pg DB) Save(n int, p architecture.Person) {
	pg[n] = p
}

//some select func to postgres db
func (pg DB) Retrieve(n int) architecture.Person {
	return pg[n]
}
