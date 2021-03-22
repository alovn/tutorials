package interpreter

import (
	"strconv"
	"strings"
)

type Expression interface {
	Interpret() int
}

//加法处理器
type AddExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewAddExpression(expr1, expr2 NumberExpression) *AddExpression {
	return &AddExpression{
		expr1: expr1,
		expr2: expr2,
	}
}
func (add AddExpression) Interpret() int {
	return add.expr1.Interpret() + add.expr2.Interpret()
}

//数字处理器
type NumberExpression struct {
	num int
}

func NewNumberExpression(expression string) *NumberExpression {
	num, _ := strconv.Atoi(strings.Trim(expression, " "))
	return &NumberExpression{
		num: num,
	}
}

func (exp NumberExpression) Interpret() int {
	return exp.num
}

//计算器
type Calculator struct {
	expression string
}

func NewCalculator(expression string) *Calculator {
	return &Calculator{expression: expression}
}

func (c Calculator) Calculate() int {
	var expr1 *NumberExpression
	var expr2 *NumberExpression
	for i, s := range c.expression {
		switch s {
		case '+':
			expr1 = NewNumberExpression(c.expression[0 : i-1])
			expr2 = NewNumberExpression(c.expression[i+1 : len(c.expression)])
		}
	}
	if expr1 != nil && expr2 != nil {
		addExpr := NewAddExpression(*expr1, *expr2)
		return addExpr.Interpret()
	}
	return 0
}
