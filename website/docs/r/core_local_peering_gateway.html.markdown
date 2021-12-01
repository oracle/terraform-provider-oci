---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_local_peering_gateway"
sidebar_current: "docs-oci-resource-core-local_peering_gateway"
description: |-
  Provides the Local Peering Gateway resource in Oracle Cloud Infrastructure Core service
---

# oci_core_local_peering_gateway
This resource provides the Local Peering Gateway resource in Oracle Cloud Infrastructure Core service.

Creates a new local peering gateway (LPG) for the specified VCN.


## Example Usage

```hcl
resource "oci_core_local_peering_gateway" "test_local_peering_gateway" {
	#Required
	compartment_id = var.compartment_id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.local_peering_gateway_display_name
	freeform_tags = {"Department"= "Finance"}
	peer_id = oci_core_local_peering_gateway.test_local_peering_gateway2.id
	route_table_id = oci_core_route_table.test_route_table.id
}
```

## Argument Reference

* Specifying a peer_id creates a connection to the specified LPG ID. `peer_id` should only be specified in one of the LPGs.
* If the specified peer_id is also a resource in the terraform config you will have do a `terraform refresh` after the `terraform apply` in order to get the latest connection information on that resource.
* To disconnect the peering connection at least one of the LPG resources in the connection will have to be destroyed, however in terraform we recommend that when one LPG is destroyed the peer should also be destroyed. If one of them is not destroyed it will have a `REVOKED` peering_status. If another LPG resource tries to connect to this LPG resource it will get a `400 Error: The Local Peering Gateway with ID X has already been connected`. To solve this you will have to run `terraform taint oci_core_local_peering_gateway.test_local_peering_gateway` on that resource or target delete it `terraform destroy -target="oci_core_local_peering_gateway.test_local_peering_gateway"`.

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the local peering gateway (LPG). 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `peer_id` - (Optional) The OCID of the LPG you want to peer with. Specifying a peer_id connects this local peering gateway (LPG) to another one in the same region. This operation must be called by the VCN administrator who is designated as the *requestor* in the peering relationship. The *acceptor* must implement an Identity and Access Management (IAM) policy that gives the requestor permission to connect to LPGs in the acceptor's compartment. Without that permission, this operation will fail. For more information, see [VCN Peering](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/VCNpeering.htm).
* `route_table_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the LPG will use.

	If you don't specify a route table here, the LPG is created without an associated route table. The Networking service does NOT automatically associate the attached VCN's default route table with the LPG.

	For information about why you would associate a route table with an LPG, see [Transit Routing: Access to Multiple VCNs in Same Region](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm). 
* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the LPG belongs to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the LPG.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The LPG's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `is_cross_tenancy_peering` - Whether the VCN at the other end of the peering is in a different tenancy.  Example: `false` 
* `peer_advertised_cidr` - The smallest aggregate CIDR that contains all the CIDR routes advertised by the VCN at the other end of the peering from this LPG. See `peerAdvertisedCidrDetails` for the individual CIDRs. The value is `null` if the LPG is not peered.  Example: `192.168.0.0/16`, or if aggregated with `172.16.0.0/24` then `128.0.0.0/1` 
* `peer_advertised_cidr_details` - The specific ranges of IP addresses available on or via the VCN at the other end of the peering from this LPG. The value is `null` if the LPG is not peered. You can use these as destination CIDRs for route rules to route a subnet's traffic to this LPG.  Example: [`192.168.0.0/16`, `172.16.0.0/24`] 
* `peer_id` - The OCID of the peered LPG
* `peering_status` - Whether the LPG is peered with another LPG. `NEW` means the LPG has not yet been peered. `PENDING` means the peering is being established. `REVOKED` means the LPG at the other end of the peering has been deleted. 
* `peering_status_details` - Additional information regarding the peering status, if applicable.
* `route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the LPG is using.

	For information about why you would associate a route table with an LPG, see [Transit Routing: Access to Multiple VCNs in Same Region](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitrouting.htm). 
* `state` - The LPG's current lifecycle state.
* `time_created` - The date and time the LPG was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN that uses the LPG.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Local Peering Gateway
	* `update` - (Defaults to 20 minutes), when updating the Local Peering Gateway
	* `delete` - (Defaults to 20 minutes), when destroying the Local Peering Gateway


## Import

LocalPeeringGateways can be imported using the `id`, e.g.

```
$ terraform import oci_core_local_peering_gateway.test_local_peering_gateway "id"
```

