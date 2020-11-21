package sessionChecker

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"log"
)

func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	notAuth := []string{"/news", "/", "/register"}
	return func(c echo.Context) error {
		for _, value := range notAuth {
			if value == c.Path() {
				return next(c)
			}
		}
		if id := GetIdFromSession(c); id == -1 {
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
