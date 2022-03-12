package events

import (
	"github.com/pellared/fluentassert/f"
	"testing"
)

func channelEmpty[T any](c chan T) bool {
	select {
	case _, ok := <-c:
		if ok {
			return false
		} else {
			return true
		}
	default:
		return true
	}
}

func TestSubscription(t *testing.T) {
	feed := Make[string]()
	ch := make(chan string, 1)
	called := false
	feed.Subscribe(func(data string) {
		called = true
		ch <- data
	})

	f.Assert(t, called).Eq(false, "subscribing shouldn't call anything")

	gotCount := feed.Send("Hello World")
	f.Assert(t, gotCount).Eq(1, "Send() should have expected subscriber count")

	f.Assert(t, <-ch).Eq("Hello World", "callback should have been called")
	f.Assert(t, called).Eq(true, "confirmed called")
}

func TestUnSubscribe(t *testing.T) {
	feed := Make[string]()
	ch1 := make(chan string)
	ch2 := make(chan string)
	feed.Subscribe(func(data string) {
		ch1 <- data
	})
	unsub := feed.Subscribe(func(data string) {
		ch2 <- data
	})

	gotCount := feed.Send("Event 1")
	f.Assert(t, gotCount).Eq(2, "Send() should have expected subscriber count")

	// Wait for response
	f.Assert(t, <-ch1).Eq("Event 1", "callback 1 should have been called") // s1
	f.Assert(t, <-ch2).Eq("Event 1", "callback 2 should have been called") // s2

	unsub.UnSubscribe()

	gotCount = feed.Send("Event 2")
	f.Assert(t, gotCount).Eq(1, "Send() should have expected subscriber count")

	// Wait for response
	f.Assert(t, <-ch1).Eq("Event 2", "callback 1 should have been called") // s1
	f.Assert(t, channelEmpty(ch2)).Eq(true, "channel 2 expected to be empty")
	go func() { ch2 <- "no more calls" }()
	f.Assert(t, <-ch2).Eq("no more calls", "no other calls should be in channel")
}

func TestSubscriptionInSubscription(t *testing.T) {
	feed := Make[string]()
	ch1 := make(chan string)
	ch2 := make(chan string)
	called := false
	feed.Subscribe(func(data string) {
		ch1 <- data
		feed.Subscribe(func(data string) {
			called = true
			ch2 <- data
		})
	})

	gotCount := feed.Send("Event 1")
	f.Assert(t, gotCount).Eq(1, "Send() should have expected subscriber count")

	// Wait for response
	f.Assert(t, <-ch1).Eq("Event 1", "callback should have been called") // first event
	f.Assert(t, called).Eq(false, "second callback shouldn't be called yet")

	gotCount = feed.Send("Event 2")
	f.Assert(t, gotCount).Eq(2, "Send() should have expected subscriber count")
	f.Assert(t, <-ch1).Eq("Event 2", "callback should have been called") // second event
	f.Assert(t, <-ch2).Eq("Event 2", "callback should have been called") // second event
}

func TestEmitOnEmit(t *testing.T) {
	feed := Make[string]()
	ch := make(chan string, 2)
	subEmit := true
	feed.Subscribe(func(data string) {
		ch <- data
		if subEmit {
			subEmit = false // toggle false as to not inf loop
			feed.Send("Event 2")
		}
	})

	feed.Send("First Event")

	firstGot := <-ch
	secondGot := <-ch

	f.Assert(t, firstGot).Eq("First Event", "callback should have been called")
	f.Assert(t, secondGot).Eq("Event 2", "callback should have been called")
}
