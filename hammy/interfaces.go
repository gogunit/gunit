package hammy

import "fmt"

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Stringy interface {
	~string
}

func Assert(isSuccessful bool, str string, args ...any) AssertionMessage {
	msg := AssertionMessage{
		IsSuccessful: isSuccessful,
	}
	if !isSuccessful {
		msg.Message = fmt.Sprintf(str, args...)
	}
	return msg
}

type AssertionMessage struct {
	IsSuccessful bool
	Message      string
}
