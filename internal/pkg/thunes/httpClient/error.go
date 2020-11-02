package httpClient

type HttpErrorModel struct {
	Errors []HttpErrorInfoModel `json:"errors"`
}

type HttpErrorInfoModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
