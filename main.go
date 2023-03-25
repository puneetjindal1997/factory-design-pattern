package main

import (
	"fmt"
	"goguru/database"
)

// Welcome to your channel go guruji

// Topic Factory design pattern?
// by road
// ocean, air

func main() {
	sql, _ := database.GetDb("sql")
	mongo, _ := database.GetDb("mongo")

	printDetails(sql)
	printDetails(mongo)
}

func printDetails(g database.FactoryMethod) {
	fmt.Printf("\n")
	fmt.Printf("db: %s", g.GetName())
	if g.GetName() == "sql" {
		fmt.Printf("\n")
		fmt.Printf("client: %v", g.GetSqlClient())
		fmt.Printf("\n")
	} else {
		fmt.Printf("\n")
		fmt.Printf("client: %v\n", g.GetMongoClient())
		fmt.Printf("\n")
	}
}
