package dhru

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

func call(serverURL string, username string, apikey string, action string) (gjson.Result, error) {
	data := url.Values{
		"username":      {username},
		"apiaccesskey":  {apikey},
		"requestformat": {"JSON"},
		"action":        {action},
	}
	client := &http.Client{}
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, serverURL, strings.NewReader(data.Encode()))
	if err != nil {
		return gjson.Result{}, fmt.Errorf("%s", err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("%s", err)
	}
	if response.StatusCode != http.StatusOK {
		switch response.StatusCode {
		case http.StatusNotFound:
			return gjson.Result{}, fmt.Errorf("StatusCode=404, %s not found", serverURL)
		default:
			return gjson.Result{}, fmt.Errorf("StatusCode=%d, %s not found", response.StatusCode, response.Body)
		}
	}
	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("%s", err)
	}
	err = response.Body.Close()
	if err != nil {
		return gjson.Result{}, fmt.Errorf("%s", err)
	}
	filteredJSON := gjson.Get(string(bodyResponse), "SUCCESS.0")
	if filteredJSON.Type != gjson.JSON {
		errorJSON := gjson.Get(string(bodyResponse), "ERROR.0.MESSAGE")
		if errorJSON.Type != gjson.String {
			return gjson.Result{}, fmt.Errorf("error in response")
		}
		return gjson.Result{}, errors.New(errorJSON.Str)
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
	accountJSON := gjson.Get(SuccessJSON.Raw, "AccountInfo")
	err = json.NewDecoder(bytes.NewReader([]byte(accountJSON.Raw))).Decode(&accountInfo)
	if err != nil {
		return AccountInfo{}, fmt.Errorf("%s", err)
	}
	slog.Debug("", "accountInfo", accountInfo)
	return accountInfo, err
}

func GetAllServices() (map[string]ServiceGroup, error) {
	var services map[string]ServiceGroup
	jsonServicesFile, err := os.ReadFile("services.json")
	if err != nil {
		return services, fmt.Errorf("%s", err)
	}
	if !gjson.Valid(string(jsonServicesFile)) {
		return services, fmt.Errorf("invalid json string")
	}
	value := gjson.Get(string(jsonServicesFile), "SUCCESS.0.LIST")
	err = json.NewDecoder(bytes.NewReader([]byte(value.Raw))).Decode(&services)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return services, nil
}
