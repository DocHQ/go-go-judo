package gogojudo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func (jp *JudoPay) ListTransactions(pageSize int, offset int, sortOrder string) (ret ListTransactions, err error) {
	var requestURL url.URL = *jp.APIUrl

	if err != nil {
		return ret, err
	}

	q := requestURL.Query()
	q.Set("pageSize", strconv.Itoa(pageSize))
	q.Set("offset", strconv.Itoa(offset))
	q.Set("sort", sortOrder)

	requestURL.RawQuery = q.Encode()

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
