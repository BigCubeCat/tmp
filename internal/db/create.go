package db

import (
	"context"
	"fmt"
	"log"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

func GenerateCreateQuery(tableName string, fields []string) string {
	indexName := tableName + "_index"
	columns := indexName + " Uint64,\n"
	for _, field := range fields {
		columns += fmt.Sprintf("    %s String,\n", field)
	}
	columns += "    PRIMARY KEY (" + indexName + ")\n"

	return fmt.Sprintf("CREATE TABLE %s (%s);", tableName, columns)

}

func CreateTable(tableName string, fields []string) error {
	conn := GetYDBConnection()
	ctx := *GetContext()

	return conn.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		createTableQuery := GenerateCreateQuery(tableName, fields)
		log.Println("---QUERY GENERATED---")
		log.Println(createTableQuery)
		log.Println("---END---")

		return s.ExecuteSchemeQuery(ctx, createTableQuery)
	})
}
