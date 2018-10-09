package main

// http://legendtkl.com/2017/07/30/understanding-golang-channel/
func main() {
	channel := make(chan int)

	close(channel)
}
