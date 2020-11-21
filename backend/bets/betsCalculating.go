package bets

import (
	"database/sql"
	"log"
)

type Bet struct {
	id           int
	amount       int
	odd          float32
	flashScoreId string
	winner       int
	status       string
	betStatus    string
	betPlayer    int
	userId       int64
	prize        float32
}

func addPrize(db *sql.DB, prize float32, userId int64) bool {
	_, err := db.Exec("update users set balance=balance+$1 where id = $2", prize, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func setPrize(db *sql.DB, id int, prize float32) bool {
	_, err := db.Exec("update bets set prize = $1 where id = $2", prize, id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func updateStatus(db *sql.DB, id int, status string) bool {
	_, err := db.Exec("update bets set status = $1 where id = $2", status, id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func calculateBetPrize(bets *[]Bet) {
	for i := range *bets {
		if (*bets)[i].status == "finished" {
			var prize float32
			if (*bets)[i].winner == (*bets)[i].betPlayer {
				prize = float32((*bets)[i].amount) * (*bets)[i].odd
				(*bets)[i].betStatus = "win"
			} else {
				prize = 0
				(*bets)[i].betStatus = "lose"
			}
			(*bets)[i].prize = prize
		}
	}
}

func getBetEvents(db *sql.DB, bets *[]Bet) bool {
	for i := range *bets {
		rows, err := db.Query("select status, winner from events where flashscore_id = $1", (*bets)[i].flashScoreId)
		if err != nil {
			log.Println(err)
			return false
		}
		for rows.Next() {
			err := rows.Scan(&(*bets)[i].status, &(*bets)[i].winner)
			if err != nil {
				log.Println(err)
				return false
			}
		}
	}
	return true
}

func calculateBets(db *sql.DB, userId int64) bool {
	rows, err := db.Query("select id, amount, odds, flashscore_id, bet_player, user_id from bets where (user_id = $1 and prize =-1)", userId)
	if err != nil {
		log.Println(err)
		return false
	}
	bets := &[]Bet{}
	for rows.Next() {
		b := Bet{}
		err := rows.Scan(&b.id, &b.amount, &b.odd, &b.flashScoreId, &b.betPlayer, &b.userId)
		if err != nil {
			log.Println(err)
			return false
		}
		*bets = append(*bets, b)
	}
	res := getBetEvents(db, bets)
	if !res {
		return false
	}
	calculateBetPrize(bets)
	for _, v := range *bets {
		if v.status == "finished" {
			if !setPrize(db, v.id, v.prize) || !updateStatus(db, v.id, v.betStatus) || !addPrize(db, v.prize, v.userId) {
				return false
			}
		}
	}
	return true
}
