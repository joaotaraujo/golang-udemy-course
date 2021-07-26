package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func oMaisRapido(url1, url2, url3 string) string {
	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	// estrutura de controle especifica para concorrencia
	select {
	case t1 := <- c1:
		return t1
	case t2 := <- c2:
		return t2
	case t3 := <- c3:
		return t3
	case <-time.After(234400 * time.Millisecond):
		return "todos perderam"
	//default:
	//	return "sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com")
	fmt.Println(campeao)
}
