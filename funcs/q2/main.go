package main

func recursiveFunction(input int) {
	if input/2 > 1 {
		recursiveFunction(input / 2)
	}
	println(input)
}

func main() {
	recursiveFunction(9)
}
