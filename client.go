package dhru

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var client = &http.Client{}

func dhruApiRequest(dhruServer Server, action Action) (ApiResponse, error) {
	values := url.Values{
		"username":      {dhruServer.Username},
		"apiaccesskey":  {dhruServer.SecretKey},
		"requestformat": {"JSON"},
		"action":        {string(action)},
	}

	request, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		dhruServer.Url,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("error creating request: %s", err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("error making request: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		switch response.StatusCode {
		case http.StatusNotFound:
			return ApiResponse{}, fmt.Errorf("StatusCode=%d, %s not found", response.StatusCode, dhruServer.Url)
		default:
			return ApiResponse{}, fmt.Errorf("StatusCode=%d, %s", response.StatusCode, response.Body)
		}
	}

	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return ApiResponse{}, fmt.Errorf("error reading response body: %s", err)
	}
	// Unmarshal JSON into the response struct
	var apiResponse ApiResponse
	if err := json.Unmarshal(bodyResponse, &apiResponse); err != nil {
		return ApiResponse{}, fmt.Errorf("error unmarshalling JSON: %s", err)
	}

	// Check if the success array is empty
	if apiResponse.Error != nil {
		return ApiResponse{}, fmt.Errorf("no success response in API response: %s", apiResponse.Error[0].Message)
	}

	return apiResponse, nil
}

// GetAccountInfo to get info
func GetAccountInfo(server Server) (DrhuAccount, error) {
	// Make the API request
	apiResponse, err := dhruApiRequest(server, ActionAccountInfo)
	if err != nil {
		return DrhuAccount{}, err
	}
	// Assign the extracted account information to the return variable
	accountInfo := apiResponse.Success[0].AccountInfo
	return accountInfo, nil
}
func GetServices(server Server) (map[string]ServiceGroup, error) {
	// Make the API request
	apiResponse, err := dhruApiRequest(server, ActionServiceList)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %s", err)
	}

	// Extract the service list from the response
	serviceList := apiResponse.Success[0].List

	return serviceList, nil
}
