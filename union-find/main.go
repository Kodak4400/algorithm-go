package main

func root(x int) int {
	return x
}

func unite(x int, y int) bool {
	x = root(x)
	y = root(y)

	if x == y {
		return false
	}
}

func main() {
	//...
}
