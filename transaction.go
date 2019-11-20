package gogojudo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

func (jp *JudoPay) Transaction(receiptID string) (ret PaymentReceiptModel, err error) {
	var requestURL url.URL = *jp.APIUrl

	if err != nil {
		return ret, err
	}

	requestURL.Path = path.Join(requestURL.Path, receiptID)

	request, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
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
