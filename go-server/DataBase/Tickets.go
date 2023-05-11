package database

import (
	"database/sql"
	"fmt"
	utils "main/Utils"
)

func AddTicket(db *sql.DB, data utils.Ticket) bool {
	res, err := db.Exec("INSERT INTO Tickets (usersid, placeid) VALUES(?, ?)", data.UserId, data.PlaceId)
	success := utils.UserError(err)
	if res == nil {
		utils.UserError(fmt.Errorf("error while adding ticket"))
	}
	return success
}

func TicketGetMounth(db *sql.DB, userId int) []map[string]string {
	res, err := db.Query("SELECT name, place, time FROM Places WHERE id IN (SELECT placeid FROM Tickets WHERE usersid = ?) AND time BETWEEN current_timestamp() AND date_add(now(), interval 1 MONTH);", userId)
	utils.UserError(err)

	var places []map[string]string

	for res.Next() {
		var name, place, time string
		err := res.Scan(&name, &place, &time)
		var resPlace = map[string]string{
			"name":  name,
			"place": place,
			"time":  time,
		}
		fmt.Println(resPlace)
		places = append(places, resPlace)
		utils.UserError(err)
	}
	fmt.Println(places)
	return places
}
