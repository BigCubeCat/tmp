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
		columns := "`index` Uint64 NOT NULL PRIMARY KEY"
		for _, field := range fields {
			columns += fmt.Sprintf(", `%s` Blob", field)
		}

		createTableQuery := fmt.Sprintf("CREATE TABLE `%s` (%s);", tableName, columns)

		log.Println(createTableQuery)

		return s.ExecuteSchemeQuery(ctx, createTableQuery)
	})
}
