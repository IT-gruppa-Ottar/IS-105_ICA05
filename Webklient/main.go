package main

import (
	"encoding/json"
	"strings"
	"io"
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
)

type currencyRatings struct {
	Base string `json:"base"`
	Date string `json:"date"`
	Rates struct {
		GBP float64 `json:"GBP"`
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
	} `json:"rates"`
}

func main(){
	api1 := loadUrl("http://api.fixer.io/latest?base=NOK")
	currency(api1)
}

func loadUrl(url string) string {
	fmt.Printf("Data from %s\n\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s", html)
	content := string(html)
	return content
}

//Code for currency
func currency(input string){
	dec := json.NewDecoder(strings.NewReader(input))
	for {
		var c currencyRatings

		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Currency rates as of %q from %q\n", c.Date, c.Base)
		fmt.Printf("USD: %.2f\nEUR: %.2f\nGBP: %.2f\n", c.Rates.USD, c.Rates.EUR, c.Rates.GBP)
	}
}