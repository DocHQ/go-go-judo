package gogojudo

import (
	"testing"
)

func TestNew(t *testing.T) {
	var jp *JudoPay

	if jp == New() {
		t.Errorf("Created object is not alike to base struct.")
	}

	t.Logf("New Passed.")

}

func TestCheckCard(t *testing.T) {
	var jp *JudoPay = New()

	var testConditions = RegisterCardModel{
		CV2:               "452",
		CardNumber:        "4976000000003436",
		ConsumerReference: "0",
		PaymentReference:  "0",
		ExpiryDate:        "12/20",
	}

	_, err := jp.CheckCard(testConditions)

	if err != nil {
		t.Errorf("CheckCard failed, error: " + err.Error())
		return
	}

	t.Logf("CheckCard Passed")
}
