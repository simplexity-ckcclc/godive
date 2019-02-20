package grammar

import "fmt"

// should print "select from closed channel. msg:"
// select from closed channel, just step into the select clause
func SelectClosedChannel() {
	channel := make(chan string)

	close(channel)

	select {
	case msg := <-channel:
		fmt.Printf("select from closed channel. msg:%s", msg)
	default:
		fmt.Println("step over to default")
	}
}
