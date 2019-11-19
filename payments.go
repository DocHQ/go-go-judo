package gogojudo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

type CardPaymentModel struct {
	// Required Fields
	CV2        string `json:"cv2"`        // CV2 from the credit card, also known as the card verification value (CVV) or security code. The 3 or 4 digit number on the back of a card
	CardNumber string `json:"cardNumber"` // The unique number printed on a credit card, should be submitted without any whitespace or non-numeric characters
	ExpiryDate string `json:"expiryDate,omitempty"`
	StartDate  string `json:"startDate,omitempty"`

	ConsumerReference string `json:"yourConsumerReference,omitempty"`
	PaymentReference  string `json:"yourPaymentReference,omitempty"`
	PaymentMetaData   []byte `json:"yourPaymentMetaData,omitempty"`

	Amount                  float64 `json:"amount,omitempty"`
	Currency                string  `json:"currency,omitempty"`
	InitialRecurringPayment bool    `json:"initialRecurringPayment,omitempty"`
	RecurringPayment        bool    `json:"recurringPayment,omitempty"`

	IssueNumber         int    `json:"issueNumber,omitempty"`
	WebPaymentReference string `json:"webPaymentReference.omitempty"`

	// Address
	CardAddress      string                 `json:"card_address,omitempty"`
	MobileNumber     string                 `json:"mobileNumber,omitempty"`
	EmailAddress     string                 `json:"emailAddress,omitempty"`
	ConsumerLocation map[string]interface{} `json:"consumerLocation,omitempty"`

	UserAgent      string                 `json:"userAgent,omitempty"`
	DeviceCategory string                 `json:"deviceCategory,omitempty"`
	AcceptHeaders  string                 `json:"acceptHeaders,omitempty"`
	ClientDetails  map[string]interface{} `json:"clientDetails,omitempty"`

	JudoID string `json:"judoId,omitempty"`
}

type PaymentsResponse struct {
	ReceiptID            string    `json:"receiptId"`
	YourPaymentReference string    `json:"yourPaymentReference"`
	Type                 string    `json:"type"`
	CreatedAt            time.Time `json:"createdAt"`
	Result               string    `json:"result"`
	Message              string    `json:"message"`
	JudoID               int       `json:"judoId"`
	MerchantName         string    `json:"merchantName"`
	AppearsOnStatementAs string    `json:"appearsOnStatementAs"`
	OriginalAmount       string    `json:"originalAmount"`
	NetAmount            string    `json:"netAmount"`
	Amount               string    `json:"amount"`
	Currency             string    `json:"currency"`
	CardDetails          struct {
		CardLastfour  string `json:"cardLastfour"`
		EndDate       string `json:"endDate"`
		CardToken     string `json:"cardToken"`
		CardType      int    `json:"cardType"`
		CardScheme    string `json:"cardScheme"`
		CardFunding   string `json:"cardFunding"`
		CardCategory  string `json:"cardCategory"`
		CardQualifier int    `json:"cardQualifier"`
		CardCountry   string `json:"cardCountry"`
		Bank          string `json:"bank"`
	} `json:"cardDetails"`
	Consumer struct {
		ConsumerToken         string `json:"consumerToken"`
		YourConsumerReference string `json:"yourConsumerReference"`
	} `json:"consumer"`
	RiskScore int `json:"riskScore"`
	Risks     struct {
		PostCodeCheck      string `json:"postCodeCheck"`
		MerchantSuggestion string `json:"merchantSuggestion"`
		MerchantStatistics struct {
		} `json:"merchantStatistics"`
	} `json:"risks"`
}

func (jp *JudoPay) Payments(rcp CardPaymentModel) (ret PaymentsResponse, err error) {
	var requestURL *url.URL = jp.APIUrl
	rcp.JudoID = jp.JudopayID
	requestBody, err := json.Marshal(rcp)

	if err != nil {
		return ret, err
	}

	requestURL.Path = path.Join(requestURL.Path, "payments")

	request, err := http.NewRequest(http.MethodPost, requestURL.String(), bytes.NewBuffer(requestBody))
	jp.SetHeaders(request)

	if err != nil {
		return ret, err
	}

	resp, err := jp.HttpClient.Do(request)

	if err != nil {
		return ret, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var jerror = &JudoError{}
		json.NewDecoder(resp.Body).Decode(&jerror)

		fmt.Printf("%+v", jerror)

		return ret, jerror.GetError()
	}

	json.NewDecoder(resp.Body).Decode(&ret)

	return ret, nil
}
