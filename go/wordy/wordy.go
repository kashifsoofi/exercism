package wordy

import (
	"github.com/alecthomas/participle"
)

type operator int

const (
	plus operator = iota
	minus
	times
	divide
)

var operatorMap = map[string]operator{"plus": plus, "minus": minus, "multiplied": times, "divided": divide}

func (o *operator) Capture(s []string) error {
	*o = operatorMap[s[0]]
	return nil
}

type term struct {
	Value *int `parser:"@('-'? Int)"`
}

type opTerm struct {
	Operator operator `@("plus" | "minus" | "multiplied" "by" | "divided" "by")`
	Term     *term    `@@`
}

type experssion struct {
	Term   *term     `@@`
	OpTerm []*opTerm `@@*`
}

type question struct {
	WhatIs       string      `"What" "is"`
	Expression   *experssion `@@`
	QuestionMark string      `"?"`
}

// Evaluation
func (o operator) eval(l, r int) int {
	switch o {
	case plus:
		return l + r
	case minus:
		return l - r
	case times:
		return l * r
	case divide:
		return l / r
	}
	panic("unsupported operator")
}

func (t *term) eval() int {
	return *t.Value
}

func (ot *opTerm) eval(l int) int {
	r := ot.Term.eval()
	return ot.Operator.eval(l, r)
}

func (e *experssion) eval() int {
	v := e.Term.eval()
	for _, ot := range e.OpTerm {
		v = ot.eval(v)
	}
	return v
}

func (q *question) eval() int {
	return q.Expression.eval()
}

// Answer evaluates question and returns result if successful
func Answer(q string) (int, bool) {
	parser, err := participle.Build(&question{})
	if err != nil {
		return 0, false
	}

	question := &question{}
	err = parser.ParseString(q, question)
	if err != nil {
		return 0, false
	}

	a := question.eval()
	return a, true
}
