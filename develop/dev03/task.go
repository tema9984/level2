package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/tema9984/dev03/internal/sortMan"
)

func main() {
	k := 0
	n := flag.Bool("n", false, "Usage string")
	r := flag.Bool("r", false, "Usage string")
	u := flag.Bool("u", false, "Usage string")
	flag.IntVar(&k, "k", -1, "Usage string")
	flag.Parse()
	buff := getFile("needSort")
	str := manSort(buff, k, *n, *r, *u)
	for key, val := range str {
		fmt.Printf("key: %d, val: %s, \n", key, val)
	}

}
func manSort(buff []string, k int, n, r, u bool) (str []string) {
	sorts := sortMan.Base{}
	//Выбор типа сортировк
	if n {
		//Сортировка по числовому значению
		typeSort := sortMan.NumerSort{}
		sorts.SetCurrentType(typeSort)
	} else {
		//Дефолтная сортировка
		typeSort := sortMan.StrSort{}
		sorts.SetCurrentType(typeSort)
	}
	//Передаем массив для сортировки
	sorts.SetNeedSort(buff)
	if u {
		//Только уникальные строки
		sorts.UniqStrings()
	}
	if k > -1 {
		//Указываем номер столбца
		sorts.SetNumberColumn(k)
	}
	if r {
		//Сортировка с ревирсом
		str = sortMan.SortStringsRevers(sorts)
	} else {
		//Обычная сортировка
		str = sortMan.SortStrings(sorts)
	}
	return str
}
func getFile(path string) []string {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}
