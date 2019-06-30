package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Point struct {
	ID             int     `db:"id"`
	UserID         int     `db:"user_id"`
	CategoryID     int     `db:"category_id"`
	Rating         int     `db:"rating"`
	EmotionalState int     `db:"emotional_state"`
	Latitude       float64 `db:"latitude"`
	Longitude      float64 `db:"longitude"`
}

const sqlSelectPoints = `
	SELECT 
		id, 
		user_id, 
		category_id, 
		rating, 
		emotional_state, 
		latitude, 
		longitude 
	FROM 
		points`

func dbGetPoints() []Point {
	fmt.Println("Получаем список точек:")
	conn, err := GetConnection()

	var points []Point
	err = conn.Select(&points, sqlSelectPoints)
	if err != nil {
		fmt.Println(err)
	}
	return points
}

const sqlCreatePoint = `
	INSERT INTO 
		points (
			user_id, 
			category_id, 
			rating,
			emotional_state,
			latitude,
			longitude
		) 
	VALUES (?, ?, ?, ?, ?, ?)`

func dbCreatePoint(point Point) bool {
	result := true
	fmt.Print(point.CategoryID, point.Rating, point.Latitude, point.Longitude)
	conn, err := GetConnection()
	res, err := conn.Exec(sqlCreatePoint, point.UserID, point.CategoryID, point.Rating, point.EmotionalState, point.Latitude, point.Longitude)
	if err != nil {
		result = false
		fmt.Println(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Create point id: ", id)
	}

	return result
}
