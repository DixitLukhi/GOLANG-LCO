package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://github.com:500/DixitLukhi/GOLANG-LCO.git?sort=true&name=min"

func main() {
	fmt.Println("Wc")

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("Type of params : %T\n", qParams)
	fmt.Println(qParams["name"])

	for _, val := range qParams {
		fmt.Println(val)
	}

	partsPfUrl := &url.URL{
		Scheme: "https",
		Host:   "applogdaiict.netlify.app",
		// Path:    "",
		// RawPath: "",
	}

	genrateUrl := partsPfUrl.String()
	fmt.Println(genrateUrl)
}
