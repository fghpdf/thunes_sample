package httpClient

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"net/http"
	"testing"
)

func TestDoWithoutKeyOrSecret(t *testing.T) {
	_, err := Do(http.MethodGet, "https://api-mt.thunes.com/connect", nil)

	if err.Error() != common.MUST_INIT_HTTP_CLIENT {
		t.Errorf("check auth info failed, now is %v", err)
	}
}
