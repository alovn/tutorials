package mediator

import "fmt"

//中介者抽象
// type Mediator interface {
// 	Register(Colleague)
// 	Reply(Colleague)
// }

//具体的中介者：房产中介
type HouseMediator struct {
	houses []HouseOwner
}

func (m *HouseMediator) Register(h HouseOwner) {
	m.houses = append(m.houses, h)
}

func (m *HouseMediator) Reply(name string) {
	fmt.Printf("我是中介，正在给您找%s\n", name)
	for _, a := range m.houses {
		if a.Name() == name {
			fmt.Print("找到了，")
			a.Rent()
		}
	}
}

//抽象同事 房主
type HouseOwner interface {
	Name() string
	Rent() //出租
}

//具体同事-小房子
type SmallHouseOwner struct {
	name string
}

func NewSmallHouseOwner() *SmallHouseOwner {
	return &SmallHouseOwner{name: "小房子"}
}
func (h SmallHouseOwner) Name() string {
	return h.name
}

func (h SmallHouseOwner) Rent() {
	fmt.Println("小房子屋主，出租小房子")
}

//具体同事-大房子出租人
type BigHouseOwner struct {
	name string
}

func NewBigHouseOwner() *BigHouseOwner {
	return &BigHouseOwner{name: "大房子"}
}
func (h BigHouseOwner) Name() string {
	return h.name
}

func (h BigHouseOwner) Rent() {
	fmt.Println("大房子屋主，出租大房子")
}

type Person struct {
	mediator *HouseMediator
}

func NewPerson(mediator *HouseMediator) *Person {
	return &Person{mediator: mediator}
}
func (p Person) FindHouse(name string) { //找房子租
	fmt.Printf("我想租%s\n", name)
	p.mediator.Reply(name)
}
