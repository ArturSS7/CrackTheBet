package events

import (
	"CrackTheBet/backend/database"
	"github.com/labstack/echo"
	"log"
)

type Event struct {
	EventType    string  `json:"event_type"`
	Player1      string  `json:"player_1"`
	Player2      string  `json:"player_2"`
	Odds1        float32 `json:"odds_1"`
	Odds2        float32 `json:"odds_2"`
	Status       string  `json:"status"`
	Time         int64   `json:"time"`
	FlashScoreId string  `json:"flashscore_id"`
}

type League struct {
	LeagueName string  `json:"league_name"`
	Events     []Event `json:"events"`
}

type AllEvents struct {
	Leagues []League `json:"leagues"`
}

type tempLeagues []string

func GetEvents(c echo.Context) error {
	t := tempLeagues{}
	cc := c.(*database.DBContext)
	rows, err := cc.Db.Query("select distinct league from events")
	if err != nil {
		return c.NoContent(500)
	}
	for rows.Next() {
		var l string
		err := rows.Scan(&l)
		if err != nil {
			return c.NoContent(500)
		}
		t = append(t, l)
	}
	all := AllEvents{}
	for _, v := range t {
		l := League{}
		rows, err := cc.Db.Query("select event_type, player1, player2, odds1, odds2, status, time, flashscore_id from events where league = $1", v)
		if err != nil {
			log.Println(err)
			return c.NoContent(500)
		}
		for rows.Next() {
			e := Event{}
			err := rows.Scan(&e.EventType, &e.Player1, &e.Player2, &e.Odds1, &e.Odds2, &e.Status, &e.Time, &e.FlashScoreId)
			if err != nil {
				log.Println(err)
				return c.NoContent(500)
			}
			l.Events = append(l.Events, e)
			l.LeagueName = v
		}
		all.Leagues = append(all.Leagues, l)
	}
	return c.JSON(200, all)
}
