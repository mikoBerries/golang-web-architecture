package main

import "fmt"

type user struct {
	firstName string
}

//dumy db
type mongoDB map[int]user

func (m mongoDB) save(n int, u user) {
	m[n] = u
}
func (m mongoDB) retrieve(n int) user {
	return m[n]
}

type storage map[int]user

func (s storage) save(n int, u user) {
	s[n] = u
}
func (s storage) retrieve(n int) user {
	return s[n]
}

type Accessor interface {
	save(n int, u user)
	retrieve(n int) user
}

func put(a Accessor, n int, u user) {
	a.save(n, u)
}
func get(a Accessor, n int) user {
	return a.retrieve(n)
}

func main() {
	mongo := mongoDB{}
	storage := storage{}

	put(mongo, 1, user{"this is user mongo DB"})
	put(storage, 1, user{"this is user storage"})

	fmt.Println(get(mongo, 1))
	fmt.Println(get(storage, 1))
}
