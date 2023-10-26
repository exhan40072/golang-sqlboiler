package main

import (
	"context"
	"golang-sqlboiler/mysql"
	"golang-sqlboiler/mysql/repository"
)

func main() {
	db := mysql.GetDBInstance()
	jets := repository.NewJets(db)
	jets.FetchAll(context.Background())
}
