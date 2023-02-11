package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

/*
HandleMain Function renders the index page when the application index route is called
*/
func HandleMain(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Status(http.StatusOK)
	c.HTML(http.StatusOK, "index.html", nil)
}

func HandileLogin(c *gin.Context, oauthConf *oauth2.Config, oauthStateString string) error {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		fmt.Printf("\n\n\nerror in handile login :%v\n\n", err)
		return err
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Printf("\n\nurl : %v\n\n", oauthConf.RedirectURL)
	c.Redirect(http.StatusTemporaryRedirect, url)
	return nil

}
