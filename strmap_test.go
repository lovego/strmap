package strmap

import (
	//	"errors"
	"reflect"
	"testing"
)

var sm = StrMap(map[string]interface{}{
	`key1`: `value1`,
	`key2`: map[interface{}]interface{}{`key`: `value`},
	`key3`: []interface{}{`value1`, `value2`, `value3`},
})

func checkPanic(t *testing.T, expect interface{}) {
	got := recover()
	if got == nil || got == expect || reflect.DeepEqual(got, expect) {
		return
	}
	t.Fatalf("unexpected Panic: %T(%#v)", got, got)
}

func TestGet1(t *testing.T) {
	defer checkPanic(t, "no key: nonexist")
	got := sm.Get(`nonexist`)
	t.Errorf("expect panic got: %v\n", got)
}

func TestGet2(t *testing.T) {
	defer checkPanic(t, `unable to cast "value1" of type string to map[string]interface{}`)
	got := sm.Get(`key1`)
	t.Errorf("expect panic got: %v\n", got)
}

func TestGet3(t *testing.T) {
	defer checkPanic(t, "no key: nonexist")
	got := sm.Get(`key2`).Get(`nonexist`)
	t.Errorf("expect panic got: %v\n", got)
}

func TestGet4(t *testing.T) {
	defer checkPanic(t, `unable to cast []interface {}{"value1", "value2", "value3"}`+
		` of type []interface {} to map[string]interface{}`)
	got := sm.Get(`key3`)
	t.Errorf("expect panic got: %v\n", got)
}

func TestGet5(t *testing.T) {
	expect := StrMap(map[string]interface{}{`key`: `value`})
	if got := sm.Get(`key2`); !reflect.DeepEqual(got, expect) {
		t.Errorf("expect: %v got: %v\n", expect, got)
	}
}

func TestGetString1(t *testing.T) {
	defer checkPanic(t, "no key: nonexist")
	got := sm.GetString(`nonexist`)
	t.Errorf("expect panic got: %v\n", got)
}

func TestGetString2(t *testing.T) {
	if expect, got := `value1`, sm.GetString(`key1`); got != expect {
		t.Errorf("expect: %v got: %v\n", expect, got)
	}
}

func TestGetString3(t *testing.T) {
	defer checkPanic(t, `unable to cast map[interface {}]interface {}{"key":"value"}`+
		` of type map[interface {}]interface {} to string`)
	got := sm.GetString(`key2`)
	t.Errorf("expect panic got: %v\n", got)
}

func TestGetString4(t *testing.T) {
	defer checkPanic(t, `unable to cast []interface {}{"value1", "value2", "value3"}`+
		` of type []interface {} to string`)
	got := sm.GetString(`key3`)
	t.Errorf("expect: panic got: %v\n", got)
}

func TestGetStringSlice1(t *testing.T) {
	defer checkPanic(t, "no key: nonexist")
	got := sm.GetStringSlice(`nonexist`)
	t.Errorf("expect panic got: %v\n", got)
}

func TestGetStringSlice2(t *testing.T) {
	if expect, got := []string{`value1`}, sm.GetStringSlice(`key1`); !reflect.DeepEqual(got, expect) {
		t.Errorf("expect: %v got: %v\n", expect, got)
	}
}

func TestGetStringSlice3(t *testing.T) {
	defer checkPanic(t, `unable to cast map[interface {}]interface {}{"key":"value"}`+
		` of type map[interface {}]interface {} to []string`)
	got := sm.GetStringSlice(`key2`)
	t.Errorf("expect panic got: %v\n", got)
}

func TestGetStringSlice4(t *testing.T) {
	expect, got := []string{`value1`, `value2`, `value3`}, sm.GetStringSlice(`key3`)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expect: %v got: %v\n", expect, got)
	}
}
