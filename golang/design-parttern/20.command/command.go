package command

import "fmt"

type Command interface {
	Execute()
}

//倒咖啡命令
type PourWaterCommand struct {
	waiter *Waiter
}

func NewPourWaterCommand(waiter *Waiter) Command {
	return &PourWaterCommand{
		waiter: waiter,
	}
}

func (c *PourWaterCommand) Execute() {
	c.waiter.PourWater()
}

//下厨烧饭命令
type CookCommand struct {
	chef *Chef
}

func NewCookCommand(chef *Chef) Command {
	return &CookCommand{
		chef: chef,
	}
}

func (c *CookCommand) Execute() {
	c.chef.Cook()
}

//结账命令
type CheckoutCommand struct {
	waiter *Waiter
}

func NewCheckoutCommand(waiter *Waiter) Command {
	return &CheckoutCommand{
		waiter: waiter,
	}
}
func (c *CheckoutCommand) Execute() {
	c.waiter.Checkout()
}

//服务员
type Waiter struct {
}

func (w Waiter) PourWater() {
	fmt.Println("服务员——倒水")
}

func (w Waiter) Checkout() {
	fmt.Println("服务员——计数账单结账")
}

//厨师
type Chef struct {
}

func (c Chef) Cook() {
	fmt.Println("厨师——烧菜")
}
