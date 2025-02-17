---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_type_group_grouped_sensitive_types"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_type_group_grouped_sensitive_types"
description: |-
  Provides the list of Sensitive Type Group Grouped Sensitive Types in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_type_group_grouped_sensitive_types
This data source provides the list of Sensitive Type Group Grouped Sensitive Types in Oracle Cloud Infrastructure Data Safe service.

Gets the list of sensitive type Ids present in the specified sensitive type group.


## Example Usage

```hcl
data "oci_data_safe_sensitive_type_group_grouped_sensitive_types" "test_sensitive_type_group_grouped_sensitive_types" {
	#Required
	sensitive_type_group_id = oci_data_safe_sensitive_type_group.test_sensitive_type_group.id

	#Optional
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
}
```

## Argument Reference

The following arguments are supported:

* `sensitive_type_group_id` - (Required) The OCID of the sensitive type group.
* `sensitive_type_id` - (Optional) A filter to return only items related to a specific sensitive type OCID.


## Attributes Reference

The following attributes are exported:

* `grouped_sensitive_type_collection` - The list of grouped_sensitive_type_collection.

### SensitiveTypeGroupGroupedSensitiveType Reference

The following attributes are exported:

* `items` - List of sensitive type id summary objects present in the sensitive type group.
	* `sensitive_type_id` - The OCID of the sensitive type.

