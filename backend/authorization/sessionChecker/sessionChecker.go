package sessionChecker

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"log"
)

type Error struct {
	Err string `json:"error"`
}

func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := GetIdFromSession(c); id == -1 {
			if c.Path() == "/api/bet" || c.Path() == "/api/bets" {
				return c.JSON(401, Error{Err: "Unauthorized"})
			}
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}

func GetIdFromSession(c echo.Context) int64 {
	sess, err := session.Get("session", c)
	if err != nil {
		log.Println(err)
	}
	if sess == nil {
		return -1
	}
	id, exists := sess.Values["id"]
	if !exists {
		return -1
	}
	return id.(int64)
}
