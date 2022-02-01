package ast

type Exp interface{}

type NilExp struct{ Line int }
type TrueExp struct{ Line int }
type FalseExp struct{ Line int }
type VarargExp struct{ Line int }
type IntegerExp struct {
	Line int
	Val  int64
}
type FloatExp struct {
	Line int
	Val  float64
}
type StringExp struct {
	Line int
	Str  string
}
type NameExp struct {
	Line int
	Name string
}

type UnopExp struct {
	Line int
	Op   int
	Exp  Exp
}
type BinopExp struct {
	Line int
	Op   int
	Exp1 Exp
	Exp2 Exp
}
type ConcatExp struct {
	Line int
	Exps []Exp
}

type TableConstructorExp struct {
	Line     int // line of '{'
	LastLine int // line of '}'
	KeyExps  []Exp
	ValExps  []Exp
}

type FuncDefExp struct {
	Line     int
	LastLine int // line of 'end'
	ParList  []string
	IsVararg bool
	Block    *Block
}

type ParentsExp struct {
	Exp Exp
}

type TableAccessExp struct {
	LastLine  int
	PrefixExp Exp
	KeyExp    Exp
}

type FuncCallExp struct {
	Line      int // line of '('
	LastLine  int // line of ')'
	PrefixExp Exp
	NameExp   *StringExp
	Args      []Exp
}
