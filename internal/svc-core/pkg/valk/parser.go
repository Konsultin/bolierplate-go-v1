package valk

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseBoolean(input interface{}) (result bool, ok bool) {
	// Assume ok to true
	ok = true

	switch v := input.(type) {
	case bool:
		result = v
	case string:
		result = strings.ToLower(v) == "true"
	case int:
		result = v > 0
	case int8:
		result = v > 0
	case int16:
		result = v > 0
	case int32:
		result = v > 0
	case int64:
		result = v > 0
	case uint:
		result = v > 0
	case uint8:
		result = v > 0
	case uint16:
		result = v > 0
	case uint32:
		result = v > 0
	case uint64:
		result = v > 0
	default:
		// Failed, set ok to false
		ok = false
	}

	return result, ok
}

func ParseBooleanFallback(input interface{}, fallbackValue bool) bool {
	result, ok := ParseBoolean(input)
	if !ok {
		return fallbackValue
	}
	return result
}

func ParseStringArray(input interface{}) (result []string, ok bool) {
	// Assume ok to true
	ok = true
	switch v := input.(type) {
	case []interface{}:
		for _, item := range v {
			result = append(result, fmt.Sprintf("%v", item))
		}
	case string:
		if len(v) == 0 {
			ok = false
		} else if strArr := strings.Split(v, ","); len(strArr) > 1 {
			result = strArr
		} else {
			result = []string{v}
		}
	default:
		// Failed, set ok to false
		ok = false
	}

	return result, ok
}

func ParseStringArrayFallback(input interface{}, fallbackValue []string) []string {
	result, ok := ParseStringArray(input)
	if !ok {
		return fallbackValue
	}
	return result
}

func ParseInt(input interface{}) (result int, ok bool) {
	// Assume ok to true
	ok = true

	switch v := input.(type) {
	case string:
		var err error
		result, err = strconv.Atoi(v)
		if err != nil {
			ok = false
		}
	case int:
		result = v
	default:
		// Failed, set ok to false
		ok = false
	}

	return result, ok
}

func ParseIntFallback(input interface{}, fallbackValue int) int {
	i, ok := ParseInt(input)
	if !ok {
		return fallbackValue
	}
	return i
}

func ParseInt64(input interface{}) (result int64, ok bool) {
	// Assume ok to true
	ok = true

	switch v := input.(type) {
	case string:
		var err error
		result, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			ok = false
		}
	case int64:
		result = v
	default:
		// Failed, set ok to false
		ok = false
	}

	return result, ok
}

func ParseInt32(input interface{}) (result int32, ok bool) {
	// Assume ok to true
	ok = true

	switch v := input.(type) {
	case string:
		var err error
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			ok = false
		}
		result = int32(i)
	case int:
		result = int32(v)
	case int8:
		result = int32(v)
	case int16:
		result = int32(v)
	case int32:
		result = v
	case int64:
		result = int32(v)
	default:
		// Failed, set ok to false
		ok = false
	}

	return result, ok
}

func ParseInt64Fallback(input interface{}, fallbackValue int64) int64 {
	i, ok := ParseInt64(input)
	if !ok {
		return fallbackValue
	}
	return i
}

func ParseString(input interface{}) (result string, ok bool) {
	ok = true
	switch v := input.(type) {
	case string:
		result = v
	case nil:
		ok = false
	default:
		result = fmt.Sprintf("%v", v)
	}
	return result, ok
}

func ParseStringFallback(input interface{}, fallbackValue string) string {
	result, ok := ParseString(input)
	if !ok || result == "" {
		return fallbackValue
	}
	return result
}

func ParseFloat64(input interface{}) (result float64, ok bool) {
	// Assume ok to true
	ok = true

	switch v := input.(type) {
	case string:
		var err error
		result, err = strconv.ParseFloat(v, 64)
		if err != nil {
			ok = false
		}
	case float64:
		result = v
	case int:
		result = float64(v)
	default:
		// Failed, set ok to false
		ok = false
	}

	return result, ok
}

func ParseFloat64Fallback(input interface{}, fallbackValue float64) float64 {
	result, ok := ParseFloat64(input)
	if !ok || result == 0 {
		return fallbackValue
	}
	return result
}

func ParseFloat32Fallback(input interface{}, fallbackValue float64) float32 {
	result, ok := ParseFloat64(input)
	if !ok || result == 0 {
		return float32(fallbackValue)
	}
	return float32(result)
}

func ParseInt32Fallback(input interface{}, fallbackValue int32) int32 {
	i, ok := ParseInt32(input)
	if !ok {
		return fallbackValue
	}
	return i
}

func SanitizeListFromString(input string) []string {
	parts := strings.Split(input, ",")
	seen := make(map[string]bool)
	valid := make([]string, 0, len(parts))

	for _, part := range parts {
		xid := strings.TrimSpace(part)
		if xid != "" && !seen[xid] {
			valid = append(valid, xid)
			seen[xid] = true
		}
	}
	return valid
}
