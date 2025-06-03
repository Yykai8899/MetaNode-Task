package main

import (
	"fmt"
	"strconv"
)

func main() {
	//返回只出现一次的数字
	nums1 := []int{2, 2, 5, 3, 4, 3, 4}
	fmt.Println(singleNumber(nums1))

	//回文数
	num1 := 1221
	fmt.Println(isPalindrome(num1))

	// 有效括号
	str1 := "({[]})"
	fmt.Println(isValid(str1))

	// 最长公共前缀
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	// 删除有序数组中的重复项
	nums2 := []int{1, 1, 1, 2, 2, 3, 3, 4, 5}
	k := removeDuplicates(nums2)
	fmt.Println(k)
	// 按网站题目要求验证断言
	// 正确数组
	expectedNums := []int{1, 2, 3, 4, 5}
	if k == len(expectedNums) {
		for i := 0; i < k; i++ {
			if expectedNums[i] != nums2[i] {
				fmt.Println("验证删除有序数组中重复项失败，与目标数组不符")
				break
			}
		}
		fmt.Println("验证删除有序数组中重复项通过")
	} else {
		fmt.Println("验证删除有序数组中重复项失败，与目标数组不符")
	}

	// 加一
	digits := []int{9}
	fmt.Println(plusOne(digits))

	// 两数之和
	nums, target := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum(nums, target))
}

// 只返回出现一次的数字
// https://leetcode.cn/problems/single-number/
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

// 回文数
// https://leetcode.cn/problems/palindrome-number/description/
func isPalindrome(num int) bool {
	// 将数字转换为字符串
	s := strconv.Itoa(num)
	// 将字符串转换为byte切片
	bytes := []byte(s)
	// 声明左右两个字符串起始下标
	left, right := 0, len(s)-1
	// 如果right小于left停止循环（证明字符对比完毕）
	for left < right {
		// 收尾两个数字进行判断不同则证明不是回文数（回文数应该是一段数字的第一位与最后一位为同样的数字）
		if bytes[left] != bytes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 有效括号
// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	strLen := len(s)
	// 合法性校验 为奇数证明此字符串必定不会有完整括号
	if strLen%2 == 1 {
		return false
	}
	pairsMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for i := 0; i < strLen; i++ {
		vlaue := s[i]
		// 出入栈判断，vlaue为左边括号则入栈，右边则出栈
		if pairsMap[vlaue] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairsMap[vlaue] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}
	// 栈长度为0表示数据正确
	return len(stack) == 0
}

// 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	// 获取字符串长度最小的长度，避免避免下标越界，同时长字符串与短字符串字符相比时与空值对比
	for _, str := range strs[1:] {
		if len(str) == 0 {
			return ""
		}
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	var resStr []byte
	// 根据获取的最小长度循环遍历字符串的每一个字符
	for i := 0; i < minLen; i++ {
		tmpStr := strs[0][i]
		// 根据字符循环数组每个字符串的每一个字符
		for _, str := range strs {
			// 判断字符不相等则直接返回结果
			if str[i] != tmpStr {
				return string(resStr)
			}

		}
		// 字符相同则将相同的字符存入resStr
		resStr = append(resStr, tmpStr)
	}
	return string(resStr)
}

// 删除有序数组中重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 与数组下一元素做对比，相同则删除下一元素
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	// 复盘：
	// 当时最开始的思路是循环遍历将当前元素与上一元素进行对比，相同则删除当前元素，执行后发现正常数字重复出现两次是可以正常执行
	// 但是出现三个以上重复就会漏排除，后面分析在切片元素删除时删除当前元素会影响后面循环，比如说当前元素删除后下标1的数据已经截取了
	// 数组长度已经减少 而循环还是按照原数组进行循环，导致部分元素会跳过循环导致漏判断
	// for i := 0; i < len(nums); i++ {
	// 	// 与数组下一元素做对比，相同则删除下一元素
	// 	if i>0 && nums[i] == nums[i-1] {
	// 		nums = append(nums[:i], nums[i+1:]...)
	//
	// 	}
	// }
	return len(nums)
}

// 加一
// https://leetcode.cn/problems/plus-one/description/
func plusOne(digits []int) []int {
	// 不是最优解 根据数组长度判断是否做进位处理，数组长度为1才做进位处理，而数组长度如果大于1则做末位加一处理，这种情况下可能会出现末位加一后结果
	// 需要进位，目前仅按照数据示例条件做处理 老师指导一下最优方法思路
	digistLen := len(digits)
	newDigits := []int{}
	if digistLen == 1 {
		onlyNum := digits[digistLen-1]
		onlyNum++
		if onlyNum%10 == 0 {
			s := strconv.Itoa(onlyNum)
			bytes := []byte(s)
			for _, v := range bytes {
				num, err := strconv.Atoi(string(v))
				if err != nil {
					panic(err)
				}
				newDigits = append(newDigits, num)
			}
		} else {
			newDigits = append(newDigits, onlyNum)
		}
	} else if digistLen > 1 {
		lastNum := digits[digistLen-1]
		for i := 0; i < digistLen; i++ {
			if i == digistLen-1 {
				lastNum++
				newDigits = append(newDigits, lastNum)
				break
			}
			newDigits = append(newDigits, digits[i])
		}
	}
	return newDigits
}

// 两数之和
// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	countMap := make(map[int]int)
	for numIndex, num := range nums {
		tmpRes := target - num
		if index, ok := countMap[tmpRes]; ok {
			return []int{index, numIndex}
		} else {
			countMap[num] = numIndex
		}

	}
	return nil
}
