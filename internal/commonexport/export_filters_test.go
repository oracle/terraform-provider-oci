// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package commonexport

import (
	"testing"
)

const (
	resourceDiscoveryTestCompartmentOcid = "ocid1.testcompartment.abc"
)

func TestUnitTypeFilterInclude(t *testing.T) {
	resourceFilter := &ResourceTypeFilter{
		ResourceTypeOperator: INCLUDE,
		ResourceType:         map[string]bool{"oci_core_vcn": true},
	}

	tests := []struct {
		testName string
		resource *OCIResource
		expected bool
	}{
		{
			"InvalidResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_resource_type1",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				}},
			false,
		},
		{
			"InvalidResourceType2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:             "ocid1.d.e.f",
					TerraformClass: "oci_resource_type2",
					TerraformName:  "type2_res1",
				},
				IsErrorResource: true},
			false,
		},
		{
			"ValidButNotFilterResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_load_balancer",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				}},
			false,
		},
		{
			"ValidButNotFilterResourceType2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_instance",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				}},
			false,
		},
		{
			"ValidFilterResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := resourceFilter.Filter(tt.resource)
			if ans != tt.expected {
				t.Errorf("got %t, want %t", ans, tt.expected)
			}
		})
	}
}

func TestUnitTypeFilterExclude(t *testing.T) {
	resourceFilter := &ResourceTypeFilter{
		ResourceTypeOperator: EXCLUDE,
		ResourceType:         map[string]bool{"oci_core_vcn": true},
	}

	tests := []struct {
		testName string
		resource *OCIResource
		expected bool
	}{
		{
			"InvalidResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_resource_type1",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
			},
			true,
		},
		{
			"InvalidResourceType2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:             "ocid1.d.e.f",
					TerraformClass: "oci_resource_type2",
					TerraformName:  "type2_res1",
				},
				IsErrorResource: true},
			true,
		},
		{
			"ValidButNotFilterResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_load_balancer",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				}},
			true,
		},
		{
			"ValidButNotFilterResourceType2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_instance",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				}},
			true,
		},
		{
			"ValidFilterResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := resourceFilter.Filter(tt.resource)
			if ans != tt.expected {
				t.Errorf("got %t, want %t", ans, tt.expected)
			}
		})
	}
}

func TestUnitMultipleValuesWithResourceTypeFilter(t *testing.T) {
	tests := []struct {
		testName  string
		resources []*OCIResource
		filter    ResourceFilter
		expected  int
	}{
		{
			"AndOperationWithExclude",
			[]*OCIResource{
				{
					CompartmentId: resourceDiscoveryTestCompartmentOcid,
					TerraformResource: TerraformResource{
						Id:                "ocid1.a.b.c",
						TerraformClass:    "oci_resource_type1",
						TerraformName:     "type1_res1",
						TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_resource_type1", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
					},
				},
				{
					CompartmentId: resourceDiscoveryTestCompartmentOcid,
					TerraformResource: TerraformResource{
						Id:                "ocid1.a.b.c",
						TerraformClass:    "oci_resource_type2",
						TerraformName:     "type1_res2",
						TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_resource_type2", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
					},
				},
			},
			&ResourceTypeFilter{
				ResourceTypeOperator: EXCLUDE,
				ResourceType:         map[string]bool{"oci_resource_type1": true, "oci_resource_type2": true},
			},
			0,
		},
		{
			"AndOperationWithExclude2",
			[]*OCIResource{
				{
					CompartmentId: resourceDiscoveryTestCompartmentOcid,
					TerraformResource: TerraformResource{
						Id:                "ocid1.a.b.c",
						TerraformClass:    "oci_resource_type1",
						TerraformName:     "type1_res1",
						TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_resource_type1", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
					},
				},
				{
					CompartmentId: resourceDiscoveryTestCompartmentOcid,
					TerraformResource: TerraformResource{
						Id:                "ocid1.a.b.c",
						TerraformClass:    "oci_resource_type2",
						TerraformName:     "type1_res2",
						TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_resource_type2", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
					},
				},
			},
			&ResourceTypeFilter{
				ResourceTypeOperator: EXCLUDE,
				ResourceType:         map[string]bool{"oci_resource_type1": true, "oci_resource_type3": true},
			},
			1,
		},
		{
			"AndOperationWithInclude",
			[]*OCIResource{
				{
					CompartmentId: resourceDiscoveryTestCompartmentOcid,
					TerraformResource: TerraformResource{
						Id:                "ocid1.a.b.c",
						TerraformClass:    "oci_resource_type1",
						TerraformName:     "type1_res1",
						TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_resource_type1", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
					},
				},
				{
					CompartmentId: resourceDiscoveryTestCompartmentOcid,
					TerraformResource: TerraformResource{
						Id:                "ocid1.a.b.c",
						TerraformClass:    "oci_resource_type2",
						TerraformName:     "type1_res2",
						TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_resource_type2", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
					},
				},
			},
			&ResourceTypeFilter{
				ResourceTypeOperator: INCLUDE,
				ResourceType:         map[string]bool{"oci_resource_type1": true, "oci_resource_type2": true},
			},
			2,
		},
		{
			"AndOperationWithInclude2",
			[]*OCIResource{
				{
					CompartmentId: resourceDiscoveryTestCompartmentOcid,
					TerraformResource: TerraformResource{
						Id:                "ocid1.a.b.c",
						TerraformClass:    "oci_resource_type1",
						TerraformName:     "type1_res1",
						TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_resource_type1", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
					},
				},
				{
					CompartmentId: resourceDiscoveryTestCompartmentOcid,
					TerraformResource: TerraformResource{
						Id:                "ocid1.a.b.c",
						TerraformClass:    "oci_resource_type2",
						TerraformName:     "type1_res2",
						TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_resource_type2", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
					},
				},
			},
			&ResourceTypeFilter{
				ResourceTypeOperator: INCLUDE,
				ResourceType:         map[string]bool{"oci_resource_type1": true, "oci_resource_type3": true},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := 0

			for _, resource := range tt.resources {
				if tt.filter.Filter(resource) {
					ans += 1
				}
			}

			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}

}

func TestUnitFieldValueFilterInclude(t *testing.T) {
	tests := []struct {
		testName string
		resource *OCIResource
		filter   ResourceFilter
		expected bool
	}{
		{
			"InvalidFieldValue1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_resource_type1",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
					},
					"display_name": "test",
				},
			},
			&FieldValueFilter{
				FieldPath:            "id",
				ResourceTypeOperator: INCLUDE,
				Values:               []string{"test"},
			},
			false,
		},
		{
			"InvalidFieldValue2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:             "ocid1.d.e.f",
					TerraformClass: "oci_resource_type2",
					TerraformName:  "type2_res1",
				},
				IsErrorResource: true,
				SourceAttributes: map[string]interface{}{
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
						"id2":          "ocid1.d.e.f",
					},
					"display_name": "test",
				},
			},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: INCLUDE,
				Values:               []string{"test1"},
			},
			false,
		},
		{
			"ValidFilterResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag": "YES",
					},
					"display_name": "test",
				}},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: INCLUDE,
				Values:               []string{"test"},
			},
			true,
		},
		{
			"ValidFilterResourceType2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": "test",
				}},
			&FieldValueFilter{
				FieldPath:            "defined_tags.myDefinedTag2",
				ResourceTypeOperator: INCLUDE,
				Values:               []string{"YES"},
			},
			true,
		},
		{
			"ValidFilterResourceTypeMultiValues1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": []string{"test", "test2"},
				}},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: INCLUDE,
				Values:               []string{"test"},
			},
			true,
		},
		{
			"ValidFilterResourceTypeMultiValues2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": []string{"test", "test2"},
				}},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: INCLUDE,
				Values:               []string{"test", "test2"},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := tt.filter.Filter(tt.resource)
			if ans != tt.expected {
				t.Errorf("got %t, want %t", ans, tt.expected)
			}
		})
	}
}

func TestUnitFieldValueFilterExclude(t *testing.T) {

	tests := []struct {
		testName string
		resource *OCIResource
		filter   ResourceFilter
		expected bool
	}{
		{
			"InvalidResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_resource_type1",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": "test",
				},
			},
			&FieldValueFilter{
				FieldPath:            "id",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"test"},
			},
			true,
		},
		{
			"InvalidResourceType2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:             "ocid1.d.e.f",
					TerraformClass: "oci_resource_type2",
					TerraformName:  "type2_res1",
				},
				IsErrorResource: true,
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": "test",
				}},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"test1"},
			},
			true,
		},
		{
			"ValidFilterResourceType1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": "test",
				}},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"test"},
			},
			false,
		},
		{
			"ValidFilterResourceType2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": "test",
				}},
			&FieldValueFilter{
				FieldPath:            "defined_tags.myDefinedTag2",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"YES"},
			},
			false,
		},
		{
			"ValidFilterResourceTypeMultiValues1",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": []string{"test", "test2"},
				}},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"test"},
			},
			false,
		},
		{
			"ValidFilterResourceTypeMultiValues2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": []string{"test", "test2"},
				}},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"test", "test2"},
			},
			false,
		},
		{
			"InvalidFilterResourceTypeMultiValues2",
			&OCIResource{
				CompartmentId: resourceDiscoveryTestCompartmentOcid,
				TerraformResource: TerraformResource{
					Id:                "ocid1.a.b.c",
					TerraformClass:    "oci_core_vcn",
					TerraformName:     "type1_res1",
					TerraformTypeInfo: &TerraformResourceHints{ResourceClass: "oci_test_parent", IgnorableRequiredMissingAttributes: map[string]bool{"test": true}},
				},
				SourceAttributes: map[string]interface{}{
					"id": "ocid1.a.b.c",
					"defined_tags": map[string]interface{}{
						"myDefinedTag":  "YES",
						"myDefinedTag2": "YES",
					},
					"display_name": []string{"test", "test2"},
				}},
			&FieldValueFilter{
				FieldPath:            "display_name",
				ResourceTypeOperator: EXCLUDE,
				Values:               []string{"test3", "test4"},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := tt.filter.Filter(tt.resource)
			if ans != tt.expected {
				t.Errorf("got %t, want %t", ans, tt.expected)
			}
		})
	}
}
