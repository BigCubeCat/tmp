package main

import (
	"fmt"
	"go_ydb_driver/internal/conf"
	"go_ydb_driver/internal/db"
	"strconv"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func main() {
	field := []string{"name", "surname", "lastname"}
	values := []string{"egor", "bit", "ivan"}
	tableName := conf.GetVar("table")
	fmt.Println("--CREATE QUERY--")
	fmt.Println(db.GenerateCreateQuery(tableName, field))
	fmt.Println("--END--")
	fmt.Println(db.GenerateInsertQuery(tableName, "table_index", field))
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
	beginTime := time.Now().Unix()
	query := db.GenerateInsertQuery(tableName, tableIndex, field)
	for i := 0; i < size; i++ {
		valuesArray := make([]table.ParameterOption, len(values))
		for j := 0; j < len(values); j++ {
			valuesArray[j] = table.ValueParam("$"+field[j], types.TextValue(values[i]))
		}
		err = db.ExecuteWithParams(query, table.NewQueryParameters(valuesArray...))
		if err != nil {
			fmt.Println("test " + strconv.Itoa(i) + " failed with error: " + err.Error())
			break
		}
	}
	endTime := time.Now().Unix()
	fmt.Println("Spent time = ", endTime-beginTime)
}
