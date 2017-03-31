package main

import (
	"fmt"
	"time"
)

func main() {


	//t := time.Date(2017, time.April, 30, 13, 21, 0, 0, time.UTC)
	//sfmt.Printf("Go launched at %s\n", t.Local())



	fmt.Println(time.Now().Format(time.RFC850))




}

