package tests

import (
	"bytes"
	"encoding/json"
	"github.com/vschettino/exfin/router"
	"net/http"
	"net/http/httptest"
)

func makeBody(body map[string]string) *bytes.Buffer{
	j, _ := json.Marshal(body)
	return bytes.NewBuffer(j)
}


func MakePOST(url string, body map[string]string)( *httptest.ResponseRecorder, error){
	r := router.Router()
	w := httptest.NewRecorder()
	bodyBuffer := makeBody(body)
	req, err := http.NewRequest("POST", url, bodyBuffer)
	if err != nil{
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w, err

}
