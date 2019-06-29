package api

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	cookiemonster "github.com/MercuryEngineering/CookieMonster"
)

// NewCookieJar create a cookijar from cookies.txt
func NewCookieJar(file string) *cookiejar.Jar {
	cookies, err := cookiemonster.ParseFile(file)
	if err != nil {
		panic(err)
	}
	cj, _ := cookiejar.New(nil)
	u, _ := url.Parse(APIBaseURL)
	cj.SetCookies(u, cookies)
	return cj
}

// BuildCookieHeader build the cookie header from a collection of cookies
func BuildCookieHeader(cookies []*http.Cookie) string {
	cookieValues := make([]string, len(cookies))
	for i, c := range cookies {
		cookieValues[i] = fmt.Sprintf("%s=%s", c.Name, c.Value)
	}
	return strings.Join(cookieValues, "; ")
}
