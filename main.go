package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//returns a runtime-constant array of all pages to scrape
func pages() [7]string {
	return [...]string{
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-503/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-506/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-509/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-514/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-533/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-551/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/F1010-20221-514/",
	}
}


func main() {
	resp, err := http.Get("http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-001/");

	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body));
	}
}