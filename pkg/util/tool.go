package util


// RemoveDuplicate id 去重
func RemoveDuplicate(list []int64) []int64 {
	result := []int64{}
	tempMap := map[int64]byte{}  // 存放不重复主键
	for _, e := range list{
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l{  // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}