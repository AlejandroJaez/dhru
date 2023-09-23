package dhru_test

import (
	"fmt"
	"github.com/AlejandroJaez/dhru"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	t.Parallel()
	accountInfo, err := dhru.GetAccountInfo("", "", "")

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	fmt.Printf("Emp: %+v\n", accountInfo)
}

func TestGetAllServices(t *testing.T) {
	t.Parallel()
	services, err := dhru.GetAllServices("", "", "")
	if len(services) == 0 || err != nil {
		t.Fatalf("error: %s", err)
	}
}

func BenchmarkGetAllServices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = dhru.GetAllServices("", "", "")
	}
}
