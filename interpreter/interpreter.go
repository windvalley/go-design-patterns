package interpreter

import "strings"

// Expression ...
type Expression interface {
	Interpret(context string) bool
}

type terminalExpression struct {
	data string
}

// NewTerminalExpression ...
func NewTerminalExpression(data string) Expression {
	return &terminalExpression{data}
}

func (t *terminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.data) {
		return true
	}

	return false
}

type orExpression struct {
	expr1 Expression
	expr2 Expression
}

// NewOrExpression ...
func NewOrExpression(expr1, expr2 Expression) Expression {
	return &orExpression{expr1, expr2}
}

func (o *orExpression) Interpret(context string) bool {
	return o.expr1.Interpret(context) || o.expr2.Interpret(context)
}

type andExpression struct {
	expr1 Expression
	expr2 Expression
}

// NewAndExpression ...
func NewAndExpression(expr1, expr2 Expression) Expression {
	return &andExpression{expr1, expr2}
}

func (a *andExpression) Interpret(context string) bool {
	return a.expr1.Interpret(context) && a.expr2.Interpret(context)
}
