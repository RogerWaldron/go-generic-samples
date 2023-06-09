func SliceToChannel[T any](slice []T) chan T {
	channel := make(chan T)
	
	go func() {
		for _, item := range slice {
			channel <- item
		}
	}()

	return channel
}