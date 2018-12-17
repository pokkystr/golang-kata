package captcha

import "strconv"

var (
	rightNumbers = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func GenCaptcha(patternType int, leftType int, rightType int) string {
	if patternType != 1 {
		return rightNumbers[rightType] + " + " + strconv.Itoa(leftType)
	} else {
		return strconv.Itoa(leftType) + " + " + rightNumbers[rightType]
	}
}
