package dhru

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var (
	serverUrl string
	username  string
	key       string
)

func TestMain(m *testing.M) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// call flag.Parse() here if TestMain uses flags
	serverUrl = os.Getenv("DHRU_SERVER_URL")
	username = os.Getenv("DHRU_USERNAME")
	key = os.Getenv("DHRU_API_KEY")
	os.Exit(m.Run())
}

func TestGetAccountInfo(t *testing.T) {
	acInfo, err := GetAccountInfo(serverUrl, username, key)
	notWant := AccountInfo{}
	if acInfo == notWant || err != nil {
		t.Fatalf("error: %s", err)
	}
}

func TestGetAllServices(t *testing.T) {
	services, err := GetAllServices()
	if len(services) == 0 || err != nil {
		t.Fatalf("error: %s", err)
	}
}
