package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type algorithm interface {
	eviction(arr *queue)
}

type firstAlg struct {
}

func (falg *firstAlg) eviction(arr *queue) {
	fmt.Println(arr.arr[0])

	arr.arr = arr.arr[1:]
}

type secondAlg struct {
}

func (secg *secondAlg) eviction(arr *queue) {
	array := []int{arr.arr[0]}
	fmt.Println(arr.arr[1])
	array = append(array, arr.arr[2:]...)
	arr.arr = array
}

type queue struct {
	arr []int
	alg algorithm
}

func (que *queue) setAlg(alg algorithm) {
	que.alg = alg
}
func (que *queue) eviction() {
	que.alg.eviction(que)
}

/*func main() {
	secAlg := secondAlg{}
	queue := queue{}
	queue.alg = &secAlg
	queue.arr = append(queue.arr, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	queue.eviction()
	fmt.Println(queue.arr)
	firAlg := firstAlg{}
	queue.alg = &firAlg
	queue.eviction()
	fmt.Println(queue.arr)
}*/
