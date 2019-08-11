package main

import "testing"
import "os"
import "fmt"
import "encoding/json"
import "io/ioutil"




type Transactions struct {
	Links
	Id string `json:"id"` 
	Currency string `json:"currency"` 
	//Links string
	Embedded string `json:"_embedded"`
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
	//fmt.Println(transactions)


	//x := result["_embedded"].(map[string]interface{})
	//y := x["transactions"].([]interface{})

	//transactions := []Transactions{}
	//err := json.Unmarshal(y, &transactions)

	//transactions := make([]interface{}*Transactions)
	//fmt.Println(transactions)
	//fmt.Println(transactions)
	//for _, p := range x {
	//	fmt.Println(p["id"])
	//}

}
