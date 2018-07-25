---
layout: "oci"
page_title: "OCI: oci_core_cross_connects"
sidebar_current: "docs-oci-datasource-core-cross_connects"
description: |-
  Provides a list of CrossConnects
---

# Data Source: oci_core_cross_connects
The CrossConnects data source allows access to the list of OCI cross_connects

Lists the cross-connects in the specified compartment. You can filter the list
by specifying the OCID of a cross-connect group.


## Example Usage

```hcl
data "oci_core_cross_connects" "test_cross_connects" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cross_connect_group_id = "${oci_core_cross_connect_group.test_cross_connect_group.id}"
	display_name = "${var.cross_connect_display_name}"
	state = "${var.cross_connect_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `cross_connect_group_id` - (Optional) The OCID of the cross-connect group.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive. 


## Attributes Reference

The following attributes are exported:

* `cross_connects` - The list of cross_connects.

### CrossConnect Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cross-connect group.
* `cross_connect_group_id` - The OCID of the cross-connect group this cross-connect belongs to (if any).
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The cross-connect's Oracle ID (OCID).
* `location_name` - The name of the FastConnect location where this cross-connect is installed.
* `port_name` - A string identifying the meet-me room port for this cross-connect.
* `port_speed_shape_name` - The port speed for this cross-connect.  Example: `10 Gbps` 
* `state` - The cross-connect's current state.
* `time_created` - The date and time the cross-connect was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

