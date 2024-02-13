package filter

import (
	"strings"
)

func extractMapKeysAsMap(inputMap map[string]interface{}) map[string]string {
	output := map[string]string{}

	for key, value := range inputMap {
		keyLower := strings.ToLower(key)
		output[keyLower] = key

		// Attempt to cast value into a map
		vAsMap, ok := value.(map[string]interface{})

		if ok {
			for subKey, subValue := range extractMapKeysAsMap(vAsMap) {
				output[keyLower+"."+subKey] = key + "." + subValue
			}
		}
	}

	return output
}

func convertValue[T any](value interface{}) (*T, bool) {
	var val T
	var ok bool

	// Return nil if nil
	if value == nil {
		return nil, true
	}

	// Perform type cast
	val, ok = value.(T)

	// Return if not ok
	if !ok {
		return nil, ok
	}

	// Return value
	return &val, ok
}
