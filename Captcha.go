package captcha

import "strconv"

var (
	numbersEng = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func GenCaptcha(patternType int, operatorType int, leftType int, rightType int) string {
	var captchaResult = ""
	var operator = ""

	if operatorType == 2 {
		operator = " - "
	} else if operatorType == 3 {
		operator = " * "
	} else {
		operator = " + "
	}

	if patternType != 1 {
		captchaResult = numbersEng[leftType] + operator + strconv.Itoa(rightType)
	} else {
		captchaResult = strconv.Itoa(leftType) + operator + numbersEng[rightType]
	}
	return captchaResult
}
