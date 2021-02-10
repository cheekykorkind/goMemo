package basic

func If1() bool {
	var v1 int = 1
	var v2 int = 2
	var v3 bool

	if v1 == v2 {
		v3 = true
	} else {
		v3 = false
	}

	return v3
}
