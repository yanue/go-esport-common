/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : slice.go
 Time    : 2018/9/14 17:37
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package util

import "strconv"

type sliceHelper struct {
}

var Slice *sliceHelper = new(sliceHelper)

/*
*@note 反转string序列
*@note s string序列
*@param index 需要删除节点的位置，0Base
*@example
*     {"1","3","2","4"} -> {"4","2","3","1"}
*/
func (this *sliceHelper) ReverseStrings(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
*@note 反转int序列
*@note s int序列
*@example
*     {1,3,2,4} -> {4,2,3,1}
*@return 返回新的slice
*/
func (this *sliceHelper) ReverseIntSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
*@note 删除一个[]int的某个节点
*@param slice 需要操作的slice
*@param index 需要删除节点的位置，0Base
*@example
*     slice = {1,2,3,4}
*     index = 2 -> slice = {1,2,4}
*     index = 4 -> slice = {1,2,3,4}
*     index = 0 -> slice = {2,3,4}
*@return 返回新的slice
 */
func (this *sliceHelper) RemoveIntSlice(slice []int, index int) []int {
	max := len(slice)
	if 0 == max {
		return slice
	}
	if index >= max {
		return slice
	}
	if index == max-1 {
		return slice[:max-1]
	}
	if 0 == index {
		return slice[1:]
	}

	return append(slice[:index], slice[index+1:]...)
}

/*
*note 删除一个[]int的某一段节点
*@param slice 需要操作的slice
*@param start 需要删除节点的第一个的位置，0Base
*@param end 需要删除节点的最后一个的位置，0Base
*@example
*  slice = {1,2,3,4,5,6,7}
*  start = 0 end = 4 -> slice = {6,7}
*  start = 6 end = 10 -> slice = {1,2,3,4,5,6}
*  start = 7 end = 10 -> slice = {1,2,3,4,5,6,7}
*@return 返回新的slice
 */
func (this *sliceHelper) RemovesIntSlice(slice []int, start, end int) []int {
	/*
	*@note 如果尾巴<开始，则不做任何操作
	*@note 如果尾巴=开始，则先对于删除一个节点
	 */
	if end < start {
		return slice
	}
	if end == start {
		return this.RemoveIntSlice(slice, start)
	}

	max := len(slice)
	if start >= max {
		return slice
	}

	if end >= max {
		return slice[:start]
	}

	if 0 == start {
		return slice[end+1:]
	}

	return append(slice[:start], slice[end+1:]...)
}

/*
*note 删除一个[]int的某一段节点
*@param slice 需要操作的slice
*@param start 需要删除节点的第一个的位置，0Base
*@param end 需要删除节点的最后一个的位置，0Base
*@example
*  slice = {1,2,3,4,5,6,7}
*  start = 0 end = 4 -> slice = {6,7}
*  start = 6 end = 10 -> slice = {1,2,3,4,5,6}
*  start = 7 end = 10 -> slice = {1,2,3,4,5,6,7}
*@return 返回新的slice
 */
func (this *sliceHelper) RemoveInt32Slice(slice []int32, index int) []int32 {
	max := len(slice)
	if 0 == max {
		return slice
	}
	if index >= max {
		return slice
	}
	if index == max-1 {
		return slice[:max-1]
	}
	if 0 == index {
		return slice[1:]
	}

	return append(slice[:index], slice[index+1:]...)
}

/*
*note 删除一个[]int的某一段节点
*@param slice 需要操作的slice
*@param start 需要删除节点的第一个的位置，0Base
*@param end 需要删除节点的最后一个的位置，0Base
*@example
*  slice = {1,2,3,4,5,6,7}
*  start = 0 end = 4 -> slice = {6,7}
*  start = 6 end = 10 -> slice = {1,2,3,4,5,6}
*  start = 7 end = 10 -> slice = {1,2,3,4,5,6,7}
*@return 返回新的slice
 */
func (this *sliceHelper) RemovesInt32Slice(slice []int32, start, end int) []int32 {
	/*
	*@note 如果尾巴<开始，则不做任何操作
	*@note 如果尾巴=开始，则先对于删除一个节点
	 */
	if end < start {
		return slice
	}
	if end == start {
		return this.RemoveInt32Slice(slice, start)
	}

	max := len(slice)
	if start >= max {
		return slice
	}

	if end >= max {
		return slice[:start]
	}

	if 0 == start {
		return slice[end+1:]
	}

	return append(slice[:start], slice[end+1:]...)
}

/*
*@note 删除一个[]string的某个节点
*@param slice 需要操作的slice
*@param index 需要删除节点的位置，0Base
*@example
*     slice = {1,2,3,4}
*     index = 2 -> slice = {1,2,4}
*     index = 4 -> slice = {1,2,3,4}
*     index = 0 -> slice = {2,3,4}
*@return 返回新的slice
 */
func (this *sliceHelper) RemoveStringSlice(slice []string, index int) []string {
	max := len(slice)
	if 0 == max {
		return slice
	}
	if index >= max {
		return slice
	}
	if index == max-1 {
		return slice[:max-1]
	}
	if 0 == index {
		return slice[1:]
	}

	return append(slice[:index], slice[index+1:]...)
}

/*
*@note 删除一个[]string的某个节点
*@param slice 需要操作的slice
*@param dest 需要删除的节点
*@example
*     slice = {1,2,3,4}
*     dest = 2 -> slice = {1,3,4}
*     dest = 4 -> slice = {1,2,3}
*     dest = 0 -> slice = {1,2,3,4}
*@return 返回新的slice
 */
func (this *sliceHelper) RemoveStringSliceEx(slice []string, dest string) []string {
	max := len(slice)
	if 0 == max {
		return slice
	}

	index := -1
	for i, v := range slice {
		if v == dest {
			index = i
		}
	}

	if index == -1 {
		return slice
	}

	if index == max-1 {
		return slice[:max-1]
	}
	if 0 == index {
		return slice[1:]
	}

	return append(slice[:index], slice[index+1:]...)
}

/*
*@note 在升序的[]string中插入一个节点
*@param slice 需要操作的slice,按数值升序的字符串数组
*@param dest 需要插入的节点
*@example
*     slice = {"1","3","5","7"}
*     dest = "0" -> slice = {"0", 1","3","5","7"}
*     dest = "1" -> slice = {"1","3","5","7"}
*     dest = "2" -> slice = {"1", "2", 3","5","7"}
*     dest = "8" -> slice = {"1","3","5","7", "8"}
*
*     slice = {}
*     dest = "1" -> slice = {"1"}
*@return 返回新的slice
 */
func (this *sliceHelper) InsertSortStringSlice(slice []string, dest string) []string {
	length := len(slice)
	if 0 == length {
		return []string{dest}
	}

	destValue, _ := strconv.ParseInt(dest, 10, 64)
	max, _ := strconv.ParseInt(slice[length-1], 10, 64)
	if destValue > max {
		slice = append(slice, dest)
		return slice

	}

	min, _ := strconv.ParseInt(slice[0], 10, 64)
	if destValue < min {
		slice = append([]string{dest}, slice...)
		return slice
	}

	for i, v := range slice {
		value, _ := strconv.ParseInt(v, 10, 64)
		if destValue == value {
			break
		}

		if destValue < value {
			newList := append(slice[:i+1], slice[i:]...)
			newList[i] = dest
			slice = newList
			break
		}
	}

	return slice
}
