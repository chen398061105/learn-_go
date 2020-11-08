package main

import "fmt"

func main() {
	// num := 3
	// switch num {
	// case 1:
	// 	fmt.Println("num : 1")
	// case 2:
	// 	fmt.Println("num : 2")
	// case 3:
	// 	fmt.Println("num : 3")
	// default:
	// 	fmt.Println("no num")
	// }
	//数组里面只能放相同类型的
	// arrs := [...]string{"one", "two", "three"}
	// for i := 0; i < len(arrs); i++ {
	// 	fmt.Println(arrs[i])
	// }

	// for _, v := range arrs {
	// 	fmt.Println(v)
	// }
	arr := [...]int{1, 3, 5, 7, 8}
	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
