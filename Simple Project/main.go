package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	webUrls := []string{
		"https://github.com/",
		"https://www.google.com",
		"https://www.youtube.com/",
		"https://go.dev/",
	}
	start := time.Now()
	for _, url := range webUrls {
		wg.Add(1)
		go getWebsiteStatus(url, wg)

	}
	wg.Add(2)
	go printJsonData(wg)
	go readCsv(wg)
	defer fmt.Println(time.Since(start))
	wg.Wait()

}
func readCsv(wg *sync.WaitGroup) {
	f, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("error occured")
	}

	// remember to close the file at the end of the program
	defer f.Close()
	defer wg.Done()

	// read csv values using csv.Reader
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
func printJsonData(wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/78")
	if err != nil {
		fmt.Println("Opps error ocurred!")
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}
func getWebsiteStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Opps error ocurred!")
	} else {
		fmt.Println(url, " status: ", res.StatusCode)
	}
}
