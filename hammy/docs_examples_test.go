package hammy_test

import (
	"errors"
	"fmt"

	a "github.com/gogunit/gunit/hammy"
)

type examplePerson struct {
	Name string
	Age  int
}

type exampleWrapper struct {
	Person examplePerson
}

type exampleError struct{}

func (exampleError) Error() string {
	return "example error"
}

func printExample(result a.AssertionMessage) {
	fmt.Println(result.IsSuccessful)
}

func ExampleNil() {
	var value any = nil
	printExample(a.Nil(value))
	// Output: true
}

func ExampleNotNil() {
	value := 42
	printExample(a.NotNil(&value))
	// Output: true
}

func ExampleTrue() {
	printExample(a.True(true))
	// Output: true
}

func ExampleFalse() {
	printExample(a.False(false))
	// Output: true
}

func ExampleError() {
	printExample(a.Error(errors.New("boom")))
	// Output: true
}

func ExampleNilError() {
	printExample(a.NilError(nil))
	// Output: true
}

func ExampleErrorIs() {
	target := errors.New("timeout")
	err := fmt.Errorf("request failed: %w", target)
	printExample(a.ErrorIs(err, target))
	// Output: true
}

func ExampleErrorAs() {
	err := fmt.Errorf("wrapped: %w", exampleError{})
	var target exampleError
	printExample(a.ErrorAs(err, &target))
	// Output: true
}

func ExampleMatch() {
	printExample(a.Match(5, a.GreaterThan(3)))
	// Output: true
}

func ExampleMatchFunc() {
	matcher := a.MatchFunc(func(actual int) a.AssertionMessage {
		return a.Assert(actual%2 == 0, "got <%d>, wanted an even number", actual)
	})
	printExample(a.Match(4, matcher))
	// Output: true
}

func ExampleNot() {
	printExample(a.Match("hello", a.Not(a.Contains("bye"))))
	// Output: true
}

func ExampleAllOf() {
	printExample(a.Match(5, a.AllOf(
		a.GreaterThan(0),
		a.LessThan(10),
	)))
	// Output: true
}

func ExampleAnyOf() {
	printExample(a.Match(10, a.AnyOf(
		a.EqualTo(7),
		a.EqualTo(10),
	)))
	// Output: true
}

func ExampleDescribe() {
	result := a.Match(2, a.Describe("age check", a.GreaterThan(18)))
	fmt.Println(result.Message)
	// Output: age check: got <2>, wanted greater than <18>
}

func ExampleEqualTo() {
	printExample(a.Match("hello", a.EqualTo("hello")))
	// Output: true
}

func ExampleGreaterThan() {
	printExample(a.Match(5, a.GreaterThan(3)))
	// Output: true
}

func ExampleGreaterOrEqual() {
	printExample(a.Match(5, a.GreaterOrEqual(5)))
	// Output: true
}

func ExampleLessThan() {
	printExample(a.Match(3, a.LessThan(5)))
	// Output: true
}

func ExampleLessOrEqual() {
	printExample(a.Match(5, a.LessOrEqual(5)))
	// Output: true
}

func ExampleZero() {
	printExample(a.Match(0, a.Zero[int]()))
	// Output: true
}

func ExampleWithin() {
	printExample(a.Match(10.1, a.Within(10.0, 0.2)))
	// Output: true
}

func ExampleContains() {
	printExample(a.Match("hello world", a.Contains("world")))
	// Output: true
}

func ExampleHasPrefix() {
	printExample(a.Match("hello world", a.HasPrefix("hello")))
	// Output: true
}

func ExampleHasSuffix() {
	printExample(a.Match("hello world", a.HasSuffix("world")))
	// Output: true
}

func ExampleEmptyString() {
	printExample(a.Match("", a.EmptyString()))
	// Output: true
}

func ExampleNotEmptyString() {
	printExample(a.Match("hello", a.NotEmptyString()))
	// Output: true
}

func ExampleEqualIgnoringCase() {
	printExample(a.Match("HeLLo", a.EqualIgnoringCase("hello")))
	// Output: true
}

func ExampleHasPrefixIgnoringCase() {
	printExample(a.Match("Hello world", a.HasPrefixIgnoringCase("heL")))
	// Output: true
}

func ExampleHasSuffixIgnoringCase() {
	printExample(a.Match("Hello world", a.HasSuffixIgnoringCase("WOrLD")))
	// Output: true
}

func ExampleMatchesRegexp() {
	printExample(a.Match("hello-42", a.MatchesRegexp(`^hello-\d+$`)))
	// Output: true
}

func ExampleEqualIgnoringWhitespace() {
	printExample(a.Match(" hello\tworld \n", a.EqualIgnoringWhitespace("hello world")))
	// Output: true
}

func ExampleEqualNormalizedWhitespace() {
	printExample(a.Match(" hello\tworld \n", a.EqualNormalizedWhitespace("hello world")))
	// Output: true
}

func ExampleNum_EqualTo() {
	printExample(a.Number(42).EqualTo(42))
	// Output: true
}

func ExampleNum_NotEqual() {
	printExample(a.Number(42).NotEqual(7))
	// Output: true
}

func ExampleNum_GreaterThan() {
	printExample(a.Number(42).GreaterThan(7))
	// Output: true
}

func ExampleNum_GreaterOrEqual() {
	printExample(a.Number(42).GreaterOrEqual(42))
	// Output: true
}

func ExampleNum_IsZero() {
	printExample(a.Number(0).IsZero())
	// Output: true
}

func ExampleNum_LessThan() {
	printExample(a.Number(7).LessThan(42))
	// Output: true
}

func ExampleNum_LessOrEqual() {
	printExample(a.Number(7).LessOrEqual(7))
	// Output: true
}

func ExampleNum_Within() {
	printExample(a.Number(10.0).Within(10.1, 0.2))
	// Output: true
}

func ExampleStr_EqualTo() {
	printExample(a.String("hello").EqualTo("hello"))
	// Output: true
}

func ExampleStr_Contains() {
	printExample(a.String("hello world").Contains("world"))
	// Output: true
}

func ExampleStr_NotContains() {
	printExample(a.String("hello world").NotContains("bye"))
	// Output: true
}

func ExampleStr_HasPrefix() {
	printExample(a.String("hello world").HasPrefix("hello"))
	// Output: true
}

func ExampleStr_HasSuffix() {
	printExample(a.String("hello world").HasSuffix("world"))
	// Output: true
}

func ExampleStr_IsEmpty() {
	printExample(a.String("").IsEmpty())
	// Output: true
}

func ExampleStr_NotEmpty() {
	printExample(a.String("hello").NotEmpty())
	// Output: true
}

func ExampleStr_ToLowerEqualTo() {
	printExample(a.String("HeLLo").ToLowerEqualTo("hello"))
	// Output: true
}

func ExampleStr_MatchesRegexp() {
	printExample(a.String("hello-42").MatchesRegexp(`^hello-\d+$`))
	// Output: true
}

func ExampleStr_EqualIgnoringCase() {
	printExample(a.String("HeLLo").EqualIgnoringCase("hello"))
	// Output: true
}

func ExampleStr_HasPrefixIgnoringCase() {
	printExample(a.String("Hello world").HasPrefixIgnoringCase("heL"))
	// Output: true
}

func ExampleStr_HasSuffixIgnoringCase() {
	printExample(a.String("Hello world").HasSuffixIgnoringCase("WOrLD"))
	// Output: true
}

func ExampleStr_EqualIgnoringWhitespace() {
	printExample(a.String(" hello\tworld \n").EqualIgnoringWhitespace("hello world"))
	// Output: true
}

func ExampleStr_EqualNormalizedWhitespace() {
	printExample(a.String(" hello\tworld \n").EqualNormalizedWhitespace("hello world"))
	// Output: true
}

func ExampleSlc_Contains() {
	printExample(a.Slice([]int{1, 2, 3}).Contains(2, 3))
	// Output: true
}

func ExampleSlc_NotContains() {
	printExample(a.Slice([]int{1, 2, 3}).NotContains(4, 5))
	// Output: true
}

func ExampleSlc_EqualTo() {
	printExample(a.Slice([]int{1, 2, 3}).EqualTo(1, 2, 3))
	// Output: true
}

func ExampleSlc_Len() {
	printExample(a.Slice([]int{1, 2, 3}).Len(3))
	// Output: true
}

func ExampleSlc_IsEmpty() {
	printExample(a.Slice([]int{}).IsEmpty())
	// Output: true
}

func ExampleSlc_NotEmpty() {
	printExample(a.Slice([]int{1}).NotEmpty())
	// Output: true
}

func ExampleSlc_ContainsExactly() {
	printExample(a.Slice([]int{3, 2, 1}).ContainsExactly(1, 2, 3))
	// Output: true
}

func ExampleMappy_WithKeys() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).WithKeys("alpha", "beta"))
	// Output: true
}

func ExampleMappy_HasKey() {
	printExample(a.Map(map[string]int{"alpha": 1}).HasKey("alpha"))
	// Output: true
}

func ExampleMappy_NotHasKey() {
	printExample(a.Map(map[string]int{"alpha": 1}).NotHasKey("beta"))
	// Output: true
}

func ExampleMappy_KeysExactly() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).KeysExactly("alpha", "beta"))
	// Output: true
}

func ExampleMappy_IsEmpty() {
	printExample(a.Map(map[string]int{}).IsEmpty())
	// Output: true
}

func ExampleMappy_NotEmpty() {
	printExample(a.Map(map[string]int{"alpha": 1}).NotEmpty())
	// Output: true
}

func ExampleMappy_WithValues() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).WithValues(1, 2))
	// Output: true
}

func ExampleMappy_NotContains() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).NotContains(3, 4))
	// Output: true
}

func ExampleMappy_WithoutKeys() {
	printExample(a.Map(map[string]int{"alpha": 1}).WithoutKeys("beta"))
	// Output: true
}

func ExampleMappy_Len() {
	printExample(a.Map(map[string]int{"alpha": 1}).Len(1))
	// Output: true
}

func ExampleMappy_WithItem() {
	printExample(a.Map(map[string]int{"alpha": 1}).WithItem("alpha", 1))
	// Output: true
}

func ExampleMappy_EqualTo() {
	printExample(a.Map(map[string]int{"alpha": 1}).EqualTo(map[string]int{"alpha": 1}))
	// Output: true
}

func ExampleEvery() {
	printExample(a.Match([]int{2, 4, 6}, a.Every(a.GreaterThan(1))))
	// Output: true
}

func ExampleHasItem() {
	printExample(a.Match([]string{"alpha", "beta"}, a.HasItem(a.HasPrefix("bet"))))
	// Output: true
}

func ExampleContainsInOrder() {
	printExample(a.Match([]string{"alpha", "beta"}, a.ContainsInOrder(
		a.EqualTo("alpha"),
		a.HasSuffix("ta"),
	)))
	// Output: true
}

func ExampleContainsInAnyOrder() {
	printExample(a.Match([]int{2, 1, 2}, a.ContainsInAnyOrder(
		a.EqualTo(2),
		a.EqualTo(2),
		a.EqualTo(1),
	)))
	// Output: true
}

func ExampleHasEntry() {
	printExample(a.Match(map[string]int{"alpha": 1, "beta": 2}, a.HasEntry(
		a.EqualTo("beta"),
		a.GreaterThan(1),
	)))
	// Output: true
}

func ExampleHasKeyMatching() {
	printExample(a.Match(map[string]int{"alpha": 1}, a.HasKeyMatching[string, int](a.HasSuffix("pha"))))
	// Output: true
}

func ExampleHasValueMatching() {
	printExample(a.Match(map[string]int{"alpha": 3}, a.HasValueMatching[string, int](a.GreaterThan(2))))
	// Output: true
}

func ExampleHaving() {
	person := examplePerson{Name: "Ada", Age: 37}
	printExample(a.Match(person, a.Having(func(actual examplePerson) int {
		return actual.Age
	}, a.GreaterThan(30))))
	// Output: true
}

func ExampleHavingField() {
	person := examplePerson{Name: "Ada", Age: 37}
	printExample(a.Match(person, a.HavingField("Name", func(actual examplePerson) string {
		return actual.Name
	}, a.EqualTo("Ada"))))
	// Output: true
}

func ExampleSt_EqualTo() {
	printExample(a.Struct(examplePerson{Name: "Ada", Age: 37}).EqualTo(examplePerson{Name: "Ada", Age: 37}))
	// Output: true
}

func ExampleSlc_Every() {
	printExample(a.Slice([]int{2, 4, 6}).Every(a.GreaterThan(1)))
	// Output: true
}

func ExampleSlc_HasItem() {
	printExample(a.Slice([]string{"alpha", "beta"}).HasItem(a.HasPrefix("bet")))
	// Output: true
}

func ExampleSlc_ContainsInOrder() {
	printExample(a.Slice([]string{"alpha", "beta"}).ContainsInOrder(
		a.EqualTo("alpha"),
		a.HasSuffix("ta"),
	))
	// Output: true
}

func ExampleSlc_ContainsInAnyOrder() {
	printExample(a.Slice([]int{2, 1, 2}).ContainsInAnyOrder(
		a.EqualTo(2),
		a.EqualTo(2),
		a.EqualTo(1),
	))
	// Output: true
}

func ExampleMappy_HasEntry() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).HasEntry(
		a.EqualTo("beta"),
		a.GreaterThan(1),
	))
	// Output: true
}

func ExampleMappy_HasKeyMatching() {
	printExample(a.Map(map[string]int{"alpha": 1}).HasKeyMatching(a.HasSuffix("pha")))
	// Output: true
}

func ExampleMappy_HasValueMatching() {
	printExample(a.Map(map[string]int{"alpha": 3}).HasValueMatching(a.GreaterThan(2)))
	// Output: true
}
