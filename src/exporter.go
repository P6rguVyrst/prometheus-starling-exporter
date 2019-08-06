package main
import "./api"
import "os"
import "log"
import "encoding/json"

type AccountResponseObject struct {
	accounts	[]interface{}
}

func alpha(body string){
	var asd AccountResponseObject
	//asd := make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &asd)
	if err != nil {
		log.Fatal(err)
	}
	log.Println([]byte(body))
	log.Println(asd)
	log.Println(asd.accounts)
}

func main() {
	api := api.TokenAPI {
		Token: os.Getenv("FOO_TOKEN"),
		Url: os.Getenv("FOO_API"),
		Version: "/v2",
	}
	x := api.Get("/accounts")
	alpha(x)
	log.Println(x)
}
