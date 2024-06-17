package main

import (
	"fmt"
	"log"
	"package/postgres"
)

// type probaSQL struct {
// 	s postgres.PostgresSQL
// }

func main() {

	dbBot, err := postgres.NewDBConnect()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(dbBot)

	c, err := dbBot.GetTable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}
