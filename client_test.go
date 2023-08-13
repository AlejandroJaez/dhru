package dhru

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("Error loading .env.test file")
	}
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestGetAccountInfo(t *testing.T) {
	url := os.Getenv("DHRU_SERVER_URL")
	username := os.Getenv("DHRU_USERNAME")
	key := os.Getenv("DHRU_API_KEY")
	acInfo, err := GetAccountInfo(url, username, key)
	notWant := AccountInfo{}
	if acInfo == notWant || err != nil {
		t.Fatalf("error: %s", err)
	}
	fmt.Printf("%#v\n", acInfo)
}
