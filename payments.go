package gogojudo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

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

		return ret, jerror.GetError()
	}

	json.NewDecoder(resp.Body).Decode(&ret)

	return ret, nil
}
