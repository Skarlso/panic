package panic

type PanicHandler func(interface{}) bool

var PanicHandlers = []PanicHandler{}

func SetPanicHandler(handler PanicHandler) {
	PanicHandlers = append(PanicHandlers, handler)
}

func HandleCrash() {
	if r := recover(); r != nil {
		reallyPanic := true
		for _, handler := range PanicHandlers {
			reallyPanic = handler(r)
		}
		if reallyPanic {
			panic(r)
		}
	}
}

func Panic(n int) {
	defer HandleCrash()
	if n == 10 {
		panic("I'm panicing")
	}
}
