package CAS

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"gopkg.in/cas.v2"
)

func SetupCASMiddleware(casURL string) (gin.HandlerFunc, *cas.Client) {
	u, err := url.Parse(casURL)
	if err != nil {
		panic("invalid CAS url: " + err.Error())
	}

	client := cas.NewClient(&cas.Options{
		URL: u,
	})

	middleware := func(c *gin.Context) {
		client.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
			if !cas.IsAuthenticated(r) {
				cas.RedirectToLogin(w, r)
				return
			}
			c.Request = r
			c.Next()
		}).ServeHTTP(c.Writer, c.Request)
	}

	return middleware, client
}

func LoginHandler(c *gin.Context) {

}

func GetUserID(c *gin.Context) {
	if cas.IsAuthenticated(c.Request) {
		username := cas.Username(c.Request)
		c.JSON(http.StatusOK, gin.H{
			"username": username,
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Unathorized"})
		return
	}
}
