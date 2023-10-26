package env

import (
	"log/slog"
	"os"
)

func MySQLHost() string   { return getEnv("MYSQL_HOST") }
func MySQLPort() string   { return getEnv("MYSQL_PORT") }
func MySQLUser() string   { return getEnv("MYSQL_USER") }
func MySQLPass() string   { return getEnv("MYSQL_PASSWORD") }
func MySQLDBName() string { return getEnv("MYSQL_DB_NAME") }

// getEnv ...
func getEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		slog.Warn("failed to get environment variable.", "name", name)
	}
	return v
}
