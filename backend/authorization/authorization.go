package authorization

import (
	"CrackTheBet/backend/database"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

type Error struct {
	Err string `json:"error"`
}

func idSession(c echo.Context, id int64) *sessions.Session {
	sess, _ := session.Get("session", c)
	sess.Values["id"] = id
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	return sess
}
func HandleAuth(c echo.Context) error {
	//if sessionChecker.GetIdFromSession(c) != -1 {
	//	return c.Redirect(301, "/profile")
	//}
	username := c.FormValue("username")
	cc := c.(*database.DBContext)

	rows, err := cc.Db.Query("select  id, password from users where username = $1", username)
	if err != nil {
		log.Println(err)
		return c.JSON(401, Error{Err: "Invalid credentials"})
	}
	var id int64
	var password string
	for rows.Next() {
		err := rows.Scan(&id, &password)
		if err != nil {
			log.Println(err)
		}
	}
	if id == 0 {
		return c.JSON(401, Error{Err: "Invalid credentials"})
	}

	if err = bcrypt.CompareHashAndPassword([]byte(password), []byte(c.FormValue("password"))); err != nil {
		return c.JSON(401, Error{Err: "Invalid credentials"})
	}

	rows, err = cc.Db.Query("select verified from users where username = $1", username)
	if err != nil {
		log.Println(err)
		return c.JSON(401, Error{Err: "Error"})
	}
	var verified bool
	for rows.Next() {
		err := rows.Scan(&verified)
		if err != nil {
			log.Println(err)
		}
	}
	if verified {
		sess := idSession(c, id)
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, Error{Err: "Something somewhere went terribly wrong"})
		}
		return c.JSON(200, Error{Err: "Success"})
	} else {
		return c.JSON(401, Error{Err: "Invalid credentials"})
	}
}
