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
	//record the start time
	start := time.Now()
	//Printing the time before ending main func
	defer fmt.Println(time.Since(start))
	//calling getWebsiteStatus for each url
	for _, url := range webUrls {
		getWebsiteStatus(url)

	}
	//Getting data for id 78 from a 100 data json api
	printJsonData()
	//Reading csv file
	readCsv()

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

/*
https://github.com/  status:  200
https://www.google.com  status:  200
https://www.youtube.com/  status:  200
https://go.dev/  status:  200
{
  "userId": 8,
  "id": 78,
  "title": "quam voluptatibus rerum veritatis",
  "body": "nobis facilis odit tempore cupiditate quia\nassumenda doloribus rerum qui ea\nillum et qui totam\naut veniam repellendus"
}
read done
6.6064231s
*/
