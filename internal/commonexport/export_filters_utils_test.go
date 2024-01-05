// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package commonexport

import (
	"testing"
)

func TestUnitGetFilterDeepCopy(t *testing.T) {
	tests := []struct {
		testName string
		filter   ResourceFilter
	}{
		{
			"ValidResourceTypeFilter",
			&ResourceTypeFilter{
				ResourceType:         map[string]bool{"oci_core_vcn": true},
				ResourceTypeOperator: INCLUDE,
			},
		},
		{
			"ValidResourceTypeFilter2",
			&ResourceTypeFilter{
				ResourceType:         map[string]bool{"oci_core_vcn": true, "oci_core_instance": true},
				ResourceTypeOperator: INCLUDE,
			},
		},
		{
			"ValidResourceTypeFilter3",
			&ResourceTypeFilter{
				ResourceType:         map[string]bool{"oci_core_vcn": true},
				ResourceTypeOperator: EXCLUDE,
			},
		},
		{
			"ValidFieldValueTypeFilter",
			&FieldValueFilter{
				FieldPath:            "defined_tags",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"example"},
			},
		},
		{
			"ValidResourceFiledValueTypeFilter",
			&ResourceFieldValueFilter{
				ResourceType:         "oci_core_vcn",
				FieldPath:            "defined_tags",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"example"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := getFilterDeepCopy(tt.filter)

			if err != nil {
				t.Errorf("unexpected error occurred in retreiving filter operator %+v", err)
				t.Fail()
			}
		})
	}
}

func TestUnitWalkAndCheckField(t *testing.T) {
	tests := []struct {
		testName  string
		fieldPath string
		resource  *OCIResource
		expected  bool
	}{
		{
			"TestTopLevelString1",
			"display_name",
			&OCIResource{
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
					},
					"display_name": "test",
				},
			},
			true,
		},
		{
			"TestTopLevelString2",
			"defined_tags.myDefinedTag",
			&OCIResource{
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
				},
			},
			true,
		},
		{
			"TestTopLevelString3",
			"defined_tags1",
			&OCIResource{
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
					},
					"display_name": "test",
				},
			},
			false,
		},
		{
			"TestTopLevelString4",
			"names",
			&OCIResource{
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
					},
					"display_name": "test",
					"names":        []string{"test01", "test02", "test03", "test04"},
				},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := WalkAndCheckField(tt.fieldPath, tt.resource.SourceAttributes)

			if ans != tt.expected {
				t.Errorf("got %t, want %t for field path %s", ans, tt.expected, tt.fieldPath)
			}
		})
	}
}

func TestUnitWalkAndGetField(t *testing.T) {
	tests := []struct {
		testName  string
		fieldPath string
		resource  *OCIResource
		expected  []interface{}
	}{
		{
			"TestTopLevelString1",
			"display_name",
			&OCIResource{
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
					},
					"display_name": "test",
				},
			},
			[]interface{}{"test"},
		},
		{
			"TestTopLevelString2",
			"defined_tags.myDefinedTag",
			&OCIResource{
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
				},
			},
			[]interface{}{"YES"},
		},
		{
			"TestTopLevelString3",
			"defined_tags1",
			&OCIResource{
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
					},
					"display_name": "test",
				},
			},
			nil,
		},
		{
			"TestTopLevelString4",
			"names",
			&OCIResource{
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
					},
					"display_name": "test",
					"names":        []string{"test01", "test02", "test03", "test04"},
				},
			},
			[]interface{}{"test01", "test02", "test03", "test04"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := WalkAndGet(tt.fieldPath, tt.resource.SourceAttributes)

			if len(ans) != len(tt.expected) {
				t.Errorf("got %t, want %t for field path %s", ans, tt.expected, tt.fieldPath)
			}
		})
	}
}
