package db

import (
	// общие импорты из стандартной библиотеки
	"context"
	"fmt"
	"go_ydb_driver/internal/conf"
	"log"
	"sync"

	// импорты пакетов ydb-go-sdk
	"github.com/ydb-platform/ydb-go-sdk/v3"
	_ "github.com/ydb-platform/ydb-go-sdk/v3/table"         // для работы с table-сервисом
	_ "github.com/ydb-platform/ydb-go-sdk/v3/table/options" // для работы с table-сервисом
	_ "github.com/ydb-platform/ydb-go-sdk/v3/table/types"   // для работы с типами YDB и значениями
)

var (
	ydbInstance *ydb.Driver
	Ctx         *context.Context
	Cancel      context.CancelFunc
	Once        sync.Once
)

const (
	database = "/local" // Путь до локальной базы данных
)

func GetContext() *context.Context {
	return Ctx
}

// GetYDBConnection returns a singleton connection to YDB
func GetYDBConnection() *ydb.Driver {
	Once.Do(func() {
		ctx_, cancel_ := context.WithCancel(context.Background())
		fmt.Println("create context")
		dsn := "grpc://" + conf.GetVar("ENDPOINT") + database

		db, err := ydb.Open(ctx_, dsn, ydb.WithAnonymousCredentials())
		if err != nil {
			log.Fatalf("Failed to connect to YDB: %v", err)
		}
		Ctx = &ctx_
		Cancel = cancel_
		ydbInstance = db
	})
	return ydbInstance
}
