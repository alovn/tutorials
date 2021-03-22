package strategy

type DiscountStrategy interface {
	//计算折扣，total为总价
	Discount(total float64) float64
}

type MemberDiscountStrategy struct{}

func (s MemberDiscountStrategy) Discount(total float64) float64 {
	//普通用户9.5折
	return total * 0.05
}

type VIPDiscountStrategy struct{}

func (s VIPDiscountStrategy) Discount(total float64) float64 {
	//VIP打9折
	return total * 0.1
}

type OverDiscountStragy struct{}

func (s OverDiscountStragy) Discount(total float64) float64 {
	//满200减40
	if total >= 200 {
		return 40
	}
	return 0
}

//折扣上下文
type DiscountContext struct {
	strategy DiscountStrategy
}

func (c *DiscountContext) SetStrategy(strategy DiscountStrategy) {
	c.strategy = strategy
}

func (c DiscountContext) CaculatePrice(total float64) float64 {
	return total - c.strategy.Discount(total)
}
