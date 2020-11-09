package main

import (
	"fmt"
)

/*
a 名字有1个e或者E的 1枚金币
b ********i  I     2
c         o   O    3
d         u   U    4
*/

var (
	coins = 5000
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	//依次拿到每个人的名字
	for _, name := range users {
		for _, c := range name {
			//根据规则分配金币
			switch c {
			case 'e', 'E':
				//每个人分的金币保存到map中
				distribution[name]++
				//记录剩下金币
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return
}

//计算单词出现次数
func main() {
	left := dispatchCoin()
	fmt.Println("剩下金币：", left)

	for name, count := range distribution {
		fmt.Printf("%s %d\n", name, count)
	}
}
