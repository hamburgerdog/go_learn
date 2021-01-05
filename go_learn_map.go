/*
练习：映射
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。
	```func Fields(s string) []string```
	Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning an array of substrings of s or an empty list if s contains only white space.
*/
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	strs := strings.Fields(s)
	for _, v := range strs {
		if elem, ok := m[v]; ok == true {
			m[v] = elem + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
