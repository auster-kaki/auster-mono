package rdb

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/infrastructure/rdb/table"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type rdb struct {
	user *table.User
}

func NewDB() (*rdb, error) {
	host := "127.0.0.1"
	if h := os.Getenv("MYSQL_HOST"); h != "" {
		host = h
	}
	dataSourceName := fmt.Sprintf("root:@tcp(%s:3306)/auster?charset=utf8mb4&parseTime=true", host)
	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqlDB, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return &rdb{
		user: table.NewUser(db),
	}, nil
}

func (r *rdb) User() repository.UserRepository {
	return r.user
}
