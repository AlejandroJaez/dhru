package dhru

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func dhruApiRequest(dhruServer Server, action Action, params Parameters) (ApiResponse, error) {
	formData := url.Values{
		"username":      {dhruServer.Username},
		"apiaccesskey":  {dhruServer.SecretKey},
		"requestformat": {"JSON"},
		"action":        {string(action)},
	}

	if action == ActionPlaceOrder {
		xmlData, err := xml.Marshal(params)
		if err != nil {
			return ApiResponse{}, err
		}
		formData.Add("parameters", string(xmlData))
	}

	response, err := http.PostForm(dhruServer.Url, formData)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("error making request: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ApiResponse{}, fmt.Errorf("StatusCode=%d, %s:%s", response.StatusCode, http.StatusText(response.StatusCode), response.Body)
	}

	var apiResponse ApiResponse
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return ApiResponse{}, fmt.Errorf("error decoding JSON: %s", err)
	}

	if len(apiResponse.Error) > 0 {
		return ApiResponse{}, fmt.Errorf("no success response in API response: %s", apiResponse.Error[0].Message)
	}

	return apiResponse, nil
}

func GetAccountInfo(server Server) (DrhuAccount, error) {
	apiResponse, err := dhruApiRequest(server, ActionAccountInfo, Parameters{})
	if err != nil {
		return DrhuAccount{}, err
	}
	return apiResponse.Success[0].AccountInfo, nil
}

func GetServices(server Server) (map[string]ServiceGroup, error) {
	apiResponse, err := dhruApiRequest(server, ActionServiceList, Parameters{})
	if err != nil {
		return nil, err
	}
	return apiResponse.Success[0].List, nil
}

func PostImeiOrder(server Server, service int32, imei int64) (ApiResponse, error) {
	parameters := Parameters{
		IMEI:        strconv.FormatInt(imei, 10),
		ID:          service,
		CustomField: base64.StdEncoding.EncodeToString([]byte(`{"SERIAL_NUMBER":"SERIAL NUMBER"}`)),
	}
	apiResponse, err := dhruApiRequest(server, ActionPlaceOrder, parameters)
	if err != nil {
		return ApiResponse{}, err
	}
	return apiResponse, nil
}
