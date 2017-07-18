package main

import (
	"testing"
)

func Test_Rand_1(t *testing.T) {
	if p := Rand0(); p != 0 && p != 1 && p != 2 && p != 3 { //try a unit test on function
		t.Error("随机函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}
