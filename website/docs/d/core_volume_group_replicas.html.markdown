---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_group_replicas"
sidebar_current: "docs-oci-datasource-core-volume_group_replicas"
description: |-
  Provides the list of Volume Group Replicas in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_group_replicas
This data source provides the list of Volume Group Replicas in Oracle Cloud Infrastructure Core service.

Lists the volume group replicas in the specified compartment. You can filter the results by volume group.
For more information, see [Volume Group Replication](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroupreplication.htm).


## Example Usage

```hcl
data "oci_core_volume_group_replicas" "test_volume_group_replicas" {
	#Required
	availability_domain = var.volume_group_replica_availability_domain
	compartment_id = var.compartment_id

	#Optional
	display_name = var.volume_group_replica_display_name
	state = var.volume_group_replica_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `volume_group_replicas` - The list of volume_group_replicas.

### VolumeGroupReplica Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the volume group replica.
* `compartment_id` - The OCID of the compartment that contains the volume group replica.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID for the volume group replica.
* `member_replicas` - Volume replicas within this volume group replica.
	* `volume_replica_id` - The volume replica ID.
* `size_in_gbs` - The aggregate size of the volume group replica in GBs.
* `state` - The current state of a volume group.
* `time_created` - The date and time the volume group replica was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_last_synced` - The date and time the volume group replica was last synced from the source volume group. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `volume_group_id` - The OCID of the source volume group.

