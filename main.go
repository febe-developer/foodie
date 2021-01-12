package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

///Consume RESTFull API provided
func main(){
	response, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s/n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}