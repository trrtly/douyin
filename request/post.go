package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// PostJSON post json 数据请求
func PostJSON(uri string, data interface{}) (response *http.Response, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)

	body := bytes.NewBuffer(jsonData)
	response, err = http.Post(uri, "application/json;charset=utf-8", body)
	return
}
