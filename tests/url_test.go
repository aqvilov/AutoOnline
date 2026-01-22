package tests

import (
	MainFuncions "FunPayAutoOnline/mainFunc"
	"testing"
)

func TestIsRealUrl(t *testing.T) {

	result := MainFuncions.IsRealUrl("https://google.com")
	if result != true {
		t.Error("https://google.com должен быть валидным URL")
	}

	result = MainFuncions.IsRealUrl("это не урл")
	if result != false {
		t.Error("'это не урл' должен быть невалидным URL")
	}

	result = MainFuncions.IsRealUrl("")
	if result != false {
		t.Error("Пустая строка должна быть невалидным URL")
	}
}
