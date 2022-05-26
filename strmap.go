package strmap

import (
	"log"

	"github.com/spf13/cast"
)

type StrMap map[string]interface{}

func (s StrMap) Get(key string) StrMap {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}

	value, err := cast.ToStringMapE(trans(s[key]))
	if err != nil {
		log.Panic(err)
	}
	return StrMap(value)
}

func trans(value interface{}) interface{} {
	var val interface{}
	switch v := value.(type) {
	case StrMap:
		val = (map[string]interface{})(v)
	default:
		val = v
	}
	return val
}

func (s StrMap) GetSlice(key string) (results []StrMap) {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	slice, err := cast.ToSliceE(s[key])
	if err != nil {
		log.Panic(err)
	}
	for _, v := range slice {
		if value, err := cast.ToStringMapE(trans(v)); err == nil {
			results = append(results, StrMap(value))
		} else {
			log.Panic(err)
		}
	}
	return
}

func (s StrMap) GetBool(key string) bool {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToBoolE(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetString(key string) string {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToStringE(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetStringSlice(key string) []string {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToStringSliceE(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetInt(key string) int {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToIntE(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetInt64(key string) int64 {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToInt64E(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetInt32(key string) int32 {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToInt32E(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetInt16(key string) int16 {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToInt16E(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetInt8(key string) int8 {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToInt8E(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetUint(key string) uint {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToUintE(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetUint64(key string) uint64 {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToUint64E(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetUint32(key string) uint32 {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToUint32E(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetUint16(key string) uint16 {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToUint16E(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetUint8(key string) uint8 {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToUint8E(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}
