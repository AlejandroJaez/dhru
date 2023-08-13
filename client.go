package dhru

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

func call(serverUrl string, username string, apikey string, action string) (*http.Response, error) {
	data := url.Values{
		"username":      {username},
		"apiaccesskey":  {apikey},
		"requestformat": {"JSON"},
		"action":        {action},
	}

	response, err := http.PostForm(serverUrl, data)

	if err != nil || response.StatusCode != 200 {
		switch response.StatusCode {
		case 404:
			return response, errors.New(fmt.Sprintf("StatusCode=404, %s not found", serverUrl))
		default:
			return response, errors.New(fmt.Sprintf("StatusCode=%d, %s not found", response.StatusCode, response.Status))
		}
	}
	return response, err
}

func GetAccountInfo(serverUrl string, username string, apikey string) (AccountInfo, error) {
	var accountInfo AccountInfo
	var err error
	response, err := call(serverUrl, username, apikey, "accountinfo")
	if err != nil {
		return AccountInfo{}, err
	}
	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return AccountInfo{}, err
	}

	filteredJSON := gjson.Get(string(bodyResponse), "SUCCESS.0.AccountInfo")
	if filteredJSON.Type != gjson.JSON {
		errorJSON := gjson.Get(string(bodyResponse), "ERROR.0.MESSAGE")
		return accountInfo, errors.New(errorJSON.Str)
	}
	err = json.NewDecoder(bytes.NewReader([]byte(filteredJSON.Raw))).Decode(&accountInfo)
	if err != nil {
		return AccountInfo{}, err
	}
	slog.Debug("", "accountInfo", accountInfo)

	return accountInfo, err
}
