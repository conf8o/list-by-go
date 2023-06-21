package lisp

type Env struct {
	Vals []map[Symbol]Val
}

type Val interface {
	Eval(*Env) Val
}

type Int32 struct {
	N int32
}

func (i Int32) Eval(_ *Env) Int32 {
	return i
}

type String struct {
	N string
}

func (s String) Eval(_ *Env) String {
	return s
}

type Optional struct {
	Some  Val
	Empty bool
}

func (o Optional) Eval(_ *Env) Optional {
	return o
}

type Symbol struct {
	N string
}

func (s Symbol) Eval(env *Env) Optional {
	for i := len(env.Vals); i > 0; i-- {
		e := env.Vals[i]
		val, exists := e[s]
		if exists {
			return Optional{Some: val, Empty: false}
		}
	}
	return Optional{Some: nil, Empty: true}
}

type Cons struct {
	Car *Val
	Cdr *Cons
}
