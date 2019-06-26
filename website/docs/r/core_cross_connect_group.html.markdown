---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect_group"
sidebar_current: "docs-oci-resource-core-cross_connect_group"
description: |-
  Provides the Cross Connect Group resource in Oracle Cloud Infrastructure Core service
---

# oci_core_cross_connect_group
This resource provides the Cross Connect Group resource in Oracle Cloud Infrastructure Core service.

Creates a new cross-connect group to use with Oracle Cloud Infrastructure
FastConnect. For more information, see
[FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).

For the purposes of access control, you must provide the OCID of the
compartment where you want the cross-connect group to reside. If you're
not sure which compartment to use, put the cross-connect group in the
same compartment with your VCN. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the cross-connect group.
It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_cross_connect_group" "test_cross_connect_group" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	customer_reference_name = "${var.cross_connect_group_customer_reference_name}"
	display_name = "${var.cross_connect_group_display_name}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment to contain the cross-connect group.
* `customer_reference_name` - (Optional) (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect group uses. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cross-connect group.
* `customer_reference_name` - A reference name or identifier for the physical fiber connection that this cross-connect group uses. 
* `display_name` - The display name of a user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The cross-connect group's Oracle ID (OCID).
* `state` - The cross-connect group's current state.
* `time_created` - The date and time the cross-connect group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Import

CrossConnectGroups can be imported using the `id`, e.g.

```
$ terraform import oci_core_cross_connect_group.test_cross_connect_group "id"
```

