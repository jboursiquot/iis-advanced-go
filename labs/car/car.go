package car

var speed = 100

func Increase(acclerate int, current int) int {
	return current + acclerate
}

func Decrease(deaccelerate int, current int) int {
	// implement code to keep speed
	// between 0 and 100
	return current - deaccelerate
}
