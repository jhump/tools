package extract

func _() {
	a := 1
	a = 5     //@mark(exSt0, "a")
	a = a + 2 //@mark(exEn0, "2")
	//@extractfunc(exSt0, exEn0)
	b := a * 2
	_ = 3 + 4
}
