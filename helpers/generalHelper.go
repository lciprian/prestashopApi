package helpers

import (
	"fmt"
	"strconv"
)

func IDtoString(data interface{}) (string, error) {
	if strVal, ok := data.(string); ok {
		return strVal, nil
	}

	if intVal, ok := data.(int); ok {
		return strconv.Itoa(intVal), nil
	}

	// Check if the interface holds an int
	if floatVal, ok := data.(float64); ok {
		return fmt.Sprintf("%.0f", floatVal), nil
	}

	return "", fmt.Errorf("invalid data type: %s", data)
}
