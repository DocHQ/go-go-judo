package gogojudo

import (
	"time"
)

type ListTransactions struct {
	ResultCount int               `json:"resultCount"`
	PageSize    int               `json:"pageSize"`
	Offset      int               `json:"offset"`
	Results     []TransactionItem `json:"results"`
	Sort        string            `json:"sort"`
}

type TransactionItem struct {
	ReceiptID            string    `json:"receiptId"`
	OriginalReceiptID    string    `json:"originalReceiptId,omitempty"`
	YourPaymentReference string    `json:"yourPaymentReference"`
	Type                 string    `json:"type"`
	CreatedAt            time.Time `json:"createdAt"`
	Result               string    `json:"result"`
	Message              string    `json:"message"`
	JudoID               int       `json:"judoId"`
	MerchantName         string    `json:"merchantName"`
	AppearsOnStatementAs string    `json:"appearsOnStatementAs"`
	OriginalAmount       string    `json:"originalAmount,omitempty"`
	NetAmount            string    `json:"netAmount"`
	Amount               string    `json:"amount"`
	Currency             string    `json:"currency"`
	CardDetails          struct {
		CardLastfour  string `json:"cardLastfour"`
		EndDate       string `json:"endDate"`
		CardType      int    `json:"cardType"`
		CardQualifier int    `json:"cardQualifier"`
	} `json:"cardDetails"`
	Consumer struct {
		ConsumerToken         string `json:"consumerToken"`
		YourConsumerReference string `json:"yourConsumerReference"`
	} `json:"consumer"`
	YourPaymentMetaData struct {
	} `json:"yourPaymentMetaData"`
	PostCodeCheckResult string `json:"postCodeCheckResult"`
}
