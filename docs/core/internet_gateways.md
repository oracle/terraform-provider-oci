# oci_core_internet_gateway

## InternetGateway Resource

### InternetGateway Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the Internet Gateway.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `enabled` - Whether the gateway is enabled. When the gateway is disabled, traffic is not routed to/from the Internet, regardless of route rules. Defaults to true. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The Internet Gateway's Oracle ID (OCID).
* `state` - The Internet Gateway's current state.
* `time_created` - The date and time the Internet Gateway was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the Internet Gateway belongs to.



### Create Operation
Creates a new Internet Gateway for the specified VCN. For more information, see
[Connectivity to the Internet](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIGs.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want the Internet
Gateway to reside. Notice that the Internet Gateway doesn't have to be in the same compartment as the VCN or
other Networking Service components. If you're not sure which compartment to use, put the Internet
Gateway in the same compartment with the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the Internet Gateway, otherwise a default is provided. It
does not have to be unique, and you can change it. Avoid entering confidential information.

For traffic to flow between a subnet and an Internet Gateway, you must create a route rule accordingly in
the subnet's route table (for example, 0.0.0.0/0 > Internet Gateway). See
[UpdateRouteTable](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/UpdateRouteTable).

You must specify whether the Internet Gateway is enabled when you create it. If it's disabled, that means no
traffic will flow to/from the internet even if there's a route rule that enables that traffic. You can later
use [UpdateInternetGateway](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/UpdateInternetGateway) to easily disable/enable
the gateway without changing the route rule.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the Internet Gateway.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `enabled` - (Optional) Whether the gateway is enabled upon creation. If this argument is not specified, the gateway will be enabled by default.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `vcn_id` - (Required) The OCID of the VCN the Internet Gateway is attached to.


### Update Operation
Updates the specified Internet Gateway. You can disable/enable it, or change its display name
or tags. Avoid entering confidential information.

If the gateway is disabled, that means no traffic will flow to/from the internet even if there's
a route rule that enables that traffic.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `enabled` - Whether the gateway is enabled upon creation.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_internet_gateway" "test_internet_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	enabled = "${var.internet_gateway_enabled}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.internet_gateway_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

# oci_core_internet_gateways

## InternetGateway DataSource

Gets a list of internet_gateways.

### List Operation
Lists the Internet Gateways in the specified VCN and the specified compartment.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `vcn_id` - (Required) The OCID of the VCN.


The following attributes are exported:

* `gateways` - The list of internet_gateways.

### Example Usage

```hcl
data "oci_core_internet_gateways" "test_internet_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.internet_gateway_display_name}"
	state = "${var.internet_gateway_state}"
}
```