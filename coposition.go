package main

import (
	"log"
	"fmt"
	"os"

)

type Job struct {
	Command string
	Logger *log.Logger
	}

type User struct {
	id int 
	name, location string
}

func (u *User) myMethod() string {
	return u.name;
}

type Player struct {
	GameId int 
	User 
}



func main() {
	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	p := Player{}
	p.id = 123456
	p.name = "iskander"
	p.GameId = 123
	p.location = "El Alia"

	p1 := Player{
		User : User{id: 42, name: "Matt", location: "LA"},
		GameId : 90404,
	}
	fmt.Println("my method ",p.myMethod())
	fmt.Println(p1)
	job.Logger.Println("test")

}
