package registration

import (
	"CrackTheBet/backend/authorization"
	"CrackTheBet/backend/database"
	"CrackTheBet/backend/mailSender"
	"CrackTheBet/backend/user"
	"github.com/labstack/echo"
	"log"
	"regexp"
	"unicode"
)

func ValidatePassword(password string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

func HandleRegistration(c echo.Context) error {
	password := c.FormValue("password")
	repeatPassword := c.FormValue("password-repeat")
	if password != repeatPassword {
		return c.Render(200, "registration.html", user.ErrorContext{Error: "Passwords don't match"})
	}
	if !ValidatePassword(password) {
		return c.Render(200, "registration.html", user.ErrorContext{Error: "Password must be minimum 8 letters long, contain numbers, upper chars and special characters"})
	}
	var result bool
	username := c.FormValue("username")
	if len(username) < 3 {
		return c.Render(200, "registration.html", user.ErrorContext{Error: "Username must be at least 3 letters long "})
	}
	email := c.FormValue("email")
	//re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	//if !re.MatchString(email)
	//err := checkmail.ValidateFormat(email)
	//if err != nil{
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		return c.Render(200, "registration.html", user.ErrorContext{Error: "Invalid email"})
	}
	cc := c.(*database.DBContext)
	rows, err := cc.Db.Query("select exists (select id from users where username = $1) and (select verified from users where username = $2)", username, username)
	if err != nil {
		log.Println(err)
		return c.String(400, "Error")
	}
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			log.Println(err)
			return c.String(400, "Error")
		}
	}
	if result {
		return c.Render(200, "registration.html", user.ErrorContext{Error: "Username already registered"})
	}
	rows, err = cc.Db.Query("select exists (select id from users where email = $1) and (select verified from users where email = $2)", email, email)
	if err != nil {
		log.Println(err)
		return c.String(400, "Error")
	}
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			log.Println(err)
			return c.String(400, "Error")
		}
	}
	if result {
		return c.Render(200, "registration.html", user.ErrorContext{Error: "Email already registered"})
	}

	password, err = authorization.HashPassword(password)
	if err != nil {
		log.Println(err)
		return c.String(400, "Error")
	}
	var id int
	rows, err = cc.Db.Query("select id from users where email = $1", email)
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
		_, err = cc.Db.Exec("update users set username=$1, password = $2, email = $3 where id = $4", username, password, email, id)
		if err != nil {
			log.Println(err)
		}
	} else {
		_, err = cc.Db.Exec("insert into users(username, password, email) values ($1, $2, $3)", username, password, email)
		if err != nil {
			log.Println(err)
		}
	}
	token := mailSender.SendConfirmationEmail(email)
	if token != "" {
		var id int
		rows, err = cc.Db.Query("select id from users where email = $1", email)
		if err != nil {
			log.Println(err)
		}
		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				log.Println(err)
			}
		}
		_, err := cc.Db.Exec("insert into verification values($1, $2) on conflict(id) do update set token = $2 where verification.id = $1", id, token)
		if err != nil {
			log.Println(err)
		}
	}
	return c.String(200, "Thank you for your registration. Confirmation email has been sent to you.")

}