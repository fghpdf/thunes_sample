package common

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"github.com/spf13/viper"
)

func NewThunesClient() *httpClient.AuthClient {
	return &httpClient.AuthClient{
		Username: viper.GetString("thunes.APIKey"),
		Password: viper.GetString("thunes.APISecret"),
		BasicUrl: viper.GetString("thunes.BasicUrl"),
	}
}
