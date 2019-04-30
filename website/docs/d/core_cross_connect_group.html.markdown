---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect_group"
sidebar_current: "docs-oci-datasource-core-cross_connect_group"
description: |-
  Provides details about a specific Cross Connect Group in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cross_connect_group
This data source provides details about a specific Cross Connect Group resource in Oracle Cloud Infrastructure Core service.

Gets the specified cross-connect group's information.

## Example Usage

```hcl
data "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	cross_connect_group_id = "${oci_core_cross_connect_group.test_cross_connect_group.id}"
}
```

## Argument Reference

The following arguments are supported:

* `cross_connect_group_id` - (Required) The OCID of the cross-connect group.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cross-connect group.
* `customer_reference_name` - A reference name or identifier for the physical fiber connection that this cross-connect group uses. 
* `display_name` - The display name of a user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The cross-connect group's Oracle ID (OCID).
* `state` - The cross-connect group's current state.
* `time_created` - The date and time the cross-connect group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

