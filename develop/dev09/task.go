package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/laher/wget-go/wget"
)

func main() {
	link := os.Args
	d, _ := os.Getwd()
	set := make(map[string]struct{})
	download(link, link[1], set, d+"/"+link[1], link[1])

}
func download(s []string, path string, set map[string]struct{}, pwd string, main string) {
	os.MkdirAll(pwd+"/"+path, os.ModePerm)
	os.Chdir(pwd + "/" + path)

	s[1] = strings.Replace(s[1], "//", "/", 1)

	err, str := wget.WgetCli(s)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	_, existErr := os.Stat(pwd + "/" + path + "/" + str)

	if existErr == nil && strings.Contains(str, ".html") {
		file := getFile(pwd + "/" + path + "/" + str)
		for _, v := range file {
			indexHref := strings.Index(v, `href="/`)
			indexSrc := -1
			if indexHref == -1 {
				indexSrc = strings.Index(v, `src="/`)
			}
			if indexHref > 0 || indexSrc > 0 {
				str := ""
				if indexHref > 0 {
					str = v[indexHref+6:]
				} else {
					str = v[indexSrc+5:]
				}
				index := strings.Index(str, `"`)
				str = str[0:index]

				if len(str) > 2 {
					ps := make([]string, 10)
					copy(ps, s)
					ps[1] = main + "/" + str
					index := strings.LastIndex(str, `/`)
					if index == -1 {
						str = str[1:]
					} else if index == 0 {
						str = str[1:]
					} else {
						str = str[1:index]
					}
					newPath := str
					_, exis := set[str]
					if !exis && len(str) > 1 {
						os.Chdir(pwd)
						set[str] = struct{}{}
						download(ps, newPath, set, pwd, main)
					}

				}

			}
		}

	}
}

func getFile(path string) []string {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}
