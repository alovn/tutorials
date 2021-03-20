package decorator

import "fmt"

//武器
type Weapon interface {
	Attack() //攻击
}

//皮肤
type Skin interface {
	Protect() //皮肤有不同的保护值
}

//枪
type Gun struct{}

func (g Gun) Attack() {
	fmt.Println("手枪-射击")
}

//大炮
type Cannon struct{}

func (c Cannon) Attack() {
	fmt.Println("加农大炮——开火")
}

//电玩小子皮肤
type GameBoySkin struct{}

func (g GameBoySkin) Protect() {
	fmt.Println("穿上电玩小子批发武力值+10，保护+10")
}

//游戏角色
type Characters interface {
	Weapon(Weapon)
	Skin(Skin)
	Attack()
}

//游戏角色-鲁班七号
type Luban struct {
	skin   Skin
	weapon Weapon
}

func (l *Luban) Weapon(w Weapon) {
	l.weapon = w
}

func (l *Luban) Skin(s Skin) {
	l.skin = s
	l.skin.Protect()
}
func (l Luban) Attack() {
	l.weapon.Attack()
}
