package hammy_test

import (
	"errors"
	"fmt"
	"math"
	"regexp"

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

type exampleGreeter interface {
	Greet() string
}

type exampleGreeterImpl struct{}

func (exampleGreeterImpl) Greet() string {
	return "hello"
}

var pointerPattern = regexp.MustCompile(`0x[0-9a-fA-F]+`)

func printExample(result a.AssertionMessage) {
	message := pointerPattern.ReplaceAllString(result.Message, "0xPTR")
	fmt.Printf("message=%q\nsuccess=%t\n", message, result.IsSuccessful)
}

func ExampleNil() {
	var value any = nil
	printExample(a.Nil(value))
	// Output:
	// message="got <<nil>>, wanted nil"
	// success=true
}

func ExampleNotNil() {
	value := 42
	printExample(a.NotNil(&value))
	// Output:
	// message="got nil, wanted <*int>"
	// success=true
}

func ExampleTrue() {
	printExample(a.True(true))
	// Output:
	// message="got false, wanted true"
	// success=true
}

func ExampleFalse() {
	printExample(a.False(false))
	// Output:
	// message="got true, wanted false"
	// success=true
}

func ExampleError() {
	printExample(a.Error(errors.New("boom")))
	// Output:
	// message="got boom, want error"
	// success=true
}

func ExampleNilError() {
	printExample(a.NilError(nil))
	// Output:
	// message="got <<nil>>, want nil error"
	// success=true
}

func ExampleErrorIs() {
	target := errors.New("timeout")
	err := fmt.Errorf("request failed: %w", target)
	printExample(a.ErrorIs(err, target))
	// Output:
	// message="got <request failed: timeout>, want error matching <timeout>"
	// success=true
}

func ExampleErrorAs() {
	err := fmt.Errorf("wrapped: %w", exampleError{})
	var target exampleError
	printExample(a.ErrorAs(err, &target))
	// Output:
	// message="got <wrapped: example error>, want error assignable to <*hammy_test.exampleError>"
	// success=true
}

func ExampleMatch() {
	printExample(a.Match(5, a.GreaterThan(3)))
	// Output:
	// message="got <5>, wanted greater than <3>"
	// success=true
}

func ExampleMatchFunc() {
	matcher := a.MatchFunc(func(actual int) a.AssertionMessage {
		return a.Assert(actual%2 == 0, "got <%d>, wanted an even number", actual)
	})
	printExample(a.Match(4, matcher))
	// Output:
	// message="got <4>, wanted an even number"
	// success=true
}

func ExampleNot() {
	printExample(a.Match("hello", a.Not(a.Contains("bye"))))
	// Output:
	// message="not(got <hello>, wanted substring <bye>)"
	// success=true
}

func ExampleAllOf() {
	printExample(a.Match(5, a.AllOf(
		a.GreaterThan(0),
		a.LessThan(10),
	)))
	// Output:
	// message="matched all 2 matchers"
	// success=true
}

func ExampleAnyOf() {
	printExample(a.Match(10, a.AnyOf(
		a.EqualTo(7),
		a.EqualTo(10),
	)))
	// Output:
	// message="matched one of 2 matchers"
	// success=true
}

func ExampleDescribe() {
	printExample(a.Match(2, a.Describe("age check", a.GreaterThan(18))))
	// Output:
	// message="age check: got <2>, wanted greater than <18>"
	// success=false
}

func ExampleEqualTo() {
	printExample(a.Match("hello", a.EqualTo("hello")))
	// Output:
	// message="got <hello>, wanted equal to <hello>"
	// success=true
}

func ExampleGreaterThan() {
	printExample(a.Match(5, a.GreaterThan(3)))
	// Output:
	// message="got <5>, wanted greater than <3>"
	// success=true
}

func ExampleGreaterOrEqual() {
	printExample(a.Match(5, a.GreaterOrEqual(5)))
	// Output:
	// message="got <5>, wanted greater or equal to <5>"
	// success=true
}

func ExampleLessThan() {
	printExample(a.Match(3, a.LessThan(5)))
	// Output:
	// message="got <3>, wanted less than <5>"
	// success=true
}

func ExampleLessOrEqual() {
	printExample(a.Match(5, a.LessOrEqual(5)))
	// Output:
	// message="got <5>, wanted less or equal to <5>"
	// success=true
}

func ExampleZero() {
	printExample(a.Match(0, a.Zero[int]()))
	// Output:
	// message="got <0>, wanted equal to zero"
	// success=true
}

func ExampleWithin() {
	printExample(a.Match(10.1, a.Within(10.0, 0.2)))
	// Output:
	// message="got <10.1>, wanted within <0.2> of <10>"
	// success=true
}

func ExampleContains() {
	printExample(a.Match("hello world", a.Contains("world")))
	// Output:
	// message="got <hello world>, wanted substring <world>"
	// success=true
}

func ExampleHasPrefix() {
	printExample(a.Match("hello world", a.HasPrefix("hello")))
	// Output:
	// message="got <hello world>, wanted prefix <hello>"
	// success=true
}

func ExampleHasSuffix() {
	printExample(a.Match("hello world", a.HasSuffix("world")))
	// Output:
	// message="got <hello world>, wanted suffix <world>"
	// success=true
}

func ExampleEmptyString() {
	printExample(a.Match("", a.EmptyString()))
	// Output:
	// message="got <>, wanted an empty string"
	// success=true
}

func ExampleNotEmptyString() {
	printExample(a.Match("hello", a.NotEmptyString()))
	// Output:
	// message="got an empty string, wanted non-empty string"
	// success=true
}

func ExampleEqualIgnoringCase() {
	printExample(a.Match("HeLLo", a.EqualIgnoringCase("hello")))
	// Output:
	// message="got <HeLLo>, wanted equal to <hello> ignoring case"
	// success=true
}

func ExampleHasPrefixIgnoringCase() {
	printExample(a.Match("Hello world", a.HasPrefixIgnoringCase("heL")))
	// Output:
	// message="got <Hello world>, wanted prefix <heL> ignoring case"
	// success=true
}

func ExampleHasSuffixIgnoringCase() {
	printExample(a.Match("Hello world", a.HasSuffixIgnoringCase("WOrLD")))
	// Output:
	// message="got <Hello world>, wanted suffix <WOrLD> ignoring case"
	// success=true
}

func ExampleMatchesRegexp() {
	printExample(a.Match("hello-42", a.MatchesRegexp(`^hello-\d+$`)))
	// Output:
	// message="got <hello-42>, wanted regexp <^hello-\\d+$>"
	// success=true
}

func ExampleEqualIgnoringWhitespace() {
	printExample(a.Match(" hello\tworld \n", a.EqualIgnoringWhitespace("hello world")))
	// Output:
	// message="got < hello\tworld \n>, wanted equal to <hello world> ignoring whitespace"
	// success=true
}

func ExampleEqualNormalizedWhitespace() {
	printExample(a.Match(" hello\tworld \n", a.EqualNormalizedWhitespace("hello world")))
	// Output:
	// message="got < hello\tworld \n>, wanted equal to <hello world> ignoring whitespace"
	// success=true
}

func ExampleCloseTo() {
	printExample(a.Match(10.0, a.CloseTo(10.1, 0.2)))
	// Output:
	// message="got <10>, wanted within <0.2> of <10.1>"
	// success=true
}

func ExampleIsNaN() {
	printExample(a.Match(math.NaN(), a.IsNaN[float64]()))
	// Output:
	// message="got <NaN>, wanted NaN"
	// success=true
}

func ExampleIsInf() {
	printExample(a.Match(math.Inf(1), a.IsInf[float64]()))
	// Output:
	// message="got <+Inf>, wanted infinity"
	// success=true
}

func ExampleIsInfSign() {
	printExample(a.Match(math.Inf(-1), a.IsInfSign[float64](-1)))
	// Output:
	// message="got <-Inf>, wanted infinity with sign <-1>"
	// success=true
}

func ExampleSamePointer() {
	value := 42
	printExample(a.Match(&value, a.SamePointer(&value)))
	// Output:
	// message="got pointer <0xPTR>, wanted same pointer as <0xPTR>"
	// success=true
}

func ExampleTypeOf() {
	var value any = examplePerson{Name: "Ada"}
	printExample(a.Match(value, a.TypeOf[examplePerson]()))
	// Output:
	// message="got dynamic type <hammy_test.examplePerson>, wanted <hammy_test.examplePerson>"
	// success=true
}

func ExampleAssignableTo() {
	var value any = exampleGreeterImpl{}
	printExample(a.Match(value, a.AssignableTo[exampleGreeter]()))
	// Output:
	// message="got dynamic type <hammy_test.exampleGreeterImpl>, wanted assignable to <hammy_test.exampleGreeter>"
	// success=true
}

func ExampleNum_EqualTo() {
	printExample(a.Number(42).EqualTo(42))
	// Output:
	// message="want <42> equal to <42>"
	// success=true
}

func ExampleNum_NotEqual() {
	printExample(a.Number(42).NotEqual(7))
	// Output:
	// message="want <42> not equal to <7>"
	// success=true
}

func ExampleNum_GreaterThan() {
	printExample(a.Number(42).GreaterThan(7))
	// Output:
	// message="want <42> greater than <7>"
	// success=true
}

func ExampleNum_GreaterOrEqual() {
	printExample(a.Number(42).GreaterOrEqual(42))
	// Output:
	// message="want <42> greater or equal to <42>"
	// success=true
}

func ExampleNum_IsZero() {
	printExample(a.Number(0).IsZero())
	// Output:
	// message="want <0> equal to zero"
	// success=true
}

func ExampleNum_LessThan() {
	printExample(a.Number(7).LessThan(42))
	// Output:
	// message="want <7> less than <42>"
	// success=true
}

func ExampleNum_LessOrEqual() {
	printExample(a.Number(7).LessOrEqual(7))
	// Output:
	// message="want <7> less or equal to <7>"
	// success=true
}

func ExampleNum_Within() {
	printExample(a.Number(10.0).Within(10.1, 0.2))
	// Output:
	// message="want <10> greater or equal to <10.1>"
	// success=true
}

func ExampleStr_EqualTo() {
	printExample(a.String("hello").EqualTo("hello"))
	// Output:
	// message="got <hello>, wanted equal to <hello>"
	// success=true
}

func ExampleStr_Contains() {
	printExample(a.String("hello world").Contains("world"))
	// Output:
	// message="got <hello world>, wanted substring <world>"
	// success=true
}

func ExampleStr_NotContains() {
	printExample(a.String("hello world").NotContains("bye"))
	// Output:
	// message="got <hello world>, wanted no substring <bye>"
	// success=true
}

func ExampleStr_HasPrefix() {
	printExample(a.String("hello world").HasPrefix("hello"))
	// Output:
	// message="got <hello world>, wanted prefix <hello>"
	// success=true
}

func ExampleStr_HasSuffix() {
	printExample(a.String("hello world").HasSuffix("world"))
	// Output:
	// message="got <hello world>, wanted suffix <world>"
	// success=true
}

func ExampleStr_IsEmpty() {
	printExample(a.String("").IsEmpty())
	// Output:
	// message="got <>, wanted an empty string"
	// success=true
}

func ExampleStr_NotEmpty() {
	printExample(a.String("hello").NotEmpty())
	// Output:
	// message="got an empty string, wanted non-empty string"
	// success=true
}

func ExampleStr_ToLowerEqualTo() {
	printExample(a.String("HeLLo").ToLowerEqualTo("hello"))
	// Output:
	// message="got <hello>, wanted equal to <hello>"
	// success=true
}

func ExampleStr_MatchesRegexp() {
	printExample(a.String("hello-42").MatchesRegexp(`^hello-\d+$`))
	// Output:
	// message="got <hello-42>, wanted regexp <^hello-\\d+$>"
	// success=true
}

func ExampleStr_EqualIgnoringCase() {
	printExample(a.String("HeLLo").EqualIgnoringCase("hello"))
	// Output:
	// message="got <HeLLo>, wanted equal to <hello> ignoring case"
	// success=true
}

func ExampleStr_HasPrefixIgnoringCase() {
	printExample(a.String("Hello world").HasPrefixIgnoringCase("heL"))
	// Output:
	// message="got <Hello world>, wanted prefix <heL> ignoring case"
	// success=true
}

func ExampleStr_HasSuffixIgnoringCase() {
	printExample(a.String("Hello world").HasSuffixIgnoringCase("WOrLD"))
	// Output:
	// message="got <Hello world>, wanted suffix <WOrLD> ignoring case"
	// success=true
}

func ExampleStr_EqualIgnoringWhitespace() {
	printExample(a.String(" hello\tworld \n").EqualIgnoringWhitespace("hello world"))
	// Output:
	// message="got < hello\tworld \n>, wanted equal to <hello world> ignoring whitespace"
	// success=true
}

func ExampleStr_EqualNormalizedWhitespace() {
	printExample(a.String(" hello\tworld \n").EqualNormalizedWhitespace("hello world"))
	// Output:
	// message="got < hello\tworld \n>, wanted equal to <hello world> ignoring whitespace"
	// success=true
}

func ExampleSlc_Contains() {
	printExample(a.Slice([]int{1, 2, 3}).Contains(2, 3))
	// Output:
	// message="got 0 unmatched items, wanted array containing the 2 items. Items at index  were missing"
	// success=true
}

func ExampleSlc_NotContains() {
	printExample(a.Slice([]int{1, 2, 3}).NotContains(4, 5))
	// Output:
	// message="got items at expected index  present in slice, wanted all absent"
	// success=true
}

func ExampleSlc_EqualTo() {
	printExample(a.Slice([]int{1, 2, 3}).EqualTo(1, 2, 3))
	// Output:
	// message="slice mismatch (-want +got):\\n"
	// success=true
}

func ExampleSlc_Len() {
	printExample(a.Slice([]int{1, 2, 3}).Len(3))
	// Output:
	// message="got len()=3, wanted 3"
	// success=true
}

func ExampleSlc_IsEmpty() {
	printExample(a.Slice([]int{}).IsEmpty())
	// Output:
	// message="got len()=0, wanted 0"
	// success=true
}

func ExampleSlc_NotEmpty() {
	printExample(a.Slice([]int{1}).NotEmpty())
	// Output:
	// message="got len()=1, wanted > 0"
	// success=true
}

func ExampleSlc_ContainsExactly() {
	printExample(a.Slice([]int{3, 2, 1}).ContainsExactly(1, 2, 3))
	// Output:
	// message="got 0 unmatched items, wanted array containing the 3 items. Items at index  were missing"
	// success=true
}

func ExampleMappy_WithKeys() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).WithKeys("alpha", "beta"))
	// Output:
	// message="got <[alpha beta]>, wanted keys <[]>"
	// success=true
}

func ExampleMappy_HasKey() {
	printExample(a.Map(map[string]int{"alpha": 1}).HasKey("alpha"))
	// Output:
	// message="got key absent <alpha>, wanted present in map"
	// success=true
}

func ExampleMappy_NotHasKey() {
	printExample(a.Map(map[string]int{"alpha": 1}).NotHasKey("beta"))
	// Output:
	// message="got key present <beta>, wanted absent from map"
	// success=true
}

func ExampleMappy_KeysExactly() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).KeysExactly("alpha", "beta"))
	// Output:
	// message="got extra keys <[]> and missing keys <[]>, wanted exact key set"
	// success=true
}

func ExampleMappy_IsEmpty() {
	printExample(a.Map(map[string]int{}).IsEmpty())
	// Output:
	// message="got len=<0>, wanted empty map"
	// success=true
}

func ExampleMappy_NotEmpty() {
	printExample(a.Map(map[string]int{"alpha": 1}).NotEmpty())
	// Output:
	// message="got len=<1>, wanted non-empty map"
	// success=true
}

func ExampleMappy_WithValues() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).WithValues(1, 2))
	// Output:
	// message="got <[1 2]>, wanted values <[]>"
	// success=true
}

func ExampleMappy_NotContains() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).NotContains(3, 4))
	// Output:
	// message="got values <[]>, wanted absent from map"
	// success=true
}

func ExampleMappy_WithoutKeys() {
	printExample(a.Map(map[string]int{"alpha": 1}).WithoutKeys("beta"))
	// Output:
	// message="got keys <[]>, wanted absent from map"
	// success=true
}

func ExampleMappy_Len() {
	printExample(a.Map(map[string]int{"alpha": 1}).Len(1))
	// Output:
	// message="got len=<1>, wanted <1>"
	// success=true
}

func ExampleMappy_WithItem() {
	printExample(a.Map(map[string]int{"alpha": 1}).WithItem("alpha", 1))
	// Output:
	// message="got value=<1> for key=<hi>, wanted <1>"
	// success=true
}

func ExampleMappy_EqualTo() {
	printExample(a.Map(map[string]int{"alpha": 1}).EqualTo(map[string]int{"alpha": 1}))
	// Output:
	// message="Map mismatch (-want +got):\n"
	// success=true
}

func ExampleEvery() {
	printExample(a.Match([]int{2, 4, 6}, a.Every(a.GreaterThan(1))))
	// Output:
	// message="all 3 items matched"
	// success=true
}

func ExampleHasItem() {
	printExample(a.Match([]string{"alpha", "beta"}, a.HasItem(a.HasPrefix("bet"))))
	// Output:
	// message="found matching item at index 1"
	// success=true
}

func ExampleContainsInOrder() {
	printExample(a.Match([]string{"alpha", "beta"}, a.ContainsInOrder(
		a.EqualTo("alpha"),
		a.HasSuffix("ta"),
	)))
	// Output:
	// message="all 2 items matched in order"
	// success=true
}

func ExampleContainsInAnyOrder() {
	printExample(a.Match([]int{2, 1, 2}, a.ContainsInAnyOrder(
		a.EqualTo(2),
		a.EqualTo(2),
		a.EqualTo(1),
	)))
	// Output:
	// message="all 3 items matched in any order"
	// success=true
}

func ExampleHasEntry() {
	printExample(a.Match(map[string]int{"alpha": 1, "beta": 2}, a.HasEntry(
		a.EqualTo("beta"),
		a.GreaterThan(1),
	)))
	// Output:
	// message="found matching entry for key <beta>"
	// success=true
}

func ExampleHasKeyMatching() {
	printExample(a.Match(map[string]int{"alpha": 1}, a.HasKeyMatching[string, int](a.HasSuffix("pha"))))
	// Output:
	// message="found matching key <alpha>"
	// success=true
}

func ExampleHasValueMatching() {
	printExample(a.Match(map[string]int{"alpha": 3}, a.HasValueMatching[string, int](a.GreaterThan(2))))
	// Output:
	// message="found matching value for key <alpha>"
	// success=true
}

func ExampleHaving() {
	person := examplePerson{Name: "Ada", Age: 37}
	printExample(a.Match(person, a.Having(func(actual examplePerson) int {
		return actual.Age
	}, a.GreaterThan(30))))
	// Output:
	// message="got <37>, wanted greater than <30>"
	// success=true
}

func ExampleHavingField() {
	person := examplePerson{Name: "Ada", Age: 37}
	printExample(a.Match(person, a.HavingField("Name", func(actual examplePerson) string {
		return actual.Name
	}, a.EqualTo("Ada"))))
	// Output:
	// message="got <Ada>, wanted equal to <Ada>"
	// success=true
}

func ExampleSt_EqualTo() {
	printExample(a.Struct(examplePerson{Name: "Ada", Age: 37}).EqualTo(examplePerson{Name: "Ada", Age: 37}))
	// Output:
	// message="Structs are not equal (+got -want):\n"
	// success=true
}

func ExampleSlc_Every() {
	printExample(a.Slice([]int{2, 4, 6}).Every(a.GreaterThan(1)))
	// Output:
	// message="all 3 items matched"
	// success=true
}

func ExampleSlc_HasItem() {
	printExample(a.Slice([]string{"alpha", "beta"}).HasItem(a.HasPrefix("bet")))
	// Output:
	// message="found matching item at index 1"
	// success=true
}

func ExampleSlc_ContainsInOrder() {
	printExample(a.Slice([]string{"alpha", "beta"}).ContainsInOrder(
		a.EqualTo("alpha"),
		a.HasSuffix("ta"),
	))
	// Output:
	// message="all 2 items matched in order"
	// success=true
}

func ExampleSlc_ContainsInAnyOrder() {
	printExample(a.Slice([]int{2, 1, 2}).ContainsInAnyOrder(
		a.EqualTo(2),
		a.EqualTo(2),
		a.EqualTo(1),
	))
	// Output:
	// message="all 3 items matched in any order"
	// success=true
}

func ExampleMappy_HasEntry() {
	printExample(a.Map(map[string]int{"alpha": 1, "beta": 2}).HasEntry(
		a.EqualTo("beta"),
		a.GreaterThan(1),
	))
	// Output:
	// message="found matching entry for key <beta>"
	// success=true
}

func ExampleMappy_HasKeyMatching() {
	printExample(a.Map(map[string]int{"alpha": 1}).HasKeyMatching(a.HasSuffix("pha")))
	// Output:
	// message="found matching key <alpha>"
	// success=true
}

func ExampleMappy_HasValueMatching() {
	printExample(a.Map(map[string]int{"alpha": 3}).HasValueMatching(a.GreaterThan(2)))
	// Output:
	// message="found matching value for key <alpha>"
	// success=true
}

func ExampleFlt_CloseTo() {
	printExample(a.Float(10.0).CloseTo(10.1, 0.2))
	// Output:
	// message="got <10>, wanted within <0.2> of <10.1>"
	// success=true
}

func ExampleFlt_IsNaN() {
	printExample(a.Float(math.NaN()).IsNaN())
	// Output:
	// message="got <NaN>, wanted NaN"
	// success=true
}

func ExampleFlt_IsInf() {
	printExample(a.Float(math.Inf(1)).IsInf())
	// Output:
	// message="got <+Inf>, wanted infinity"
	// success=true
}

func ExampleFlt_IsInfSign() {
	printExample(a.Float(math.Inf(-1)).IsInfSign(-1))
	// Output:
	// message="got <-Inf>, wanted infinity with sign <-1>"
	// success=true
}
