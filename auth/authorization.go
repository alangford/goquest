package auth

import (
	"log"
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"os"
	"io/ioutil"
)

func loginURI(c string) url.Values {
	cid := os.Getenv("AUTH_CLIENT_ID")
	rrURI := os.Getenv("AUTH_REDIRECT_URI")
	cs := os.Getenv("AUTH_CLIENT_SECRET")
	v := url.Values{}
  	v.Set("grant_type", "authorization_code")
  	v.Add("client_id", cid)
  	v.Add("code", c)
		v.Add("redirect_uri", rrURI)
		v.Add("resource", "https://graph.windows.net")
		v.Add("client_secret", cs)
	return v
}

func Login () func(http.ResponseWriter, *http.Request) {
	fmt.Println("attempting to create issuer")
	return func (w http.ResponseWriter, r *http.Request) {
		
		defer r.Body.Close()
		var azure Azure
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&azure)
		if err != nil {
			panic(err)
	}
		tenent := os.Getenv("AUTH_TENENT")
		resp, err := http.PostForm(tenent + "/token", loginURI(azure.Code))
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

	}
}