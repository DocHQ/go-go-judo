package gogojudo

import (
	"time"
)

type PaymentReceiptModel struct {
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
	
	Result  string `json:"result,omitempty"`
	Message string `json:"message,omitempty"`
	AcsURL  string `json:"acsUrl,omitempty"`
	MD      string `json:"md,omitempty"`
	PaReq   string `json:"paReq,omitempty"`
	TermURL string `json:"termUrl,omitempty"`
	
	// 3D Secure
	PaymentRequires3DSecure struct {
		Result  string `json:"result,omitempty"`
		Message string `json:"message,omitempty"`
		AcsURL  string `json:"acsUrl,omitempty"`
		MD      string `json:"md,omitempty"`
		PaReq   string `json:"paReq,omitempty"`
		TermURL string `json:"termUrl,omitempty"`
	} `json:"paymentRequires3DSecure"`
}
