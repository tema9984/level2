package sortMan

import "strings"

type StrSort struct {
}

//Сортировка по чарам
func (ns StrSort) compLower(first, second string) bool {
	if strings.ToLower(first) < strings.ToLower(second) {
		return true
	}
	return false

}
func (ns StrSort) equalsLower(first, second string) bool {
	if strings.ToLower(first) == strings.ToLower(second) {
		return true
	}
	return false
}
func (ns StrSort) comp(first, second string) bool {
	if first > second {
		return true
	}
	return false
}
