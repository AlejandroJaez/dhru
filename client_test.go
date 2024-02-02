package dhru_test

import (
	"github.com/AlejandroJaez/dhru"
	"testing"
)

var credentials = dhru.Server{
	Url:       "https://www.movilunlock.com/api/index.php",
	Username:  "taoog",
	SecretKey: "FTX-6RT-VG5-MKB-K3P-5AS-KQZ-03T",
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
