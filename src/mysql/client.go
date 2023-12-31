package mysql

import (
	"database/sql"
	"fmt"
	"golang-sqlboiler/mysql/env"
	"log/slog"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	instance *sql.DB
	once     sync.Once
)

// getDBInstance returns a new DBConnection.
func getDBInstance() *sql.DB {
	once.Do(func() {
		loc, err := time.LoadLocation("Local")
		if err != nil {
			slog.Warn("failed to exec time.LoadLocation().", "error", err)
		}
		c := mysql.Config{
			User:                 env.MySQLUser(),
			Passwd:               env.MySQLPass(),
			Net:                  "tcp",
			Addr:                 fmt.Sprintf("%s:%s", env.MySQLHost(), env.MySQLPort()),
			DBName:               env.MySQLDBName(),
			ParseTime:            true,
			Collation:            "utf8mb4_unicode_ci",
			Loc:                  loc,
			AllowNativePasswords: true,
		}
		db, err := sql.Open("mysql", c.FormatDSN())
		if err != nil {
			slog.Error("failed to exec sql.Open().", err)
			panic(err)
		}
		if err := db.Ping(); err != nil {
			slog.Error("failed to exec db.Ping().", err)
			panic(err)
		}
		boil.SetDB(db)
		instance = db
	})
	return instance
}
