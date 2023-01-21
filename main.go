package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"springscan/lib"
	"strings"
	"sync"
)

var (
	u string
	f string
	t int
	h bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&u, "u", "", "a target url(Please add http or https)")
	flag.StringVar(&f, "f", "", "a target filename")
	flag.IntVar(&t, "t", runtime.NumCPU(), "thread Num")

	// 改变默认的 Usage
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `springscan version: 1.0.0
Usage:  [-uft] [-u url] [-f filename] [-t thread] [-h help]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	lib.Logo()
	flag.Parse()

	if h {
		flag.Usage()
	}

	res1 := strings.Contains(u, "http://")
	res2 := strings.Contains(u, "https://")

	if u == "" {

	} else {
		if !res1 && !res2 {
			fmt.Println(lib.Red("[-] Please add http or https for url !!!"))
			os.Exit(0)
		} else {
			fmt.Println(lib.Yellow("target: " + u))
			var wg sync.WaitGroup
			lib.Envscan(u, t)
			fmt.Println(lib.Yellow("\nSpring漏洞扫描开始："))
			wg.Add(5)
			go lib.SpringElRce(u, &wg)
			go lib.H2databaseRce(u, &wg)
			go lib.SpringFramework(u, &wg)
			go lib.SpringDataCommonsRce(u, &wg)
			go lib.SnakeYaml(u, &wg)
			wg.Wait()

		}
	}

}
