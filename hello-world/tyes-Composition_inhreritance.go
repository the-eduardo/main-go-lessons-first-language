package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

type Player struct {
	User
	GameId int
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}
func main() {
	p := Player{} //initializing
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v\n", p)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(p.Greetings())
}
