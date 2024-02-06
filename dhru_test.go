package dhru_test

import (
	"github.com/AlejandroJaez/dhru"
	"testing"
)

var credentials = dhru.Server{
	Url:       "https://fakedhru.alesoft.workers.dev/api/index.php",
	Username:  "testuser",
	SecretKey: "testkey",
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
	_, err := dhru.PostImeiOrder(&credentials, 1, 342424234243235)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
}

func BenchmarkIsValidIMEI(b *testing.B) {
	imei := 342424234243235
	for i := 0; i < b.N; i++ {
		dhru.IsValidIMEI(int64(imei))
	}
}
