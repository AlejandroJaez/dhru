package dhru

import (
	"strconv"
	"strings"
)

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

type Service struct {
	ServiceId   int               `json:"SERVICEID"`
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

type serviceGroup struct {
	GroupName string             `json:"GROUPNAME"`
	GroupType string             `json:"GROUPTYPE"`
	Services  map[string]Service `json:"SERVICES"`
}
