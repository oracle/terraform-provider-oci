// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

// Not supplying filters should not restrict results
// issue-routing-tag: terraform/default
func TestUnitApplyFilters_passThrough(t *testing.T) {
	items := []map[string]interface{}{
		{},
		{},
		{},
	}
	testSchema := map[string]*schema.Schema{
		"letter": {
			Type: schema.TypeString,
		},
	}

	res := ApplyFilters(nil, items, testSchema)
	if len(res) != 3 {
		t.Errorf("Expected 3 results, got %d", len(res))
	}
}

// Filtering against a nonexistent property should throw no errors and return no results
// issue-routing-tag: terraform/default
func TestUnitApplyFilters_nonExistentProperty(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
	}
	testSchema := map[string]*schema.Schema{
		"letter": {
			Type: schema.TypeString,
		},
	}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "number",
		"values": []interface{}{"1"},
	})

	res := ApplyFilters(filters, items, testSchema)
	if len(res) > 0 {
		t.Errorf("Expected 0 results, got %d", len(res))
	}
}

// Filtering against an empty resource set should not throw errors
// issue-routing-tag: terraform/default
func TestUnitApplyFilters_noResources(t *testing.T) {
	items := []map[string]interface{}{}

	testSchema := map[string]*schema.Schema{
		"number": {
			Type: schema.TypeString,
		},
	}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "number",
		"values": []interface{}{"1"},
	})

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 0 {
		t.Errorf("Expected 0 results, got %d", len(res))
	}
}

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_basic(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "b"},
		{"letter": "c"},
	}

	testSchema := map[string]*schema.Schema{
		"letter": {
			Type: schema.TypeString,
		},
	}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"b"},
	})

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_duplicates(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "a"},
		{"letter": "c"},
	}
	testSchema := map[string]*schema.Schema{
		"letter": {
			Type: schema.TypeString,
		},
	}

	filters := &schema.Set{F: func(v interface{}) int {
		return schema.HashString(v.(map[string]interface{})["name"])
	}}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"a"},
	})

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 2 {
		t.Errorf("Expected 2 results, got %d", len(res))
	}
}

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_OR(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "b"},
		{"letter": "c"},
	}

	testSchema := map[string]*schema.Schema{
		"letter": {
			Type: schema.TypeString,
		},
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

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 2 {
		t.Errorf("Expected 2 results, got %d", len(res))
	}
}

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_cascadeAND(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "b"},
		{"letter": "c"},
	}
	testSchema := map[string]*schema.Schema{
		"letter": {
			Type: schema.TypeString,
		},
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

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 0 {
		t.Errorf("Expected 0 results, got %d", len(res))
	}
}

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_regex(t *testing.T) {
	items := []map[string]interface{}{
		{"string": "xblx:PHX-AD-1"},
		{"string": "xblx:PHX-AD-2"},
		{"string": "xblx:PHX-AD-3"},
	}

	testSchema := map[string]*schema.Schema{
		"string": {
			Type: schema.TypeString,
		},
	}

	filters := &schema.Set{F: func(v interface{}) int {
		return schema.HashString(v.(map[string]interface{})["name"])
	}}
	filters.Add(map[string]interface{}{
		"name":   "string",
		"values": []interface{}{"\\w*:PHX-AD-2"},
		"regex":  true,
	})

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}

// Filters should test against an array of strings
// issue-routing-tag: terraform/default
func TestUnitApplyFilters_arrayOfStrings(t *testing.T) {
	items := []map[string]interface{}{
		{"letters": []string{"a"}},
		{"letters": []string{"b", "c"}},
		{"letters": []string{"c", "d", "e"}},
		{"letters": []string{"e", "f"}},
	}

	testSchema := map[string]*schema.Schema{
		"letters": {
			Type: schema.TypeList,
			Elem: schema.TypeString,
		},
	}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "letters",
		"values": []interface{}{"a", "c"},
	})

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 3 {
		t.Errorf("Expected 3 result, got %d", len(res))
	}

	filters = &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "letters",
		"values": []interface{}{"a", "f"},
	})

	res = ApplyFilters(filters, items, testSchema)
	if len(res) != 2 {
		t.Errorf("Expected 2 result, got %d", len(res))
	}
}

type CustomStringTypeA string
type CustomStringTypeB CustomStringTypeA
type CustomEnumType oci_core.VcnLifecycleStateEnum

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_underlyingStringTypes(t *testing.T) {
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

	testSchema := map[string]*schema.Schema{
		"letters": {
			Type: schema.TypeList,
			Elem: schema.TypeString,
		},
		"number": {
			Type: schema.TypeString,
		},
		"state": {
			Type: schema.TypeString,
		},
		"custom": {
			Type: schema.TypeString,
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

	res := ApplyFilters(filters, items, testSchema)
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

	res = ApplyFilters(filters1, items, testSchema)
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

	res = ApplyFilters(filters2, items, testSchema)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}

// Test fields that aren't supported: list of non-strings or structured objects
// issue-routing-tag: terraform/default
func TestUnitApplyFilters_unsupportedTypes(t *testing.T) {
	items := []map[string]interface{}{
		{
			"nums": []int{1, 2, 3},
		},
		{
			"nums": []int{3, 4, 5},
		},
		{
			"nums": []int{5, 6, 7},
		},
	}

	testSchema := map[string]*schema.Schema{
		"nums": {
			Type: schema.TypeList,
			Elem: schema.TypeInt,
		},
	}

	filters := &schema.Set{
		F: func(v interface{}) int {
			return schema.HashString(v.(map[string]interface{})["name"])
		},
	}

	intArrayFilter := map[string]interface{}{
		"name":   "nums",
		"values": []interface{}{"1", "3", "5"},
	}
	filters.Add(intArrayFilter)

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 0 {
		t.Errorf("Expected 0 result, got %d", len(res))
	}
}

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_booleanTypes(t *testing.T) {
	items := []map[string]interface{}{
		{
			"enabled": true,
		},
		{
			"enabled": "true",
		},
		{
			"enabled": "1",
		},
		{
			"enabled": false,
		},
		{
			"enabled": "false",
		},
		{
			"enabled": "0",
		},
	}

	testSchema := map[string]*schema.Schema{
		"enabled": {
			Type: schema.TypeBool,
		},
	}

	filters := &schema.Set{
		F: func(v interface{}) int {
			return schema.HashString(v.(map[string]interface{})["name"])
		},
	}

	truthyBooleanFilter := map[string]interface{}{
		"name":   "enabled",
		"values": []interface{}{"true", "1"}, // while we can pass an actual boolean true here in the test, terraform
		// doesnt, so keep coercion logic simple in filters.go
	}
	filters.Add(truthyBooleanFilter)

	res := ApplyFilters(filters, items, testSchema)

	for _, i := range res {
		switch enabled := i["enabled"].(type) {
		case bool:
			if !enabled {
				t.Errorf("Expected a truthy value, got %t", enabled)
			}
		case string:
			enabledBool, _ := strconv.ParseBool(enabled)
			if !enabledBool {
				t.Errorf("Expected a truthy value, got %s", enabled)
			}
		}
	}

	if len(res) != 3 {
		t.Errorf("Expected 3 results, got %d", len(res))
	}
	filters.Remove(truthyBooleanFilter)

	falsyBooleanFilter := map[string]interface{}{
		"name":   "enabled",
		"values": []interface{}{"false", "0"},
	}
	filters.Add(falsyBooleanFilter)

	res = ApplyFilters(filters, items, testSchema)

	for _, i := range res {
		switch enabled := i["enabled"].(type) {
		case bool:
			if enabled {
				t.Errorf("Expected a falsy value, got %t", enabled)
			}
		case string:
			enabledBool, _ := strconv.ParseBool(enabled)
			if enabledBool {
				t.Errorf("Expected a falsy value, got %s", enabled)
			}
		}
	}

	if len(res) != 3 {
		t.Errorf("Expected 3 results, got %d", len(res))
	}
	filters.Remove(falsyBooleanFilter)
}

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_numberTypes(t *testing.T) {
	items := []map[string]interface{}{
		{
			"integer": 1,
			"float":   1.1,
		},
		{
			"integer": 2,
			"float":   2.2,
		},
		{
			"integer": 3,
			"float":   3.3,
		},
	}

	testSchema := map[string]*schema.Schema{
		"integer": {
			Type: schema.TypeInt,
		},
		"float": {
			Type: schema.TypeFloat,
		},
	}

	filters := &schema.Set{
		F: func(v interface{}) int {
			return schema.HashString(v.(map[string]interface{})["name"])
		},
	}

	// int filter with single target value
	intFilter := map[string]interface{}{
		"name":   "integer",
		"values": []interface{}{"2"},
	}
	filters.Add(intFilter)

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
	filters.Remove(intFilter)

	// test filter with multiple target value
	intsFilter := map[string]interface{}{
		"name":   "integer",
		"values": []interface{}{"1", "3"},
	}
	filters.Add(intsFilter)

	res = ApplyFilters(filters, items, testSchema)
	if len(res) != 2 {
		t.Errorf("Expected 2 results, got %d", len(res))
	}
	filters.Remove(intsFilter)

	// test float filter
	floatFilter := map[string]interface{}{
		"name":   "float",
		"values": []interface{}{"1.1", "3.3"},
	}
	filters.Add(floatFilter)

	res = ApplyFilters(filters, items, testSchema)
	if len(res) != 2 {
		t.Errorf("Expected 2 results, got %d", len(res))
	}
	filters.Remove(floatFilter)
}

// issue-routing-tag: terraform/default
func TestUnitApplyFilters_multiProperty(t *testing.T) {
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

	testSchema := map[string]*schema.Schema{
		"letter": {
			Type: schema.TypeString,
		},
		"number": {
			Type: schema.TypeInt,
		},
		"symbol": {
			Type: schema.TypeString,
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

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
}

// Test to validate that the Apply filters do not impact the original item order
// issue-routing-tag: terraform/default
func TestUnitApplyFilters_ElementOrder(t *testing.T) {
	items := []map[string]interface{}{
		{"letter": "a"},
		{"letter": "b"},
		{"letter": "c"},
		{"letter": "d"},
	}

	testSchema := map[string]*schema.Schema{
		"letter": {
			Type: schema.TypeString,
		},
	}

	filters := &schema.Set{F: func(interface{}) int { return 1 }}
	filters.Add(map[string]interface{}{
		"name":   "letter",
		"values": []interface{}{"a", "d"},
	})

	res := ApplyFilters(filters, items, testSchema)
	if len(res) != 2 {
		t.Errorf("Expected 2 result, got %d", len(res))
	}
	if res[0]["letter"] != "a" || res[1]["letter"] != "d" {
		t.Errorf("Expected sort order not retained, got %s %s", res[0]["letter"], res[1]["letter"])
	}

}

// issue-routing-tag: terraform/default
func TestUnitGetValue_EmptyMap(t *testing.T) {
	item := map[string]interface{}{}

	_, singleLevelGetOk := getValueFromPath(item, []string{"path"})
	_, multiLevelGetOk := getValueFromPath(item, []string{"path", "to", "target"})

	if singleLevelGetOk || multiLevelGetOk {
		t.Error("Expected non OK result")
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetValue_MultiLevelMap(t *testing.T) {
	item := map[string]interface{}{
		"level1": map[string]interface{}{
			"level2": map[string]interface{}{
				"level3": "value",
			},
		},
	}

	singleLevelGet, singleLevelGetOk := getValueFromPath(item, []string{"level1"})
	multiLevelGet, multiLevelGetOk := getValueFromPath(item, []string{"level1", "level2", "level3"})

	if !singleLevelGetOk || !multiLevelGetOk {
		t.Errorf("Expected OK result for topLevel %v multi level %v", singleLevelGetOk, multiLevelGetOk)
	}

	if multiLevelGet != "value" {
		t.Errorf("Expected = value, Got = %s", multiLevelGet)
	}

	if len(singleLevelGet.(map[string]interface{})) != 1 {
		t.Error("Expected size of map is 1")
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetPathElements_EmptyFilterName(t *testing.T) {
	if _, error := getFieldPathElements(CoreInstanceResource().Schema, ""); error == nil {
		t.Error("expected non nil error")
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetPathElements_NonExistentPropertyTopLevel(t *testing.T) {
	if _, error := getFieldPathElements(CoreInstanceResource().Schema, "non_existent"); error == nil {
		t.Error("expected non nil error")
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetPathElements_NonExistentPropertyNestedLevel(t *testing.T) {
	if _, error := getFieldPathElements(CoreInstanceResource().Schema, "create_vnic_details.non_existent"); error == nil {
		t.Error("expected non nil error")
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetPathElements_TopLevelPrimitive(t *testing.T) {
	if path, error := getFieldPathElements(CoreInstanceResource().Schema, "boot_volume_id"); error != nil || !reflect.DeepEqual(path, []string{"boot_volume_id"}) {
		t.Errorf("unexpected path value %s found", path)
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetPathElements_MultiLevelMap(t *testing.T) {
	if path, error := getFieldPathElements(CoreInstanceResource().Schema, "create_vnic_details.defined_tags.namespace.key"); error != nil || !reflect.DeepEqual(path, []string{"create_vnic_details", "defined_tags", "namespace.key"}) {
		t.Errorf("unexpected path value %s found", path)
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetPathElements_MultiLevelNonMap(t *testing.T) {
	if path, error := getFieldPathElements(CoreInstanceResource().Schema, "launch_options.firmware"); error != nil || !reflect.DeepEqual(path, []string{"launch_options", "firmware"}) {
		t.Errorf("unexpected path value %s found", path)
	}
	if _, error := getFieldPathElements(CoreInstanceResource().Schema, "launch_options.firmware.XYZ"); error == nil {
		t.Errorf("Expected Error")
	}
}

// issue-routing-tag: terraform/default
func TestUnitNestedMap(t *testing.T) {
	item := map[string]interface{}{
		"level1": map[string]interface{}{
			"level2": map[string]interface{}{
				"level3":   "value",
				"level3_1": []string{"A", "B", "C"},
				"level3_2": []int{2, 3, 4},
				"level3_3": []float64{2.1, 3.1, 4.1},
				"level3_4": []interface{}{2, 3, 4},
			},
		},
	}
	services := GenericMapToJsonMap(item)

	if len(services) != 1 {
		t.Errorf("unexpected number of values returned in map")
	}
}

// This is for unit test input only
func CoreInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"async": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"create_vnic_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"assign_private_dns_record": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"assign_public_ip": {
							// Change type from boolean to string because TF doesn't handle default
							// values for boolean nested objects correctly.
							Type:     schema.TypeString,
							Optional: true,
							Default:  "true",
							ValidateFunc: func(v interface{}, k string) ([]string, []error) {
								// Verify that we can parse the string value as a bool value.
								var es []error
								if _, err := strconv.ParseBool(v.(string)); err != nil {
									es = append(es, fmt.Errorf("%s: cannot parse 'assign_public_ip' as bool: %v", k, err))
								}
								return nil, es
							},
						},
						"defined_tags": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: DefinedTagsDiffSuppressFunction,
							Elem:             schema.TypeString,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"skip_source_dest_check": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			"agent_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"are_all_plugins_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_management_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_monitoring_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"plugins_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"desired_state": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"availability_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_live_migration_preferred": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"recovery_action": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"capacity_reservation_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dedicated_vm_host_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"instance_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"are_legacy_imds_endpoints_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"ipxe_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_pv_encryption_in_transit_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"launch_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"boot_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"firmware": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_consistent_volume_naming_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remote_data_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_maintenance_reboot_due": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Legacy custom computed convenience values
			"public_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Computed
			// Add this computed boot_volume_id field even though it's not part of the API specs. This will make it easier to
			// discover the attached boot volume's ID; to preserve it for reattachment.
			"boot_volume_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
