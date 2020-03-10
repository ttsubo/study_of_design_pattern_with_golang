package chainOfResponsibility

import "fmt"

// SupportInterface is interface
type SupportInterface interface {
	resolve(trouble Trouble) bool
	Handle(support SupportInterface, trouble Trouble)
}

// Support is struct
type Support struct {
	Name string
	next SupportInterface
}

// SetNext func for relating to next Supporter
func (s *Support) SetNext(next SupportInterface) {
	s.next = next
}

// Handle func for handling trouble
func (s *Support) Handle(support SupportInterface, trouble Trouble) {
	if support.resolve(trouble) {
		s.done(trouble)
	} else if s.next != nil {
		s.next.Handle(s.next, trouble)
	} else {
		s.fail(trouble)
	}
}

func (s *Support) print() string {
	return fmt.Sprintf("[%s]", s.Name)
}

func (s *Support) done(trouble Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble.print(), s.print())
}

func (s *Support) fail(trouble Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble.print())
}

// NoSupport is struct
type NoSupport struct {
	*Support
}

func (n *NoSupport) resolve(trouble Trouble) bool {
	return false
}

// NewNoSupport func for initializing NoSupport
func NewNoSupport(name string) *NoSupport {
	return &NoSupport{
		Support: &Support{Name: name},
	}
}

// LimitSupport is struct
type LimitSupport struct {
	*Support
	Limit int
}

// NewLimitSupport func for initializing LimitSupport
func NewLimitSupport(name string, limit int) *LimitSupport {
	return &LimitSupport{
		Support: &Support{Name: name},
		Limit:   limit,
	}
}

func (l *LimitSupport) resolve(trouble Trouble) bool {
	if trouble.getNumber() < l.Limit {
		return true
	}
	return false
}

// OddSupport is struct
type OddSupport struct {
	*Support
}

// NewOddSupport func for initializing OddSupport
func NewOddSupport(name string) *OddSupport {
	return &OddSupport{
		Support: &Support{Name: name},
	}
}

func (o *OddSupport) resolve(trouble Trouble) bool {
	if trouble.getNumber()%2 == 1 {
		return true
	}
	return false
}

// SpecialSupport is struct
type SpecialSupport struct {
	*Support
	Number int
}

// NewSpecialSupport func for initializing SpecialSupport
func NewSpecialSupport(name string, number int) *SpecialSupport {
	return &SpecialSupport{
		Support: &Support{Name: name},
		Number:  number,
	}
}

func (s *SpecialSupport) resolve(trouble Trouble) bool {
	if trouble.getNumber() == s.Number {
		return true
	}
	return false
}
