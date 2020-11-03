package bets

import (
	"CrackTheBet/backend/authorization/sessionChecker"
	"CrackTheBet/backend/database"
	"github.com/labstack/echo"
	"log"
)

type BetRes struct {
	EventType string  `json:"event_type"`
	Player1   string  `json:"player_1"`
	Player2   string  `json:"player_2"`
	Odds      float32 `json:"odds"`
	Time      int64   `json:"time"`
	BetTime   int64   `json:"bet_time"`
	BetPlayer int     `json:"bet_player"`
	Prize     float32 `json:"prize"`
	Status    string  `json:"status"`
	Amount    int     `json:"amount"`
}

type BetsRes struct {
	Bets []BetRes `json:"bets"`
}

func GetBets(c echo.Context) error {
	id := sessionChecker.GetIdFromSession(c)
	cc := c.(*database.DBContext)
	if calculateBets(cc.Db, id) {
		rows, err := cc.Db.Query("select event_type, player1, player2, odds, time, bet_time, bet_player, amount, prize, status from bets where user_id = $1", id)
		if err != nil {
			log.Println(err)
			return c.NoContent(500)
		}
		bets := BetsRes{}
		for rows.Next() {
			b := BetRes{}
			err := rows.Scan(&b.EventType, &b.Player1, &b.Player2, &b.Odds, &b.Time, &b.BetTime, &b.BetPlayer, &b.Amount, &b.Prize, &b.Status)
			if err != nil {
				log.Println(err)
				return c.NoContent(500)
			}
			bets.Bets = append(bets.Bets, b)
		}
		return c.JSON(200, bets)
	}
	return c.NoContent(500)
}