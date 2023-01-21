package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDnslog() (string, string) {
	url := "http://www.dnslog.cn/getdomain.php"
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var reC []string = response.Header["Set-Cookie"]
	cookie := reC[0][:36]
	return cookie, string(body)

}
func GetDnsResult(cookie string) string {
	url := "http://www.dnslog.cn/getrecords.php"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", cookie)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("internet failed,err:%v\n\n", err)
	}
	return string(b)

}
