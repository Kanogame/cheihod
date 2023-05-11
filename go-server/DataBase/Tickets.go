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

/*
func TicketUserQuery(db *sql.DB, userId string) []map[string]string {

}*/
