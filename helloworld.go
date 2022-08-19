package main

import(
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	res, _:=http.Get("https://api.ipfy.org")
	ip, _:=ioutil.ReadAll(res.Body)
	os.Stdout.Write(ip)
}