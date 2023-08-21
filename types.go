package dhru

type AccountInfo struct {
	Credit    string
	Creditraw StringToFloat
	Mail      string
	Currency  string
}

type ServiceGroup struct {
	GroupName string    `json:"GROUPNAME"`
	GroupType string    `json:"GROUPTYPE"`
	Services  []Service `json:"SERVICES"`
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

type serviceGroupUnmarshall struct {
	GroupName string             `json:"GROUPNAME"`
	GroupType string             `json:"GROUPTYPE"`
	Services  map[string]Service `json:"SERVICES"`
}
