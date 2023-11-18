package dhru_test

import (
	"fmt"
	"github.com/AlejandroJaez/dhru"
	"testing"
)

var server = "https://www.movilunlock.com/api/index.php"
var username = "alejandrojaez"
var apikey = "57M-T5S-RC1-TZ4-OA0-D02-M8Q-W0U"

func TestServices(t *testing.T) {
	t.Parallel()
	services, err := dhru.Services("", "", "")
	if len(services) == 0 || err != nil {
		t.Fatalf("error: %s", err)
	}
}

func TestAccountInfo(t *testing.T) {
	t.Parallel()
	account, err := dhru.AccountInfo(server, username, apikey)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	fmt.Printf("%v", account)
}

func BenchmarkServices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = dhru.Services("", "", "")
	}
}
