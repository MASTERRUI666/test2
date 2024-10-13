package main

import "fmt"

/*******************************************************************************
* 函数名称: smallerNumbersThanCurrent
* 函数功能: 统计数组中比当前数字小的数字的数量
* 函数参数: 3个
* 参数名称:          类型         输入/输出      描述
* anums             []int        输入       输入的数组
* u32anumsSize      int          输入       输入数组的有效元素个数
* aret              []int        输出       输出的数组，已申请255个元素的内存
*
* 返回值: int，用于返回函数执行结果，0表示执行成功，-1表示执行失败
*******************************************************************************/
func smallerNumbersThanCurrent(anums []int, u32anumsSize int, aret []int) int {
	if u32anumsSize > 255 {
		return -1 // 输入超过最大有效元素个数
	}

	// 创建计数数组，用于统计每个数字出现的次数
	count := make([]int, 101) // 假设 nums[i] 的范围是 0 到 100
	for i := 0; i < u32anumsSize; i++ {
		count[anums[i]]++
	}

	// 计算小于当前数字的数量
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// 填充结果数组
	for i := 0; i < u32anumsSize; i++ {
		if anums[i] > 0 {
			aret[i] = count[anums[i]-1]
		} else {
			aret[i] = 0
		}
	}

	return 0 // 执行成功
}

func main() {
	anums := []int{8, 1, 2, 2, 3}
	aret := make([]int, 255) // 申请255个元素的内存
	result := smallerNumbersThanCurrent(anums, len(anums), aret)

	if result == 0 {
		fmt.Println(aret[:len(anums)]) // 输出: [4, 0, 1, 1, 3]
	} else {
		fmt.Println("执行失败")
	}
}
