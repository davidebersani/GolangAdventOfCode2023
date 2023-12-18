func SafeGet[K, V comparable](m map[K]V, key K, defaultReturnValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultReturnValue
}