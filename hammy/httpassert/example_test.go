package httpassert_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/httpassert"
)

func ExampleResp_Status() {
	printExample(httpassert.Response(exampleResponse(http.StatusCreated, nil, "")).Status(http.StatusCreated))
	// Output:
	// message="got status <201>, wanted <201>"
	// success=true
}

func ExampleResp_StatusInRange() {
	printExample(httpassert.Response(exampleResponse(http.StatusNoContent, nil, "")).StatusInRange(200, 299))
	// Output:
	// message="got status <204>, wanted in range <200..299>"
	// success=true
}

func ExampleResp_Header() {
	resp := exampleResponse(http.StatusOK, http.Header{"Content-Type": {"application/json"}}, "")
	printExample(httpassert.Response(resp).Header("Content-Type", "application/json"))
	// Output:
	// message="got header <Content-Type>=<application/json>, wanted <application/json>"
	// success=true
}

func ExampleResp_HeaderContains() {
	resp := exampleResponse(http.StatusOK, http.Header{"Content-Type": {"application/json; charset=utf-8"}}, "")
	printExample(httpassert.Response(resp).HeaderContains("Content-Type", "application/json"))
	// Output:
	// message="got header <Content-Type>=<application/json; charset=utf-8>, wanted containing <application/json>"
	// success=true
}

func ExampleResp_BodyEqual() {
	printExample(httpassert.Response(exampleResponse(http.StatusOK, nil, "hello world")).BodyEqual("hello world"))
	// Output:
	// message="got body <hello world>, wanted equal to <hello world>"
	// success=true
}

func ExampleResp_BodyContains() {
	printExample(httpassert.Response(exampleResponse(http.StatusOK, nil, "hello world")).BodyContains("world"))
	// Output:
	// message="got body <hello world>, wanted containing <world>"
	// success=true
}

func ExampleResp_BodyMatchesRegexp() {
	printExample(httpassert.Response(exampleResponse(http.StatusOK, nil, "status 204")).BodyMatchesRegexp(`status \d+`))
	// Output:
	// message="got body <status 204>, wanted regexp <status \\d+>"
	// success=true
}

func exampleResponse(status int, headers http.Header, body string) *http.Response {
	recorder := httptest.NewRecorder()
	for key, values := range headers {
		for _, value := range values {
			recorder.Header().Add(key, value)
		}
	}
	recorder.WriteHeader(status)
	_, _ = recorder.WriteString(body)
	return recorder.Result()
}

func printExample(result hammy.AssertionMessage) {
	fmt.Printf("message=%q\nsuccess=%t\n", result.Message, result.IsSuccessful)
}
