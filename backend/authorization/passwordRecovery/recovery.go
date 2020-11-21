package passwordRecovery

import (
	"CrackTheBet/backend/authorization"
	"CrackTheBet/backend/authorization/registration"
	"CrackTheBet/backend/database"
	"CrackTheBet/backend/mailSender"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

func RecoverPassword(c echo.Context) error {
	email := c.FormValue("email")
	cc := c.(*database.DBContext)
	var result bool
	rows, err := cc.Db.Query("select exists (select id from users where email = $1) and (select verified from users where email = $1)", email)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			log.Println(err)
		}
	}
	if result {
		fmt.Println(result)
		token := mailSender.SendPasswordRecoveryEmail(email)
		if token != "" {
			fmt.Println(token)
			var id int
			rows, err := cc.Db.Query("select id from users where email = $1", email)
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				err := rows.Scan(&id)
				if err != nil {
					log.Println(err)
				}
			}
			_, err = cc.Db.Exec("insert into password_recovery values($1, $2) on conflict(id) do update set token = $2 where password_recovery.id = $1", id, token)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return c.String(200, "An email with information for password recovery has been sent to you")
}

func CheckPasswordToken(c echo.Context) error {
	token := c.QueryParam("token")
	if token == "" {
		return c.String(200, "Invalid token")
	}
	var result bool
	cc := c.(*database.DBContext)
	rows, err := cc.Db.Query("select exists(select id from password_recovery where token = $1)", token)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			log.Println(err)
		}
	}
	if result {
		return c.Render(200, "recovery.html", token)
	}
	return c.String(200, "Invalid token")
}

func UpdatePassword(c echo.Context) error {
	token := c.FormValue("token")
	if token == "" {
		return c.String(200, "invalid token")
	}
	password := c.FormValue("password")
	repeatPassword := c.FormValue("repeat-password")
	if password != repeatPassword {
		return c.String(200, "Passwords don't match")
	}
	if !registration.ValidatePassword(password) {
		return c.String(200, "Password must be minimum 8 letters long, contain numbers, upper chars and special characters")
	}
	var id int
	cc := c.(*database.DBContext)
	rows, err := cc.Db.Query("select id from password_recovery where token = $1", token)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
		}
	}
	if id != 0 {
		password, err = authorization.HashPassword(password)
		if err != nil {
			log.Println(err)
		}
		_, err := cc.Db.Exec("update users set password = $1 where id = $2", password, id)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(id)
		_, err = cc.Db.Exec("delete from password_recovery where id = $1", id)
		if err != nil {
			log.Println(err)
		}
		return c.String(200, "Password successfully changed. You can now login")
	}
	return c.String(200, "Invalid token")
}
