package model

// URLShortenerRequest represents request for URLShortenerHandler
type URLShortenerRequest struct {
	URL string `json:"url"`
}

// URLShortenerResponse represents response for URLShortenerHandler
type URLShortenerResponse struct {
	ShortURL string `json:"shortUrl"`
	Result   string `json:"result"`
}
