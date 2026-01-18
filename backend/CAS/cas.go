package CAS

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var nginxURL = "http://10.37.50.87:70"

var adminList = map[string]struct{}{
	"13283": {},
}

type Session struct {
	Login   string
	IsAdmin bool
}

var sessions = struct {
	sync.RWMutex
	data map[string]Session
}{
	data: make(map[string]Session),
}

func AddSession(login string) string {
	sessions.Lock()
	defer sessions.Unlock()

	_, isAdmin := adminList[login]

	sessionID := uuid.NewString()
	sessions.data[sessionID] = Session{
		Login:   login,
		IsAdmin: isAdmin,
	}

	return sessionID
}

func GetSession(sessionID string) (Session, bool) {
	sessions.RLock()
	defer sessions.RUnlock()

	s, ok := sessions.data[sessionID]
	return s, ok
}

func LoginHandler(c *gin.Context) {
	casURL := "https://cas.al.edu.pl/cas/login?service=" +
		url.QueryEscape(nginxURL+"/api/callback")

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
		url.QueryEscape(nginxURL+"/api/callback") +
		"&ticket=" + url.QueryEscape(ticket)

	resp, err := http.Get(serviceValidateURL)
	if err != nil || resp.StatusCode != 200 {
		return ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result casResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return ""
	}

	if result.Success == nil {
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

	sessionID := AddSession(login)

	c.SetCookie(
		"session_id",
		sessionID,
		3600*24,
		"/api",
		"",
		false,
		true,
	)

	c.Redirect(http.StatusFound, nginxURL+"/kalendarium")
}

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
			c.Abort()
			return
		}

		session, ok := GetSession(sessionID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			c.Abort()
			return
		}

		if !session.IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin only"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserID(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}

	session, ok := GetSession(sessionID)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": session.Login,
		"isAdmin":  session.IsAdmin,
	})
}

func DeleteSession(sessionID string) {
	sessions.Lock()
	defer sessions.Unlock()
	delete(sessions.data, sessionID)
}

func LogoutHandler(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err == nil {
		DeleteSession(sessionID)
	}

	c.SetCookie(
		"session_id",
		"",
		-1,
		"/api",
		"",
		false,
		true,
	)

	c.Status(http.StatusNoContent)
}
