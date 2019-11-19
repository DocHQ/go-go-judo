package gogojudo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RegisterCardModel struct {

	// Required Fields
	CV2               string `json:"cv2"`        // CV2 from the credit card, also known as the card verification value (CVV) or security code. The 3 or 4 digit number on the back of a card
	CardNumber        string `json:"cardNumber"` // The unique number printed on a credit card, should be submitted without any whitespace or non-numeric characters
	ConsumerReference string `json:"yourConsumerReference,omitempty"`
	PaymentReference  string `json:"yourPaymentReference,omitempty"`

	// Dates
	ExpiryDate string `json:"expiryDate,omitempty"`
	StartDate  string `json:"startDate,omitempty"`

	IssueNumber int `json:"issue_number,omitempty"`

	// Address
	CardAddress string `json:"card_address,omitempty"`

	// Other Details
	ClientDetails map[string]interface{} `json:"client_details,omitempty"`
	JudoID        string                 `json:"judo_id,omitempty"`
	Currency      string                 `json:"currency,omitempty"`
}

type RegisterCardResponse struct {
}

func (jp *JudoPay) RegisterCard(rcp RegisterCardModel) (ret RegisterCardResponse, err error) {
	requestBody, err := json.Marshal(rcp)

	if err != nil {
		return ret, err
	}

	request, err := http.NewRequest(http.MethodPost, jp.APIUrl.String(), bytes.NewBuffer(requestBody))
	jp.SetHeaders(request)

	if err != nil {
		return ret, err
	}

	resp, err := jp.HttpClient.Do(request)

	if err != nil {
		return ret, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&ret)

	return ret, nil
}
