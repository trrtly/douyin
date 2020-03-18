package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

//Get get 请求
func Get(uri string, queryData map[string]string) (respData []byte, err error) {
	params := url.Values{}
	for k, v := range queryData {
		params.Set(k, v)
	}

	URL, err := url.Parse(uri)
	if err != nil {
		return
	}
	URL.RawQuery = params.Encode()

	response, err := http.Get(URL.String())
	if err != nil {
		return
	}

	defer response.Body.Close()
	respData, err = ioutil.ReadAll(response.Body)
	return
}
