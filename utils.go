package dhru

import (
	"regexp"
	"strconv"
)

const (
	ActionAccountInfo Action = "accountinfo"
	ActionServiceList Action = "imeiservicelist"
	ActionPlaceOrder  Action = "placeimeiorder"
)

func isValidIMEI(imei string) bool {
	// Check if IMEI is 15 digits long
	if len(imei) != 15 {
		return false
	}

	// Check if IMEI contains only digits
	match, _ := regexp.MatchString("^[0-9]*$", imei)
	if !match {
		return false
	}

	// Convert IMEI to integers
	var digits [15]int
	for i := 0; i < 15; i++ {
		digits[i], _ = strconv.Atoi(string(imei[i]))
	}

	// Perform Luhn algorithm check
	sum := 0
	double := false
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
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
