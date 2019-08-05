package main
import "./api"
import "os"
import "log"


func main() {
	api := api.TokenAPI {
		Token: os.Getenv("FOO_TOKEN"),
		Url: os.Getenv("FOO_API"),
		Version: "/v2",
	}
	x := api.Get("/accounts")
	log.Println(x)
}
