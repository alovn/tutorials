package decorator

import "testing"

func TestLuban_Attack(t *testing.T) {
	//武器
	var gun Weapon = &Gun{}
	var cannon Weapon = &Cannon{}

	//皮肤
	var gameBoySkin Skin = &GameBoySkin{}

	//游戏角色
	var luban Characters = &Luban{}
	luban.Weapon(gun)
	luban.Attack()

	luban.Skin(gameBoySkin) //穿上皮肤
	luban.Weapon(cannon)    //使用加农大炮
	luban.Attack()          //攻击
}
