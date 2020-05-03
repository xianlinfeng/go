package Fibonacci

// 打印前10位斐波那契数列
import (
	"testing"
)

// 方法一：用for循环
func TestFibList1(t *testing.T) {
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		t.Log(a)
		a, b = b, a+b
	}
}

// 方法二：使用闭包
func TestFibList2(t *testing.T) {
	f := fibClosures() // 用闭包打印斐波那契数列
	for i := 0; i < 10; i++ {
		t.Log(f())
	}
}

func fibClosures() func() int { //返回斐波那契数列
	a := 0
	b := 1
	return func() int { // 该函数可以访问外部的a，b值，随着每次call该函数，更新外部的a，b值。
		c := a
		a = b
		b = b + c
		return c
	}
}

// 方法三：递归
func fib(n int) int { // get the fibonacci sequence
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func TestFibList3(t *testing.T) {
	for i := 0; i <= 10; i++ {
		t.Log("  ", fib(i))
	}
}
