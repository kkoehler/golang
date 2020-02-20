package main

//go:generate gogen-avro --package=generator  .. ../user.avsc

import (
	"fmt"
	"os"

	"golang.source-fellows.com/generator"
)

func main() {

	user := generator.NewUser()
	user.Name = "Peter"
	user.Favorite_color = "red"

	//write to file
	f, err := os.Create("user.dat")
	if err != nil {
		fmt.Errorf("could not create file because of %v", err)
	}
	defer f.Close()

	err = user.Serialize(f)

	if err != nil {
		fmt.Errorf("could not serialize data because of %v", err)
	}

	//read from file
	fi, err := os.Open("user.dat")
	defer fi.Close()
	user2, err := generator.DeserializeUser(fi)
	if err != nil {
		fmt.Errorf("could not deserialize User because of %v", err)
	}
	fmt.Printf("And here is the user %v\n", user2)

}
