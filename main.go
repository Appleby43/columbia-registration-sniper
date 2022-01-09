package main

import (
	"fmt"

	"github.com/Appleby43/columbia-registration-sniper/course"
	"github.com/Appleby43/columbia-registration-sniper/fetch"
)

//returns a runtime-constant array of all pages to scrape
func pages() [8]string {
	return [...]string{
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-503/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-506/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-509/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-514/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-533/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-551/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/F1010-20221-514/",
		"http://www.columbia.edu/cu/bulletin/uwb/subj/ENGL/C1010-20221-058/",
	}
}


func main() {
	for _, url := range pages() {
		fmt.Println("------------")

		html, err := fetch.FromUrl(url)

		if err != nil {
			fmt.Println(err)
			continue
		}

		course, err := course.StripFrom(html)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(&course)
	}
}