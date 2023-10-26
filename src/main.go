package main

import (
	"context"
	"fmt"
	"golang-sqlboiler/mysql"
	"golang-sqlboiler/mysql/repository"
)

func main() {
	db := mysql.NewDB()
	jets := repository.NewJets(db)
	str := jets.FetchAll(context.Background())
	fmt.Println(str)
}
