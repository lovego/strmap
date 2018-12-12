package strmap

import (
	"fmt"
)

func Convert(data map[interface{}]interface{}, path string) (
	map[string]interface{}, error,
) {
	result := make(map[string]interface{})
	for k, v := range data {
		key, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("key: %s%v is not a string.")
		}
		switch value := v.(type) {
		case map[interface{}]interface{}:
			if strMap, err := Convert(value, path+key+"."); err == nil {
				result[key] = strMap
			} else {
				return nil, err
			}
		case []interface{}:
			if slice, err := ConvertSlice(value, path+key); err == nil {
				result[key] = slice
			} else {
				return nil, err
			}
		default:
			result[key] = v
		}
	}
	return result, nil
}

func ConvertSlice(slice []interface{}, path string) (result []interface{}, err error) {
	for i, v := range slice {
		switch value := v.(type) {
		case map[interface{}]interface{}:
			if strMap, err := Convert(value, fmt.Sprintf("%s[%d]", path, i)); err == nil {
				result = append(result, strMap)
			} else {
				return nil, err
			}
		case []interface{}:
			if slice, err := ConvertSlice(value, fmt.Sprintf("%s[%d]", path, i)); err == nil {
				result = append(result, slice)
			} else {
				return nil, err
			}
		default:
			result = append(result, v)
		}
	}
	return
}
