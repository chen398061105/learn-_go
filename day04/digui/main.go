package main

import "fmt"

//递归必须有一个退出条件，适合问题相同，问题越来越下场景
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

func taijie(num uint64) uint64 {
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 2
	}
	return taijie(num-1) + taijie(num-2)
}
func main() {
	// ret := f(5)
	// fmt.Println(ret)
	ret := taijie(4)
	fmt.Println(ret)
}
