package env

import (
	"os"
)

type ApiToken struct {
	API_URL   string
	API_TOKEN string
}

func GetApiToken() *ApiToken {
	token := os.Getenv("KOMITER_API_TOKEN")
	apiUrl := os.Getenv("KOMITER_API_URL")

	return &ApiToken{
		API_URL:   apiUrl,
		API_TOKEN: token,
	}
}
