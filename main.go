package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
)

type Profile struct {
	Id       string `xorm:"pk not null autoincr"`
	Name     string
	Company  string
	Position string
}

func main() {
	engine := ConnectDatabase()
	defer engine.Close()

	CreateTable(engine)

	CreateProfile(engine)
	WaitForCtrlC("Press Ctrl+C to Read the Profile. . . .")

	ReadProfile(engine)
	WaitForCtrlC("Press Ctrl+C to Update the Profile. . . .")

	UpdateProfile(engine)
	WaitForCtrlC("Press Ctrl+C to Delete the Profile. . . .")

	DeleteProfile(engine)

}

func ConnectDatabase() *xorm.Engine {
	engine, err := xorm.NewEngine("postgres", "host=localhost port=5432 user=fahim password=1234 dbname=test sslmode=disable")
	printError(err)

	engine.SetMapper(core.SnakeMapper{})
	engine.SetTableMapper(core.SameMapper{})
	engine.SetColumnMapper(core.SnakeMapper{})

	err = engine.Ping()
	printError(err)

	log.Println("Successfully Connected to Database. . . . .")

	return engine
}

func CreateTable(engine *xorm.Engine) {
	err := engine.Sync2(new(Profile))
	printError(err)
}

func CreateProfile(engine *xorm.Engine) {
	_, err := engine.Insert(Profile{
		Id:       "1",
		Name:     "fahim",
		Company:  "AppsCode",
		Position: "Software Engineer",
	})
	printError(err)
	fmt.Println("Profile Created.")
}

func ReadProfile(engine *xorm.Engine) {
	var profile = Profile{Id: "1"}
	_, err := engine.Get(&profile)
	printError(err)
	fmt.Println("Created Profile = \n", prettyStruct(profile))
}

func UpdateProfile(engine *xorm.Engine) {
	_, err := engine.Update(Profile{
		Id:      "1",
		Name:    "Fahim Abrar",
		Company: "AppsCode Inc.",
	})
	printError(err)

	profile := Profile{Id: "1"}
	_, err = engine.Get(&profile)
	printError(err)
	fmt.Println("Updated Profile = \n", prettyStruct(profile))
}

func DeleteProfile(engine *xorm.Engine) {
	_, err := engine.Delete(&Profile{Id: "1"})
	printError(err)
	fmt.Println("\nProfile Deleted Successfully.")
}

func WaitForCtrlC(msg string) {
	log.Println(msg)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
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
