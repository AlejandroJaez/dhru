package dhru

import (
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	acInfo, err := GetAccountInfo("https://www.movilunlock.com/api/index.php", "taoog", "FTX-6RT-VG5-MKB-K3P-5AS-KQZ-03T")
	notWant := AccountInfo{}
	if acInfo == notWant || err != nil {
		t.Fatalf("error: %s", err)
	}
}
