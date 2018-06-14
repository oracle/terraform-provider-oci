# oci_core_remote_peering_connection

## RemotePeeringConnection Resource

### RemotePeeringConnection Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the RPC.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The OCID of the DRG that this RPC belongs to.
* `id` - The OCID of the RPC.
* `is_cross_tenancy_peering` - Whether the VCN at the other end of the peering is in a different tenancy.  Example: `false` 
* `peer_id` - If this RPC is peered, this value is the OCID of the other RPC. 
* `peer_region_name` - If this RPC is peered, this value is the region that contains the other RPC.  Example: `us-ashburn-1` 
* `peer_tenancy_id` - If this RPC is peered, this value is the OCID of the other RPC's tenancy. 
* `peering_status` - Whether the RPC is peered with another RPC. `NEW` means the RPC has not yet been peered. `PENDING` means the peering is being established. `REVOKED` means the RPC at the other end of the peering has been deleted. 
* `state` - The RPC's current lifecycle state.
* `time_created` - The date and time the RPC was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new remote peering connection (RPC) for the specified DRG.

* Specifying a `peer_id` and a `peer_region_name` creates a connection to the specified RPC ID. Both `peer_id` and `peer_region_name` are required for the connection to succeed.
* If the specified peer_id is also a resource in the terraform config you will have do a `terraform refresh` after the `terraform apply` in order to get the latest connection information on that resource.
* To disconnect the peering connection at least one of the RPC resources in the connection will have to be destroyed, however in terraform we recommend that when one RPC is destroyed the peer should also be destroyed. If one of them is not destroyed it will have a `REVOKED` peering_status. If another RPC resource tries to connect to this RPC resource the peering_status on the requestor will be `INVALID`. To solve this you will have to run `terraform taint oci_core_remote_peering_connection.test_remote_peering_connection` on the acceptor resource or target delete it `terraform destroy -target="oci_core_remote_peering_connection.test_remote_peering_connection"`.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the RPC.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - (Required) The OCID of the DRG the RPC belongs to.
* `peer_id` - (Optional) The OCID of the RPC you want to peer with.
* `peer_region_name` - (Optional) The name of the region that contains the RPC you want to peer with.  Example: `us-ashburn-1`

### Update Operation
Updates the specified remote peering connection (RPC).


The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_remote_peering_connection" "test_remote_peering_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	drg_id = "${oci_core_drg.test_drg.id}"

	#Optional
	display_name = "${var.remote_peering_connection_display_name}"
	peer_id = "${oci_core_remote_peering_connection.test_remote_peering_connection2.id}"
	peer_region_name = "${var.remote_peering_connection_peer_region_name}"
}
```

# oci_core_remote_peering_connections

## RemotePeeringConnection DataSource

Gets a list of remote_peering_connections.

### List Operation
Lists the remote peering connections (RPCs) for the specified DRG and compartment
(the RPC's compartment).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `drg_id` - (Optional) The OCID of the DRG.


The following attributes are exported:

* `remote_peering_connections` - The list of remote_peering_connections.

### Example Usage

```hcl
data "oci_core_remote_peering_connections" "test_remote_peering_connections" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	drg_id = "${oci_core_drg.test_drg.id}"
}
```