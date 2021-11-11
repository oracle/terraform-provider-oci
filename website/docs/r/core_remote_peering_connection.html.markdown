---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_remote_peering_connection"
sidebar_current: "docs-oci-resource-core-remote_peering_connection"
description: |-
  Provides the Remote Peering Connection resource in Oracle Cloud Infrastructure Core service
---

# oci_core_remote_peering_connection
This resource provides the Remote Peering Connection resource in Oracle Cloud Infrastructure Core service.

Creates a new remote peering connection (RPC) for the specified DRG.


## Example Usage

```hcl
resource "oci_core_remote_peering_connection" "test_remote_peering_connection" {
	#Required
	compartment_id = var.compartment_id
	drg_id = oci_core_drg.test_drg.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.remote_peering_connection_display_name
	freeform_tags = {"Department"= "Finance"}
	peer_id = oci_core_remote_peering_connection.test_remote_peering_connection2.id
	peer_region_name = var.remote_peering_connection_peer_region_name
}
```

## Argument Reference

* Specifying a `peer_id` and a `peer_region_name` creates a connection to the specified RPC ID. Both `peer_id` and `peer_region_name` are optional for creating the resource but are required for the connection to succeed. If only one of them is present the connection will not succeed.
* If the specified peer_id is also a resource in the terraform config you will have do a `terraform refresh` after the `terraform apply` in order to get the latest connection information on that resource.
* To disconnect the peering connection at least one of the RPC resources in the connection will have to be destroyed, however in terraform we recommend that when one RPC is destroyed the peer should also be destroyed. If one of them is not destroyed it will have a `REVOKED` peering_status. If another RPC resource tries to connect to this RPC resource the peering_status on the requestor will be `INVALID`. To solve this you will have to run `terraform taint oci_core_remote_peering_connection.test_remote_peering_connection` on the acceptor resource or target delete it `terraform destroy -target="oci_core_remote_peering_connection.test_remote_peering_connection"`.

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the RPC.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the RPC belongs to.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `peer_id` - (Optional) The OCID of the RPC you want to peer with.
* `peer_region_name` - (Optional) The name of the region that contains the RPC you want to peer with.  Example: `us-ashburn-1`


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the RPC.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG that this RPC belongs to.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the RPC.
* `is_cross_tenancy_peering` - Whether the VCN at the other end of the peering is in a different tenancy.  Example: `false` 
* `peer_id` - If this RPC is peered, this value is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the other RPC. 
* `peer_region_name` - If this RPC is peered, this value is the region that contains the other RPC.  Example: `us-ashburn-1` 
* `peer_tenancy_id` - If this RPC is peered, this value is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the other RPC's tenancy. 
* `peering_status` - Whether the RPC is peered with another RPC. `NEW` means the RPC has not yet been peered. `PENDING` means the peering is being established. `REVOKED` means the RPC at the other end of the peering has been deleted. 
* `state` - The RPC's current lifecycle state.
* `time_created` - The date and time the RPC was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Remote Peering Connection
	* `update` - (Defaults to 20 minutes), when updating the Remote Peering Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Remote Peering Connection


## Import

RemotePeeringConnections can be imported using the `id`, e.g.

```
$ terraform import oci_core_remote_peering_connection.test_remote_peering_connection "id"
```

