package testdata

type B struct{}

type T struct {
	private string  `props:"*"`
	Public  string  `props:"*"`
	star    *string `props:"*"`
	noting  string

	setOnly string `props:"set"`
	getOnly string `props:"get"`

	namedSetter string `props:"set=SetNamedSetter,get=GetNamedSetter"`
	namedGetter string `props:"set=SetNamedGetter,get=GetNamedGetter"`

	chanType chan string `props:"*"`

	arrType     []string  `props:"*"`
	arrStarType []*string `props:"*"`

	mapType map[string]string `props:"*"`

	structType     B  `props:"*"`
	structStarType *B `props:"*"`

	interfaceType interface{} `props:"*"`
}
