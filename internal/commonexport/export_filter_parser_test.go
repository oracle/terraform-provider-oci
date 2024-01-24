// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package commonexport

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnitGetFilterOperator(t *testing.T) {
	tests := []struct {
		testName     string
		filterString string
		expected     FilterOperator
		err          error
	}{
		{
			"ValidIncludeFilterOperator",
			"=",
			INCLUDE,
			nil,
		},
		{
			"ValidIncludeFilterOperator",
			"!=",
			EXCLUDE,
			nil,
		},
		{
			"InvalidIncludeFilterOperator",
			"==",
			INCLUDE,
			fmt.Errorf(""),
		},
		{
			"InvalidIncludeFilterOperator",
			"=!",
			INCLUDE,
			fmt.Errorf(""),
		},
		{
			"InvalidIncludeFilterOperator",
			"!!",
			INCLUDE,
			fmt.Errorf(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans, err := getFilterOperator(tt.filterString)

			if err != nil && tt.err == nil {
				t.Errorf("unexpected error occurred in retreiving filter operator %+v", err)
			}

			if ans != tt.expected {
				t.Errorf("got %+v, want %+v", ans, tt.expected)
			}
		})
	}

}

func TestUnitParseFilter(t *testing.T) {
	var filterFlag Filter
	tests := []struct {
		testName     string
		filterString string
		filterType   reflect.Type
		err          error
	}{
		{
			"ValidResourceTypeFilter",
			"Type=oci_core_vcn",
			reflect.TypeOf(&ResourceTypeFilter{}),
			nil,
		},
		{
			"ValidResourceTypeFilter2",
			"Type=oci_core_vcn,oci_core_instance",
			reflect.TypeOf(&ResourceTypeFilter{}),
			nil,
		},
		{
			"ValidResourceTypeFilter",
			"Type!=oci_core_vcn",
			reflect.TypeOf(&ResourceTypeFilter{}),
			nil,
		},
		{
			"ValidResourceTypeFilter3",
			"Type!=oci_core_vcn,oci_core_instance",
			reflect.TypeOf(&ResourceTypeFilter{}),
			nil,
		},
		{
			"ValidFieldValueFilter",
			"AttrName=defined_tags.example-namespace.example-key;Value=example-value",
			reflect.TypeOf(&FieldValueFilter{}),
			nil,
		},
		{
			"ValidFieldValueFilter2",
			"AttrName=defined_tags.example-namespace;Value!=example-value,example-value1",
			reflect.TypeOf(&FieldValueFilter{}),
			nil,
		},
		{
			"ValidFieldValueFilter3",
			"AttrName=defined_tags.example-namespace.example-key;Value=example-value, example-value1",
			reflect.TypeOf(&FieldValueFilter{}),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans, err := filterFlag.ParseFilter(tt.filterString)

			if err != nil && tt.err == nil {
				t.Errorf("unexpected error occurred in retreiving filter operator %+v", err)
				t.Fail()
			}

			if reflect.TypeOf(ans) != tt.filterType {
				t.Errorf("got %+v, want %+v", reflect.TypeOf(ans), tt.filterType)
			}
		})
	}
}
