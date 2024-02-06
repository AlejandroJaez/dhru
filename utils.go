package dhru

func IsValidIMEI(imei int64) bool {

	if imei < 100000000000000 || imei > 999999999999999 {
		return false
	}

	sum := 0
	double := false
	for i := 15; i > 0; i-- {
		digit := int(imei % 10)
		imei /= 10
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
