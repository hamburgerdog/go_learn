/*
练习：斐波纳契闭包
让我们用函数做些好玩的事情。

实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
*/

package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	n := 0
	x := 0
	y := 1
	return func() int {
		if n == 0 {
			n += 1
			return x
		} else if n == 1 {
			n += 1
			return y
		} else {
			tmp := y
			y += x
			x = tmp
			return y
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
