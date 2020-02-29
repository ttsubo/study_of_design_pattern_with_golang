package bridge

type displayImplInterface interface {
	rawOpen()
	rawPrint()
	rawClose()
}
