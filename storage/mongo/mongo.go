//mongo databse connection
package mongo

import architecture "github.com/golang-web-architecture"

//dummy database connection
type DB map[int]architecture.Person

//some save func to postgres db
func (m DB) Save(n int, p architecture.Person) {
	m[n] = p
}

//some select func to postgres db
func (m DB) Retrieve(n int) architecture.Person {
	return m[n]
}
