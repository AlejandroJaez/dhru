package dhru

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

func debugRequest(action string) (gjson.Result, error) {
	var JSONSuccess gjson.Result
	value, debug := os.LookupEnv("DHRU_DEBUG")
	if debug == true && value == "TRUE" {
		fileInfo, err := os.Stat(fmt.Sprintf("%s.json", action))
		if err == nil {
			CachedFileJSON, _ := os.ReadFile(fileInfo.Name())
			JSONSuccess = gjson.Get(string(CachedFileJSON), "SUCCESS.0")
			return JSONSuccess, nil
		}
	}
	return JSONSuccess, nil
}

func request(serverURL string, username string, apikey string, action string) (gjson.Result, error) {
	var JSONSuccess gjson.Result

	// if not debug, call the api
	client := &http.Client{}
	request, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		serverURL,
		strings.NewReader(url.Values{
			"username":      {username},
			"apiaccesskey":  {apikey},
			"requestformat": {"JSON"},
			"action":        {action},
		}.Encode()))
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
			return gjson.Result{}, fmt.Errorf("StatusCode=%d, %s", response.StatusCode, response.Body)
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
	if !gjson.Valid(string(bodyResponse)) {
		return gjson.Result{}, fmt.Errorf("response from remote server is not a valid json: %q", bodyResponse)
	}

	JSONSuccess = gjson.Get(string(bodyResponse), "SUCCESS.0")
	if JSONSuccess.Type != gjson.JSON {
		errorJSON := gjson.Get(string(bodyResponse), "ERROR.0.MESSAGE")
		if errorJSON.Type != gjson.String {
			return gjson.Result{}, fmt.Errorf("error in response")
		}
		return gjson.Result{}, errors.New(errorJSON.Str)
	}

	return JSONSuccess, err
}

func AccountInfo(serverUrl string, username string, apikey string) (DrhuAccount, error) {
	var accountInfo DrhuAccount
	responseJSON, err := request(serverUrl, username, apikey, "accountinfo")
	if err != nil {
		return DrhuAccount{}, err
	}
	accountJSON := gjson.Get(responseJSON.Raw, "AccountInfo")
	err = json.NewDecoder(bytes.NewReader([]byte(accountJSON.Raw))).Decode(&accountInfo)
	if err != nil {
		return DrhuAccount{}, fmt.Errorf("%s", err)
	}
	return accountInfo, err
}

func Services(serverUrl string, username string, apikey string) (map[string]ServiceGroup, error) {
	var services map[string]ServiceGroup
	responseJSON, err := request(serverUrl, username, apikey, "imeiservicelist")
	if err != nil {
		return services, fmt.Errorf("%s", err)
	}
	servicesJSON := gjson.Get(responseJSON.Raw, "LIST")
	err = json.NewDecoder(bytes.NewReader([]byte(servicesJSON.Raw))).Decode(&services)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return services, nil
}
