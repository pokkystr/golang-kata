package captcha

import "strconv"

var (
	numbersEng = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func GenCaptcha(patternType int, leftType int, rightType int) string {
	if patternType != 1 {
		return numbersEng[leftType] + " + " + strconv.Itoa(rightType)
	} else {
		return strconv.Itoa(leftType) + " + " + numbersEng[rightType]
	}
}
