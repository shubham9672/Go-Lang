package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	webUrls := []string{
		"https://github.com/",
		"https://www.google.com",
		"https://www.youtube.com/",
		"https://go.dev/",
	}
	start := time.Now()

	for _, url := range webUrls {
		getWebsiteStatus(url)

	}
	printJsonData()
	readCsv()
	defer fmt.Println(time.Since(start))

}
func readCsv() {
	f, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("error occured")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	flag := true
	for {
		_, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			flag = false
		}
	}
	if flag {
		fmt.Println("read done")
	}
}
func printJsonData() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/78")
	if err != nil {
		fmt.Println("Opps error ocurred!")
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}
func getWebsiteStatus(url string) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Opps error ocurred!")
	} else {
		fmt.Println(url, " status: ", res.StatusCode)
	}
}
