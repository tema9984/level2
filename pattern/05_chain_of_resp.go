package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type move interface {
	execute(*user)
	setNext(move)
}

type user struct {
	name string
	age  int8
}

type move0plus struct {
	next move
}

func (m3p *move0plus) execute(us *user) {
	fmt.Println("Смотрим развивающие мультики")
	if us.age < 3 {
		return
	} else {
		m3p.next.execute(us)
	}

}

func (m3p *move0plus) setNext(mv move) {
	m3p.next = mv
}

type move3plus struct {
	next move
}

func (m3p *move3plus) execute(us *user) {
	fmt.Println("Смотрим смешные мультики")
	if us.age < 12 {
		return
	} else {
		m3p.next.execute(us)
	}

}

func (m3p *move3plus) setNext(mv move) {
	m3p.next = mv
}

type move12plus struct {
	next move
}

func (m3p *move12plus) execute(us *user) {
	fmt.Println("Смотрим фильмы для детей")
	if us.age < 16 {
		return
	} else {
		m3p.next.execute(us)
	}

}

func (m3p *move12plus) setNext(mv move) {
	m3p.next = mv
}

type move16plus struct {
	next move
}

func (m3p *move16plus) execute(us *user) {
	fmt.Println("Смотрим добрые фильмы")
	if us.age < 18 {
		return
	} else {
		m3p.next.execute(us)
	}

}

func (m3p *move16plus) setNext(mv move) {
	m3p.next = mv
}

type move18plus struct {
	next move
}

func (m3p *move18plus) execute(us *user) {
	fmt.Println("Смотрим любые фильмы")
}

func (m3p *move18plus) setNext(mv move) {
	m3p.next = mv
}

/*
func main() {
	user := user{name: "petr", age: 8}

	move18plus := move18plus{}
	move16plus := move16plus{}
	move16plus.setNext(&move18plus)

	move12plus := move12plus{}
	move12plus.setNext(&move16plus)

	move3plus := move3plus{}
	move3plus.setNext(&move12plus)

	move0plus := move0plus{}
	move0plus.setNext(&move3plus)

	move0plus.execute(&user)
}
*/
