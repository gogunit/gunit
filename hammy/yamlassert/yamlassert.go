package yamlassert

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/gogunit/gunit/hammy"
	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v3"
)

func Equal(actual, expected string) hammy.AssertionMessage {
	return EqualBytes([]byte(actual), []byte(expected))
}

func EqualWithOptions(actual, expected string, opts ...Option) hammy.AssertionMessage {
	return EqualBytesWithOptions([]byte(actual), []byte(expected), opts...)
}

func EqualReader(actual, expected io.Reader) hammy.AssertionMessage {
	actualBytes, result := readYAML("actual", actual)
	if !result.IsSuccessful {
		return result
	}

	expectedBytes, result := readYAML("expected", expected)
	if !result.IsSuccessful {
		return result
	}

	return EqualBytes(actualBytes, expectedBytes)
}

func EqualBytes(actual, expected []byte) hammy.AssertionMessage {
	return EqualBytesWithOptions(actual, expected)
}

func EqualBytesWithOptions(actual, expected []byte, opts ...Option) hammy.AssertionMessage {
	actualYAML, err := parseYAML(actual)
	if err != nil {
		return hammy.Assert(false, "actual YAML invalid: %v", err)
	}

	expectedYAML, err := parseYAML(expected)
	if err != nil {
		return hammy.Assert(false, "expected YAML invalid: %v", err)
	}

	actualYAML, expectedYAML, err = applyOptions(actualYAML, expectedYAML, opts...)
	if err != nil {
		return hammy.Assert(false, "%v", err)
	}

	diff := cmp.Diff(expectedYAML, actualYAML)
	return hammy.Assert(diff == "", "YAML mismatch (-want +got):\n%s", diff)
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
	actualBytes, result := readYAML("actual", actual)
	if !result.IsSuccessful {
		return result
	}
	return ValidBytes(actualBytes)
}

func ValidBytes(actual []byte) hammy.AssertionMessage {
	if _, err := parseYAMLDocuments(actual); err != nil {
		return hammy.Assert(false, "YAML invalid: %v", err)
	}
	return hammy.Assert(true, "got valid YAML")
}

func Contains(actual, expected string) hammy.AssertionMessage {
	return ContainsBytes([]byte(actual), []byte(expected))
}

func ContainsBytes(actual, expected []byte) hammy.AssertionMessage {
	actualYAML, err := parseYAML(actual)
	if err != nil {
		return hammy.Assert(false, "actual YAML invalid: %v", err)
	}

	expectedYAML, err := parseYAML(expected)
	if err != nil {
		return hammy.Assert(false, "expected YAML invalid: %v", err)
	}

	if ok, message := containsYAML(actualYAML, expectedYAML, "$"); !ok {
		return hammy.Assert(false, "%s", message)
	}
	return hammy.Assert(true, "YAML contained expected subset")
}

func PathExists(actual, path string) hammy.AssertionMessage {
	actualYAML, err := parseYAML([]byte(actual))
	if err != nil {
		return hammy.Assert(false, "actual YAML invalid: %v", err)
	}

	if _, found, err := lookupYAMLPath(actualYAML, path); err != nil {
		return hammy.Assert(false, "invalid YAML path <%s>: %v", path, err)
	} else if !found {
		return hammy.Assert(false, "YAML path <%s> missing", path)
	}
	return hammy.Assert(true, "YAML path <%s> exists", path)
}

func PathMissing(actual, path string) hammy.AssertionMessage {
	actualYAML, err := parseYAML([]byte(actual))
	if err != nil {
		return hammy.Assert(false, "actual YAML invalid: %v", err)
	}

	if _, found, err := lookupYAMLPath(actualYAML, path); err != nil {
		return hammy.Assert(false, "invalid YAML path <%s>: %v", path, err)
	} else if found {
		return hammy.Assert(false, "YAML path <%s> exists, wanted missing", path)
	}
	return hammy.Assert(true, "YAML path <%s> missing", path)
}

func PathEqual(actual, path, expected string) hammy.AssertionMessage {
	return PathEqualBytes([]byte(actual), path, []byte(expected))
}

func PathEqualBytes(actual []byte, path string, expected []byte) hammy.AssertionMessage {
	actualYAML, err := parseYAML(actual)
	if err != nil {
		return hammy.Assert(false, "actual YAML invalid: %v", err)
	}

	expectedYAML, err := parseYAML(expected)
	if err != nil {
		return hammy.Assert(false, "expected YAML invalid: %v", err)
	}

	actualValue, found, err := lookupYAMLPath(actualYAML, path)
	if err != nil {
		return hammy.Assert(false, "invalid YAML path <%s>: %v", path, err)
	}
	if !found {
		return hammy.Assert(false, "YAML path <%s> missing", path)
	}

	diff := cmp.Diff(expectedYAML, actualValue)
	return hammy.Assert(diff == "", "YAML path <%s> mismatch (-want +got):\n%s", path, diff)
}

func ArrayContains(actual, path, expectedElement string) hammy.AssertionMessage {
	return ArrayContainsBytes([]byte(actual), path, []byte(expectedElement))
}

func ArrayContainsBytes(actual []byte, path string, expectedElement []byte) hammy.AssertionMessage {
	actualYAML, err := parseYAML(actual)
	if err != nil {
		return hammy.Assert(false, "actual YAML invalid: %v", err)
	}

	expectedYAML, err := parseYAML(expectedElement)
	if err != nil {
		return hammy.Assert(false, "expected YAML invalid: %v", err)
	}

	actualValue, found, err := lookupYAMLPath(actualYAML, path)
	if err != nil {
		return hammy.Assert(false, "invalid YAML path <%s>: %v", path, err)
	}
	if !found {
		return hammy.Assert(false, "YAML path <%s> missing", path)
	}

	actualArray, ok := actualValue.([]any)
	if !ok {
		return hammy.Assert(false, "got YAML path <%s> type <%T>, wanted array", path, actualValue)
	}
	for i, item := range actualArray {
		if cmp.Equal(item, expectedYAML) {
			return hammy.Assert(true, "found matching element at YAML path <%s> index <%d>", path, i)
		}
	}
	return hammy.Assert(false, "got no matching element at YAML path <%s>", path)
}

func DocumentCount(actual string, expected int) hammy.AssertionMessage {
	return DocumentCountBytes([]byte(actual), expected)
}

func DocumentCountBytes(actual []byte, expected int) hammy.AssertionMessage {
	documents, err := parseYAMLDocuments(actual)
	if err != nil {
		return hammy.Assert(false, "YAML invalid: %v", err)
	}
	return hammy.Assert(len(documents) == expected, "got YAML document count <%d>, wanted <%d>", len(documents), expected)
}

func DocumentEqual(actual string, index int, expected string, opts ...Option) hammy.AssertionMessage {
	return DocumentEqualBytes([]byte(actual), index, []byte(expected), opts...)
}

func DocumentEqualBytes(actual []byte, index int, expected []byte, opts ...Option) hammy.AssertionMessage {
	documents, err := parseYAMLDocuments(actual)
	if err != nil {
		return hammy.Assert(false, "actual YAML invalid: %v", err)
	}
	if index < 0 || index >= len(documents) {
		return hammy.Assert(false, "got YAML document index <%d> out of range for count <%d>", index, len(documents))
	}

	expectedYAML, err := parseYAML(expected)
	if err != nil {
		return hammy.Assert(false, "expected YAML invalid: %v", err)
	}

	actualYAML, expectedYAML, err := applyOptions(documents[index], expectedYAML, opts...)
	if err != nil {
		return hammy.Assert(false, "%v", err)
	}

	diff := cmp.Diff(expectedYAML, actualYAML)
	return hammy.Assert(diff == "", "YAML document <%d> mismatch (-want +got):\n%s", index, diff)
}

func DocumentContains(actual string, index int, expected string) hammy.AssertionMessage {
	return DocumentContainsBytes([]byte(actual), index, []byte(expected))
}

func DocumentContainsBytes(actual []byte, index int, expected []byte) hammy.AssertionMessage {
	documents, err := parseYAMLDocuments(actual)
	if err != nil {
		return hammy.Assert(false, "actual YAML invalid: %v", err)
	}
	if index < 0 || index >= len(documents) {
		return hammy.Assert(false, "got YAML document index <%d> out of range for count <%d>", index, len(documents))
	}

	expectedYAML, err := parseYAML(expected)
	if err != nil {
		return hammy.Assert(false, "expected YAML invalid: %v", err)
	}

	if ok, message := containsYAML(documents[index], expectedYAML, "$"); !ok {
		return hammy.Assert(false, "%s", message)
	}
	return hammy.Assert(true, "YAML document <%d> contained expected subset", index)
}

func readYAML(name string, reader io.Reader) ([]byte, hammy.AssertionMessage) {
	if reader == nil {
		return nil, hammy.Assert(false, "%s YAML reader is nil", name)
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, hammy.Assert(false, "%s YAML read error: %v", name, err)
	}
	return data, hammy.Assert(true, "%s YAML read", name)
}

func applyOptions(actual, expected any, opts ...Option) (any, any, error) {
	var options compareOptions
	for _, opt := range opts {
		opt(&options)
	}

	for _, path := range options.ignorePaths {
		steps, err := parsePath(path)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid YAML path <%s>: %w", path, err)
		}
		actual = deleteYAMLPath(actual, steps)
		expected = deleteYAMLPath(expected, steps)
	}

	for _, path := range options.unorderedArrayPaths {
		steps, err := parsePath(path)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid YAML path <%s>: %w", path, err)
		}
		if err := sortYAMLArrayAtPath(actual, steps, path); err != nil {
			return nil, nil, err
		}
		if err := sortYAMLArrayAtPath(expected, steps, path); err != nil {
			return nil, nil, err
		}
	}

	return actual, expected, nil
}

func deleteYAMLPath(value any, steps []pathStep) any {
	if len(steps) == 0 {
		return nil
	}

	parent, found := lookupParsedYAMLPath(value, steps[:len(steps)-1])
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

func sortYAMLArrayAtPath(value any, steps []pathStep, path string) error {
	target, found := lookupParsedYAMLPath(value, steps)
	if !found {
		return nil
	}

	items, ok := target.([]any)
	if !ok {
		return fmt.Errorf("got YAML path <%s> type <%T>, wanted array", path, target)
	}
	sort.SliceStable(items, func(i, j int) bool {
		return canonicalYAMLKey(items[i]) < canonicalYAMLKey(items[j])
	})
	return nil
}

func canonicalYAMLKey(value any) string {
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
	case taggedScalar:
		return "tagged:" + typed.Tag + ":" + strconv.Quote(typed.Value)
	case []any:
		parts := make([]string, 0, len(typed))
		for _, item := range typed {
			parts = append(parts, canonicalYAMLKey(item))
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
			parts = append(parts, strconv.Quote(key)+":"+canonicalYAMLKey(typed[key]))
		}
		return "object:{" + strings.Join(parts, ",") + "}"
	default:
		return fmt.Sprintf("%T:%v", typed, typed)
	}
}

func containsYAML(actual, expected any, path string) (bool, string) {
	switch expectedValue := expected.(type) {
	case map[string]any:
		actualValue, ok := actual.(map[string]any)
		if !ok {
			return false, fmt.Sprintf("got YAML path <%s> type <%T>, wanted object", path, actual)
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
				return false, fmt.Sprintf("YAML path <%s> missing", itemPath)
			}
			if ok, message := containsYAML(actualItem, expectedValue[key], itemPath); !ok {
				return false, message
			}
		}
		return true, ""
	case []any:
		if diff := cmp.Diff(expected, actual); diff != "" {
			return false, fmt.Sprintf("YAML path <%s> mismatch (-want +got):\n%s", path, diff)
		}
		return true, ""
	default:
		if !cmp.Equal(actual, expected) {
			return false, fmt.Sprintf("YAML path <%s> mismatch (-want +got):\n%s", path, cmp.Diff(expected, actual))
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

func lookupYAMLPath(value any, path string) (any, bool, error) {
	steps, err := parsePath(path)
	if err != nil {
		return nil, false, err
	}
	value, found := lookupParsedYAMLPath(value, steps)
	return value, found, nil
}

func lookupParsedYAMLPath(value any, steps []pathStep) (any, bool) {
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

func parseYAML(data []byte) (any, error) {
	documents, err := parseYAMLDocuments(data)
	if err != nil {
		return nil, err
	}
	if len(documents) == 0 {
		return nil, nil
	}
	if len(documents) > 1 {
		return nil, fmt.Errorf("multiple YAML documents")
	}
	return documents[0], nil
}

func parseYAMLDocuments(data []byte) ([]any, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var documents []any
	for {
		var node yaml.Node
		if err := decoder.Decode(&node); err != nil {
			if err == io.EOF {
				return documents, nil
			}
			return nil, err
		}

		normalized, err := normalizeYAMLNode(&node)
		if err != nil {
			return nil, err
		}
		documents = append(documents, normalized)
	}
}

func normalizeYAMLNode(node *yaml.Node) (any, error) {
	if node == nil {
		return nil, nil
	}

	switch node.Kind {
	case 0:
		return nil, nil
	case yaml.DocumentNode:
		if len(node.Content) == 0 {
			return nil, nil
		}
		return normalizeYAMLNode(node.Content[0])
	case yaml.AliasNode:
		if node.Alias == nil {
			return nil, fmt.Errorf("unresolved YAML alias <%s>", node.Value)
		}
		return normalizeYAMLNode(node.Alias)
	case yaml.SequenceNode:
		normalized := make([]any, len(node.Content))
		for i, item := range node.Content {
			normalizedItem, err := normalizeYAMLNode(item)
			if err != nil {
				return nil, err
			}
			normalized[i] = normalizedItem
		}
		return normalized, nil
	case yaml.MappingNode:
		normalized := make(map[string]any, len(node.Content)/2)
		for i := 0; i < len(node.Content); i += 2 {
			key, err := normalizeYAMLKey(node.Content[i])
			if err != nil {
				return nil, err
			}
			if _, ok := normalized[key]; ok {
				return nil, fmt.Errorf("duplicate YAML key <%s>", key)
			}

			value, err := normalizeYAMLNode(node.Content[i+1])
			if err != nil {
				return nil, err
			}
			normalized[key] = value
		}
		return normalized, nil
	case yaml.ScalarNode:
		return normalizeYAMLScalar(node)
	default:
		return nil, fmt.Errorf("unsupported YAML node kind <%d>", node.Kind)
	}
}

func normalizeYAMLKey(node *yaml.Node) (string, error) {
	if node.Kind == yaml.ScalarNode && node.Tag == "!!str" {
		return node.Value, nil
	}

	value, err := normalizeYAMLNode(node)
	if err != nil {
		return "", err
	}
	return canonicalYAMLKey(value), nil
}

type normalizedNumber struct {
	Value string
}

type taggedScalar struct {
	Tag   string
	Value string
}

func normalizeYAMLScalar(node *yaml.Node) (any, error) {
	switch node.Tag {
	case "!!null":
		return nil, nil
	case "!!bool":
		value, err := strconv.ParseBool(strings.ToLower(node.Value))
		if err != nil {
			return nil, fmt.Errorf("invalid YAML bool %q", node.Value)
		}
		return value, nil
	case "!!int", "!!float":
		if isSpecialFloat(node.Value) {
			return taggedScalar{Tag: node.Tag, Value: strings.ToLower(node.Value)}, nil
		}
		return normalizeNumber(node.Value)
	case "!!str":
		return node.Value, nil
	default:
		return taggedScalar{Tag: node.Tag, Value: node.Value}, nil
	}
}

func isSpecialFloat(value string) bool {
	normalized := strings.ToLower(strings.TrimPrefix(value, "+"))
	return normalized == ".inf" || normalized == "-.inf" || normalized == ".nan"
}

func normalizeNumber(raw string) (normalizedNumber, error) {
	rat, err := parseYAMLNumber(raw)
	if err != nil {
		return normalizedNumber{}, err
	}
	return normalizedNumber{Value: rat.RatString()}, nil
}

func parseYAMLNumber(raw string) (*big.Rat, error) {
	cleaned := strings.ReplaceAll(raw, "_", "")
	cleaned = strings.TrimPrefix(cleaned, "+")

	mantissa := cleaned
	exponent := 0
	if exponentIndex := strings.IndexAny(cleaned, "eE"); exponentIndex >= 0 {
		parsedExponent, err := strconv.Atoi(cleaned[exponentIndex+1:])
		if err != nil {
			return nil, fmt.Errorf("unsupported YAML number exponent %q: %w", raw, err)
		}
		exponent = parsedExponent
		mantissa = cleaned[:exponentIndex]
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
		return nil, fmt.Errorf("invalid YAML number %q", raw)
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
