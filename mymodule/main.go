package main 

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchCatFacts() string {

	getURL := "https://cat-fact.herokuapp.com/facts"
	response, err := http.Get(getURL)
	if err != nil {
		fmt.Println("Error making get request: ", err)
		return "Error making request"
	}

	defer response.Body.Close()

	// read and print the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return "Error reading response"
	}
	return string(body)

}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		catFacts := fetchCatFacts()
		fmt.Fprintf(w, catFacts)
	})

	fmt.Println("Serving on port 5555")
	http.ListenAndServe(":5555", nil)
}