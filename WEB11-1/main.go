package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "Lee", Email: "Lee@naver.com", Age: 25}
	user2 := User{Name: "aaa", Email: "aaa@gmail.com", Age: 40}

	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl")
	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user)
	tmpl.ExecuteTemplate(os.Stderr, "tmpl1.tmpl", user2)
}
