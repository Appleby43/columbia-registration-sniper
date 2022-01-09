package fetch

import (
	"io/ioutil"
	"net/http"
)

//FromUrl fetches html via network from a url.
func FromUrl(url string) (html string, err error) {
	resp, err := http.Get(url);

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	return string(body), nil
}