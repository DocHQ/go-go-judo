package gogojudo

import (
	"strconv"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

var JP *JudoPay

func TestNew(t *testing.T) {
	JP = New()

	if JP == New() {
		t.Errorf("Created object is not alike to base struct.")
	}

	t.Logf("New Passed.")
}

func TestCheckCard(t *testing.T) {
	t.Log("Test skipped")
	return
	var testConditions = RegisterCardModel{
		CV2:               "452",
		CardNumber:        "4976000000003436",
		ConsumerReference: "0",
		PaymentReference:  "0",
		ExpiryDate:        "12/20",
	}

	_, err := JP.CheckCard(testConditions)

	if err != nil {
		t.Errorf("CheckCard failed, error: " + err.Error())
		return
	}

	t.Logf("CheckCard Passed")
}

func TestPayments(t *testing.T) {
	ti := time.Now()

	var testPayment = CardPaymentModel{
		CV2:               "452",
		CardNumber:        "4976000000003436",
		ConsumerReference: "0",
		PaymentReference:  strconv.FormatInt(ti.Unix(), 10),
		ExpiryDate:        "12/20",

		Amount: 0.01,
	}

	res, err := JP.Payments(testPayment)

	if err != nil {
		t.Error("Payment failed, error: " + err.Error())
		return
	}

	pretty.Print(res)

	t.Logf("Payment Passed, Receipt ID: " + res.ReceiptID)

}
