package gogojudo

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"os"
)

type JudoPay struct {
	HttpClient    *http.Client
	APIUrl        *url.URL
	Authorization string
}

// funcction New returns an instacce of the Judopay struct
func New() *JudoPay {
	api_url, err := url.Parse(os.Getenv("JUDOPAY_URL"))

	if err != nil {
		panic(err)
	}

	return &JudoPay{
		&http.Client{},
		api_url,
		base64.StdEncoding.EncodeToString(
			[]byte(
				os.Getenv("JUDOPAY_TOKEN") + ":" + os.Getenv("JUDOPAY_SECRET"),
			),
		),
	}
}

func (jp *JudoPay) SetHeaders(req *http.Request) error {

	req.Header.Set("Authorization", "Basic "+jp.Authorization)
	req.Header.Set("API-Version", os.Getenv("JUDOPAY_API_VERSION"))
	req.Header.Set("Content-Type", "application/json")

	return nil
}
