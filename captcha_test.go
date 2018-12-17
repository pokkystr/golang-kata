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

func TestPattern1Left1Right3(t *testing.T) {
	var expect = "1 + three"
	var result = captcha.GenCaptcha(1, 1, 3)

	if expect != result {
		t.Errorf("expect not equals result!")
	}
}

func TestPattern1Left1Right4(t *testing.T) {
	var expect = "1 + four"
	var result = captcha.GenCaptcha(1, 1, 4)

	if expect != result {
		t.Errorf("expect not equals result!")
	}
}

func TestPattern1Left2Right1(t *testing.T) {
	var expect = "2 + one"
	var result = captcha.GenCaptcha(1, 2, 1)

	if expect != result {
		t.Errorf("expect not equals result!")
	}
}

func TestPattern1Left3Right1(t *testing.T) {
	var expect = "3 + one"
	var result = captcha.GenCaptcha(1, 3, 1)

	if expect != result {
		t.Errorf("expect not equals result!")
	}
}

func TestPattern1Left3Right2(t *testing.T) {
	var expect = "3 + two"
	var result = captcha.GenCaptcha(1, 3, 2)

	if expect != result {
		t.Errorf("expect not equals result!")
	}
}
