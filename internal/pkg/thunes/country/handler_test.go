package country

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/httpClient"
	"net/http"
	"testing"
)

func TestList(t *testing.T) {
	server := common.ServerMock("/v2/money-transfer/countries", listSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	params := &ListParams{
		Page:    0,
		PerPage: 50,
	}

	res, err := List(authClient, params)
	if err != nil {
		t.Error(err)
	}

	if res == nil || len(*res) != 3 {
		t.Errorf("expected 3 countries but got %d\n", len(*res))
	}

	if (*res)[0].IsoCode != "KEN" {
		t.Errorf("expected KEN but got %s\n", (*res)[0].IsoCode)
	}
}

func listSuccessMock(w http.ResponseWriter, r *http.Request) {
	countries := &[3]Model{
		{
			Name:    "Kenya",
			IsoCode: "KEN",
		},
		{
			Name:    "China",
			IsoCode: "CHN",
		},
		{
			Name:    "Singapore",
			IsoCode: "SGP",
		},
	}

	res, _ := json.Marshal(countries)
	_, _ = w.Write(res)
}
