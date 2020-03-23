package bridge

type displayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}
