package main

import (
	"fmt"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var numbers = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

func main() {
	str, err := unpacking(`"3a40bc\*5n2d5e5"`)
	fmt.Println("i am str:"+str, "i am err:", err)
}

func unpacking(str string) (string, error) {
	var count int
	var number int
	var newString string
	var escape bool
	//Добавляем конечный символ
	str += "%"
	for key, val := range str {
		//завершаем работу если дошли до конца
		if key+1 >= len(str) && string(val) == "%" {
			break
		}
		//проверка на эскейп
		if string(val) == `\` && escape == false {
			escape = true
			continue
		}
		//проверяем является ли символ цифрой
		mnoj, exist := numbers[string(val)]
		//если это ескейп символ, то цифра считается обычным символом
		if escape == true {
			exist = false
			escape = false
		}
		//Если это не цифра, и последний символ то добавить к итоговой строке
		if key+1 >= len(str) && !exist {
			newString += string(val)
			continue
		}
		//Если это не цифра, и некст тоже, то добавить к строке
		_, eixstNext := numbers[string(str[key+1])]
		if exist == false && eixstNext == false {
			newString += string(val)
			continue
		}
		// Если это цифра, то добавить к числу для анпакинга
		if exist == true {
			number = number*10 + mnoj
			count++
		}
		//Если число закончилось, то распаковать символ
		if exist && !eixstNext {
			if key-count < 0 {
				return "", fmt.Errorf("Incorrect string")

			}
			unpack := strings.Repeat(string(str[key-count]), number)
			number, count = 0, 0
			newString += unpack
		}

	}
	return newString, nil
}
