package httpassert_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/httpassert"
)

func ExampleResp_HasStatus() {
	printExample(httpassert.Response(exampleResponse(http.StatusCreated, nil, "")).HasStatus(http.StatusCreated))
	// Output:
	// message="got status <201>, wanted <201>"
	// success=true
}

func ExampleResp_HasStatusInRange() {
	printExample(httpassert.Response(exampleResponse(http.StatusNoContent, nil, "")).HasStatusInRange(200, 299))
	// Output:
	// message="got status <204>, wanted in range <200..299>"
	// success=true
}

func ExampleResp_HeaderEqualTo() {
	resp := exampleResponse(http.StatusOK, http.Header{"Content-Type": {"application/json"}}, "")
	printExample(httpassert.Response(resp).HeaderEqualTo("Content-Type", "application/json"))
	// Output:
	// message="got header <Content-Type>=<application/json>, wanted <application/json>"
	// success=true
}

func ExampleResp_HasHeaderContaining() {
	resp := exampleResponse(http.StatusOK, http.Header{"Content-Type": {"application/json; charset=utf-8"}}, "")
	printExample(httpassert.Response(resp).HasHeaderContaining("Content-Type", "application/json"))
	// Output:
	// message="got header <Content-Type>=<application/json; charset=utf-8>, wanted containing <application/json>"
	// success=true
}

func ExampleResp_BodyEqualTo() {
	printExample(httpassert.Response(exampleResponse(http.StatusOK, nil, "hello world")).BodyEqualTo("hello world"))
	// Output:
	// message="got body <hello world>, wanted equal to <hello world>"
	// success=true
}

func ExampleResp_HasBodyContaining() {
	printExample(httpassert.Response(exampleResponse(http.StatusOK, nil, "hello world")).HasBodyContaining("world"))
	// Output:
	// message="got body <hello world>, wanted containing <world>"
	// success=true
}

func ExampleResp_BodyMatches() {
	printExample(httpassert.Response(exampleResponse(http.StatusOK, nil, "status 204")).BodyMatches(`status \d+`))
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
