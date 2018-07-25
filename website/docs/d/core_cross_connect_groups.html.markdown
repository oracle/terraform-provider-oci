---
layout: "oci"
page_title: "OCI: oci_core_cross_connect_groups"
sidebar_current: "docs-oci-datasource-core-cross_connect_groups"
description: |-
  Provides a list of CrossConnectGroups
---

# Data Source: oci_core_cross_connect_groups
The `oci_core_cross_connect_groups` data source allows access to the list of OCI cross_connect_groups

Lists the cross-connect groups in the specified compartment.


## Example Usage

```hcl
data "oci_core_cross_connect_groups" "test_cross_connect_groups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.cross_connect_group_display_name}"
	state = "${var.cross_connect_group_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive. 


## Attributes Reference

The following attributes are exported:

* `cross_connect_groups` - The list of cross_connect_groups.

### CrossConnectGroup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cross-connect group.
* `display_name` - The display name of A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The cross-connect group's Oracle ID (OCID).
* `state` - The cross-connect group's current state.
* `time_created` - The date and time the cross-connect group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

