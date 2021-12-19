package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	//*-
	f := ""
	//-
	d := "\t"
	s := flag.Bool("s", false, "separated - только строки с разделителем")
	flag.StringVar(&f, "f", "-1", "fields - выбрать поля (колонки)")
	flag.StringVar(&d, "d", "\t", "delimiter - использовать другой разделитель")
	flag.Parse()
	cutStart(f, d, *s)
}
func cutStart(f string, d string, s bool) {
	if f == "-1" {
		fmt.Println("укажите флаг -f")
		return
	}
	arrFs := strings.Split(f, ",")
	arrF := []int{}
	for _, v := range arrFs {
		F, _ := strconv.Atoi(v)
		arrF = append(arrF, F)
	}

	cut(arrF, d, s)
}
func cut(f []int, d string, s bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		str := serparat(text, f, d, s)
		if str == "" && s {
			continue
		}
		fmt.Println(str)
	}
}
func serparat(buff string, i []int, delimiter string, s bool) (res string) {
	split := strings.Split(string(buff), delimiter)
	if len(split) == 1 && s {
		return ""
	}
	for _, v := range i {
		if v >= len(split) {
			return
		}
		if res != "" {
			res += delimiter
		}
		res += split[v]
	}
	return
}
