package database

import (
	"database/sql"
	"fmt"
	utils "main/Utils"
)

func AddPlace(db *sql.DB, data utils.AddPlace) bool {
	res, err := db.Exec(fmt.Sprintf("INSERT INTO Places (name, place, time, capacity) VALUES('%v', '%v', '%v', '%v')", data.Name, data.Place, data.Date+" "+data.Time, data.Capacity))
	success := utils.UserError(err)
	if res == nil {
		utils.UserError(fmt.Errorf("error while adding place"))
	}
	return success
}

func GetPlacesMounth(db *sql.DB) []map[string]string {
	fmt.Println("GetPlacesMounth")
	res, err := db.Query("SELECT name, place, time, capacity FROM Places WHERE time BETWEEN date_sub(now(), interval 1 MONTH) AND date_add(now(), interval 1 MONTH);")
	utils.UserError(err)

	var places []map[string]string

	for res.Next() {
		var name, place, time, capacity string
		err := res.Scan(&name, &place, &time, &capacity)
		var resPlace = map[string]string{
			"name":     name,
			"place":    place,
			"time":     time,
			"capacity": capacity,
		}
		fmt.Println(resPlace)
		places = append(places, resPlace)
		utils.UserError(err)
	}
	fmt.Println(places)
	return places
}
