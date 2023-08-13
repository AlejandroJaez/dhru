package dhru

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type AccountInfo struct {
	Credit    string
	Creditraw string
	Mail      string
	Currency  string
}

func call(serverUrl string, username string, apikey string, action string) *http.Response {
	data := url.Values{
		"username":      {username},
		"apiaccesskey":  {apikey},
		"requestformat": {"JSON"},
		"action":        {action},
	}
	response, err := http.PostForm(serverUrl, data)
	if err != nil {
		slog.Error("", err)
	}
	return response
}

func GetAccountInfo(serverUrl string, username string, apikey string) (AccountInfo, error) {
	var accountInfo AccountInfo
	var err error
	response := call(serverUrl, username, apikey, "accountinfo")
	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return AccountInfo{}, err
	}

	filteredJSON := gjson.Get(string(bodyResponse), "SUCCESS.0.AccountInfo")
	err = json.NewDecoder(bytes.NewReader([]byte(filteredJSON.Raw))).Decode(&accountInfo)
	if err != nil {
		return AccountInfo{}, err
	}
	slog.Debug("", "accountInfo", accountInfo)

	return accountInfo, err
}
