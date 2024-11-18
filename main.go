package main

import (
	"fmt"
	"go_ydb_driver/internal/conf"
	"go_ydb_driver/internal/db"
	"strconv"
)

func main() {
	field := []string{"name", "surname", "lastname"}
	values := []string{"egor", "bit", "ivan"}
	tableName := conf.GetVar("table")
	fmt.Println("--CREATE QUERY--")
	fmt.Println(db.GenerateCreateQuery(tableName, field))
	fmt.Println("--END--")
	size, err := strconv.Atoi(conf.GetVar("COUNT"))
	if err != nil {
		fmt.Println("no count insrertions")
		panic(err)
	}

	err = db.CreateTable(tableName, field)
	if err != nil {
		panic(err)
	}
	tableIndex := tableName + "_index"
	for i := 0; i < size; i++ {
		query := db.GenerateInsertQuery(tableName, tableIndex, i, field, values)
		err = db.Execute(query)
		if err != nil {
			fmt.Println("test " + strconv.Itoa(i) + " failed with error: " + err.Error())
			break
		}
	}
}
