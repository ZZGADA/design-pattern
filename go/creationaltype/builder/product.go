package builder

import "fmt"

type Product interface {
	Describe()
}

type ProductA struct {
	name    string
	radio   string
	speaker string
}

func (p *ProductA) Describe() {
	fmt.Println("it is Product A my radio is ", p.radio, "  and my speaker is ", p.speaker)
}

type ProductB struct {
	name    string
	radio   string
	speaker string
}

func (p *ProductB) Describe() {
	fmt.Println("it is Product B my radio is ", p.radio, "  and my speaker is ", p.speaker)
}
