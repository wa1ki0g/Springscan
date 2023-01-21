package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func H2databaseRce(urlstring string, wg *sync.WaitGroup) {
	wg.Done()
	res1, _ := url(string(urlstring) + "/h2-console/")
	res2, _ := url(string(urlstring) + "/login.jsp")
	if res1 {
		fmt.Println(Green("[+] H2databaseRCE 存在！"))
		return
	}
	if res2 {
		fmt.Println(Green("[+] H2databaseRCE 存在！"))
		return

	}
	if !res1 && !res2 {
		fmt.Println(Red("[-] H2databaseRCE 不存在"))
	}

}

func url(url string) (bool, string) {
	resp, err := http.Get(string(url))
	if err != nil {
		fmt.Println(err)
		return false, "error"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res1 := strings.Contains(string(body), "Welcome to H2")
	res2 := strings.Contains(string(body), "H2 Console")
	if resp.StatusCode == 200 {
		if res1 || res2 {
			return true, url
		} else {
			return false, url
		}
	} else {
		return false, url
	}

}
