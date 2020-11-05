package main

import (
	"CrackTheBet/backend/authorization"
	"CrackTheBet/backend/authorization/emailConfirmation"
	"CrackTheBet/backend/authorization/passwordRecovery"
	"CrackTheBet/backend/authorization/registration"
	"CrackTheBet/backend/authorization/sessionChecker"
	"CrackTheBet/backend/bets"
	"CrackTheBet/backend/database"
	"CrackTheBet/backend/events"
	"CrackTheBet/backend/user"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}
	secret := []byte("ggunjXx8$SKYe3twGqz%")
	fmt.Println(authorization.HashPassword("kek"))
	db := database.Connect()
	e := echo.New()
	//e.Use(sessionChecker.CheckSession) for some reason doesn't work
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &database.DBContext{Context: c, Db: db}
			return h(cc)
		}
	})
	e.Renderer = t
	e.Use(session.Middleware(sessions.NewCookieStore(secret)))
	e.Static("/static", "static")
	e.GET("/", user.GetIndex)
	e.GET("/profile", user.GetProfile, sessionChecker.CheckSession)
	//e.GET("/login", authorization.LoginPage)
	e.POST("/login", authorization.HandleAuth)
	e.GET("/register", user.GetRegistration)
	e.POST("/register", registration.HandleRegistration)
	e.GET("/logout", user.LogOut)
	e.GET("/confirm-email/", emailConfirmation.ConfirmEmail)
	e.GET("/forgot-password", user.ForgotPassword)
	e.POST("/forgot-password", passwordRecovery.RecoverPassword)
	e.GET("/recovery", passwordRecovery.CheckPasswordToken)
	e.POST("/recovery", passwordRecovery.UpdatePassword)

	e.GET("/api/events", events.GetEvents)
	e.POST("/api/bet", bets.MakeBet, sessionChecker.CheckSession)
	e.GET("/api/bets", bets.GetBets, sessionChecker.CheckSession)
	e.Debug = true
	e.Logger.Fatal(e.Start(":666"))
}
