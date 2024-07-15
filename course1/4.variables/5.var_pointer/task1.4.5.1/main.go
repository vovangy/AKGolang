package main

import "fmt"

func changeInt(a *int) {
	if a == nil {
		return
	}
	*a = 20
}

func changeFloat(b *float64) {
	if b == nil {
		return
	}
	*b = 6.28
}

func changeString(c *string) {
	if c == nil {
		return
	}
	*c = "Goodbye, world!"
}

func changeBool(d *bool) {
	if d == nil {
		return
	}
	*d = false
}

func main() {
	var a int = 10
	var b float64 = 0.3
	var c string = "AIUUU"
	var d bool = false

	changeInt(&a)
	changeFloat(&b)
	changeString(&c)
	changeBool(&d)

	fmt.Println(a, b, c, d)
	return
}
