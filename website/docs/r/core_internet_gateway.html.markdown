---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_internet_gateway"
sidebar_current: "docs-oci-resource-core-internet_gateway"
description: |-
  Provides the Internet Gateway resource in Oracle Cloud Infrastructure Core service
---

# oci_core_internet_gateway
This resource provides the Internet Gateway resource in Oracle Cloud Infrastructure Core service.

Creates a new internet gateway for the specified VCN. For more information, see
[Access to the Internet](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIGs.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want the Internet
Gateway to reside. Notice that the internet gateway doesn't have to be in the same compartment as the VCN or
other Networking Service components. If you're not sure which compartment to use, put the Internet
Gateway in the same compartment with the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the internet gateway, otherwise a default is provided. It
does not have to be unique, and you can change it. Avoid entering confidential information.

For traffic to flow between a subnet and an internet gateway, you must create a route rule accordingly in
the subnet's route table (for example, 0.0.0.0/0 > internet gateway). See
[UpdateRouteTable](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/RouteTable/UpdateRouteTable).

You must specify whether the internet gateway is enabled when you create it. If it's disabled, that means no
traffic will flow to/from the internet even if there's a route rule that enables that traffic. You can later
use [UpdateInternetGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/InternetGateway/UpdateInternetGateway) to easily disable/enable
the gateway without changing the route rule.


## Example Usage

```hcl
resource "oci_core_internet_gateway" "test_internet_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	enabled = "${var.internet_gateway_enabled}"
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.internet_gateway_display_name}"
	enabled = "${var.internet_gateway_enabled}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the internet gateway.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `enabled` - (Optional) (Updatable) Whether the gateway is enabled upon creation.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `vcn_id` - (Required) The OCID of the VCN the internet gateway is attached to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the internet gateway.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `enabled` - Whether the gateway is enabled. When the gateway is disabled, traffic is not routed to/from the Internet, regardless of route rules. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The internet gateway's Oracle ID (OCID).
* `state` - The internet gateway's current state.
* `time_created` - The date and time the internet gateway was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the internet gateway belongs to.

## Import

InternetGateways can be imported using the `id`, e.g.

```
$ terraform import oci_core_internet_gateway.test_internet_gateway "id"
```

