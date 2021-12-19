package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type visitor interface {
	visitGym(gym gym)
	visitArt(art art)
}

type gym struct {
	name string
}

func (gm gym) accept(vis visitor) {
	vis.visitGym(gm)
}

type art struct {
	name string
}

type man struct {
	name string
}

func (man *man) visitArt(art art) {
	fmt.Println(man.name, " делает умный вид смотря на ", art.name)
}
func (man *man) visitGym(gym gym) {
	fmt.Println("Алло, не могу сейчас говорить, я в зале ", gym.name)
}

type woman struct {
	name string
}

func (woman *woman) visitArt(art art) {
	fmt.Println(woman.name, " любуется искусством ", art.name)
}
func (woman *woman) visitGym(gym gym) {
	fmt.Println(woman.name, " бегает на дорожке в зале ", gym.name)
}

/*
func main() {
	gym := gym{name: "GOld "}
	visitorMan := man{name: "El"}
	gym.accept(&visitorMan)
}
*/
