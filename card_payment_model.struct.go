package gogojudo

// CardPaymentModel is a base structure that provides everything needed to
// submit a payment request
type CardPaymentModel struct {
	// Required card details
	CV2         string `json:"cv2,omitempty"`
	CardNumber  string `json:"cardNumber,omitempty"`
	CardToken   string `json:"cardToken,omitempty"`
	OneUseToken string `json:"oneUseToken,omitempty"`
	ExpiryDate  string `json:"expiryDate,omitempty"`
	StartDate   string `json:"startDate,omitempty"`

	// Reference numbers
	ConsumerReference    string `json:"yourConsumerReference,omitempty"`
	YourPaymentReference string `json:"yourPaymentReference,omitempty"`
	PaymentMetaData      []byte `json:"yourPaymentMetaData,omitempty"`

	// Payment details
	Amount                  float64 `json:"amount,omitempty"`
	Currency                string  `json:"currency,omitempty"`
	InitialRecurringPayment bool    `json:"initialRecurringPayment,omitempty"`
	RecurringPayment        bool    `json:"recurringPayment,omitempty"`
	RecurringPaymentType    string  `json:"recurringPaymentType,omitempty"`

	// Other reference information
	IssueNumber         int    `json:"issueNumber,omitempty"`
	WebPaymentReference string `json:"webPaymentReference.omitempty"`

	// Cardholder information
	CardAddress      string                 `json:"card_address,omitempty"`
	MobileNumber     string                 `json:"mobileNumber,omitempty"`
	EmailAddress     string                 `json:"emailAddress,omitempty"`
	ConsumerLocation map[string]interface{} `json:"consumerLocation,omitempty"`

	// Misc. information
	UserAgent      string                 `json:"userAgent,omitempty"`
	DeviceCategory string                 `json:"deviceCategory,omitempty"`
	AcceptHeaders  string                 `json:"acceptHeaders,omitempty"`
	ClientDetails  map[string]interface{} `json:"clientDetails,omitempty"`

	// Merchant JudoID
	JudoID string `json:"judoId,omitempty"
}
