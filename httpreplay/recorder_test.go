// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// Run with a command something like:
//   go test -run TestUnmarshal

package httpreplay

import "testing"

func TestUnmarshal(t *testing.T) {
	t.Run("Unmarshal Array", func(t *testing.T) {
		tests := []struct {
			input    string
			desc     string
			elements int
		}{
			{`[]`, "empty array", 0},
			{`[{}, {}]`, "array of two empty objects", 2},
			{`[{"key1":"value1","key2":"value2"}]`, "array with a single object", 1},
		}

		for _, test := range tests {
			if result, err := unmarshal([]byte(test.input)); err != nil {
				t.Errorf("Unable to unmarshal %v : %v", test.desc, err)
			} else {
				if arr, ok := result.(jsonArr); !ok {
					t.Errorf("After unmarshalling %v, expected jsonArr, got %T", test.desc, result)
				} else {
					if len(arr) != test.elements {
						t.Errorf("After unmarshalling %v, expected %v elements, got %v", test.desc, test.elements, len(arr))
					}
				}
			}
		}
	})
	t.Run("Unmarshal Object", func(t *testing.T) {
		tests := []struct {
			input    string
			desc     string
			elements []string
		}{
			{`{}`, "empty object", []string{}},
			{`{"key":"string"}`, "object", []string{"key"}},
			{`{"key":3}`, "object containing an int", []string{"key"}},
			{`{"key":["one","two","three"]}`, "object containing an array", []string{"key"}},
			{`{"key":{"one":1,"two":3,"three":2}}`, "object containing an object", []string{"key"}},
			{`{"key1":"string","key2":3,"key3":["one","two","three"],"key4":{"one":1,"two":3,"three":2}}`, "object containing multiple types", []string{"key1", "key2", "key3", "key4"}},
		}
		for _, test := range tests {
			if result, err := unmarshal([]byte(test.input)); err != nil {
				t.Errorf("Unable to unmarshal %v: %v", test.desc, err)
			} else {
				if obj, ok := result.(jsonObj); !ok {
					t.Errorf("After unmarshalling %v, expected jsonObj, got %T", test.desc, result)
				} else {
					for _, key := range test.elements {
						if _, ok := obj[key]; !ok {
							t.Errorf("After unmarshalling %v, missing key %v", test.desc, key)
						}
					}
				}
			}
		}
	})
}
