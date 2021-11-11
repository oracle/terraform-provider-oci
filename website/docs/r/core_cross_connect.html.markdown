---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect"
sidebar_current: "docs-oci-resource-core-cross_connect"
description: |-
  Provides the Cross Connect resource in Oracle Cloud Infrastructure Core service
---

# oci_core_cross_connect
This resource provides the Cross Connect resource in Oracle Cloud Infrastructure Core service.

Creates a new cross-connect. Oracle recommends you create each cross-connect in a
[CrossConnectGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CrossConnectGroup) so you can use link aggregation
with the connection.

After creating the `CrossConnect` object, you need to go the FastConnect location
and request to have the physical cable installed. For more information, see
[FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the
compartment where you want the cross-connect to reside. If you're
not sure which compartment to use, put the cross-connect in the
same compartment with your VCN. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the cross-connect.
It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_cross_connect" "test_cross_connect" {
	#Required
	compartment_id = var.compartment_id
	location_name = var.cross_connect_location_name
	port_speed_shape_name = var.cross_connect_port_speed_shape_name

	#Optional
	cross_connect_group_id = oci_core_cross_connect_group.test_cross_connect_group.id
	customer_reference_name = var.cross_connect_customer_reference_name
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.cross_connect_display_name
	far_cross_connect_or_cross_connect_group_id = oci_core_cross_connect_group.test_cross_connect_group.id
	freeform_tags = {"Department"= "Finance"}
	near_cross_connect_or_cross_connect_group_id = oci_core_cross_connect_group.test_cross_connect_group.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment to contain the cross-connect.
* `cross_connect_group_id` - (Optional) The OCID of the cross-connect group to put this cross-connect in. 
* `customer_reference_name` - (Optional) (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect uses. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `far_cross_connect_or_cross_connect_group_id` - (Optional) If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on a different router (for the purposes of redundancy), provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that existing cross-connect or cross-connect group. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `is_active` - (Optional) (Updatable) Set to true to activate the cross-connect. You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up. Activation indicates to Oracle that the physical connection is ready.
* `location_name` - (Required) The name of the FastConnect location where this cross-connect will be installed. To get a list of the available locations, see [ListCrossConnectLocations](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CrossConnectLocation/ListCrossConnectLocations).  Example: `CyrusOne, Chandler, AZ` 
* `near_cross_connect_or_cross_connect_group_id` - (Optional) If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on the same router, provide the OCID of that existing cross-connect or cross-connect group.
* `port_speed_shape_name` - (Required) The port speed for this cross-connect. To get a list of the available port speeds, see [ListCrossConnectPortSpeedShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CrossConnectPortSpeedShape/ListCrossconnectPortSpeedShapes).  Example: `10 Gbps` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cross-connect group.
* `cross_connect_group_id` - The OCID of the cross-connect group this cross-connect belongs to (if any). 
* `customer_reference_name` - A reference name or identifier for the physical fiber connection that this cross-connect uses. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The cross-connect's Oracle ID (OCID).
* `location_name` - The name of the FastConnect location where this cross-connect is installed. 
* `port_name` - A string identifying the meet-me room port for this cross-connect.
* `port_speed_shape_name` - The port speed for this cross-connect.  Example: `10 Gbps` 
* `state` - The cross-connect's current state.
* `time_created` - The date and time the cross-connect was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cross Connect
	* `update` - (Defaults to 20 minutes), when updating the Cross Connect
	* `delete` - (Defaults to 20 minutes), when destroying the Cross Connect


## Import

CrossConnects can be imported using the `id`, e.g.

```
$ terraform import oci_core_cross_connect.test_cross_connect "id"
```

