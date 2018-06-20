# oci_core_local_peering_gateway

## LocalPeeringGateway Resource

### LocalPeeringGateway Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the Local Peering Gateway (LPG).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The LPG's Oracle ID (OCID).
* `is_cross_tenancy_peering` - Whether the VCN at the other end of the peering is in a different tenancy.  Example: `false` 
* `peer_advertised_cidr` - The range of IP addresses available on the VCN at the other end of the peering from this LPG. The value is `null` if the LPG is not peered. You can use this as the destination CIDR for a route rule to route a subnet's traffic to this LPG.  Example: `192.168.0.0/16` 
* `peering_status` - Whether the LPG is peered with another LPG. `NEW` means the LPG has not yet been peered. `PENDING` means the peering is being established. `REVOKED` means the LPG at the other end of the peering has been deleted. 
* `peering_status_details` - Additional information regarding the peering status, if applicable.
* `state` - The LPG's current lifecycle state.
* `time_created` - The date and time the LPG was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the LPG belongs to.



### Create Operation
Creates a new local peering gateway (LPG) for the specified VCN.

* Specifying a peer_id creates a connection to the specified LPG ID. 
* If the specified peer_id is also a resource in the terraform config you will have do a `terraform refresh` after the `terraform apply` in order to get the latest connection information on that resource.
* To disconnect the peering connection at least one of the LPG resources in the connection will have to be destroyed, however in terraform we recommend that when one LPG is destroyed the peer should also be destroyed. If one of them is not destroyed it will have a `REVOKED` peering_status. If another LPG resource tries to connect to this LPG resource it will get a `400 Error: The Local Peering Gateway with ID X has already been connected`. To solve this you will have to run `terraform taint oci_core_local_peering_gateway.test_local_peering_gateway` on that resource or target delete it `terraform destroy -target="oci_core_local_peering_gateway.test_local_peering_gateway"`.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the local peering gateway (LPG).
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `peer_id` - (Optional) The OCID of the LPG you want to peer with. Specifying a peer_id connects this local peering gateway (LPG) to another one in the same region. This operation must be called by the VCN administrator who is designated as the *requestor* in the peering relationship. The *acceptor* must implement an Identity and Access Management (IAM) policy that gives the requestor permission to connect to LPGs in the acceptor's compartment. Without that permission, this operation will fail. For more information, see [VCN Peering](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/VCNpeering.htm).
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `vcn_id` - (Required) The OCID of the VCN the LPG belongs to.


### Update Operation
Updates the specified local peering gateway (LPG).


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_local_peering_gateway" "test_local_peering_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.local_peering_gateway_display_name}"
	freeform_tags = {"Department"= "Finance"}
	peer_id = "${oci_core_local_peering_gateway.test_local_peering_gateway2.id}"
}
```

# oci_core_local_peering_gateways

## LocalPeeringGateway DataSource

Gets a list of local_peering_gateways.

### List Operation
Lists the local peering gateways (LPGs) for the specified VCN and compartment
(the LPG's compartment).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vcn_id` - (Required) The OCID of the VCN.


The following attributes are exported:

* `local_peering_gateways` - The list of local_peering_gateways.

### Example Usage

```hcl
data "oci_core_local_peering_gateways" "test_local_peering_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```
