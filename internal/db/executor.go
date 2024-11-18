package db

import (
	"context"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

func Execute(query string) error {
	conn := GetYDBConnection()
	ctx := *GetContext()

	return conn.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		return s.ExecuteSchemeQuery(ctx, query)
	})

}
