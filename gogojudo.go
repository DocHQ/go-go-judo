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
	JudopayID     string
}

// funcction New returns an instacce of the Judopay struct
func New() *JudoPay {
	api_url, err := url.Parse(getEnv("JUDOPAY_URL", "https://gw1.judopay-sandbox.com/transactions"))

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
		os.Getenv("JUDOPAY_ID"),
	}
}

func (jp *JudoPay) SetHeaders(req *http.Request) error {

	req.Header.Set("Authorization", "Basic "+jp.Authorization)
	req.Header.Set("API-Version", getEnv("JUDOPAY_API_VERSION", "5.6"))
	req.Header.Set("Content-Type", "application/json")

	return nil
}

func getEnv(name string, fallback string) string {
	val := os.Getenv(name)

	if len(val) == 0 {
		return fallback
	}

	return val
}
