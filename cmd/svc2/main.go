package main

import (
	"fmt"
	"os"

	"github.com/tengis617/gomonorepo"
)

func main() {
	u := &gomonorepo.User{
		ID:   "user1",
		Name: "Bob",
		Age:  50,
	}

	u2 := gomonorepo.NewUser("user2", "Alice", 40)
	fmt.Printf("Hello svc2, from %s \n", u.Name)
	fmt.Printf("Greetings from %s! \n", u2.Name)
	os.Exit(1)
}
