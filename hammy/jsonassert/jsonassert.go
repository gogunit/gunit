package jsonassert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"

	"github.com/gogunit/gunit/hammy"
	"github.com/google/go-cmp/cmp"
)

func Equal(actual, expected string) hammy.AssertionMessage {
	return EqualBytes([]byte(actual), []byte(expected))
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
	actualJSON, err := parseJSON(actual)
	if err != nil {
		return hammy.Assert(false, "actual JSON invalid: %v", err)
	}

	expectedJSON, err := parseJSON(expected)
	if err != nil {
		return hammy.Assert(false, "expected JSON invalid: %v", err)
	}

	diff := cmp.Diff(expectedJSON, actualJSON)
	return hammy.Assert(diff == "", "JSON mismatch (-want +got):\n%s", diff)
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
