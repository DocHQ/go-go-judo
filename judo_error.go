package gogojudo

type JudoError struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	Category int    `json:"category"`

	Details []struct {
		Code      int    `json:"code"`
		FieldName string `json:"fieldName"`
		Message   string `json:"message"`
	} `json:"details"`
}

func (je *JudoError) GetError() error {
	return ErrorMap[je.Code]
}
