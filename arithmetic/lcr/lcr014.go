package lcr

import "fmt"

// https://leetcode.cn/problems/MPnaiL/

// LCR 014. 字符串的排列

// 分析特点:
// 1. s1 的其中一个排列是 s2 的子串，意味着只要 s1 的每个字符的个数，和 s2 子串里每个字符的个数相等，则满足

func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}

	var cn1, cn2 [26]int
	// 计算 s1 中每个字符的个数
	for i, c := range s1 {
		cn1[c-'a']++
		cn2[s2[i]-'a']++ // 这里应该算是个特殊优化，判断s2的首个s1长度的子串是否和s1的字符数相等，也可以不判断
		// 如果不判断，则下面的滑动窗口部分，边界需要重新确定
	}
	if cn1 == cn2 {
		return true
	}

	for i := n; i < m; i++ {
		cn2[s2[i]-'a']++   // 此乃当前窗口最后一个字符
		cn2[s2[i-n]-'a']-- // 此乃当前窗口第一个字符
		if cn1 == cn2 {
			return true
		}
	}

	return false
}

func MainForCheckInclusion() {
	fmt.Println(checkInclusion("bak", "jfkjabcki"))
}

// 总结一下:
// 我的思维缺陷:
// 循规蹈矩，总是只能顺序的解决问题: 先把字符串进行全排列，然后利用滑动窗口，把问题转换为判断子串

// 优化思路: 先观察规律，再判断写法
