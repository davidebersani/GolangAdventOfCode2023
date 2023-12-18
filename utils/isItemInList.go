func isItemInList[T comparable](item T, list []T) bool {
	for _, o := range list {
		if o == item {
			return true
		}
	}
	return false
}