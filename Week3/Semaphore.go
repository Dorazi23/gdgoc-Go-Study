func main() {
	semaphore := make(chan struct{})
	go func() {
		semaphore <- struct{}{}
		fmt.Println("signalling")
	}()
	<-semaphore
	fmt.Println("exiting")
}