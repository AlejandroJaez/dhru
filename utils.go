package dhru

import "strconv"

const (
	ActionAccountInfo    Action = "accountinfo"
	ActionServiceList    Action = "imeiservicelist"
	ActionPlaceOrder     Action = "placeimeiorder"
	ActionPlaceOrderBulk Action = "placeimeiorderbulk"
	ActionGetOrder       Action = "getimeiorder"
	ActionGetOrderBulk   Action = "getimeiorderbulk"
)

func IsValidIMEI[T string | int | int64](imei T) bool {
	var imeiInt int64
	if v, ok := any(imei).(string); ok {
		imeiInt64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return false
		}
		imeiInt = imeiInt64
	} else {
		return false
	}

	if imeiInt < 100000000000000 || imeiInt > 999999999999999 {
		return false
	}

	sum := 0
	double := false
	for i := 15; i > 0; i-- {
		digit := int(imeiInt % 10)
		imeiInt /= 10
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}
