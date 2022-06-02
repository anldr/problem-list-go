package main

import (
	. "../common_utils"
	"fmt"
)

func main() {
	fmt.Println(largestVariance("ababbbbb"))
}

/*
暴力枚举两个字符（不能相同），以字符 aa 和 bb 为例，我们把字符串中的所有 aa 视为 11，所有的 bb 视为 -1−1，其它为 00。这样，只需要求出转换后的 最大子数组和，即可得所有子串中 \text{count}(a) - \text{count}(b)count(a)−count(b) 的最大值。

但是，还需要注意一个条件，就是子字符串中必须存在 bb。那么，稍微更改一下求 最大子数组和 的思路即可。设

dp_{i,0}=dp
i,0
​
 = 以字符串的第 ii 个字符结尾的，转换后的最大子数组和；
dp_{i,1}=dp
i,1
​
 = 以字符串的第 ii 个字符结尾的，且包含至少一个 bb 的最大子数组和。

那么，状态转移为：

dp_{i,0} = \max(dp_{i-1,0} + v, v)dp
i,0
​
 =max(dp
i−1,0
​
 +v,v)，其中，v = \left\{ \begin{aligned} 1&, s[i] = a\\ -1&,s[i] = b \\ 0&, \text{otherwise}\\ \end{aligned} \right.v=
⎩
⎪
⎪
⎨
⎪
⎪
⎧
​

1
−1
0
​

,s[i]=a
,s[i]=b
,otherwise
​
 。
（这里套用了最大子数组和的 dpdp）。

dp_{i,1} =\left\{ \begin{aligned} &\max(dp_{i-1,1} + v,\ dp_{i-1,0} + v,\ v) &, s[i] = b \\ &dp_{i-1,1} + v&, s[i] \ne b \end{aligned} \right.dp
i,1
​
 ={
​

max(dp
i−1,1
​
 +v, dp
i−1,0
​
 +v, v)
dp
i−1,1
​
 +v
​

,s[i]=b
,s[i]

​
 =b
​
  。
注意当 s[i] = bs[i]=b 时和 s[i] \ne bs[i]

​
 =b 时的状态转移的不同。

这样我们暴力枚举完所有的字符对之后，总有一个是正确答案，其它答案都比正确答案小。因此答案就是之前求的所有的 dp_{i,1}dp
i,1
​
  的最大值。

时间复杂度：O(n\cdot 26^2)O(n⋅26
2
 )

代码实现要注意 dp_0dp
0
​
  和 dp_1dp
1
​
  的更新顺序。
*/
func largestVariance(s string) int {
	var ret = 0

	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			var dp0 = -999999999
			var dp1 = -999999999
			for i := range s {
				var val = 0
				if rune(s[i]) == a {
					val = 1
				} else if rune(s[i]) == b {
					val = -1
				}

				if rune(s[i]) == b {
					dp1 = MaxInt32(val, MaxInt32(dp0+val, dp1+val))
				} else {
					dp1 = dp1 + val
				}

				dp0 = MaxInt32(dp0+val, val)

				ret = MaxInt32(ret, dp1)
			}
		}
	}

	return ret
}
