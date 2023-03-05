package mongo

import architecture "github.com/golang-web-architecture"

//dumy db
type mongoDB map[int]architecture.User

func (m mongoDB) save(n int, u architecture.User) {
	m[n] = u
}
func (m mongoDB) retrieve(n int) architecture.User {
	return m[n]
}
