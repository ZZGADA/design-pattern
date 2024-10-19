package proxy

type Proxy struct {
	subject Subject
}

func NewProxy(subject string) *Proxy {
	switch subject {
	case "real":
		return &Proxy{subject: &RealSubject{}}
	default:
		return &Proxy{}
	}
}

func (p *Proxy) Say() {
	p.subject.say()
}

func (p *Proxy) Run() {
	p.subject.run()
}
