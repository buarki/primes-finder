package main

import (
	"syscall/js"
)

func findPrimes() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		limit := args[0].Int()
		x := findPrimesUpTo(limit)

		var res []interface{}
		for _, p := range x {
			res = append(res, p)
		}
		return res
	})
}

func main() {
	js.Global().Set("findPrimes", findPrimes())
	<-make(chan bool)
}
