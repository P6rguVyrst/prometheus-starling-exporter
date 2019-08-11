package main

import (
	"os"
	"fmt"
	//"encoding/json"
	"io/ioutil"
	"testing"
)
import "github.com/tidwall/gjson"


func TestHello(t *testing.T) {
	// t.Fatal("not implemented")
	transactions_data := "data/transactions.json"
	jsonFile, err := os.Open(transactions_data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened", transactions_data)
	
	defer jsonFile.Close()
	fmt.Println(jsonFile)	
	byteValue, _ := ioutil.ReadAll(jsonFile)
	
	data := string([]byte(byteValue))
	transaction_count := gjson.Get(data, "_embedded.transactions.#")
	transactions := gjson.Get(data, "_embedded.transactions")

	fmt.Println(transaction_count)

	transactions.ForEach(func(key, value gjson.Result) bool {
		x := gjson.Get(value.String(), "amount") 
		fmt.Println(x)
		return true
	})

}
