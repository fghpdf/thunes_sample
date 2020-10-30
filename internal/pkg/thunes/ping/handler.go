package ping

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"github.com/spf13/viper"
	"net/http"
)

// check connectivity
func Send() (*Model, error) {
	url := viper.GetString("thunes.basicUrl") + "/ping"
	response, err := httpClient.Do(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	pingModel := &Model{}
	err = json.NewDecoder(response.Body).Decode(pingModel)
	if err != nil {
		return nil, err
	}
	return pingModel, nil
}
