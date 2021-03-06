package main

// Dependencies
import (
	//"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	//"os"
)

// User struct to represent data about users
type User struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Day             int    `json:"day"`
	Month           int    `json:"month"`
	Year            int    `json:"year"`
}

var userLogged = []User{}

// Users slice to seed record users data.
var users = []User{
	{Username: "xGiovanni", Password: "nike4545", ConfirmPassword: "nike4545", Firstname: "Giovanni", Lastname: "Mosquera", Day: 27, Month: 07, Year: 2001},
	{Username: "Juseros9", Password: "contraseña", ConfirmPassword: "contraseña", Firstname: "Sebastián", Lastname: "Rodriguez", Day: 16, Month: 11, Year: 2002},
}

// Creating the handler functions
func main() {
	router := gin.Default()
	router.LoadHTMLFiles("login.html", "register.html", "view.html")
	router.GET("/", defaultRedirect)
	router.GET("/users", loadViewLogin)
	router.POST("/users", loginUser)
	router.GET("/users/register", loadViewRegister)
	router.POST("/users/register", userRegister)
	router.GET("/users/logout", logoutUser)

	router.Run("localhost:8080")
}

func defaultRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/users")
}

func loadViewLogin(c *gin.Context) {
	if len(userLogged) != 0 {
		c.HTML(http.StatusOK, "view.html", gin.H{
			"user":  userLogged,
			"users": users,
		})
		return
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": " ",
		})
	}
}

func loginUser(c *gin.Context) {
	username := c.PostForm("Username")
	password := c.PostForm("Password")

	if len(username) > 0 && len(password) > 0 {
		for _, a := range users {
			if a.Username == username {
				if a.Password == password {
					userLogged := a
					c.HTML(http.StatusOK, "view.html", gin.H{
						"username": userLogged.Username,
						"users":    users,
					})
					return
				}
				c.HTML(http.StatusOK, "login.html", gin.H{
					"message": "Incorrect Password",
				})
				return
			}
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": "This user doesn't exist",
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": "Fill the camps for access",
		})
	}
}

func loadViewRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"message": " ",
	})
}

func userRegister(c *gin.Context) {
	username := c.PostForm("Username")
	password := c.PostForm("Password")
	confirm_password := c.PostForm("ConfirmPassword")
	firstname := c.PostForm("Firstname")
	lastname := c.PostForm("Lastname")
	day := c.PostForm("Day")
	month := c.PostForm("Month")
	year := c.PostForm("Year")
	if len(username) > 0 && len(password) > 0 && len(confirm_password) > 0 && len(firstname) > 0 && len(lastname) > 0 && len(day) > 0 && len(month) > 0 && len(year) > 0 {
		if password == confirm_password {
			dayP, _ := strconv.Atoi(day)
			monthP, _ := strconv.Atoi(month)
			yearP, _ := strconv.Atoi(year)
			newUser := []User{
				{Username: username, Password: password, ConfirmPassword: confirm_password, Firstname: firstname, Lastname: lastname, Day: dayP, Month: monthP, Year: yearP},
			}
			users = append(users, newUser...)
			c.HTML(http.StatusOK, "register.html", gin.H{
				"message": "Te has registrado exitosamente!",
			})
		} else {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"message": "Passwords doesn't match",
			})
		}
	} else {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"message": "Fill all the camps below",
		})
	}
}

func logoutUser(c *gin.Context) {
	userLogged = []User{}
	c.Redirect(http.StatusMovedPermanently, "/users")
}
