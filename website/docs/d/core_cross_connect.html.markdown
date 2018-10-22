---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect"
sidebar_current: "docs-oci-datasource-core-cross_connect"
description: |-
  Provides details about a specific Cross Connect in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cross_connect
This data source provides details about a specific Cross Connect resource in Oracle Cloud Infrastructure Core service.

Gets the specified cross-connect's information.

## Example Usage

```hcl
data "oci_core_cross_connect" "test_cross_connect" {
	#Required
	cross_connect_id = "${oci_core_cross_connect.test_cross_connect.id}"
}
```

## Argument Reference

The following arguments are supported:

* `cross_connect_id` - (Required) The OCID of the cross-connect.


## Attributes Reference

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

