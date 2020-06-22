package wrapper

import (
	"fmt"
	"net/http"
)

func ExchangeCodeForToken(data *ExchangeParams) (*ExchangeResult, *error) {
	response := &ExchangeResult{}
	header := &http.Header{}

	accessTokenURL := "https://api.instagram.com/oauth/access_token"
	formData, contentType := ExchangeFormData(data)

	header.Set("Content-Type", contentType)

	err := Call(
		"POST",
		accessTokenURL,
		header,
		formData,
		data,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func QueryUserMedia(accessToken string) (*QueryMediaResult, *error) {
	response := &QueryMediaResult{}
	header := &http.Header{}

	queryMediaURL := fmt.Sprintf(
		"https://graph.instagram.com/me/media?fields=id,caption&access_token=%s", accessToken)

	err := Call(
		"GET",
		queryMediaURL,
		header,
		nil,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func QueryMediaNode(mediaID string, accessToken string) (*QueryMediaNodeResult, *error) {
	response := &QueryMediaNodeResult{}
	header := &http.Header{}

	queryMediaURL := fmt.Sprintf(
		"https://graph.instagram.com/%s?fields=id,media_type,media_url,username,timestamp&access_token=%s",
		mediaID, accessToken,
	)

	err := Call(
		"GET",
		queryMediaURL,
		header,
		nil,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
