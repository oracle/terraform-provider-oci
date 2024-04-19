---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume"
sidebar_current: "docs-oci-resource-core-volume"
description: |-
  Provides the Volume resource in Oracle Cloud Infrastructure Core service
---

# oci_core_volume
This resource provides the Volume resource in Oracle Cloud Infrastructure Core service.

Creates a new volume in the specified compartment. Volumes can be created in sizes ranging from
50 GB (51200 MB) to 32 TB (33554432 MB), in 1 GB (1024 MB) increments. By default, volumes are 1 TB (1048576 MB).
For general information about block volumes, see
[Overview of Block Volume Service](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm).

A volume and instance can be in separate compartments but must be in the same availability domain.
For information about access control and compartments, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about
availability domains, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the `ListAvailabilityDomains` operation
in the Identity and Access Management Service API.

You may optionally specify a *display name* for the volume, which is simply a friendly name or
description. It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_volume" "test_volume" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	autotune_policies {
		#Required
		autotune_type = var.volume_autotune_policies_autotune_type

		#Optional
		max_vpus_per_gb = var.volume_autotune_policies_max_vpus_per_gb
	}
	availability_domain = var.volume_availability_domain
	backup_policy_id = data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id
	block_volume_replicas {
		#Required
		availability_domain = var.volume_block_volume_replicas_availability_domain

		#Optional
		display_name = var.volume_block_volume_replicas_display_name
	}
	cluster_placement_group_id = oci_identity_group.test_group.id
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.volume_display_name
	freeform_tags = {"Department"= "Finance"}
	is_auto_tune_enabled = var.volume_is_auto_tune_enabled
	kms_key_id = oci_kms_key.test_key.id
	size_in_gbs = var.volume_size_in_gbs
	size_in_mbs = var.volume_size_in_mbs
	source_details {
		#Required
		id = var.volume_source_details_id
		type = var.volume_source_details_type
	}
	vpus_per_gb = var.volume_vpus_per_gb
    block_volume_replicas_deletion = true
}
```

## Argument Reference

The following arguments are supported:

* `autotune_policies` - (Optional) (Updatable) The list of autotune policies to be enabled for this volume.
	* `autotune_type` - (Required) (Updatable) This specifies the type of autotunes supported by OCI.
	* `max_vpus_per_gb` - (Required when autotune_type=PERFORMANCE_BASED) (Updatable) This will be the maximum VPUs/GB performance level that the volume will be auto-tuned temporarily based on performance monitoring. 
* `availability_domain` - (Optional) The availability domain of the volume. Omissible for cloning a volume. The new volume will be created in the availability domain of the source volume.  Example: `Uocm:PHX-AD-1` 
* `backup_policy_id` - (Optional) If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. This field is deprecated. Use the `oci_core_volume_backup_policy_assignments` instead to assign a backup policy to a volume.
* `block_volume_replicas` - (Optional) (Updatable) The list of block volume replicas to be enabled for this volume in the specified destination availability domains. 
	* `availability_domain` - (Required) (Updatable) The availability domain of the block volume replica.  Example: `Uocm:PHX-AD-1` 
	* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `cluster_placement_group_id` - (Optional) The clusterPlacementGroup Id of the volume for volume placement.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the volume.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_auto_tune_enabled` - (Optional) (Updatable) Specifies whether the auto-tune performance is enabled for this volume. This field is deprecated. Use the `DetachedVolumeAutotunePolicy` instead to enable the volume for detached autotune. 
* `kms_key_id` - (Optional) (Updatable) The OCID of the Vault service key to assign as the master encryption key for the volume. 
* `size_in_gbs` - (Optional) (Updatable) The size of the volume in GBs.
* `size_in_mbs` - (Optional) The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Use `size_in_gbs` instead. 
* `source_details` - (Optional) 
	* `id` - (Required) The OCID of the block volume replica.
	* `type` - (Required) The type can be one of these values: `blockVolumeReplica`, `volume`, `volumeBackup`
* `volume_backup_id` - (Optional) The OCID of the volume backup from which the data should be restored on the newly created volume. This field is deprecated. Use the sourceDetails field instead to specify the backup for the volume. 
* `vpus_per_gb` - (Optional) (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

	Allowed values:
	* `0`: Represents Lower Cost option.
	* `10`: Represents Balanced option.
	* `20`: Represents Higher Performance option.
	* `30`-`120`: Represents the Ultra High Performance option.

	For performance autotune enabled volumes, it would be the Default(Minimum) VPUs/GB. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Volume
	* `update` - (Defaults to 20 minutes), when updating the Volume
	* `delete` - (Defaults to 20 minutes), when destroying the Volume


## Import

Volumes can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume.test_volume "id"
```

