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

	ERROR_CREATE_TRANSACTION = HttpErrorModel{
		Code:    "100100003",
		Message: "create transaction error",
	}

	ERROR_CREATE_TRANSACTION_BIND = HttpErrorModel{
		Code:    "100100004",
		Message: "bind create transaction params error",
	}

	ERROR_CREATE_QUOTATION = HttpErrorModel{
		Code:    "100100005",
		Message: "create quotation error",
	}

	ERROR_CREATE_QUOTATION_BIND = HttpErrorModel{
		Code:    "100100006",
		Message: "bind create quotation params error",
	}
)
