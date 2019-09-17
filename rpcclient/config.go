package rpcclient

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func flatRepeatedList(dataValue reflect.Value, result map[string]string, prefix string) (err error) {
	if !dataValue.IsValid() {
		return
	}

	dataType := dataValue.Type()
	if dataType.Kind().String() == "slice" {
		err = handleRepeatedParams(dataValue, result, prefix)
		if err != nil {
			return
		}
	} else if dataType.Kind().String() == "map" {
		err = handleMap(dataValue, result, prefix)
		if err != nil {
			return
		}
	} else {
		result[prefix] = fmt.Sprintf("%v", dataValue.Interface())
	}
	return
}

func handleRepeatedParams(repeatedFieldValue reflect.Value, result map[string]string, prefix string) (err error) {
	if repeatedFieldValue.IsValid() && !repeatedFieldValue.IsNil() {
		for m := 0; m < repeatedFieldValue.Len(); m++ {
			elementValue := repeatedFieldValue.Index(m)
			key := prefix + "." + strconv.Itoa(m+1)
			fieldValue := reflect.ValueOf(elementValue.Interface())
			if fieldValue.Kind().String() == "map" {
				err = handleMap(fieldValue, result, key)
				if err != nil {
					return
				}
			} else {
				result[key] = fmt.Sprintf("%v", fieldValue.Interface())
			}
		}
	}
	return nil
}

func handleMap(valueField reflect.Value, result map[string]string, prefix string) (err error) {
	if valueField.IsValid() && valueField.String() != "" {
		valueFieldType := valueField.Type()
		if valueFieldType.Kind().String() == "map" {
			var byt []byte
			byt, err = json.Marshal(valueField.Interface())
			if err != nil {
				return
			}
			cache := make(map[string]interface{})
			err = json.Unmarshal(byt, &cache)
			if err != nil {
				return
			}
			for key, value := range cache {
				pre := prefix + "." + key
				fieldValue := reflect.ValueOf(value)
				err = flatRepeatedList(fieldValue, result, pre)
				if err != nil {
					return
				}
			}
		}
	}
	return nil
}
