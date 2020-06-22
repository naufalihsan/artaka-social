package wrapper

import (
	"bytes"
	"mime/multipart"
	"strings"
)

func RemoveTrail(code string) string {
	return strings.Replace(code, "#_", "", -1)
}

func ExchangeFormData(data *ExchangeParams) (*bytes.Buffer, string) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("client_id", data.ClientID)
	_ = writer.WriteField("client_secret", data.ClientSecret)
	_ = writer.WriteField("grant_type", data.GrantType)
	_ = writer.WriteField("redirect_uri", data.RedirectURI)
	_ = writer.WriteField("code", data.Code)
	err := writer.Close()

	contentType := writer.FormDataContentType()

	if err != nil {
		panic(err)
	}

	return payload, contentType
}
