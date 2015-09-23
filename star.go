package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mikeqian/log"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var cookie string
var message = make(chan int)

type Config struct {
	Cookie string
}

type Ing struct {
	Content    string
	PublicFlag int32
}

func getLastIng() (text string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://ing.cnblogs.com/ajax/ing/GetIngList?IngListType=my&PageIndex=1&PageSize=1&Tag=&_=1441948524646", nil)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}

func insertIng(i int) {
	ing := Ing{}
	ing.PublicFlag = 1

	rand.Seed(time.Now().Unix())

	ing.Content = "mm" + strconv.Itoa(rand.Intn(100))
	text, _ := json.Marshal(ing)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://ing.cnblogs.com/ajax/ing/Publish", bytes.NewReader(text))
	if err != nil {
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	message <- 1
}

func deleteIng(ing string) {
	begin := strings.Index(ing, "DelIng(")
	l := len("DelIng(")
	ingId := ing[begin+l : begin+l+6]

	fmt.Println("{ingId:" + ingId + "}")

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://ing.cnblogs.com/ajax/ing/del", strings.NewReader("{ingId:"+ingId+"}"))

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func getconfig() {
	r, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	decoder := json.NewDecoder(r)
	var c Config
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalln(err)
	}
	cookie = c.Cookie
}

func main() {
	getconfig()
	for i := 1; i < 20; i++ {
		go insertIng(i)
		time.Sleep(10 * time.Minute)

		done := <-message

		log.Println(done)
		ing := getLastIng()
		if !strings.Contains(ing, "幸运闪") {
			go deleteIng(ing)
		}

		time.Sleep(5 * time.Minute)
	}
}
