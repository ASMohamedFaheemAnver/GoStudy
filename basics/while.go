package main

func main() {
	var i = 0
	for {
		i++
		println("While Loop Iteration:", i)
		if i >= 5000 {
			break
		}
	}
}
