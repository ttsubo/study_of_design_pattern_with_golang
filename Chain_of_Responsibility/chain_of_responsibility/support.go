package chainOfResponsibility

import "fmt"

type supportInterface interface {
	resolve(trouble *Trouble) bool
	Handle(trouble *Trouble)
	SetNext(next supportInterface) supportInterface
}

type support struct {
	name string
	own  supportInterface
	next supportInterface
}

// SetNext func for relating to next Supporter
func (s *support) SetNext(next supportInterface) supportInterface {
	s.next = next
	return next
}

// Handle func for handling trouble
func (s *support) Handle(trouble *Trouble) {
	if s.own.resolve(trouble) {
		s.done(trouble)
	} else if s.next != nil {
		s.next.Handle(trouble)
	} else {
		s.fail(trouble)
	}
}

func (s *support) print() string {
	return fmt.Sprintf("[%s]", s.name)
}

func (s *support) done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble.print(), s.print())
}

func (s *support) fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble.print())
}

// NoSupport is struct
type NoSupport struct {
	*support
}

// NewNoSupport func for initializing NoSupport
func NewNoSupport(name string) *NoSupport {
	noSupport := &NoSupport{
		support: &support{
			name: name,
		},
	}
	noSupport.own = noSupport
	return noSupport
}

func (n *NoSupport) resolve(trouble *Trouble) bool {
	return false
}

// LimitSupport is struct
type LimitSupport struct {
	*support
	limit int
}

// NewLimitSupport func for initializing LimitSupport
func NewLimitSupport(name string, limit int) *LimitSupport {
	limitSupport := &LimitSupport{
		support: &support{
			name: name,
		},
		limit: limit,
	}
	limitSupport.own = limitSupport
	return limitSupport
}

func (l *LimitSupport) resolve(trouble *Trouble) bool {
	if trouble.getNumber() < l.limit {
		return true
	}
	return false
}

// OddSupport is struct
type OddSupport struct {
	*support
}

// NewOddSupport func for initializing OddSupport
func NewOddSupport(name string) *OddSupport {
	oddSupport := &OddSupport{
		support: &support{
			name: name,
		},
	}
	oddSupport.own = oddSupport
	return oddSupport
}

func (o *OddSupport) resolve(trouble *Trouble) bool {
	if trouble.getNumber()%2 == 1 {
		return true
	}
	return false
}

// SpecialSupport is struct
type SpecialSupport struct {
	*support
	number int
}

// NewSpecialSupport func for initializing SpecialSupport
func NewSpecialSupport(name string, number int) *SpecialSupport {
	specialSupport := &SpecialSupport{
		support: &support{
			name: name,
		},
		number: number,
	}
	specialSupport.own = specialSupport
	return specialSupport
}

func (s *SpecialSupport) resolve(trouble *Trouble) bool {
	if trouble.getNumber() == s.number {
		return true
	}
	return false
}
