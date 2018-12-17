package captcha_test

import (
	"captcha"
	"testing"
)

func TestPattern1Left1Right1(t *testing.T) {
	var expect = "1 + one"
	var result = captcha.GenCaptcha(1, 1, 1)

	if expect != result {
		t.Errorf("expect not equals result!")
	}
}

func TestPattern1Left1Right2(t *testing.T) {
	var expect = "1 + two"
	var result = captcha.GenCaptcha(1, 1, 2)

	if expect != result {
		t.Errorf("expect not equals result!")
	}
}
