package gogojudo

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var JP *JudoPay

func TestNew(t *testing.T) {
	JP = New()

	if JP == New() {
		t.Errorf("Created object is not alike to base struct.")
	}

	t.Logf("New Passed.")
}

func TestPayments(t *testing.T) {
	// Using the current unit nanosecond as a rand seed since
	// rand.Intn() uses the same default seed so you get duplicated
	// in sequestial test runs
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var testPayment = CardPaymentModel{
		CV2:                  "452",
		CardNumber:           "4976000000003436",
		ConsumerReference:    "0",
		YourPaymentReference: strconv.Itoa(r1.Intn(10000)),
		ExpiryDate:           "12/20",

		Amount: 1,
	}

	payment, err := JP.Payments(testPayment)

	if err != nil {
		t.Error("Payment failed, error: " + err.Error())
		return
	}

	t.Logf("Payment Passed, Receipt ID: " + payment.ReceiptID)

	// Get transaction test
	t.Run("Get", func(t *testing.T) {
		res, err := JP.Transaction(payment.ReceiptID)

		if err != nil {
			t.Errorf("Failed, error: " + err.Error())
			return
		}

		if res.ReceiptID != "" {
			t.Log("Pass, receipt ID: " + res.ReceiptID)
		}

	})

	// Refund half the amount
	t.Run("Partial Refund", func(t *testing.T) {
		res, err := JP.Refund(RefundModel{
			ReceiptID:            payment.ReceiptID,
			YourPaymentReference: strconv.Itoa(r1.Intn(10000)),
			Amount:               0.50,
		})

		if err != nil {
			t.Errorf("Failed, error: " + err.Error())
			return
		}

		if res.ReceiptID != "" {
			t.Log("Pass, receipt ID: " + res.ReceiptID)
		}

	})

	// Refund the remaining amount for a full refund
	t.Run("Finish Refund", func(t *testing.T) {
		res, err := JP.Refund(RefundModel{
			ReceiptID:            payment.ReceiptID,
			YourPaymentReference: strconv.Itoa(r1.Intn(10000)),
			Amount:               0.50,
		})

		if err != nil {
			t.Errorf("Failed, error: " + err.Error())
			return
		}

		if res.ReceiptID != "" {
			t.Log("Pass, receipt ID: " + res.ReceiptID)
		}

	})

}

func TestTransactions(t *testing.T) {

	res, err := JP.ListTransactions(10, 0, TimeAscending)

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Pass: %v Transactions", res.ResultCount)
}
