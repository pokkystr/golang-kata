package captcha

var (
	rightNumbers = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func GenCaptcha(patternType int, leftType int, rightType int) string {
	return "1 + " + rightNumbers[rightType]
}
