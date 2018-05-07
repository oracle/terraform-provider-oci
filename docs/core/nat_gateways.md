# oci_core_nat_gateway

## NatGateway Resource

### NatGateway Reference

The following attributes are exported:

* `block_traffic` - Whether the NAT gateway blocks traffic through it. The default is `false`.  Example: `false` 
* `compartment_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the NAT gateway. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the NAT gateway.
* `nat_ip` - The IP address associated with the NAT gateway. 
* `state` - The NAT gateway's current state.
* `time_created` - The date and time the NAT gateway was created, in the format defined by RFC3339.  Example: '2016-08-25T21:10:29.600Z' 
* `vcn_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VCN the NAT gateway belongs to. 



### Create Operation
Creates a new NAT gateway for the specified VCN. You must also set up a route rule with the
NAT gateway as the rule's target. See [Route Table](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/).


The following arguments are supported:

* `block_traffic` - (Optional) Whether the NAT gateway blocks traffic through it. The default is `false`.  Example: `false` 
* `compartment_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the NAT gateway. 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `vcn_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VCN the gateway belongs to. 


### Update Operation
Updates the specified NAT gateway.


The following arguments support updates:
* `block_traffic` - Whether the NAT gateway blocks traffic through it. The default is `false`.  Example: `false` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_core_nat_gateway" "test_nat_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	block_traffic = "${var.nat_gateway_block_traffic}"
	display_name = "${var.nat_gateway_display_name}"
}
```

# oci_core_nat_gateways

## NatGateway DataSource

Gets a list of nat_gateways.

### List Operation
Lists the NAT gateways in the specified compartment. You may optionally specify a VCN OCID
to filter the results by VCN.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 
* `vcn_id` - (Optional) The OCID of the VCN.


The following attributes are exported:

* `nat_gateways` - The list of nat_gateways.

### Example Usage

```
data "oci_core_nat_gateways" "test_nat_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.nat_gateway_display_name}"
	state = "${var.nat_gateway_state}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```