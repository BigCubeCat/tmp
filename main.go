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
	field := []string{"name"}
	values := []string{"egor"}
	tableName := "test_table"
	tableIndex := tableName + "_index"
	query := db.GenerateInsertQuery(tableName, tableIndex, field)
	fmt.Println(query)

	size, err := strconv.Atoi(conf.GetVar("COUNT"))
	if err != nil {
		fmt.Println("no count insrertions")
		panic(err)
	}

	beginTime := time.Now().Unix()
	for i := uint64(0); i < uint64(size); i++ {
		valuesArray := make([]table.ParameterOption, len(values)+1)
		valuesArray[0] = table.ValueParam("$"+tableIndex, types.Uint64Value(i))
		for j := 0; j < len(values); j++ {
			fmt.Println("$"+field[j], values[j])
			valuesArray[j+1] = table.ValueParam("$"+field[j], types.TextValue(values[j]))
		}
		fmt.Println("valuesArray=", valuesArray)
		err = db.ExecuteWithParams(query, table.NewQueryParameters(valuesArray...))
		if err != nil {
			fmt.Println("test " + strconv.Itoa(int(i)) + " failed with error: " + err.Error())
			break
		}
	}
	endTime := time.Now().Unix()
	fmt.Println("Spent time = ", endTime-beginTime)
}
