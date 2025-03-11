---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_type_group"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_type_group"
description: |-
  Provides details about a specific Sensitive Type Group in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_type_group
This data source provides details about a specific Sensitive Type Group resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified sensitive type group.


## Example Usage

```hcl
data "oci_data_safe_sensitive_type_group" "test_sensitive_type_group" {
	#Required
	sensitive_type_group_id = oci_data_safe_sensitive_type_group.test_sensitive_type_group.id
}
```

## Argument Reference

The following arguments are supported:

* `sensitive_type_group_id` - (Required) The OCID of the sensitive type group.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the sensitive type group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the sensitive type group.
* `display_name` - The display name of the sensitive type group.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the sensitive type group.
* `sensitive_type_count` - The number of sensitive types in the specified sensitive type group.
* `state` - The current state of the sensitive type group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the sensitive type group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the sensitive type group was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

