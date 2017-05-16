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

/**
Struckt 1 - Viser konversjosnraten fra norsk kr til Britisk pund, US dollar og euro
 */
type CurrencyRatings struct {
	Base string `json:"base"`
	Date string `json:"date"`
	Rates struct {
		GBP float64 `json:"GBP"`
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
	} `json:"rates"`
}

/**
Struckt 2 - Viser antall påloggede Steam-spillere
 */
type Steam struct {
	Response struct {
		PlayerCount int `json:"player_count"`
		Result int `json:"result"`
	} `json:"response"`
}

/**
Struckt 3 - Viser antall påloggede spiller på Battlefield
og det største antallet som har vært pålogget de siste 24t.
 */
type Battlefield struct {
	Response struct {
		PC int `json:"count"`
		Peak24 int `json:"peak24"`
	} `json:"pc"`
}

/**
Struckt 4 - Viser hvor mange mennesker som befinner seg i verdensrommet
 */
type SpaceType struct {
	People []struct {
		Craft string `json:"craft"`
		Name string `json:"name"`
	} `json:"people"`
	Message string `json:"message"`
	Number int `json:"number"`
}

/**
Struckt 5 - Viser informasjon om Batman filmen "The Dark Knight"
 */
type DarkKnight struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
}



/**
Main - Laster inn dataene igjennom "loadUrl" og lagrer de i variabler som brukes for å kjøre printefunksjonene
 */
func main(){
	api1 := loadUrl("http://api.fixer.io/latest?base=NOK")
	Currency(api1)
	fmt.Println("---------------------------------------------------------------")

	api2 := loadUrl("https://api.steampowered.com/ISteamUserStats/GetNumberOfCurrentPlayers/v0001/?format=json&appid=0")
	SteamPlayers(api2)
	fmt.Println("---------------------------------------------------------------")

	api3 := loadUrl("http://api.bf4stats.com/api/onlinePlayers?output=json")
	BattlefieldPlayers(api3)
	fmt.Println("---------------------------------------------------------------")

	api4 := loadUrl("http://api.open-notify.org/astros.json")
	Space(api4)
	fmt.Println("---------------------------------------------------------------")

	api5 := loadUrl("http://www.omdbapi.com/?t=the+dark+knight&y=&plot=short&r=json")
	TheDarkKnight(api5)
	fmt.Println("---------------------------------------------------------------")
}

/**
loadUrl
@param String - En url som inneholder json
@return String - Returner innholdet i url-en i en string som kan brukes til å printe data med
 */
func loadUrl(url string) string {
	//Printer ut url-en som er brukt
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

/**
Funksjonene under bruker samme "mal", men har sine egne tilpasninger for å printe ut "relevant" infromajson istedenfor json.
Funksjonene henter informasjonen via struckt-ene.
@param String - Stringen som returneres fra "loadUrl"
 */
func Currency(input string){
	dec := json.NewDecoder(strings.NewReader(input))
	for {
		var c CurrencyRatings

		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Currency rates as of %q from %q\n", c.Date, c.Base)
		fmt.Printf("USD: %.2f\nEUR: %.2f\nGBP: %.2f\n", c.Rates.USD, c.Rates.EUR, c.Rates.GBP)
	}
}

func SteamPlayers(input string){
	dec := json.NewDecoder(strings.NewReader(input))
	for {
		var c Steam

		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Steam Players: ")
		fmt.Println(c.Response.PlayerCount)
	}
}

func BattlefieldPlayers(input string){
	dec := json.NewDecoder(strings.NewReader(input))
	for {
		var c Battlefield

		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Battlefield Players: ")
		fmt.Println( c.Response.PC)
		fmt.Print("Battlefield peak last 24h: ")
		fmt.Println( c.Response.Peak24)
	}
}

func TheDarkKnight(input string){
	dec := json.NewDecoder(strings.NewReader(input))
	for {
		var c DarkKnight

		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Movie Information")
		fmt.Println("Movie Title: ", c.Title)
		fmt.Println("Rated: ", c.Rated)
		fmt.Println("Release Date: ", c.Released)
		fmt.Println("Release Year: ", c.Year)
	}
}

func Space(input string){
	dec := json.NewDecoder(strings.NewReader(input))
	for {
		var c SpaceType

		if err := dec.Decode(&c); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Print("People in space: ")
		fmt.Println( c.Number)
	}
}