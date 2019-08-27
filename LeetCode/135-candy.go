package main

import "fmt"

/**
135. 分发糖果

老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:
输入: [1,0,2]
输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。

示例 2:
输入: [1,2,2]
输出: 4
解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/candy
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	n := []int{1, 0, 2}
	fmt.Println(candy(n))

	n = []int{1, 2, 2}
	fmt.Println(candy(n))
}

func candy(ratings []int) int {
	l := len(ratings)
	// 使用数组c记录每人发的糖果数量
	c, sum := make([]int, l), 0
	for i := 0; i < l; i++ {
		c[i] = 1 // 每人至少一颗糖果
	}

	// 从左到右两两比较，相邻分数高的比低的同学多一颗糖果
	// 需要注意的是，如果分数高的同学糖果已经多于分数低的同学，则不应该再多发
	for i := 0; i < l-1; i++ {
		if ratings[i] < ratings[i+1] && c[i+1] <= c[i] {
			c[i+1] = c[i] + 1
		}
	}

	// 从右到左两两比较，相邻分数高的比低的同学多一颗糖果
	// 需要注意的是，如果分数高的同学糖果已经多于分数低的同学，则不应该再多发
	for i := l - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && c[i-1] <= c[i] {
			c[i-1] = c[i] + 1
		}
	}

	// 求和
	for _, v := range c {
		sum += v
	}

	return sum
}
