package sortMan

import (
	"strconv"
	"strings"
)

//Сортировка по числам
type NumerSort struct {
}

func (ns NumerSort) compLower(first, second string) bool {
	f, _ := strconv.Atoi(strings.ToLower(first))
	s, _ := strconv.Atoi(strings.ToLower(second))
	if f < s {
		return true
	}
	return false
}
func (ns NumerSort) equalsLower(first, second string) bool {
	f, _ := strconv.Atoi(strings.ToLower(first))
	s, _ := strconv.Atoi(strings.ToLower(second))
	if f == s {
		return true
	}
	return false
}
func (ns NumerSort) comp(first, second string) bool {
	f, _ := strconv.Atoi(first)
	s, _ := strconv.Atoi(second)
	if f > s {
		return true
	}
	return false
}
