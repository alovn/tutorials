package strategy

import (
	"fmt"
	"testing"
)

func TestDiscountContext_CaculatePrice(t *testing.T) {
	ctx := &DiscountContext{}
	ctx.SetStrategy(MemberDiscountStrategy{})
	fmt.Printf("使用普通用户策略后最终价格：%.2f\n", ctx.CaculatePrice(300))

	ctx.SetStrategy(VIPDiscountStrategy{})
	fmt.Printf("使用会员策略后最终价格：%.2f\n", ctx.CaculatePrice(300))

	ctx.SetStrategy(OverDiscountStragy{})
	fmt.Printf("使用满减策略后最终价格：%.2f\n", ctx.CaculatePrice(300))
}
