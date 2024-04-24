---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volumes"
sidebar_current: "docs-oci-datasource-core-volumes"
description: |-
  Provides the list of Volumes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volumes
This data source provides the list of Volumes in Oracle Cloud Infrastructure Core service.

Lists the volumes in the specified compartment and availability domain.


## Example Usage

```hcl
data "oci_core_volumes" "test_volumes" {

	#Optional
	availability_domain = var.volume_availability_domain
	cluster_placement_group_id = oci_identity_group.test_group.id
	compartment_id = var.compartment_id
	display_name = var.volume_display_name
	state = var.volume_state
	volume_group_id = oci_core_volume_group.test_volume_group.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `cluster_placement_group_id` - (Optional) A filter to return only resources that match the given cluster placement group Id exactly. 
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 
* `volume_group_id` - (Optional) The OCID of the volume group.


## Attributes Reference

The following attributes are exported:

* `volumes` - The list of volumes.

### Volume Reference

The following attributes are exported:

* `auto_tuned_vpus_per_gb` - The number of Volume Performance Units per GB that this volume is effectively tuned to. 
* `autotune_policies` - The list of autotune policies enabled for this volume.
	* `autotune_type` - This specifies the type of autotunes supported by OCI.
	* `max_vpus_per_gb` - This will be the maximum VPUs/GB performance level that the volume will be auto-tuned temporarily based on performance monitoring. 
* `availability_domain` - The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
* `block_volume_replicas` - The list of block volume replicas of this volume.
	* `availability_domain` - The availability domain of the block volume replica.  Example: `Uocm:PHX-AD-1` 
	* `block_volume_replica_id` - The block volume replica's Oracle ID (OCID).
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `cluster_placement_group_id` - The clusterPlacementGroup Id of the volume for volume placement.
* `compartment_id` - The OCID of the compartment that contains the volume.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume.
* `is_auto_tune_enabled` - Specifies whether the auto-tune performance is enabled for this volume. This field is deprecated. Use the `DetachedVolumeAutotunePolicy` instead to enable the volume for detached autotune. 
* `is_hydrated` - Specifies whether the cloned volume's data has finished copying from the source volume or backup. 
* `kms_key_id` - The OCID of the Vault service key which is the master encryption key for the volume. 
* `size_in_gbs` - The size of the volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. This field is deprecated. Use `size_in_gbs` instead.
* `source_details` - 
	* `id` - The OCID of the block volume replica.
	* `type` - The type can be one of these values: `blockVolumeReplica`, `volume`, `volumeBackup`
* `state` - The current state of a volume.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the volume was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `volume_group_id` - The OCID of the source volume group.
* `vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

	Allowed values:
	* `0`: Represents Lower Cost option.
	* `10`: Represents Balanced option.
	* `20`: Represents Higher Performance option.
	* `30`-`120`: Represents the Ultra High Performance option.

	For performance autotune enabled volumes, It would be the Default(Minimum) VPUs/GB. 

