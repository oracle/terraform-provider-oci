---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_boot_volume"
sidebar_current: "docs-oci-datasource-core-boot_volume"
description: |-
  Provides details about a specific Boot Volume in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_boot_volume
This data source provides details about a specific Boot Volume resource in Oracle Cloud Infrastructure Core service.

Gets information for the specified boot volume.

## Example Usage

```hcl
data "oci_core_boot_volume" "test_boot_volume" {
	#Required
	boot_volume_id = oci_core_boot_volume.test_boot_volume.id
}
```

## Argument Reference

The following arguments are supported:

* `boot_volume_id` - (Required) The OCID of the boot volume.


## Attributes Reference

The following attributes are exported:

* `auto_tuned_vpus_per_gb` - The number of Volume Performance Units per GB that this boot volume is effectively tuned to. 
* `autotune_policies` - The list of autotune policies enabled for this volume.
	* `autotune_type` - This specifies the type of autotunes supported by OCI.
	* `max_vpus_per_gb` - This will be the maximum VPUs/GB performance level that the volume will be auto-tuned temporarily based on performance monitoring. 
* `availability_domain` - The availability domain of the boot volume.  Example: `Uocm:PHX-AD-1` 
* `boot_volume_replicas` - The list of boot volume replicas of this boot volume
	* `availability_domain` - The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1` 
	* `boot_volume_replica_id` - The boot volume replica's Oracle ID (OCID).
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `cluster_placement_group_id` - The clusterPlacementGroup Id of the volume for volume placement.
* `compartment_id` - The OCID of the compartment that contains the boot volume.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The boot volume's Oracle ID (OCID).
* `image_id` - The image OCID used to create the boot volume.
* `is_auto_tune_enabled` - Specifies whether the auto-tune performance is enabled for this boot volume. This field is deprecated. Use the `DetachedVolumeAutotunePolicy` instead to enable the volume for detached autotune. 
* `is_hydrated` - Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup. 
* `kms_key_id` - The OCID of the Vault service master encryption key assigned to the boot volume.
* `size_in_gbs` - The size of the boot volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`. 
* `source_details` - 
	* `id` - The OCID of the boot volume replica.
	* `type` - The type can be one of these values: `bootVolume`, `bootVolumeBackup`, `bootVolumeReplica`
* `state` - The current state of a boot volume.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the boot volume was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `volume_group_id` - The OCID of the source volume group.
* `vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this boot volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

	Allowed values:
	* `10`: Represents Balanced option.
	* `20`: Represents Higher Performance option.
	* `30`-`120`: Represents the Ultra High Performance option.

	For performance autotune enabled volumes, it would be the Default(Minimum) VPUs/GB. 

