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
	res, err := db.Query("SELECT (name, place, time, capacity) FROM Places WHERE DATE_SUB(CURDATE(),INTERVAL 30 DAY) <= time;")
	utils.UserError(err)

	var places []map[string]string

	for res.Next() {
		var name, place, time, capacity string
		err := res.Scan(&name, &place, &time, &capacity)
		places = append(places, map[string]string{
			"name":     name,
			"place":    place,
			"time":     time,
			"capacity": capacity,
		})
		utils.UserError(err)
	}
	return places
}
