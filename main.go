package main

import (
	"fmt"
	"go_ydb_driver/internal/db"
)

func main() {
	fmt.Println("hi")
	var field []string
	field = append(field, "name")
	field = append(field, "surname")
	err := db.CreateTable("test_table", field)
	if err != nil {
		panic(err)
	}
}
