package pattern

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Employee struct {
	name          string
	taskCompleted int
	works         bool
	orders        []int
}

func (eml *Employee) work(t int) {
	eml.works = true
	count := 1
	rand.Seed(time.Now().UnixNano())
	for {
		//time.Sleep(time.Millisecond)
		if count == rand.Intn(100) {
			count++
			if count >= 10 {
				break
			}
		}
	}
	fmt.Println(eml.name, "i am done at work", t)
	eml.taskCompleted = t
}

type Kitchen struct {
	orders map[int]int
	load   int
}

func (kit *Kitchen) getOrder() int {
	order := len(kit.orders)
	kit.orders[order] = order
	return order
}

type Fridge struct {
	stocks int
}

func (fr *Fridge) getFood() {
	if fr.stocks <= 0 {
		fr.replenish()
	}
	fr.stocks--
}
func (fr *Fridge) replenish() {
	fr.stocks = fr.stocks + 100
}

type Facade struct {
	employee []Employee
	kitchen  *Kitchen
	fridge   *Fridge
}

func NewFacade(numEmpl int) *Facade {
	employees := []Employee{}
	employees = append(employees, Employee{name: "Jon", taskCompleted: -1, works: false})

	orders := make(map[int]int)
	facade := &Facade{employee: employees, kitchen: &Kitchen{orders: orders}, fridge: &Fridge{stocks: 100}}
	return facade
}
func (facade *Facade) MakeOrder() int {
	facade.fridge.getFood()
	order := facade.kitchen.getOrder()
	facade.employee[0].work(order)
	return order
}
func (facade *Facade) GetOrder(id int) bool {
	if facade.employee[0].taskCompleted == id {
		facade.employee[0].taskCompleted = -1
		fmt.Println("Ваш заказ")
		return true
	} else {
		fmt.Printf("заказ %v еще не готов\n", id)
		return false
	}
}

/*func main() {
	facade := NewFacade(5)
	id := facade.MakeOrder()
	facade.GetOrder(id)
	time.Sleep(1 * time.Second)
	facade.GetOrder(id)
}
*/
