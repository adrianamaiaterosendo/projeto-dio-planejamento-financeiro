package utils

import (
	"testing"
)

func TestStringToTime(t *testing.T) {

	convertedTime, err := StringToTime("2020-02-12T10:10:10")

	if convertedTime.Year() != 2020 {
		t.Errorf("Espera que o ano seja 2020")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Espera que o mês seja 02")
	}

	if convertedTime.Day() != 12 {
		t.Errorf("Espera que o dia seja 12")
	}

	if err != nil {
		t.Errorf("Espera que não haja erro")
	}

	_, err = StringToTime("5020-52-110:10:10")

	if err == nil {
		t.Errorf("Espera que haja erro")
	}

}
