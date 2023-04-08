package mapString

// CheckKeyExist 检查key是否存在
func CheckKeyExist(targetMap map[string]string, targetKey string) bool {
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

// InitMap 初始化map
func InitMap(targetMap map[string]string) map[string]string {
	return make(map[string]string)
}
