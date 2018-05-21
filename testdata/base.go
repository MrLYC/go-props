package testdata

import (
	"go/ast"
)

type B struct{}

type T struct {
	private string  `props:"+"`
	Public  string  `props:"+"`
	star    *string `props:"+"`
	noting  string

	setOnly string `props:"set"`
	getOnly string `props:"get"`

	namedSetter string `props:"set=SetNamedSetter,get=GetNamedSetter"`
	namedGetter string `props:"set=SetNamedGetter,get=GetNamedGetter"`

	chanType          chan string   `props:"+"`
	chanReadOnlyType  <-chan string `props:"+"`
	chanWriteOnlyType chan<- string `props:"+"`

	arrType     []string  `props:"+"`
	arrStarType []*string `props:"+"`

	mapType map[string]string `props:"+"`

	structType     B  `props:"+"`
	structStarType *B `props:"+"`

	interfaceType interface{} `props:"+"`

	emptyFuncType   func()                          `props:"+"`
	paramFuncType   func(int)                       `props:"+"`
	resultFuncType  func() int                      `props:"+"`
	resultsFuncType func() (int, string)            `props:"+"`
	fullFuncType    func(int, string) (int, string) `props:"+"`

	selectorType     ast.SelectorExpr  `props:"+"`
	selectorStarType *ast.SelectorExpr `props:"+"`

	setterEx string `props:"set=SetEx,get"`
	getterEx string `props:"set,get=GetEx"`
	bothEx   string `props:"set=SetEx,get=GetEx"`

	toPtr   string `props:"+,to=*"`
	fromPtr string `props:"+,from=*"`
	asPtr   string `props:"+,to=*,from=*"`
}

func (t *T) SetEx(v string) {

}

func (t T) GetEx() string {
	return ""
}
