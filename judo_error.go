package gogojudo

type JudoError struct {
	Message  string
	Code     int
	Category int
}

func (je *JudoError) GetError() error {
	return ErrorMap[je.Code]
}
