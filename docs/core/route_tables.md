# oci_core_route_table

## RouteTable Resource

For more information on configuring a VCN's default route table, see [Managing Default VCN Resources](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Managing%20Default%20Resources.md)

### RouteTable Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the route table.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The route table's Oracle ID (OCID).
* `route_rules` - The collection of rules for routing destination IPs to network devices.
	* `cidr_block` - A destination IP address range in CIDR notation. Matching packets will be routed to the indicated network entity (the target).  Example: `0.0.0.0/0` 
	* `network_entity_id` - The OCID for the route rule's target. For information about the type of targets you can specify, see [Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm). 
* `state` - The route table's current state.
* `time_created` - The date and time the route table was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the route table list belongs to.



### Create Operation
Creates a new route table for the specified VCN. In the request you must also include at least one route
rule for the new route table. For information on the number of rules you can have in a route table, see
[Service Limits](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/servicelimits.htm). For general information about route
tables in your VCN and the types of targets you can use in route rules,
see [Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want the route
table to reside. Notice that the route table doesn't have to be in the same compartment as the VCN, subnets,
or other Networking Service components. If you're not sure which compartment to use, put the route
table in the same compartment as the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the route table, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the route table.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `route_rules` - (Required) The collection of rules used for routing destination IPs to network devices.
	* `cidr_block` - (Required) A destination IP address range in CIDR notation. Matching packets will be routed to the indicated network entity (the target).  Example: `0.0.0.0/0` 
	* `network_entity_id` - (Required) The OCID for the route rule's target. For information about the type of targets you can specify, see [Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm). 
* `vcn_id` - (Required) The OCID of the VCN the route table belongs to.


### Update Operation
Updates the specified route table's display name or route rules.
Avoid entering confidential information.

Note that the `routeRules` object you provide replaces the entire existing set of rules.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `route_rules` - The collection of rules used for routing destination IPs to network devices.
	* `cidr_block` - A destination IP address range in CIDR notation. Matching packets will be routed to the indicated network entity (the target).  Example: `0.0.0.0/0` 
	* `network_entity_id` - The OCID for the route rule's target. For information about the type of targets you can specify, see [Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		cidr_block = "${var.route_table_route_rules_cidr_block}"
		network_entity_id = "${oci_core_internet_gateway.test_internet_gateway.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.route_table_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

# oci_core_route_tables

## RouteTable DataSource

Gets a list of route_tables.

### List Operation
Lists the route tables in the specified VCN and specified compartment. The response
includes the default route table that automatically comes with each VCN, plus any route tables
you've created.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `vcn_id` - (Required) The OCID of the VCN.


The following attributes are exported:

* `route_tables` - The list of route_tables.

### Example Usage

```hcl
data "oci_core_route_tables" "test_route_tables" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.route_table_display_name}"
	state = "${var.route_table_state}"
}
```