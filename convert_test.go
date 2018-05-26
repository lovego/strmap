package strmap

import (
	"reflect"
	"testing"

	"github.com/lovego/deep"
)

func TestConvert(t *testing.T) {
	input := map[interface{}]interface{}{
		"properties": map[interface{}]interface{}{
			"host":     map[interface{}]interface{}{"type": "keyword"},
			"query":    map[interface{}]interface{}{"type": "text"},
			"status":   map[interface{}]interface{}{"type": "keyword"},
			"req_body": map[interface{}]interface{}{"type": "integer"},
			"res_body": map[interface{}]interface{}{"type": "integer"},
			"agent":    map[interface{}]interface{}{"type": "text"},
			"at":       map[interface{}]interface{}{"type": "date"},
			"method":   map[interface{}]interface{}{"type": "keyword"},
			"path": map[interface{}]interface{}{
				"type": "text",
				"fields": map[interface{}]interface{}{
					"raw": map[interface{}]interface{}{"type": "keyword"},
				},
			},
			"ip":       map[interface{}]interface{}{"type": "ip"},
			"refer":    map[interface{}]interface{}{"type": "text"},
			"proto":    map[interface{}]interface{}{"type": "keyword"},
			"duration": map[interface{}]interface{}{"type": "float"},
		},
		"dynamic_templates": []interface{}{
			map[interface{}]interface{}{
				"query": map[interface{}]interface{}{
					"path_match": "query.*",
					"mapping": map[interface{}]interface{}{
						"type": "text",
						"fields": map[interface{}]interface{}{
							"raw": map[interface{}]interface{}{"type": "keyword"},
						},
					},
				},
			},
		},
	}
	expect := map[string]interface{}{
		"properties": map[string]interface{}{
			"host":     map[string]interface{}{"type": "keyword"},
			"query":    map[string]interface{}{"type": "text"},
			"status":   map[string]interface{}{"type": "keyword"},
			"req_body": map[string]interface{}{"type": "integer"},
			"res_body": map[string]interface{}{"type": "integer"},
			"agent":    map[string]interface{}{"type": "text"},
			"at":       map[string]interface{}{"type": "date"},
			"method":   map[string]interface{}{"type": "keyword"},
			"path": map[string]interface{}{
				"type": "text",
				"fields": map[string]interface{}{
					"raw": map[string]interface{}{"type": "keyword"},
				},
			},
			"ip":       map[string]interface{}{"type": "ip"},
			"refer":    map[string]interface{}{"type": "text"},
			"proto":    map[string]interface{}{"type": "keyword"},
			"duration": map[string]interface{}{"type": "float"},
		},
		"dynamic_templates": []interface{}{
			map[string]interface{}{
				"query": map[string]interface{}{
					"path_match": "query.*",
					"mapping": map[string]interface{}{
						"type": "text",
						"fields": map[string]interface{}{
							"raw": map[string]interface{}{"type": "keyword"},
						},
					},
				},
			},
		},
	}

	got, err := Convert(input, "")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("\ninput: %s\n diff: %+v\n", input, deep.Equal(got, expect))
	}
}
