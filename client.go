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
	"strconv"
	"strings"
)

type StringToFloat float64

type AccountInfo struct {
	Credit    string
	Creditraw StringToFloat
	Mail      string
	Currency  string
}

func (foe *StringToFloat) UnmarshalJSON(data []byte) error {
	if string(data) == "\"\"" {
		if foe != nil {
			*foe = 0
		}
		return nil
	}
	num := strings.ReplaceAll(string(data), "\"", "")
	n, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return err
	}
	*foe = StringToFloat(n)
	return nil
}

func call(serverUrl string, username string, apikey string, action string) (gjson.Result, error) {
	//parsedURL, err := url.ParseRequestURI(serverUrl)
	//if parsedURL.Scheme != "https" && parsedURL.Scheme != "http" {
	//	return gjson.Result{}, errors.New(fmt.Sprintf("invalid url: %s", parsedURL.String()))
	//}

	//if err != nil {
	//	return gjson.Result{}, err
	//}

	data := url.Values{
		"username":      {username},
		"apiaccesskey":  {apikey},
		"requestformat": {"JSON"},
		"action":        {action},
	}

	//response, err := http.PostForm(serverUrl, data)
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, serverUrl, strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return gjson.Result{}, err
	}
	if response.StatusCode != 200 {
		switch response.StatusCode {
		case 404:
			return gjson.Result{}, errors.New(fmt.Sprintf("StatusCode=404, %s not found", serverUrl))
		default:
			return gjson.Result{}, errors.New(fmt.Sprintf("StatusCode=%d, %s not found", response.StatusCode, response.Status))
		}
	}

	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return gjson.Result{}, err
	}

	filteredJSON := gjson.Get(string(bodyResponse), "SUCCESS.0")
	if filteredJSON.Type != gjson.JSON {
		errorJSON := gjson.Get(string(bodyResponse), "ERROR.0.MESSAGE")
		return errorJSON, errors.New(errorJSON.Str)
	}
	return filteredJSON, err
}

func GetAccountInfo(serverUrl string, username string, apikey string) (AccountInfo, error) {
	var accountInfo AccountInfo
	var err error

	SuccessJSON, err := call(serverUrl, username, apikey, "accountinfo")
	if err != nil {
		return AccountInfo{}, err
	}
	accountJson := gjson.Get(SuccessJSON.Raw, "AccountInfo")

	err = json.NewDecoder(bytes.NewReader([]byte(accountJson.Raw))).Decode(&accountInfo)
	if err != nil {
		return AccountInfo{}, err
	}
	slog.Debug("", "accountInfo", accountInfo)

	return accountInfo, err
}
