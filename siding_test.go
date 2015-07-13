package siding

import (
	"flag"
	"fmt"
	"testing"
)

var (
	username, password string
	courseID           uint
)

func init() {
	flag.StringVar(&username, "user", "", "Siding username")
	flag.StringVar(&password, "pass", "", "Siding password")
	flag.UintVar(&courseID, "course", 1000, "Course ID to test announcements")
}

func ValidSession() *Siding {
	return &Siding{Username: username, Password: password}
}

func VerifySession(siding *Siding, t *testing.T) {
	if session := siding.sessionCookie().Name; session != "PHPSESSID" {
		t.Error("Expected session cookie to be named PHPSESSID, got", session)
	}
}

func TestSessionArguments(t *testing.T) {
	if username == "" || password == "" {
		t.Error(`Must declare session credentials to begin testing.
			Example: go test --user USERNAME -pass PASSWORD`)
	}
}

func TestPostArguments(t *testing.T) {
	siding := ValidSession()

	args := siding.postArguments()

	if a := args.Get("login"); a != username {
		t.Error("Expected ", username, ", got ", a)
	}

	if a := args.Get("passwd"); a != password {
		t.Error("Expected ", password, ", got ", a)
	}

	for _, key := range []string{"sw", "sh", "cd"} {
		if a := args.Get(key); a != "" {
			t.Error("Expected ", key, "to be empty")
		}
	}
}

func TestLogin(t *testing.T) {
	siding := ValidSession()
	siding.Login()

	VerifySession(siding, t)
}

func TestClient(t *testing.T) {
	siding := ValidSession()

	_, err := siding.Client()
	if err != nil {
		t.Error("Expected valid logged session on client, got error", err)
	}

	VerifySession(siding, t)
}

func TestAnnouncements(t *testing.T) {
	siding := ValidSession()

	resp, err := siding.Announcements(courseID)

	if err != nil {
		t.Error("Got error", err)
	}

	html, err := ReadResponse(resp)
	if err != nil {
		t.Error("Got error while getting html: ", err)
	}

	fmt.Println(html)
}
