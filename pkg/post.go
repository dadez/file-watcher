package post

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// post json body
// func main() {
// 	reqBody, err := json.Marshal(map[string]string{
// 		"username": "foo",
// 		"email":    "foo@bar.com",
// 	})
// 	if err != nil {
// 		print(err)
// 	}
// 	resp, err := http.Post("https://httpbin.org/post",
// 		"application/json", bytes.NewBuffer(reqBody))
// 	if err != nil {
// 		print(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		print(err)
// 	}
// 	fmt.Println(string(body))
// }

// post file
func PostFile() {

	url := "https://httpbin.org/post"

	file, err := os.Open("/tmp/testfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(file))
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

// another way
// func main() {

// 	url := "https://httpbin.org/post"

// 	file, err := os.Open("/tmp/testfile")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	resp, err := http.Post(url,
// 		"application/json", file)
// 	if err != nil {
// 		print(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		print(err)
// 	}
// 	fmt.Println(string(body))
// }
