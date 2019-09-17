package rpcclient

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_flatRepeatedList(t *testing.T) {
	filter := map[string]interface{}{
		"client":  "test",
		"version": "1",
		"null":    nil,
		"slice": []interface{}{
			map[string]interface{}{
				"map": "valid",
			},
			6,
		},
		"map": map[string]interface{}{
			"value": "ok",
		},
	}

	result := make(map[string]string)
	for key, value := range filter {
		filterValue := reflect.ValueOf(value)
		err := flatRepeatedList(filterValue, result, key)
		assert.Nil(t, err)
	}
}
