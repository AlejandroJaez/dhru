package dhru_test

import (
	"github.com/AlejandroJaez/dhru"
	"testing"
)

var credentials = dhru.Server{
	Url:       "https://www.movilunlock.com/api/index.php",
	Username:  "alejandrojaez",
	SecretKey: "P89-M5G-FWX-3YS-YLH-MY-XDH-JXX",
}

func TestServices(t *testing.T) {
	t.Parallel()
	_, err := dhru.GetServices(credentials)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
}

func TestAccountInfo(t *testing.T) {
	t.Parallel()
	_, err := dhru.GetAccountInfo(credentials)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
}
