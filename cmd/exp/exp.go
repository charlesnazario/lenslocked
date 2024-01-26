package main

import (
	"html/template"
	"os"
)

type User struct {
	Name     string
	Children map[string]int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Charlie",
		Children: map[string]int{
			"Charlotte": 8,
			"Amelia":    7,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
