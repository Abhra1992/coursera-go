package api

import (
	"net/http/cookiejar"
	"net/url"

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
