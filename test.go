package main

import "fmt"

func main() {
	var l, r int
	fmt.Scanf("%d,%d", &l, &r)
	temp := 0
	num := 0
	for i := l; i <= r; i++ {
		temp += i
		if temp%3 == 0 {
			num += 1
		}
	}
	fmt.Println(num)

}
