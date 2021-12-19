package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения + +
-B - "before" печатать +N строк до совпадения + +
-C - "context" (A+B) печатать ±N строк вокруг совпадения + +
-c - "count" (количество строк)+ +
-i - "ignore-case" (игнорировать регистр)+ +
-v - "invert" (вместо совпадения, исключать)+ +
-F - "fixed", точное совпадение со строкой, не паттерн + +
-n - "line num", печатать номер строки + +

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type grep struct {
	str      []string
	ind      []int //строки для вывода A,B,C
	inIndex  map[int]struct{}
	finded   []int //найденные строки
	inFinded map[int]struct{}
}

//-v - "invert" (вместо совпадения, исключать)
func (gr *grep) invert() {
	invertIndex := []int{}
	for k := range gr.str {
		_, exist := gr.inIndex[k]
		_, existF := gr.inFinded[k]
		if !exist && !existF {
			invertIndex = append(invertIndex, k)
		}

	}
	gr.ind = invertIndex
}

//-c - "count" (количество строк)
func (gr *grep) count() int {
	return len(gr.finded)
}

//-n - "line num", печатать номер строки
func (gr *grep) lineNum() (ret []string) {

	if len(gr.ind) < len(gr.finded) {
		for _, v := range gr.finded {
			ret = append(ret, strconv.Itoa(v)+":"+gr.str[v])
		}
	} else {
		for _, v := range gr.ind {
			ret = append(ret, strconv.Itoa(v)+":"+gr.str[v])
		}
	}
	return
}
func (gr *grep) getStr() (ret []string) {

	if len(gr.ind) < len(gr.finded) {
		for _, v := range gr.finded {
			ret = append(ret, gr.str[v])
		}
	} else {
		for _, v := range gr.ind {
			ret = append(ret, gr.str[v])
		}
	}
	return
}

//-F - "fixed", точное совпадение со строкой, не паттерн
func (gr *grep) findStrFix(findStr string) {
	for k, v := range gr.str {
		if strings.Contains(v, findStr) {
			gr.finded = append(gr.finded, k)
			gr.inFinded[k] = struct{}{}
		}
	}
}

//-F - "fixed", точное совпадение со строкой, не паттерн, игнорирование регистра
func (gr *grep) findStrFixIgnore(findStr string) {
	for k, v := range gr.str {
		if strings.Contains(strings.ToLower(v), strings.ToLower(findStr)) {
			gr.finded = append(gr.finded, k)
			gr.inFinded[k] = struct{}{}
		}
	}
}
func (gr *grep) findStr(findStr string) {
	for k, v := range gr.str {
		matched, err := regexp.MatchString(findStr, v)
		if err != nil {
			break
		}
		if matched {
			gr.finded = append(gr.finded, k)
			gr.inFinded[k] = struct{}{}
		}
	}
}

// игнорирование регистра
func (gr *grep) findStrIgnore(findStr string) {
	for k, v := range gr.str {
		matched, err := regexp.MatchString(strings.ToLower(findStr), strings.ToLower(v))
		if err != nil {
			break
		}
		if matched {
			gr.finded = append(gr.finded, k)
			gr.inFinded[k] = struct{}{}
		}
	}
}

//-A - "after" печатать +N строк после совпадения
func (gr *grep) after(a int) {
	for _, ind := range gr.finded {
		for i := ind; i <= ind+a && i < len(gr.str); i++ {
			_, exist := gr.inIndex[i]
			if !exist {
				gr.ind = append(gr.ind, i)
				gr.inIndex[i] = struct{}{}
			}

		}
	}

}

//-B - "before" печатать +N строк до совпадения
func (gr *grep) before(b int) {
	for _, ind := range gr.finded {
		for i := ind - b; i <= ind; i++ {
			_, exist := gr.inIndex[i]
			if !exist && i >= 0 {
				gr.ind = append(gr.ind, i)
				gr.inIndex[i] = struct{}{}
			}

		}
	}
}

//-C - "context" (A+B) печатать ±N строк вокруг совпадения
func (gr *grep) context(ctx int) {
	for _, ind := range gr.finded {
		for i := ind - ctx; i <= ind+ctx && i < len(gr.str); i++ {
			_, exist := gr.inIndex[i]
			if !exist && i >= 0 {
				gr.ind = append(gr.ind, i)
				gr.inIndex[i] = struct{}{}
			}

		}
	}
}

func main() {
	//*-A - "after" печатать +N строк после совпадения + +
	A := 0
	//-B - "before" печатать +N строк до совпадения + +
	B := 0
	//-C - "context" (A+B) печатать ±N строк вокруг совпадения + +
	C := 0
	//-c - "count" (количество строк)+ +
	count := flag.Bool("c", false, "количество строк")
	//-i - "ignore-case" (игнорировать регистр)+ +
	i := flag.Bool("i", false, "игнорировать регистр")
	//-v - "invert" (вместо совпадения, исключать)+ +
	v := flag.Bool("v", false, "вместо совпадения, исключать")
	//-F - "fixed", точное совпадение со строкой, не паттерн + +
	F := flag.Bool("F", false, "точное совпадение со строкой")
	//-n - "line num", печатать номер строки + +
	n := flag.Bool("n", false, "печатать номер строки")
	flag.IntVar(&A, "A", -1, "печатать +N строк после совпадения")
	flag.IntVar(&B, "B", -1, "печатать +N строк до совпадения")
	flag.IntVar(&C, "C", -1, "печатать ±N строк вокруг совпадения")
	flag.Parse()

	buff := getFile("file")
	s := manGreep(buff, `ads.*h`, A, B, C, *count, *i, *v, *F, *n)
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}
func manGreep(buff []string, str string, A, B, C int, count, i, v, F, n bool) []string {
	gr := grep{str: buff, inFinded: map[int]struct{}{}, inIndex: map[int]struct{}{}}
	//блок поиска
	if F {
		if i {
			gr.findStrFixIgnore(str)
		} else {
			gr.findStrFix(str)
		}
	} else {
		if i {
			gr.findStrIgnore(str)
		} else {
			gr.findStr(str)
		}

	}
	//блок контекста
	if A > 0 && B <= 0 && C <= 0 {
		gr.after(A)
	}
	if B > 0 && A <= 0 && C <= 0 {
		gr.before(B)
	}
	if C > 0 && A <= 0 && B <= 0 {
		gr.context(C)
	}
	//invers
	if v {
		gr.invert()
	}
	//блок вывода
	if n {
		return gr.lineNum()
	} else if count {
		return []string{strconv.Itoa(gr.count())}
	} else {
		return gr.getStr()
	}
}
func getFile(path string) []string {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}
