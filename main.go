package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-001/");

	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body));
	}
}