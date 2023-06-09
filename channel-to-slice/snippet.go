func ChannelToSlice[T any] (channel chan T) (slice []T) {
	for item := range channel {
		slice = append(slice, item)
	}
	return slice
}