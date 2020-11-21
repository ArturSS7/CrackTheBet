package emailConfirmation

import (
	"CrackTheBet/backend/database"
	"github.com/labstack/echo"
	"log"
)

func ConfirmEmail(c echo.Context) error {
	token := c.QueryParam("token")
	cc := c.(*database.DBContext)
	rows, err := cc.Db.Query("select id from verification where token = $1", token)
	if err != nil {
		log.Println(err)
	}
	if rows.Next() != false {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
		}
		_, err = cc.Db.Exec("delete from verification where id = $1", id)
		if err != nil {
			log.Println(err)
		}
		_, err = cc.Db.Exec("update users set verified = true where id = $1", id)
		if err != nil {
			log.Println(err)
		}
		return c.String(200, "Email confirmed, now you can login.")
	} else {
		return c.String(200, "Invalid confirmation token")
	}
}
