package CAS

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
)

var adminList = map[string]struct{}{
	"jan.kowalski": {},
	"anna.nowak":   {},
}

var sessions = struct {
	sync.RWMutex
	users map[string]bool // login -> isAdmin
}{users: make(map[string]bool)}

func AddSession(login string) {
	sessions.Lock()
	defer sessions.Unlock()
	_, isAdmin := adminList[login]
	sessions.users[login] = isAdmin
}
func CheckSession(login string) (exists bool, isAdmin bool) {
	sessions.RLock()
	defer sessions.RUnlock()
	isAdmin, exists = sessions.users[login]
	return
}

func LoginHandler(c *gin.Context) {
	casURL := "https://cas.al.edu.pl/cas/login?service=" +
		url.QueryEscape("https://10.37.50.87/api/callback") // <- замени на реальный URL
	c.Redirect(http.StatusFound, casURL)
}

type casResponse struct {
	XMLName xml.Name `xml:"serviceResponse"`
	Success *struct {
		User string `xml:"user"`
	} `xml:"authenticationSuccess"`
}

func ValidateCASTicket(ticket string) string {
	serviceValidateURL := "https://cas.al.edu.pl/cas/serviceValidate?service=" +
		url.QueryEscape("https://10.37.50.87/api/callback") +
		"&ticket=" + url.QueryEscape(ticket)

	resp, err := http.Get(serviceValidateURL)
	if err != nil || resp.StatusCode != 200 {
		return ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result casResponse
	if err := xml.Unmarshal(body, &result); err != nil || result.Success == nil {
		return ""
	}
	return result.Success.User
}

func CallbackHandler(c *gin.Context) {
	ticket := c.Query("ticket")
	if ticket == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ticket missing"})
		return
	}

	login := ValidateCASTicket(ticket)
	if login == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid ticket"})
		return
	}

	AddSession(login)

	c.SetCookie("login", login, 3600*24, "/", "", false, true)

	c.Redirect(http.StatusFound, "/")
}

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		login, err := c.Cookie("login")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			c.Abort()
			return
		}

		exists, isAdmin := CheckSession(login)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
			c.Abort()
			return
		}
		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin only"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserID(c *gin.Context) {
	login, err := c.Cookie("login")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"username": login})
}
