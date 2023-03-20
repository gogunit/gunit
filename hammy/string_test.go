package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
)

func Test_string_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.String("hi").EqualTo("by"))
	aSpy.HadErrorContaining(t, "got <hi>, wanted equal to <by>")
}

func Test_string_EqualTo_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.String("hi").EqualTo("hi"))
}

func Test_String_ToLowerEqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.String("hi").ToLowerEqualTo("BYE"))
	aSpy.HadErrorContaining(t, "got <hi>, wanted equal to <bye>")
}

func Test_String_ToLowerEqualTo_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.String("hI").ToLowerEqualTo("Hi"))
}

func Test_string_Contains_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.String("hello world").Contains("goodbye"))
	aSpy.HadErrorContaining(t, "got <hello world>, wanted substring <goodbye>")
}

func Test_string_Contains_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.String("hello world").Contains("world"))
}

func Test_string_HasPrefix_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.String("hello world").HasPrefix("good"))
	aSpy.HadErrorContaining(t, "got <hello world>, wanted prefix <good>")
}

func Test_string_HasPrefix_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.String("hello world").HasPrefix("hel"))
}

func Test_string_HasSuffix_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.String("hello world").HasSuffix("good"))
	aSpy.HadErrorContaining(t, "got <hello world>, wanted suffix <good>")
}

func Test_string_HasSuffix_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.String("hello world").HasSuffix("world"))
}

func Test_string_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.String("hello world").IsEmpty())
	aSpy.HadErrorContaining(t, "got <hello world>, wanted an empty string")
}

func Test_string_IsEmpty_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.String("").IsEmpty())
}
