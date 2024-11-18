package db

import (
	"context"
	"fmt"
	"log"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

func CreateTable(tableName string, fields []string) error {
	conn := GetYDBConnection()
	ctx := *GetContext()

	return conn.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		columns := "`index` Uint64"
		for _, field := range fields {
			columns += fmt.Sprintf("    `%s` String,", field)
		}
		columns += "    PRIMARY_KEY (index)"

		createTableQuery := fmt.Sprintf("CREATE TABLE `%s` (%s);", tableName, columns)

		log.Println("---QUERY GENERATED---")
		log.Println(createTableQuery)
		log.Println("---END---")

		return s.ExecuteSchemeQuery(ctx, createTableQuery)
	})
}
