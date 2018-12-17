package captcha

import "strconv"

var (
	numbersEng = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func GenCaptcha(patternType int, leftType int, rightType int) string {
	var captchaResult = ""
	if patternType != 1 {
		captchaResult = numbersEng[leftType] + " + " + strconv.Itoa(rightType)
	} else {
		captchaResult = strconv.Itoa(leftType) + " + " + numbersEng[rightType]
	}
	return captchaResult
}
