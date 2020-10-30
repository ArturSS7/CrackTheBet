package user

import (
	"CrackTheBet/backend/authorization/sessionChecker"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"net/http"
)

type ErrorContext struct {
	Error string
}

type LoginContext struct {
	Login bool
}

func GetProfile(c echo.Context) error {
	return c.Render(200, "profile.html", nil)
}

func GetIndex(c echo.Context) error {
	if sessionChecker.GetIdFromSession(c) == -1 {
		return c.Render(200, "index.html", LoginContext{Login: true})
	} else {
		return c.Render(200, "index.html", nil)
	}
}

func GetRegistration(c echo.Context) error {
	return c.Render(200, "registration.html", nil)
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
