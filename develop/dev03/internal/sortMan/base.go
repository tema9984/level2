package sortMan

import "strings"

type forSort interface {
	compLower(first, second int) bool
	equalsLower(first, second int) bool
	comp(first, second int) bool
	getSorted() []string
	getLen() int
}
type TypeSort interface {
	compLower(first, second string) bool
	equalsLower(first, second string) bool
	comp(first, second string) bool
}
type Base struct {
	needSort     []string
	sortIsColumn bool
	numberColumn int
	currentType  TypeSort
}

func (ns *Base) SetNeedSort(str []string) {
	ns.needSort = str
}
func (ns *Base) SetNumberColumn(nubCol int) {
	ns.numberColumn = nubCol
	ns.sortIsColumn = true
}
func (ns *Base) SetCurrentType(curType TypeSort) {
	ns.currentType = curType
}
func (ns *Base) UniqStrings() {
	mapin := make(map[string]struct{})
	for _, val := range ns.needSort {
		mapin[val] = struct{}{}
	}
	uniqArr := []string{}
	for key := range mapin {
		uniqArr = append(uniqArr, key)
	}
	ns.needSort = uniqArr
}
func (ns Base) compLower(first, second int) bool {
	if ns.sortIsColumn {
		f := strings.Split(ns.needSort[first], " ")
		s := strings.Split(ns.needSort[second], " ")
		return ns.currentType.compLower(f[ns.numberColumn], s[ns.numberColumn])
	} else {
		return ns.currentType.compLower(ns.needSort[first], ns.needSort[second])
	}
}
func (ns Base) equalsLower(first, second int) bool {
	if ns.sortIsColumn {
		f := strings.Split(ns.needSort[first], " ")
		s := strings.Split(ns.needSort[second], " ")
		return ns.currentType.equalsLower(f[ns.numberColumn], s[ns.numberColumn])
	} else {
		return ns.currentType.equalsLower(ns.needSort[first], ns.needSort[second])
	}
}
func (ns Base) comp(first, second int) bool {
	if ns.sortIsColumn {
		f := strings.Split(ns.needSort[first], " ")
		s := strings.Split(ns.needSort[second], " ")
		return ns.currentType.comp(f[ns.numberColumn], s[ns.numberColumn])
	} else {
		return ns.currentType.comp(ns.needSort[first], ns.needSort[second])
	}
}

func (ns Base) getSorted() []string {
	return ns.needSort
}
func (ns Base) getLen() int {
	return len(ns.needSort)
}
