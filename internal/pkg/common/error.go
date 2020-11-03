package common

type HttpErrorModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	ERROR_LIST_COUNTRY = HttpErrorModel{
		Code:    "100100001",
		Message: "list country error",
	}

	ERROR_LIST_PAYER = HttpErrorModel{
		Code:    "100100002",
		Message: "list payer error",
	}
)
