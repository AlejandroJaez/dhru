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
	"os"
	"strconv"
	"strings"
)

type StringToFloat float64

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

func PlaceSingleOrder(serverUrl string, username string, apikey string) {

}

func keyExists(decoded map[string]interface{}, key string) bool {
	val, ok := decoded[key]
	return ok && val != nil
}

func GetAllServicesFromFile() ([]ServiceGroup, error) {
	var groupsMap map[string]interface{}
	jsonServicesFile, err := os.ReadFile("services.json")
	if err != nil {
		return nil, err
	}
	if !gjson.Valid(string(jsonServicesFile)) {
		return nil, errors.New("error in json file")
	}

	value := gjson.Get(string(jsonServicesFile), "SUCCESS.0.LIST")
	err = json.NewDecoder(bytes.NewReader([]byte(value.Raw))).Decode(&groupsMap)
	if err != nil {
		return nil, err
	}
	groupSlice := make([]ServiceGroup, 0, len(groupsMap))

	for _, group := range groupsMap {
		servicesMap := group.(map[string]interface{})["SERVICES"].(map[string]interface{})
		serviceSlice := make([]Service, 0, len(servicesMap))
		for _, service := range servicesMap {
			serviceMap := service.(map[string]interface{})

			var customField map[string]string
			if keyExists(serviceMap, "CUSTOM") {
				customField = serviceMap["CUSTOM"].(map[string]string)
			}

			serviceSlice = append(serviceSlice, Service{
				ServiceId:   0,
				ServiceType: serviceMap["SERVICETYPE"].(string),
				Qnt:         serviceMap["QNT"].(string),
				Server:      serviceMap["SERVER"].(string),
				MinQnt:      serviceMap["MINQNT"].(string),
				MaxQnt:      serviceMap["MAXQNT"].(string),
				Custom:      customField,
				ServiceName: serviceMap["SERVICENAME"].(string),
				Credit:      serviceMap["CREDIT"].(string),
				Time:        serviceMap["TIME"].(string),
				Info:        serviceMap["INFO"].(string),
			})
		}
		groupSlice = append(groupSlice, ServiceGroup{
			GroupName: group.(map[string]interface{})["GROUPNAME"].(string),
			GroupType: group.(map[string]interface{})["GROUPTYPE"].(string),
			Services:  serviceSlice,
		})

	}

	return groupSlice, nil
}

func GetAllServices() (map[string]serviceGroupUnmarshall, error) {
	var services map[string]serviceGroupUnmarshall

	jsonServicesFile, err := os.ReadFile("services.json")
	if !gjson.Valid(string(jsonServicesFile)) {
		println("Error in json file")
	}
	value := gjson.Get(string(jsonServicesFile), "SUCCESS.0.LIST")
	err = json.NewDecoder(bytes.NewReader([]byte(value.Raw))).Decode(&services)
	if err != nil {
		return nil, err
	}

	return services, nil
}
