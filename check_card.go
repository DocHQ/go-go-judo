package gogojudo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type CheckCardResponse struct {
}

func (jp *JudoPay) CheckCard(rcp RegisterCardModel) (ret CheckCardResponse, err error) {
	var requestURL *url.URL = jp.APIUrl
	requestBody, err := json.Marshal(rcp)

	if err != nil {
		return ret, err
	}

	requestURL.Path = path.Join(requestURL.Path, "checkcard")

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

		fmt.Printf("%+v", jerror)

		return ret, jerror.GetError()
	}

	var response map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&response)

	fmt.Printf("\n\n%+v\n", response)

	return ret, nil
}
