package main

import "syscall/js"

func main() {
	done := make(chan int, 0)
	js.Global().Set("php_strip", js.FuncOf(StripFunc))
	<-done
}

func StripFunc(this js.Value, args []js.Value) interface{} {
	res, err := Strip([]byte(args[0].String()))
	if err != nil {
		return js.ValueOf(err.Error())
	}

	return js.ValueOf(string(res))
}
