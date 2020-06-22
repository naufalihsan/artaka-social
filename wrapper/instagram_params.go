package wrapper

type ExchangeParams struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
	Code         string `json:"code"`
}

type ExchangeResult struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
}

type QueryMediaResult struct {
	Data   []MediaItem     `json:"data"`
	Paging MediaPagination `json:"paging"`
}

type QueryMediaNodeResult struct {
	ID        string `json:"id"`
	MediaType string `json:"media_type"`
	MediaURL  string `json:"media_url"`
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
}

type MediaItem struct {
	ID      string `json:"id"`
	Caption string `json:"caption"`
}

type MediaPagination struct {
	Cursors PaginationCursor `json:"cursors"`
	Next    string           `json:"next"`
}

type PaginationCursor struct {
	After  string `json:"after"`
	Before string `json:"before"`
}
