package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Pizza struct {
	Diameter       int
	sauce          string
	cheese         int
	sausages       int
	salami         int
	ham            int
	smokedSausage  int
	olives         int
	onion          int
	mushrooms      int
	tomatoes       int
	chicken        int
	icebergLettuce int
}

type Builder struct {
	Diameter       int
	sauce          string
	cheese         int
	sausages       int
	salami         int
	ham            int
	smokedSausage  int
	olives         int
	onion          int
	mushrooms      int
	tomatoes       int
	chicken        int
	icebergLettuce int
}

func (b *Builder) setDiameter(diam int) {
	b.Diameter = diam
}
func (b *Builder) setSauce(sauce string) {
	b.sauce = sauce
}
func (b *Builder) setCheese(cheese int) {
	b.cheese = cheese
}
func (b *Builder) setSausages(sausag int) {
	b.sausages = sausag
}
func (b *Builder) setSalami(salami int) {
	b.salami = salami
}
func (b *Builder) setHam(ham int) {
	b.ham = ham
}
func (b *Builder) setSmokedSaus(smokedSaus int) {
	b.smokedSausage = smokedSaus
}
func (b *Builder) setOlives(olives int) {
	b.olives = olives
}
func (b *Builder) setOnion(onion int) {
	b.onion = onion
}
func (b *Builder) setMushrooms(mush int) {
	b.mushrooms = mush
}
func (b *Builder) setTomatoes(tomato int) {
	b.tomatoes = tomato
}
func (b *Builder) setChicken(chick int) {
	b.chicken = chick
}
func (b *Builder) setIceberg(iceberg int) {
	b.icebergLettuce = iceberg
}

func (b *Builder) getPizza() Pizza {

	return Pizza{
		Diameter: b.Diameter, sauce: b.sauce, cheese: b.cheese,
		sausages: b.sausages, salami: b.salami, ham: b.ham,
		smokedSausage: b.smokedSausage, olives: b.olives,
		onion: b.onion, mushrooms: b.mushrooms, tomatoes: b.tomatoes,
		chicken: b.chicken, icebergLettuce: b.icebergLettuce}
}

type Director struct {
	build Builder
}

func newDirector(b Builder) *Director {
	return &Director{build: b}
}

func (dir *Director) BuildFarmPizza33() Pizza {
	dir.build.setDiameter(33)
	dir.build.setSauce("farm")
	dir.build.setCheese(100)
	dir.build.setSalami(80)
	dir.build.setHam(80)
	dir.build.setSmokedSaus(80)
	dir.build.setOlives(50)
	dir.build.setOnion(50)
	return dir.build.getPizza()
}
func (dir *Director) BuildFarmPizza23() Pizza {
	dir.build.setDiameter(23)
	dir.build.setSauce("farm")
	dir.build.setCheese(80)
	dir.build.setSalami(60)
	dir.build.setHam(60)
	dir.build.setSmokedSaus(60)
	dir.build.setOlives(35)
	dir.build.setOnion(35)
	return dir.build.getPizza()
}
func (dir *Director) BuildCaesarPizza33() Pizza {
	dir.build.setDiameter(33)
	dir.build.setSauce("Caesar")
	dir.build.setChicken(150)
	dir.build.setTomatoes(100)
	dir.build.setIceberg(80)
	dir.build.setCheese(80)
	return dir.build.getPizza()
}

/*func main() {
	build := Builder{}
	dir := newDirector(build)
	fmt.Println(dir.BuildCaesarPizza33())
}*/
