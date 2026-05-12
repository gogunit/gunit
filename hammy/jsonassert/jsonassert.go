package jsonassert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/gogunit/gunit/hammy"
	"github.com/google/go-cmp/cmp"
)

func Equal(actual, expected string) hammy.AssertionMessage {
	return EqualBytes([]byte(actual), []byte(expected))
}

func EqualWithOptions(actual, expected string, opts ...Option) hammy.AssertionMessage {
	return EqualBytesWithOptions([]byte(actual), []byte(expected), opts...)
}

func EqualReader(actual, expected io.Reader) hammy.AssertionMessage {
	actualBytes, result := readJSON("actual", actual)
	if !result.IsSuccessful {
		return result
	}

	expectedBytes, result := readJSON("expected", expected)
	if !result.IsSuccessful {
		return result
	}

	return EqualBytes(actualBytes, expectedBytes)
}

func EqualBytes(actual, expected []byte) hammy.AssertionMessage {
	return EqualBytesWithOptions(actual, expected)
}

func EqualBytesWithOptions(actual, expected []byte, opts ...Option) hammy.AssertionMessage {
	actualJSON, err := parseJSON(actual)
	if err != nil {
		return hammy.Assert(false, "actual JSON invalid: %v", err)
	}

	expectedJSON, err := parseJSON(expected)
	if err != nil {
		return hammy.Assert(false, "expected JSON invalid: %v", err)
	}

	actualJSON, expectedJSON, err = applyOptions(actualJSON, expectedJSON, opts...)
	if err != nil {
		return hammy.Assert(false, "%v", err)
	}

	diff := cmp.Diff(expectedJSON, actualJSON)
	return hammy.Assert(diff == "", "JSON mismatch (-want +got):\n%s", diff)
}

type Option func(*compareOptions)

func IgnorePaths(paths ...string) Option {
	return func(options *compareOptions) {
		options.ignorePaths = append(options.ignorePaths, paths...)
	}
}

func UnorderedArraysAt(paths ...string) Option {
	return func(options *compareOptions) {
		options.unorderedArrayPaths = append(options.unorderedArrayPaths, paths...)
	}
}

type compareOptions struct {
	ignorePaths         []string
	unorderedArrayPaths []string
}

func Valid(actual string) hammy.AssertionMessage {
	return ValidBytes([]byte(actual))
}

func ValidReader(actual io.Reader) hammy.AssertionMessage {
	actualBytes, result := readJSON("actual", actual)
	if !result.IsSuccessful {
		return result
	}
	return ValidBytes(actualBytes)
}

func ValidBytes(actual []byte) hammy.AssertionMessage {
	if _, err := parseJSON(actual); err != nil {
		return hammy.Assert(false, "JSON invalid: %v", err)
	}
	return hammy.Assert(true, "got valid JSON")
}

func Contains(actual, expected string) hammy.AssertionMessage {
	return ContainsBytes([]byte(actual), []byte(expected))
}

func ContainsBytes(actual, expected []byte) hammy.AssertionMessage {
	actualJSON, err := parseJSON(actual)
	if err != nil {
		return hammy.Assert(false, "actual JSON invalid: %v", err)
	}

	expectedJSON, err := parseJSON(expected)
	if err != nil {
		return hammy.Assert(false, "expected JSON invalid: %v", err)
	}

	if ok, message := containsJSON(actualJSON, expectedJSON, "$"); !ok {
		return hammy.Assert(false, "%s", message)
	}
	return hammy.Assert(true, "JSON contained expected subset")
}

func PathExists(actual, path string) hammy.AssertionMessage {
	actualJSON, err := parseJSON([]byte(actual))
	if err != nil {
		return hammy.Assert(false, "actual JSON invalid: %v", err)
	}

	if _, found, err := lookupJSONPath(actualJSON, path); err != nil {
		return hammy.Assert(false, "invalid JSON path <%s>: %v", path, err)
	} else if !found {
		return hammy.Assert(false, "JSON path <%s> missing", path)
	}
	return hammy.Assert(true, "JSON path <%s> exists", path)
}

func PathMissing(actual, path string) hammy.AssertionMessage {
	actualJSON, err := parseJSON([]byte(actual))
	if err != nil {
		return hammy.Assert(false, "actual JSON invalid: %v", err)
	}

	if _, found, err := lookupJSONPath(actualJSON, path); err != nil {
		return hammy.Assert(false, "invalid JSON path <%s>: %v", path, err)
	} else if found {
		return hammy.Assert(false, "JSON path <%s> exists, wanted missing", path)
	}
	return hammy.Assert(true, "JSON path <%s> missing", path)
}

func PathEqual(actual, path, expected string) hammy.AssertionMessage {
	return PathEqualBytes([]byte(actual), path, []byte(expected))
}

func PathEqualBytes(actual []byte, path string, expected []byte) hammy.AssertionMessage {
	actualJSON, err := parseJSON(actual)
	if err != nil {
		return hammy.Assert(false, "actual JSON invalid: %v", err)
	}

	expectedJSON, err := parseJSON(expected)
	if err != nil {
		return hammy.Assert(false, "expected JSON invalid: %v", err)
	}

	actualValue, found, err := lookupJSONPath(actualJSON, path)
	if err != nil {
		return hammy.Assert(false, "invalid JSON path <%s>: %v", path, err)
	}
	if !found {
		return hammy.Assert(false, "JSON path <%s> missing", path)
	}

	diff := cmp.Diff(expectedJSON, actualValue)
	return hammy.Assert(diff == "", "JSON path <%s> mismatch (-want +got):\n%s", path, diff)
}

func ArrayContains(actual, path, expectedElement string) hammy.AssertionMessage {
	return ArrayContainsBytes([]byte(actual), path, []byte(expectedElement))
}

func ArrayContainsBytes(actual []byte, path string, expectedElement []byte) hammy.AssertionMessage {
	actualJSON, err := parseJSON(actual)
	if err != nil {
		return hammy.Assert(false, "actual JSON invalid: %v", err)
	}

	expectedJSON, err := parseJSON(expectedElement)
	if err != nil {
		return hammy.Assert(false, "expected JSON invalid: %v", err)
	}

	actualValue, found, err := lookupJSONPath(actualJSON, path)
	if err != nil {
		return hammy.Assert(false, "invalid JSON path <%s>: %v", path, err)
	}
	if !found {
		return hammy.Assert(false, "JSON path <%s> missing", path)
	}

	actualArray, ok := actualValue.([]any)
	if !ok {
		return hammy.Assert(false, "got JSON path <%s> type <%T>, wanted array", path, actualValue)
	}
	for i, item := range actualArray {
		if cmp.Equal(item, expectedJSON) {
			return hammy.Assert(true, "found matching element at JSON path <%s> index <%d>", path, i)
		}
	}
	return hammy.Assert(false, "got no matching element at JSON path <%s>", path)
}

func readJSON(name string, reader io.Reader) ([]byte, hammy.AssertionMessage) {
	if reader == nil {
		return nil, hammy.Assert(false, "%s JSON reader is nil", name)
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, hammy.Assert(false, "%s JSON read error: %v", name, err)
	}
	return data, hammy.Assert(true, "%s JSON read", name)
}

func applyOptions(actual, expected any, opts ...Option) (any, any, error) {
	var options compareOptions
	for _, opt := range opts {
		opt(&options)
	}

	for _, path := range options.ignorePaths {
		steps, err := parsePath(path)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid JSON path <%s>: %w", path, err)
		}
		actual = deleteJSONPath(actual, steps)
		expected = deleteJSONPath(expected, steps)
	}

	for _, path := range options.unorderedArrayPaths {
		steps, err := parsePath(path)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid JSON path <%s>: %w", path, err)
		}
		if err := sortJSONArrayAtPath(actual, steps, path); err != nil {
			return nil, nil, err
		}
		if err := sortJSONArrayAtPath(expected, steps, path); err != nil {
			return nil, nil, err
		}
	}

	return actual, expected, nil
}

func deleteJSONPath(value any, steps []pathStep) any {
	if len(steps) == 0 {
		return nil
	}

	parent, found := lookupParsedJSONPath(value, steps[:len(steps)-1])
	if !found {
		return value
	}

	last := steps[len(steps)-1]
	if last.isIndex {
		items, ok := parent.([]any)
		if ok && last.index >= 0 && last.index < len(items) {
			items[last.index] = nil
		}
		return value
	}

	fields, ok := parent.(map[string]any)
	if ok {
		delete(fields, last.key)
	}
	return value
}

func sortJSONArrayAtPath(value any, steps []pathStep, path string) error {
	target, found := lookupParsedJSONPath(value, steps)
	if !found {
		return nil
	}

	items, ok := target.([]any)
	if !ok {
		return fmt.Errorf("got JSON path <%s> type <%T>, wanted array", path, target)
	}
	sort.SliceStable(items, func(i, j int) bool {
		return canonicalJSONKey(items[i]) < canonicalJSONKey(items[j])
	})
	return nil
}

func canonicalJSONKey(value any) string {
	switch typed := value.(type) {
	case nil:
		return "null"
	case bool:
		if typed {
			return "bool:true"
		}
		return "bool:false"
	case string:
		return "string:" + strconv.Quote(typed)
	case normalizedNumber:
		return "number:" + typed.Value
	case []any:
		parts := make([]string, 0, len(typed))
		for _, item := range typed {
			parts = append(parts, canonicalJSONKey(item))
		}
		return "array:[" + strings.Join(parts, ",") + "]"
	case map[string]any:
		keys := make([]string, 0, len(typed))
		for key := range typed {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		parts := make([]string, 0, len(keys))
		for _, key := range keys {
			parts = append(parts, strconv.Quote(key)+":"+canonicalJSONKey(typed[key]))
		}
		return "object:{" + strings.Join(parts, ",") + "}"
	default:
		return fmt.Sprintf("%T:%v", typed, typed)
	}
}

func containsJSON(actual, expected any, path string) (bool, string) {
	switch expectedValue := expected.(type) {
	case map[string]any:
		actualValue, ok := actual.(map[string]any)
		if !ok {
			return false, fmt.Sprintf("got JSON path <%s> type <%T>, wanted object", path, actual)
		}

		keys := make([]string, 0, len(expectedValue))
		for key := range expectedValue {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			actualItem, ok := actualValue[key]
			itemPath := joinPath(path, key)
			if !ok {
				return false, fmt.Sprintf("JSON path <%s> missing", itemPath)
			}
			if ok, message := containsJSON(actualItem, expectedValue[key], itemPath); !ok {
				return false, message
			}
		}
		return true, ""
	case []any:
		if diff := cmp.Diff(expected, actual); diff != "" {
			return false, fmt.Sprintf("JSON path <%s> mismatch (-want +got):\n%s", path, diff)
		}
		return true, ""
	default:
		if !cmp.Equal(actual, expected) {
			return false, fmt.Sprintf("JSON path <%s> mismatch (-want +got):\n%s", path, cmp.Diff(expected, actual))
		}
		return true, ""
	}
}

func joinPath(path, key string) string {
	if path == "$" {
		return "$." + key
	}
	return path + "." + key
}

type pathStep struct {
	key     string
	index   int
	isIndex bool
}

func lookupJSONPath(value any, path string) (any, bool, error) {
	steps, err := parsePath(path)
	if err != nil {
		return nil, false, err
	}
	value, found := lookupParsedJSONPath(value, steps)
	return value, found, nil
}

func lookupParsedJSONPath(value any, steps []pathStep) (any, bool) {
	current := value
	for _, step := range steps {
		if step.isIndex {
			items, ok := current.([]any)
			if !ok || step.index < 0 || step.index >= len(items) {
				return nil, false
			}
			current = items[step.index]
			continue
		}

		fields, ok := current.(map[string]any)
		if !ok {
			return nil, false
		}
		next, ok := fields[step.key]
		if !ok {
			return nil, false
		}
		current = next
	}
	return current, true
}

func parsePath(path string) ([]pathStep, error) {
	if path == "" {
		return nil, fmt.Errorf("path is empty")
	}
	if path == "$" {
		return nil, nil
	}
	path = strings.TrimPrefix(path, "$.")

	var steps []pathStep
	for i := 0; i < len(path); {
		switch path[i] {
		case '.':
			return nil, fmt.Errorf("unexpected dot at offset %d", i)
		case '[':
			step, next, err := parseIndexStep(path, i)
			if err != nil {
				return nil, err
			}
			steps = append(steps, step)
			i = next
		default:
			start := i
			for i < len(path) && path[i] != '.' && path[i] != '[' && path[i] != ']' {
				i++
			}
			if start == i {
				return nil, fmt.Errorf("empty field at offset %d", start)
			}
			if i < len(path) && path[i] == ']' {
				return nil, fmt.Errorf("unexpected closing bracket at offset %d", i)
			}
			steps = append(steps, pathStep{key: path[start:i]})
		}

		for i < len(path) && path[i] == '[' {
			step, next, err := parseIndexStep(path, i)
			if err != nil {
				return nil, err
			}
			steps = append(steps, step)
			i = next
		}

		if i < len(path) {
			if path[i] != '.' {
				return nil, fmt.Errorf("unexpected character %q at offset %d", path[i], i)
			}
			i++
			if i == len(path) {
				return nil, fmt.Errorf("path ends with dot")
			}
		}
	}
	return steps, nil
}

func parseIndexStep(path string, start int) (pathStep, int, error) {
	end := strings.IndexByte(path[start:], ']')
	if end < 0 {
		return pathStep{}, 0, fmt.Errorf("unclosed index at offset %d", start)
	}
	end += start
	rawIndex := path[start+1 : end]
	if rawIndex == "" {
		return pathStep{}, 0, fmt.Errorf("empty index at offset %d", start)
	}
	index, err := strconv.Atoi(rawIndex)
	if err != nil || index < 0 {
		return pathStep{}, 0, fmt.Errorf("invalid index <%s>", rawIndex)
	}
	return pathStep{index: index, isIndex: true}, end + 1, nil
}

func parseJSON(data []byte) (any, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	var value any
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}

	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		if err == nil {
			return nil, fmt.Errorf("multiple JSON values")
		}
		return nil, err
	}

	return normalizeJSON(value)
}

func normalizeJSON(value any) (any, error) {
	switch typed := value.(type) {
	case map[string]any:
		normalized := make(map[string]any, len(typed))
		for key, item := range typed {
			normalizedItem, err := normalizeJSON(item)
			if err != nil {
				return nil, err
			}
			normalized[key] = normalizedItem
		}
		return normalized, nil
	case []any:
		normalized := make([]any, len(typed))
		for i, item := range typed {
			normalizedItem, err := normalizeJSON(item)
			if err != nil {
				return nil, err
			}
			normalized[i] = normalizedItem
		}
		return normalized, nil
	case json.Number:
		return normalizeNumber(typed)
	default:
		return typed, nil
	}
}

type normalizedNumber struct {
	Value string
}

func normalizeNumber(number json.Number) (normalizedNumber, error) {
	rat, err := parseJSONNumber(number.String())
	if err != nil {
		return normalizedNumber{}, err
	}
	return normalizedNumber{Value: rat.RatString()}, nil
}

func parseJSONNumber(raw string) (*big.Rat, error) {
	mantissa := raw
	exponent := 0
	if exponentIndex := strings.IndexAny(raw, "eE"); exponentIndex >= 0 {
		parsedExponent, err := strconv.Atoi(raw[exponentIndex+1:])
		if err != nil {
			return nil, fmt.Errorf("unsupported JSON number exponent %q: %w", raw, err)
		}
		exponent = parsedExponent
		mantissa = raw[:exponentIndex]
	}

	negative := strings.HasPrefix(mantissa, "-")
	mantissa = strings.TrimPrefix(mantissa, "-")

	fractionDigits := 0
	if dotIndex := strings.IndexByte(mantissa, '.'); dotIndex >= 0 {
		fractionDigits = len(mantissa) - dotIndex - 1
		mantissa = mantissa[:dotIndex] + mantissa[dotIndex+1:]
	}

	mantissa = strings.TrimLeft(mantissa, "0")
	if mantissa == "" {
		return new(big.Rat), nil
	}

	numerator := new(big.Int)
	if _, ok := numerator.SetString(mantissa, 10); !ok {
		return nil, fmt.Errorf("invalid JSON number %q", raw)
	}
	if negative {
		numerator.Neg(numerator)
	}

	scale := fractionDigits - exponent
	if scale <= 0 {
		numerator.Mul(numerator, pow10(-scale))
		return new(big.Rat).SetInt(numerator), nil
	}

	return new(big.Rat).SetFrac(numerator, pow10(scale)), nil
}

func pow10(exponent int) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exponent)), nil)
}
