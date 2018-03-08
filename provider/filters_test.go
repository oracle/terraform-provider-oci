package provider

import (
	"testing"

	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

// Filter function should select an item for which the compare function returns true
func TestFilter(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "b"},
		{"letter": "c"},
	}

	res := filter(items, func(item map[string]interface{}) bool {
		return item["letter"] == "b"
	})

	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}

// Not supplying filters should not restrict results
func TestApplyFilters_passThrough(t *testing.T) {
	items := []map[string]interface{}{
		{},
		{},
		{},
	}

	res := ApplyFilters(nil, items)
	if len(res) != 3 {
		t.Errorf("Expected 3 results, got %d", len(res))
	}
}

// Filtering against a nonexistent property should throw no errors and return no results
func TestApplyFilters_nonExistentProperty(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
	}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "number",
		"values": []interface{}{"1"},
	})

	res := ApplyFilters(filters, items)
	if len(res) > 0 {
		t.Errorf("Expected 0 results, got %d", len(res))
	}
}

// Filtering against an empty resource set should not throw errors
func TestApplyFilters_noResources(t *testing.T) {
	items := []map[string]interface{}{}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "number",
		"values": []interface{}{"1"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 0 {
		t.Errorf("Expected 0 results, got %d", len(res))
	}
}

func TestApplyFilters_basic(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "b"},
		{"letter": "c"},
	}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"b"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}

func TestApplyFilters_duplicates(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "a"},
		{"letter": "c"},
	}

	filters := &schema.Set{F: func(v interface{}) int {
		return schema.HashString(v.(map[string]interface{})["name"])
	}}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"a"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 2 {
		t.Errorf("Expected 2 results, got %d", len(res))
	}
}

func TestApplyFilters_OR(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "b"},
		{"letter": "c"},
	}

	filters := &schema.Set{
		F: func(v interface{}) int {
			elems := v.(map[string]interface{})["values"].([]interface{})
			res := make([]string, len(elems))
			for i, v := range elems {
				res[i] = v.(string)
			}
			return schema.HashString(strings.Join(res, ""))
		},
	}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"a", "b"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 2 {
		t.Errorf("Expected 2 results, got %d", len(res))
	}
}

func TestApplyFilters_cascadeAND(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "b"},
		{"letter": "c"},
	}

	filters := &schema.Set{
		F: func(v interface{}) int {
			elems := v.(map[string]interface{})["values"].([]interface{})
			res := make([]string, len(elems))
			for i, v := range elems {
				res[i] = v.(string)
			}
			return schema.HashString(strings.Join(res, ""))
		},
	}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"a", "b"},
	})
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"c"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 0 {
		t.Errorf("Expected 0 results, got %d", len(res))
	}
}

func TestApplyFilters_regex(t *testing.T) {
	items := []map[string]interface{}{
		{"string": "xblx:PHX-AD-1"},
		{"string": "xblx:PHX-AD-2"},
		{"string": "xblx:PHX-AD-3"},
	}

	filters := &schema.Set{F: func(v interface{}) int {
		return schema.HashString(v.(map[string]interface{})["name"])
	}}
	filters.Add(map[string]interface{}{
		"name":   "string",
		"values": []interface{}{"\\w*:PHX-AD-2"},
		"regex":  true,
	})

	res := ApplyFilters(filters, items)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}

// Invalid regex should throw an error
func TestApplyFilters_regexPanic(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("Expected regex compile error was not thrown")
		} else {
			if err.(error).Error() != `Invalid regular expression ")(" for "string" filter` {
				t.Errorf("Unexpected regex compile error:\n%s", err)
			}
		}
	}()

	items := []map[string]interface{}{
		{"string": "xblx:PHX-AD-1"},
	}

	filters := &schema.Set{F: func(v interface{}) int {
		return schema.HashString(v.(map[string]interface{})["name"])
	}}
	filters.Add(map[string]interface{}{
		"name":   "string",
		"values": []interface{}{")("},
		"regex":  true,
	})

	ApplyFilters(filters, items)
}

// Filters should test against an array of strings
func TestApplyFilters_arrayOfStrings(t *testing.T) {
	items := []map[string]interface{}{
		{"letters": []string{"a"}},
		{"letters": []string{"b", "c"}},
		{"letters": []string{"c", "d", "e"}},
		{"letters": []string{"e", "f"}},
	}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "letters",
		"values": []interface{}{"a", "c"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 3 {
		t.Errorf("Expected 3 result, got %d", len(res))
	}

	filters = &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "letters",
		"values": []interface{}{"a", "f"},
	})

	res = ApplyFilters(filters, items)
	if len(res) != 2 {
		t.Errorf("Expected 2 result, got %d", len(res))
	}
}

type CustomStringTypeA string
type CustomStringTypeB CustomStringTypeA
type CustomEnumType oci_core.VcnLifecycleStateEnum

func TestApplyFilters_underlyingStringTypes(t *testing.T) {
	items := []map[string]interface{}{
		{
			"letters": []CustomStringTypeA{"a"},
			"number":  CustomStringTypeB("1"),
			"state":   oci_core.SecurityListLifecycleStateAvailable,
			"custom":  CustomEnumType(oci_core.VcnLifecycleStateTerminated),
		},
		{
			"letters": []CustomStringTypeA{"a"},
			"number":  CustomStringTypeB("1"),
			"state":   oci_core.SecurityListLifecycleStateProvisioning,
			"custom":  CustomEnumType(oci_core.VcnLifecycleStateTerminating),
		},
		{
			"letters": []CustomStringTypeA{"b", "c"},
			"number":  CustomStringTypeB("2"),
			"state":   oci_core.SecurityListLifecycleStateTerminating,
			"custom":  CustomEnumType(oci_core.VcnLifecycleStateProvisioning),
		},
		{
			"letters": []CustomStringTypeA{"c", "d", "e"},
			"number":  CustomStringTypeB("3"),
			"state":   oci_core.SecurityListLifecycleStateTerminated,
			"custom":  CustomEnumType(oci_core.VcnLifecycleStateAvailable),
		},
		{
			"letters": []CustomStringTypeA{"e", "f"},
			"number":  CustomStringTypeB("5"),
			"state":   oci_core.SecurityListLifecycleStateAvailable,
			"custom":  CustomEnumType(oci_core.VcnLifecycleStateTerminated),
		},
	}

	filters := &schema.Set{
		F: func(v interface{}) int {
			return schema.HashString(v.(map[string]interface{})["name"])
		},
	}
	filters.Add(map[string]interface{}{
		"name":   "letters",
		"values": []interface{}{"a", "c"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 4 {
		t.Errorf("Expected 4 result, got %d", len(res))
	}

	filters1 := &schema.Set{
		F: func(v interface{}) int {
			return schema.HashString(v.(map[string]interface{})["name"])
		},
	}
	filters1.Add(map[string]interface{}{
		"name":   "letters",
		"values": []interface{}{"a", "b", "e"},
	})
	filters1.Add(map[string]interface{}{
		"name":   "number",
		"values": []interface{}{"1", "notANumber"},
	})

	res = ApplyFilters(filters1, items)
	if len(res) != 2 {
		t.Errorf("Expected 2 result, got %d", len(res))
	}

	filters2 := &schema.Set{
		F: func(v interface{}) int {
			return schema.HashString(v.(map[string]interface{})["name"])
		},
	}
	filters2.Add(map[string]interface{}{
		"name":   "letters",
		"values": []interface{}{"a", "b", "e"},
	})
	filters2.Add(map[string]interface{}{
		"name":   "number",
		"values": []interface{}{"1", "2", "3", "5"},
	})
	filters2.Add(map[string]interface{}{
		"name":   "state",
		"values": []interface{}{string(oci_core.SecurityListLifecycleStateAvailable), string(oci_core.SecurityListLifecycleStateTerminating)},
	})
	filters2.Add(map[string]interface{}{
		"name":   "custom",
		"values": []interface{}{string(oci_core.VcnLifecycleStateProvisioning)},
	})

	res = ApplyFilters(filters2, items)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}

// Test various fields that aren't strings. Non-string filters should result in item being filtered out.
func TestApplyFilters_nonString(t *testing.T) {
	items := []map[string]interface{}{
		{
			"letter":  "a",
			"number":  1,
			"enabled": true,
			"nums":    []int{1, 2, 3},
		},
		{
			"letter":  "b",
			"number":  2,
			"enabled": false,
			"nums":    []int{3, 4, 5},
		},
		{
			"letter":  "c",
			"number":  2,
			"enabled": true,
			"nums":    []int{5, 6, 7},
		},
	}

	filters := &schema.Set{
		F: func(v interface{}) int {
			return schema.HashString(v.(map[string]interface{})["name"])
		},
	}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"a", "b", "d"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 2 {
		t.Errorf("Expected 2 result, got %d", len(res))
	}

	numberFilter := map[string]interface{}{
		"name":   "number",
		"values": []interface{}{"1", "2", "3"},
	}
	filters.Add(numberFilter)

	res = ApplyFilters(filters, items)
	if len(res) != 0 {
		t.Errorf("Expected 0 result, got %d", len(res))
	}
	filters.Remove(numberFilter)

	booleanFilter := map[string]interface{}{
		"name":   "enabled",
		"values": []interface{}{"true", "false", "1", "0"},
	}
	filters.Add(booleanFilter)

	res = ApplyFilters(filters, items)
	if len(res) != 0 {
		t.Errorf("Expected 0 result, got %d", len(res))
	}
	filters.Remove(booleanFilter)

	intArrayFilter := map[string]interface{}{
		"name":   "nums",
		"values": []interface{}{"1", "3", "5"},
	}
	filters.Add(intArrayFilter)

	res = ApplyFilters(filters, items)
	if len(res) != 0 {
		t.Errorf("Expected 0 result, got %d", len(res))
	}
}

func TestApplyFilters_multiProperty(t *testing.T) {
	items := []map[string]interface{}{
		{
			"letter": "a",
			"number": "1",
			"symbol": "!",
		},
		{
			"letter": "b",
			"number": "2",
			"symbol": "@",
		},
		{
			"letter": "c",
			"number": "3",
			"symbol": "#",
		},
		{
			"letter": "d",
			"number": "4",
			"symbol": "$",
		},
	}

	filters := &schema.Set{
		F: func(v interface{}) int {
			return schema.HashString(v.(map[string]interface{})["name"])
		},
	}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"a", "b", "c"},
	})
	filters.Add(map[string]interface{}{
		"name":   "number",
		"values": []interface{}{"2", "3", "4"},
	})
	filters.Add(map[string]interface{}{
		"name":   "symbol",
		"values": []interface{}{"#", "$"},
	})

	res := ApplyFilters(filters, items)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}
