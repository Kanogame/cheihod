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
