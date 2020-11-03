package country

// ViewModel controller return object
type ViewModel struct {
	Name     string `json:"name"`
	Currency string `json:"currency"`
	Flag     string `json:"flag"` // emoji, eg. 🇯🇵
}
