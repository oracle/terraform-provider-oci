---
layout: "oci"
page_title: "OCI: oci_core_cross_connect_group"
sidebar_current: "docs-oci-datasource-core-cross_connect_group"
description: |-
Provides details about a specific CrossConnectGroup
---

# Data Source: oci_core_cross_connect_group
The CrossConnectGroup data source provides details about a specific CrossConnectGroup

Gets the specified cross-connect group's information.

## Example Usage

```hcl
data "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	cross_connect_group_id = "${var.cross_connect_group_cross_connect_group_id}"
}
```

## Argument Reference

The following arguments are supported:

* `cross_connect_group_id` - (Required) The OCID of the cross-connect group.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cross-connect group.
* `display_name` - The display name of A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The cross-connect group's Oracle ID (OCID).
* `state` - The cross-connect group's current state.
* `time_created` - The date and time the cross-connect group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

