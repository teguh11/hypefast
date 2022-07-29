package models

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenData struct {
	URL           string `json:"url"`
	ShortenValue  string `json:"shorten_value"`
	CreatedAt     string `json:"created_at"`
	RedirectCount int    `json:"redirect_count"`
}
