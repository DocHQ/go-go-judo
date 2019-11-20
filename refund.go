package gogojudo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

type RefundModel struct {
	// Payment identification
	ReceiptID            string `json:"receiptId"`
	YourPaymentReference string `json:"yourPaymentReference"`

	// Information relating to the transaction
	PaymentMetaData string                 `json:"yourPaymentMetaData"`
	ClientDetails   map[string]interface{} `json:"clientDetails"`

	// Finantial information
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`

	// Since it needs to be included
	JudoID string `json:"judoId"`
}

func (jp *JudoPay) Refund(rcp RefundModel) (ret PaymentReceiptModel, err error) {
	var requestURL url.URL = *jp.APIUrl
	rcp.JudoID = jp.JudopayID
	requestBody, err := json.Marshal(rcp)

	if err != nil {
		return ret, err
	}

	requestURL.Path = path.Join(requestURL.Path, "refunds")

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

		return ret, jerror.GetError()
	}

	json.NewDecoder(resp.Body).Decode(&ret)

	return ret, nil
}
