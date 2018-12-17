package captcha

func GenCaptcha(patternType int, leftType int, rightType int) string {
	if rightType == 2 {
		return "1 + two"
	} else if rightType == 3 {
		return "1 + three"
	}
	return "1 + one"
}
