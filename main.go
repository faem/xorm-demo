package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"os"
	"os/signal"
)

type Profile struct {
	ID       string
	Name     string
	Company  string
	Position string
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:1234@/test")
	printError(err)
	defer engine.Close()

	engine.SetMapper(core.SnakeMapper{})
	engine.SetTableMapper(core.SameMapper{})
	engine.SetColumnMapper(core.SnakeMapper{})

	err = engine.Sync2(new(Profile))
	printError(err)

	//Create Profile
	_, err = engine.Insert(Profile{
		ID:       "1",
		Name:     "fahim",
		Company:  "AppsCode",
		Position: "Software Engineer",
	})
	printError(err)
	fmt.Println("Profile Created.")

	log.Println("Press Ctrl+C to Read the Profile. . . .")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	//Read Profile
	var profile = Profile{ID: "1"}
	_, err = engine.Get(&profile)
	printError(err)
	fmt.Println("Created Profile = \n", prettyStruct(profile))

	log.Println("Press Ctrl+C to Update the Profile. . . .")
	ch = make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	//Update Profile
	_, err = engine.Update(Profile{
		ID:      "1",
		Name:    "Fahim Abrar",
		Company: "AppsCode Inc.",
	})
	printError(err)

	profile = Profile{ID: "1"}
	_, err = engine.Get(&profile)
	printError(err)
	fmt.Println("Updated Profile = \n", prettyStruct(profile))

	log.Println("Press Ctrl+C to Delete the Profile. . . .")
	ch = make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	//Delete Profile
	_, err = engine.Delete(&Profile{ID: "1"})
	printError(err)
	fmt.Println("\nProfile Deleted Successfully.")
}

func printError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func prettyStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
