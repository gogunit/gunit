package hammy

import "time"

func Eventually(assertion func() AssertionMessage, waitFor, tick time.Duration) AssertionMessage {
	tick = normalizePollInterval(tick)
	deadline := time.Now().Add(waitFor)
	attempts := 0
	var last AssertionMessage

	for {
		attempts++
		result := assertion()
		if result.IsSuccessful {
			return Assert(true, "condition succeeded after %d attempts", attempts)
		}
		last = result

		if !time.Now().Before(deadline) {
			return Assert(false, "condition did not succeed within <%v> after %d attempts: %s", waitFor, attempts, last.Message)
		}
		sleepUntilNextPoll(deadline, tick)
	}
}

func Never(assertion func() AssertionMessage, waitFor, tick time.Duration) AssertionMessage {
	tick = normalizePollInterval(tick)
	deadline := time.Now().Add(waitFor)
	attempts := 0
	var last AssertionMessage

	for {
		attempts++
		result := assertion()
		if result.IsSuccessful {
			return Assert(false, "condition succeeded within <%v> after %d attempts: %s", waitFor, attempts, result.Message)
		}
		last = result

		if !time.Now().Before(deadline) {
			return Assert(true, "condition stayed unsuccessful for <%v> after %d attempts: %s", waitFor, attempts, last.Message)
		}
		sleepUntilNextPoll(deadline, tick)
	}
}

func Consistently(assertion func() AssertionMessage, duration, tick time.Duration) AssertionMessage {
	tick = normalizePollInterval(tick)
	deadline := time.Now().Add(duration)
	attempts := 0

	for {
		attempts++
		result := assertion()
		if !result.IsSuccessful {
			return Assert(false, "condition failed within <%v> after %d attempts: %s", duration, attempts, result.Message)
		}

		if !time.Now().Before(deadline) {
			return Assert(true, "condition stayed successful for <%v> after %d attempts", duration, attempts)
		}
		sleepUntilNextPoll(deadline, tick)
	}
}

func normalizePollInterval(tick time.Duration) time.Duration {
	if tick <= 0 {
		return time.Millisecond
	}
	return tick
}

func sleepUntilNextPoll(deadline time.Time, tick time.Duration) {
	remaining := time.Until(deadline)
	if remaining <= 0 {
		return
	}
	if remaining < tick {
		time.Sleep(remaining)
		return
	}
	time.Sleep(tick)
}
