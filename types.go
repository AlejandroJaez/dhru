package dhru

import (
	"fmt"
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

type Credentials struct {
	ServerURL string
	Username  string
	ApiKey    string
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
		return fmt.Errorf("%s", err)
	}
	*foe = StringToFloat(n)
	return nil
}

type Service struct {
	ServiceID   int               `json:"SERVICEID"`
	ServiceType string            `json:"SERVICETYPE"`
	Qnt         string            `json:"QNT"`
	Server      string            `json:"SERVER"`
	MinQnt      string            `json:"MINQNT"`
	MaxQnt      string            `json:"MAXQNT"`
	Custom      map[string]string `json:"CUSTOM"`
	ServiceName string            `json:"SERVICENAME"`
	Credit      string            `json:"CREDIT"`
	Time        string            `json:"TIME"`
	Info        string            `json:"INFO"`
}

type ServiceGroup struct {
	GroupName string             `json:"GROUPNAME"`
	GroupType string             `json:"GROUPTYPE"`
	Services  map[string]Service `json:"SERVICES"`
}
