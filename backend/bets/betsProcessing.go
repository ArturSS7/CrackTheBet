package bets

import (
	"CrackTheBet/backend/authorization/sessionChecker"
	"CrackTheBet/backend/database"
	"CrackTheBet/backend/events"
	"database/sql"
	"github.com/labstack/echo"
	"log"
	"strconv"
	"time"
)

func checkBalance(db *sql.DB, userId int64, amount int) (bool, int) {
	var balance int
	rows, err := db.Query("select balance from users where id = $1", userId)
	if err != nil {
		log.Println(err)
		return false, -1
	}
	for rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			log.Println(err)
			return false, -1
		}
	}
	if amount <= balance {
		return true, balance
	}
	return false, -1
}

func checkEvent(db *sql.DB, id string) bool {
	var exists bool
	rows, err := db.Query("select exists(select id from events where flashscore_id = $1)", id)
	if err != nil {
		log.Println(err)
		return false
	}
	for rows.Next() {
		err := rows.Scan(&exists)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	if exists {
		return true
	}
	return false
}

func checkEventStatus(db *sql.DB, id string) bool {
	var status string
	rows, err := db.Query("select status from events where flashscore_id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	for rows.Next() {
		err := rows.Scan(&status)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	if status != "finished" {
		return true
	}
	return false
}

func updateBalance(db *sql.DB, userId int64, amount int) bool {
	_, err := db.Exec("update users set balance=balance-$1 where id=$2", amount, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func getEventData(db *sql.DB, player int, id string) (bool, events.Event) {
	e := events.Event{}
	if player == 1 {
		rows, err := db.Query("select event_type, player1, player2, time, odds1 from events where flashscore_id = $1", id)
		if err != nil {
			log.Println(err)
			return false, e
		}
		for rows.Next() {
			err := rows.Scan(&e.EventType, &e.Player1, &e.Player2, &e.Time, &e.Odds1)
			if err != nil {
				log.Println(err)
				return false, e
			}
		}
	} else if player == 2 {
		rows, err := db.Query("select event_type, player1, player2, time, odds2 from events where flashscore_id = $1", id)
		if err != nil {
			log.Println(err)
			return false, e
		}
		for rows.Next() {
			err := rows.Scan(&e.EventType, &e.Player1, &e.Player2, &e.Time, &e.Odds2)
			if err != nil {
				log.Println(err)
				return false, e
			}
		}
	}
	return true, e
}

func addBet(db *sql.DB, userId int64, player int, amount int, id string) bool {
	res, e := getEventData(db, player, id)
	if !res {
		return false
	}
	var odd float32
	if e.Odds1 != 0 {
		odd = e.Odds1
	} else if e.Odds2 != 0 {
		odd = e.Odds2
	}
	_, err := db.Exec("insert into bets (event_type, player1, player2, bet_player, time, odds, bet_time, amount, user_id, flashscore_id) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", e.EventType, e.Player1, e.Player2, player, e.Time, odd, time.Now().Unix(), amount, userId, id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func processBet(db *sql.DB, userId int64, player int, amount int, balance int, id string) bool {
	if updateBalance(db, userId, amount) {
		if addBet(db, userId, player, amount, id) {
			return true
		}
		return false
	}
	return false
}

func MakeBet(c echo.Context) error {
	userId := sessionChecker.GetIdFromSession(c)
	id := c.FormValue("id")
	player, err := strconv.Atoi(c.FormValue("player"))
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}
	if player != 1 && player != 2 {
		return c.String(200, "Incorrect player")
	}
	amount, err := strconv.Atoi(c.FormValue("amount"))
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}
	if amount <= 1 {
		return c.String(200, "Incorrect amount")
	}
	cc := c.(*database.DBContext)
	if checkEvent(cc.Db, id) {
		if checkEventStatus(cc.Db, id) {
			if res, balance := checkBalance(cc.Db, userId, amount); res {
				res := processBet(cc.Db, userId, player, amount, balance, id)
				if res {
					return c.String(200, "Bet processed")
				}
				return c.String(200, "Error processing bet")
			}
			return c.String(200, "Not enough cash")
		}
		return c.String(200, "Event has already finished")
	}
	return c.String(200, "Event doesn't exist")
}
