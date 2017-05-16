package main

import (
	"html/template"
	"net/http"
	"path"
	"log"
	"io/ioutil"
	"encoding/json"
)

type currencyRatings struct {
	Base string `json:"base"`
	Date string `json:"date"`
	Rates struct {
		AUD float64 `json:"AUD"`
		BGN float64 `json:"BGN"`
		BRL float64 `json:"BRL"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		CNY float64 `json:"CNY"`
		CZK float64 `json:"CZK"`
		DKK float64 `json:"DKK"`
		GBP float64 `json:"GBP"`
		HKD float64 `json:"HKD"`
		HRK float64 `json:"HRK"`
		HUF float64 `json:"HUF"`
		IDR float64 `json:"IDR"`
		ILS float64 `json:"ILS"`
		INR float64 `json:"INR"`
		JPY float64 `json:"JPY"`
		KRW float64 `json:"KRW"`
		MXN float64 `json:"MXN"`
		MYR float64 `json:"MYR"`
		NZD float64 `json:"NZD"`
		PHP float64 `json:"PHP"`
		PLN float64 `json:"PLN"`
		RON float64 `json:"RON"`
		RUB float64 `json:"RUB"`
		SEK float64 `json:"SEK"`
		SGD float64 `json:"SGD"`
		THB float64 `json:"THB"`
		TRY float64 `json:"TRY"`
		USD float64 `json:"USD"`
		ZAR float64 `json:"ZAR"`
		EUR float64 `json:"EUR"`
	} `json:"rates"`
}

func main() {
	go http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	i := GetCurrency()

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCurrency() *currencyRatings{
	url , err := http.Get("http://api.fixer.io/latest?base=NOK")
	if err != nil{
		log.Fatal(err)
	}

	byte, err := ioutil.ReadAll(url.Body)
	if err != nil{
		log.Fatal(err)
	}

	clubs := &currencyRatings{}
	if err := json.Unmarshal(byte, &clubs); err != nil{
		log.Fatal(err)
	}
	return clubs
}