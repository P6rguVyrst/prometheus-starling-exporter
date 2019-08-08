package main

import "testing"
import "os"
import "fmt"
import "encoding/json"
import "io/ioutil"


type Transactions struct {
	Id string `json:"id"` 
	Currency string `json:"currency"` 
	//Links string
	//Embedded string `json:"_embedded"`
}


func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	transactions_data := "data/transactions.json"
	jsonFile, err := os.Open(transactions_data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened", transactions_data)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	x := result["_embedded"].(map[string]interface{})
	fmt.Println(x)	
	transactions := make(map[string][]*Transactions)
	fmt.Println(transactions)
	for _, p := range x {
		fmt.Println(p["id"])
	}

	//fmt.Println(result["_links"])
	//fmt.Println(result["_embedded"]["transactions"])
	//Map := make(map[string]interface{})	
	//hm := Map["_embedded"].([]Transactions)
	//fmt.Println(hm[0])
}
