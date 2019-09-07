package devto

import (
	"io/ioutil"
	"net/http"
)

func decodeResponse(r *http.Response) []byte {
	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte("")
	}
	defer r.Body.Close()
	return c
}
