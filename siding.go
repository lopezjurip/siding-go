package siding

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

/*
Siding base URL
*/
const BASEURL = "https://intrawww.ing.puc.cl/siding/dirdes/ingcursos/cursos/index.phtml"

/*
Login url
*/
const LOGINBASEURL = "https://intrawww.ing.puc.cl/siding/index.phtml"

/*
ReadResponse to get HTML string from response
*/
func ReadResponse(resp *http.Response) (html string, err error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
Siding class
*/
type Siding struct {
	Username, Password string
	client             *http.Client
	// cookies            []*http.Cookie
}

/*
Login to Siding
*/
func (s *Siding) Login() (resp *http.Response, err error) {
	s.client = &http.Client{
		Jar: newCookieJar(),
	}

	response, err := s.client.PostForm(LOGINBASEURL, s.postArguments())

	if err != nil {
		return nil, err
	}
	return response, err
}

/*
Client for HTTP usage
*/
func (s *Siding) Client() (client *http.Client, err error) {
	cookie := s.sessionCookie()

	if cookie == nil || cookie.Expires.After(time.Now()) {
		if _, err := s.Login(); err != nil {
			return nil, err
		}
	}
	return s.client, nil
}

/*
Get request operation to URL
*/
func (s *Siding) Get(url string) (resp *http.Response, err error) {
	client, err := s.Client()
	if err != nil {
		return nil, err
	}
	return client.Get(url)
}

/*
Announcements for course with id
*/
func (s *Siding) Announcements(id uint) (resp *http.Response, err error) {
	return s.Get(
		fmt.Sprintf("%s?accion_curso=avisos&acc_aviso=mostrar&id_curso_ic=%d", BASEURL, id),
	)
}

func (s Siding) sessionCookie() *http.Cookie {
	url, _ := url.Parse(LOGINBASEURL)
	if s.client != nil && len(s.client.Jar.Cookies(url)) > 0 {
		return s.client.Jar.Cookies(url)[0]
	}
	return nil
}

func (s Siding) postArguments() url.Values {
	// "login={0}&passwd={1}&sw=&sh=&cd="
	return url.Values{
		"login":  {s.Username},
		"passwd": {s.Password},
		"sw":     {""},
		"sh":     {""},
		"cd":     {""},
	}
}

func newCookieJar() *cookiejar.Jar {
	jar, _ := cookiejar.New(nil)
	return jar
}
