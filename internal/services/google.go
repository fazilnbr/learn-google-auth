package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/callback-gl",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateStringGl = ""
)

/*
InitializeOAuthGoogle Function
*/
func InitializeOAuthGoogle() {
	oauthConfGl.ClientID = viper.GetString("google.clientID")
	oauthConfGl.ClientSecret = viper.GetString("google.clientSecret")
	oauthStateStringGl = viper.GetString("oauthStateString")
	fmt.Printf("\n\n%v\n\n", oauthConfGl)
}

// handile login

func GoogleLogin(c *gin.Context) {
	HandileLogin(c, oauthConfGl, oauthStateStringGl)
}

// callback from google
func CallBackFromGoogle(c *gin.Context) {
	c.Request.ParseForm()
	state := c.Request.FormValue("state")

	if state != oauthStateStringGl {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Request.FormValue("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, "Code Not Found to provide AccessToken..\n")

		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			c.JSON(http.StatusBadRequest, "User has denied Permission..")
		}
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			return
		}
		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		type date struct {
			id             string
			email          string
			verified_email bool
			picture        string
			// data           string
		}
		var any date
		json.Unmarshal(response, &any)
		fmt.Printf("\n\ndata :%v\n\n", string(response))
		fmt.Printf("\n\ndata :%v\n\n", any)

		c.JSON(http.StatusOK, "Hello, I'm protected\n")
		c.JSON(http.StatusOK, string(response))
		return
	}
}
