---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_remote_peering_connections"
sidebar_current: "docs-oci-datasource-core-remote_peering_connections"
description: |-
  Provides the list of Remote Peering Connections in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_remote_peering_connections
This data source provides the list of Remote Peering Connections in Oracle Cloud Infrastructure Core service.

Lists the remote peering connections (RPCs) for the specified DRG and compartment
(the RPC's compartment).


## Example Usage

```hcl
data "oci_core_remote_peering_connections" "test_remote_peering_connections" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	drg_id = oci_core_drg.test_drg.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `drg_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.


## Attributes Reference

The following attributes are exported:

* `remote_peering_connections` - The list of remote_peering_connections.

### RemotePeeringConnection Reference

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

