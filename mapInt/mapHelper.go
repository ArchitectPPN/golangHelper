package mapInt

// checkKeyExist 检查key是否存在
func checkKeyExist(targetMap map[int]string, targetKey int) bool {
	// map为nil,返回false
	if targetMap == nil {
		return false
	}

	_, exist := targetMap[targetKey]
	if !exist {
		return false
	}

	//默认返回false
	return true
}

// initMap 初始化map
func initMap(targetMap map[int]string) map[int]string {
	return make(map[int]string)
}
