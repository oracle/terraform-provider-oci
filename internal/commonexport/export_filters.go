// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package commonexport

type FilterOperator string

const (
	INCLUDE FilterOperator = "="
	EXCLUDE FilterOperator = "!="
)

// Type=oci_core_vcn
// Type=oci_core_vcn, oci_core_instance
// Type!=oci_core_vcn
// Type!=oci_core_vcn, oci_core_instance
// Type=oci_core_vcn;AttrName=defined_tags;Value!=example-namespace.example-key:example-value
// Type=oci_core_vcn;AttrName=defined_tags;Value!=example-namespace.example-key:example-value,example-namespace.example-key1:example-value1
// Type=oci_core_vcn;AttrName=defined_tags;Value=example-namespace.example-key:example-value
// AttrName=defined_tags;Value=example-namespace.example-key:example-value
// AttrName=defined_tags;Value!=example-namespace.example-key:example-value
// AttrName=defined_tags;Value!=example-namespace.example-key:example-value,example-namespace.example-key1:example-value1
// Upon parsing of above combinations:
// if multiple Types are passed, they will be treated as separate Type filters
// if one type, one attribute and array values is passed, it will be one filter
// AttrName will not accept a list of names
// Type will accept one value when used with AttrName and values
// filter method returns true if the resource satisfies filter criteria with INCLUDE operator
// filter method returns true if the resource doesn't satisfy filter criteria with EXCLUDE operator
type ResourceFilter interface {
	Filter(resource *OCIResource) bool
}

type ResourceTypeFilter struct {
	ResourceType         map[string]bool
	ResourceTypeOperator FilterOperator
}

type FieldValueFilter struct {
	FieldPath            string
	ResourceTypeOperator FilterOperator
	Values               []string
}

type ResourceFieldValueFilter struct {
	ResourceType         string
	ResourceTypeOperator FilterOperator
	FieldPath            string
	Values               []string
}

// filter method returns true if the resource satisfies filter criteria with INCLUDE operator
// filter method returns true if the resource doesn't satisfy filter criteria with EXCLUDE operator
// If filter is not applicable, then also filter returns true
func (rtf *ResourceTypeFilter) Filter(resource *OCIResource) bool {
	// since resource type is empty, we don't have enough information for filtering
	if rtf.ResourceType == nil || resource == nil {
		return false
	}

	if rtf.ResourceTypeOperator == INCLUDE {
		if _, ok := rtf.ResourceType[resource.TerraformClass]; ok {
			return true
		}
	}

	if rtf.ResourceTypeOperator == EXCLUDE {
		if _, ok := rtf.ResourceType[resource.TerraformClass]; !ok {
			return true
		}
	}
	return false
}

// filter method returns true if the resource satisfies filter criteria with INCLUDE operator
// filter method returns true if the resource doesn't satisfy filter criteria with EXCLUDE operator
func (fv *FieldValueFilter) Filter(resource *OCIResource) bool {
	if resource == nil {
		return false
	}

	vals := WalkAndGet(fv.FieldPath, resource.SourceAttributes)
	// if no values are found
	if len(vals) == 0 {
		return false
	}

	switch fv.ResourceTypeOperator {
	case EXCLUDE:
		for _, val := range vals {
			for _, unacceptableValue := range fv.Values {
				if val == unacceptableValue {
					return false
				}
			}
		}
		return true
	case INCLUDE:
		for _, val := range vals {
			for _, acceptableValue := range fv.Values {
				if val == acceptableValue {
					return true
				}
			}
		}
	}
	return false
}

// this filter will support only top level fields in a resource schema and limited to `defined_tags` `time_created`..
func (rfv *ResourceFieldValueFilter) Filter(resource *OCIResource) bool {
	// TODO:implement this filter
	return false
}
