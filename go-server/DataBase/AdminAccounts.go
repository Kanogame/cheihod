package database

import (
	"database/sql"
	"fmt"
	utils "main/Utils"
)

func AddPlace(db *sql.DB, data utils.AddPlace) {
	res, err := db.Exec(fmt.Sprintf("INSERT INTO Places (name, place, time, capacity) VALUES('%v', '%v', '%v', '%v')", data.name, data.place, data.date+" "+data.time, capacity))
	success := utils.UserError(err)
	if res == nil {
		utils.UserError(fmt.Errorf("error while adding place"))
	}
	return success
}
