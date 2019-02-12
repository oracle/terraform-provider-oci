---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect_groups"
sidebar_current: "docs-oci-datasource-core-cross_connect_groups"
description: |-
  Provides the list of Cross Connect Groups in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cross_connect_groups
This data source provides the list of Cross Connect Groups in Oracle Cloud Infrastructure Core service.

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
* `customer_reference_name` - A reference name or identifier for the physical fiber connection that this cross-connect group uses. 
* `display_name` - The display name of A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The cross-connect group's Oracle ID (OCID).
* `state` - The cross-connect group's current state.
* `time_created` - The date and time the cross-connect group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

