package dhru_test

import (
	"dhru"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var (
	serverURL string
	username  string
	key       string
)

func TestMain(m *testing.M) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// call flag.Parse() here if TestMain uses flags
	serverURL = os.Getenv("DHRU_SERVER_URL")
	username = os.Getenv("DHRU_USERNAME")
	key = os.Getenv("DHRU_API_KEY")
	os.Exit(m.Run())
}

func TestGetAccountInfo(t *testing.T) {
	t.Parallel()
	acInfo, err := dhru.GetAccountInfo(serverURL, username, key)
	notWant := dhru.AccountInfo{}
	if acInfo == notWant || err != nil {
		t.Fatalf("error: %s", err)
	}
}

func TestGetAllServices(t *testing.T) {
	t.Parallel()
	services, err := dhru.GetAllServices()
	if len(services) == 0 || err != nil {
		t.Fatalf("error: %s", err)
	}
}
