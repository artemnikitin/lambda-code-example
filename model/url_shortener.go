package model

// URLShortenerRequest represents request for URLShortenerHandler
// JSON looks like:
// {
//   "url": "http://example.com/vvv/bbb/ccc/1212/xxx.html"
// }
type URLShortenerRequest struct {
	URL string `json:"url"`
}

// URLShortenerResponse represents response for URLShortenerHandler
// JSON looks like:
// {
//   "shortUrl": "http://sh.cc/wX2S1qz",
//   "result": "success"
// }
type URLShortenerResponse struct {
	ShortURL string `json:"shortUrl"`
	Result   string `json:"result"`
}
