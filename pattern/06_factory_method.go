package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type gadGet interface {
	goOnline()
	setSim()
}

type gadget struct {
	sim    bool
	online bool
}

func (gt *gadget) goOnline() {
	if gt.sim {
		gt.online = true
		fmt.Println("You went online")
	} else {
		fmt.Println("You need to set SIM")
	}

}
func (gt *gadget) setSim() {
	fmt.Println("i set sim")
	gt.sim = true
}

type phone struct {
	gadget
}

func newPhone() gadGet {
	return &phone{
		gadget: gadget{
			sim:    false,
			online: false,
		},
	}
}

type tablet struct {
	gadget
}

func newTablet() gadGet {
	return &tablet{
		gadget: gadget{
			sim:    false,
			online: false,
		},
	}
}

func getGadget(gadgetType string) gadGet {
	if gadgetType == "tablet" {
		return newTablet()
	}
	return newPhone()
}

/*
func main() {
	phone := getGadget("something")
	phone.setSim()
	phone.goOnline()
	tablet := getGadget("tablet")
	tablet.goOnline()
}
*/
