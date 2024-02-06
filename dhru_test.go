package dhru_test

import (
	"github.com/AlejandroJaez/dhru"
	"testing"
)

var credentials = dhru.Server{
	Url:       "https://fakedhru.alesoft.workers.dev/api/index.php",
	Username:  "alejandrojaez",
	SecretKey: "P89-M5G-FWX-3YS-YLH-MY-XDH-JXX",
}

func TestServices(t *testing.T) {
	t.Parallel()
	list, err := dhru.GetServices(&credentials)
	if err != nil && list != nil {
		t.Fatalf("error: %s", err)
	}
}

func TestAccountInfo(t *testing.T) {
	t.Parallel()
	account, err := dhru.GetAccountInfo(&credentials)
	if err != nil && account != (dhru.DrhuAccount{}) {
		t.Fatalf("error: %s", err)
	}
}

func TestPostImeiOrder(t *testing.T) {
	t.Parallel()
	_, err := dhru.PostImeiOrder(&credentials, 1, 309777379750368)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
}

func BenchmarkIsValidIMEI(b *testing.B) {
	imei := 309777379750368
	for i := 0; i < b.N; i++ {
		dhru.IsValidIMEI(imei)
	}
}
