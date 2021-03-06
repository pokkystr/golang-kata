package captcha_test

import (
	"captcha"
	"testing"
)

func TestPattern1Left1Right1(t *testing.T) {
	var expect = "1 + one"
	var result = captcha.GenCaptcha(1, 1, 1, 1)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Left1Right2(t *testing.T) {
	var expect = "1 + two"
	var result = captcha.GenCaptcha(1, 1, 1, 2)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Left1Right3(t *testing.T) {
	var expect = "1 + three"
	var result = captcha.GenCaptcha(1, 1, 1, 3)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Left1Right4(t *testing.T) {
	var expect = "1 + four"
	var result = captcha.GenCaptcha(1, 1, 1, 4)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Left2Right1(t *testing.T) {
	var expect = "2 + one"
	var result = captcha.GenCaptcha(1, 1, 2, 1)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Left3Right1(t *testing.T) {
	var expect = "3 + one"
	var result = captcha.GenCaptcha(1, 1, 3, 1)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Left3Right2(t *testing.T) {
	var expect = "3 + two"
	var result = captcha.GenCaptcha(1, 1, 3, 2)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern2Left1Right1(t *testing.T) {
	var expect = "one + 1"
	var result = captcha.GenCaptcha(2, 1, 1, 1)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern2Left5Right9(t *testing.T) {
	var expect = "five + 9"
	var result = captcha.GenCaptcha(2, 1, 5, 9)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern2Left9Right7(t *testing.T) {
	var expect = "nine + 7"
	var result = captcha.GenCaptcha(2, 1, 9, 7)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Operator1Left1Right1(t *testing.T) {
	var expect = "1 + one"
	var result = captcha.GenCaptcha(1, 1, 1, 1)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Operator2Left1Right1(t *testing.T) {
	var expect = "1 - one"
	var result = captcha.GenCaptcha(1, 2, 1, 1)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}

func TestPattern1Operator3Left1Right1(t *testing.T) {
	var expect = "1 * one"
	var result = captcha.GenCaptcha(1, 3, 1, 1)

	if expect != result {
		t.Errorf("expect not equals result! " + result)
	}
}
