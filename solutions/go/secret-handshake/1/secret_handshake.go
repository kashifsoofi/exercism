// Package secret implements a simple library that implements
// method to generate sequence of events from secret code
package secret

const (
	wink          uint = 1 << iota // wink
	doubleBlink                    // double wink
	closeYourEyes                  // close your eyes
	jump                           // jump
	reverse                        // reverse
)

var operations = []struct {
	key   uint
	event string
}{
	{wink, "wink"},
	{doubleBlink, "double blink"},
	{closeYourEyes, "close your eyes"},
	{jump, "jump"},
}

// Handshake given a code returns sequence of events
func Handshake(code uint) []string {
	var events []string
	for _, op := range operations {
		if op.key&code == op.key {
			events = append(events, op.event)
		}
	}

	if code&reverse == reverse {
		for i, j := 0, len(events)-1; i < j; i, j = i+1, j-1 {
			events[i], events[j] = events[j], events[i]
		}
	}
	return events
}
