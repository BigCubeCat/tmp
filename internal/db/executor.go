package db

import (
	"context"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

func ExecuteWithParams(query string, params *table.QueryParameters) error {
	conn := GetYDBConnection()
	ctx := *GetContext()

	return conn.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		_, _, err := s.Execute(
			ctx,
			table.TxControl(
				table.BeginTx(table.WithSerializableReadWrite()),
				table.CommitTx(),
			), // Управление транзакцией
			query,
			params, // Передача параметров
		)
		return err
	})
}

func Execute(query string) error {
	conn := GetYDBConnection()
	ctx := *GetContext()

	return conn.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		_, _, err := s.Execute(
			ctx,
			table.TxControl(
				table.BeginTx(table.WithSerializableReadWrite()),
				table.CommitTx(),
			), // Управление транзакцией
			query,
			nil,
		)
		return err
	})

}
