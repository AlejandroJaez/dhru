package dhru_test

import (
	"github.com/AlejandroJaez/dhru"
	"testing"
)

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
