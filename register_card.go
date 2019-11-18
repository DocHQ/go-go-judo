package gogojudo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RegisterCardModel struct {

	// Required Fields
	CV2               string // CV2 from the credit card, also known as the card verification value (CVV) or security code. The 3 or 4 digit number on the back of a card
	CardNumber        string // The unique number printed on a credit card, should be submitted without any whitespace or non-numeric characters
	ConsumerReference string
	PaymentReference  string

	// Dates
	ExpiryDate string
	StartDate  string

	IssueNumber int

	// Address
	CardAddress struct{}

	// Other Details
	ClientDetails map[string]interface{}
	JudoID        string
	Currency      string
}

type RegisterCardResponse struct {
}

func (jp *JudoPay) RegisterCard(rcp RegisterCardModel) (ret RegisterCardResponse, err error) {
	requestBody, err := json.Marshal(rcp)

	if err != nil {
		return ret, err
	}

	request, err := http.NewRequest(http.MethodPost, jp.APIUrl.String(), bytes.NewBuffer(requestBody))
	request.Header.Set("Authorization", "Basic"+jp.Authorization)

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
