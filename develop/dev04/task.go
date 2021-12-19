package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"пятак", "листок", "пятка", "тяпка", "слиток", "столик", "авб", "Go", "OG", "Ы", "РОММ", "Один", "морМ", "бав", "абв", "ваб", "актяп"}
	an := findAnagrams(words)
	for k, v := range an {
		fmt.Printf("key:%s, val:%v\n", k, v)
	}
}

func findAnagrams(arrayStrings []string) map[string][]string {
	anagrams := make(map[string][]string)
	letters := make(map[string]map[rune]int)
	//Цикл по всем словам
	for _, v := range arrayStrings {
		//приводим к нижнему регистру
		v := strings.ToLower(v)
		isAnagram := false
		//цикл по мапе с словами и подсчитанными буквами
		for key, value := range letters {
			//сравниваем слово и набор
			isAnagram = checkAnagram(value, v)
			if isAnagram {
				_, exist := anagrams[key]
				if exist {
					app := true
					for _, str := range anagrams[key] {
						if str == v {
							app = false
							break
						}
					}
					if app {
						anagrams[key] = append(anagrams[key], v)
					}

				} else {
					anagrams[key] = []string{key, v}
				}
				break
			}
		}
		//Если совпадений не было, то добавляем новый набор
		if !isAnagram {

			mapWords := make(map[rune]int)
			for _, val := range v {
				mapWords[val]++
			}
			letters[v] = mapWords
		}
	}
	//Сортируем массивы перед выходом
	for _, val := range anagrams {
		sort.Slice(val, func(i, j int) bool {
			if val[i] < val[j] {
				return true
			}
			return false
		})
	}
	return anagrams
}

func checkAnagram(letters map[rune]int, str2 string) bool {
	let := map[rune]int{}
	for _, v := range str2 {
		//подсчитываем количество букв в слове
		let[v]++
	}
	for k := range let {
		//Проверяем что кол-во букв совпадает
		if let[k] != letters[k] {
			return false
		}
	}
	return true
}
