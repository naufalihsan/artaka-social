package wrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

func Call(method string, url string, header *http.Header, formData *bytes.Buffer, body interface{}, result interface{}) *error {
	reqBody := []byte("")
	var err error
	var payload *bytes.Buffer
	var client = &http.Client{}

	isParamsNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())
	if !isParamsNil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			panic(err)
		}
	}

	if formData != nil {
		payload = formData
	} else {
		payload = bytes.NewBuffer(reqBody)
	}

	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		panic(err)
	}

	if header != nil {
		request.Header = *header
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		panic(err)
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		panic(err)
	}

	return nil

}
