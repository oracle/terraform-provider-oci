---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_attribute_set"
sidebar_current: "docs-oci-datasource-data_safe-attribute_set"
description: |-
  Provides details about a specific Attribute Set in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_attribute_set
This data source provides details about a specific Attribute Set resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified attribute set.

## Example Usage

```hcl
data "oci_data_safe_attribute_set" "test_attribute_set" {
	#Required
	attribute_set_id = oci_data_safe_attribute_set.test_attribute_set.id
}
```

## Argument Reference

The following arguments are supported:

* `attribute_set_id` - (Required) OCID of an attribute set.


## Attributes Reference

The following attributes are exported:

* `attribute_set_type` - The type of attribute set.
* `attribute_set_values` - The list of values in an attribute set
* `compartment_id` - The OCID of the compartment where the attribute set is stored.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of an attribute set.
* `display_name` - The display name of an attribute set. The name does not have to be unique, and is changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of an attribute set.
* `in_use` - Indicates whether the attribute set is in use by other resource.
* `is_user_defined` - A boolean flag indicating to list user defined or seeded attribute sets.
* `state` - The current state of an attribute set.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time an attribute set was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time an attribute set was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

