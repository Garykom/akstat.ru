package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

const sqlSelectCategories = `
	SELECT 
		id, 
		name
	FROM 
		categories`

func dbGetCategories() []Category {
	fmt.Println("Получаем список точек:")
	conn, err := GetConnection()

	var categories []Category
	err = conn.Select(&categories, sqlSelectCategories)
	if err != nil {
		fmt.Println(err)
	}
	return categories
}
