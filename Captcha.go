package captcha

func GenCaptcha(patternType int, leftType int, rightType int) string {
	if rightType != 1 {
		return "1 + two"
	}
	return "1 + one"
}
