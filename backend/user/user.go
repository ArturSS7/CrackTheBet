package user

import (
	"CrackTheBet/backend/authorization/sessionChecker"
	"CrackTheBet/backend/database"
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"html"
	"log"
	"net/http"
)

type ErrorContext struct {
	Error string
}

type LoginContext struct {
	Login bool
}

func GetProfile(c echo.Context) error {
	cc := c.(*database.DBContext)
	id := sessionChecker.GetIdFromSession(c)
	err, d := getUserData(cc.Db, id)
	if err != nil {
		return c.Render(200, "profile.html", nil)
	}
	d.Username = html.EscapeString(d.Username)
	return c.Render(200, "profile.html", d)
}

type Data struct {
	Username string
	Balance  float32
	Logged   bool
}

func getUserData(db *sql.DB, id int64) (error, Data) {
	d := Data{}
	rows, err := db.Query("select username, balance from users where id = $1", id)
	if err != nil {
		return err, d
	}
	for rows.Next() {
		err := rows.Scan(&d.Username, &d.Balance)
		if err != nil {
			return err, d
		}
	}
	return nil, d
}

func GetIndex(c echo.Context) error {
	id := sessionChecker.GetIdFromSession(c)
	if id != -1 {
		cc := c.(*database.DBContext)
		err, d := getUserData(cc.Db, id)
		if err != nil {
			log.Println(err)
			return c.Render(200, "index.html", Data{Logged: false})
		}
		d.Logged = true
		d.Username = html.EscapeString(d.Username)
		return c.Render(200, "index.html", d)
	} else {
		return c.Render(200, "index.html", Data{Logged: false})
	}
}

func LogOut(c echo.Context) error {
	sess, err := c.Cookie("session")
	if err != nil {
		return err
	}
	if sess != nil {
		cookie := new(http.Cookie)
		cookie.Name = "session"
		cookie.Value = ""
		c.SetCookie(cookie)
		return c.Redirect(302, "/")
	} else {
		return c.Redirect(302, "/")
	}
}

func ForgotPassword(c echo.Context) error {
	return c.Render(200, "forgot-password.html", nil)
}
