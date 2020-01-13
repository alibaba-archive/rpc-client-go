package service

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
		flatRepeatedList(filterValue, result, key)
	}
	assert.Equal(t, result["slice.1.map"], "valid")
	assert.Equal(t, result["slice.2"], "6")
	assert.Equal(t, result["map.value"], "ok")
	assert.Equal(t, result["client"], "test")
	assert.Equal(t, result["slice.1.map"], "valid")
}

func Test_Prettify(t *testing.T) {
	tmp := map[string]string{
		"rpc": "ok",
	}
	str := Prettify(tmp)
	assert.Equal(t, str, "{\n   \"rpc\": \"ok\"\n}")
}

func Test_Get(t *testing.T) {
	num := getIntValue(10)
	assert.Equal(t, num, 10)

	val := getBoolValue(nil)
	assert.Equal(t, val, false)

	val = getBoolValue(true)
	assert.Equal(t, val, true)
}
