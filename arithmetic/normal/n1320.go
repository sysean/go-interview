package normal

// https://leetcode.cn/problems/minimum-distance-to-type-a-word-using-two-fingers/

// 1320. 二指输入的的最小距离

// 这里的位置为指向的字母编号，例如 A 对应 0，B 对应 1，以此类推，而非字母在键盘上的位置。

// 如果左手在 word[i - 1] 的位置，那么在输入第 i 个字母时，左手从 word[i - 1] 移动至 word[i]，状态转移方程为：
// dp[i][l = word[i]][r] = dp[i - 1][l0 = word[i - 1]][r] + dist(word[i - 1], word[i])

func minimumDistance(word string) int {
	//var dp [][]int
	//for i := 0; i < 26; {
	//	for j := 0; j < 6 && i < 26; j++ {
	//		dp[]
	//	}
	//}
	return 0
}
