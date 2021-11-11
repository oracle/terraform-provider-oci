---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_boot_volume"
sidebar_current: "docs-oci-resource-core-boot_volume"
description: |-
  Provides the Boot Volume resource in Oracle Cloud Infrastructure Core service
---

# oci_core_boot_volume
This resource provides the Boot Volume resource in Oracle Cloud Infrastructure Core service.

Creates a new boot volume in the specified compartment from an existing boot volume or a boot volume backup.
For general information about boot volumes, see [Boot Volumes](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumes.htm).
You may optionally specify a *display name* for the volume, which is simply a friendly name or
description. It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_boot_volume" "test_boot_volume" {
	#Required
	compartment_id = var.compartment_id
	source_details {
		#Required
		id = var.boot_volume_source_details_id
		type = var.boot_volume_source_details_type
	}

	#Optional
	availability_domain = var.boot_volume_availability_domain
	backup_policy_id = data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id
	boot_volume_replicas {
		#Required
		availability_domain = var.boot_volume_boot_volume_replicas_availability_domain

		#Optional
		display_name = var.boot_volume_boot_volume_replicas_display_name
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.boot_volume_display_name
	freeform_tags = {"Department"= "Finance"}
	is_auto_tune_enabled = var.boot_volume_is_auto_tune_enabled
	kms_key_id = oci_kms_key.test_key.id
	size_in_gbs = var.boot_volume_size_in_gbs
	vpus_per_gb = var.boot_volume_vpus_per_gb
    boot_volume_replicas_deletion = true
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The availability domain of the volume. Omissible for cloning a volume. The new volume will be created in the availability domain of the source volume.  Example: `Uocm:PHX-AD-1` 
* `backup_policy_id` - (Optional) If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned. 
* `boot_volume_replicas` - (Optional) (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains. 
	* `availability_domain` - (Required) (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1` 
	* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the boot volume.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_auto_tune_enabled` - (Optional) (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume. 
* `kms_key_id` - (Optional) (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume. 
* `size_in_gbs` - (Optional) (Updatable) The size of the volume in GBs.
* `source_details` - (Required) 
	* `id` - (Required) The OCID of the boot volume replica.
	* `type` - (Required) The type can be one of these values: `bootVolume`, `bootVolumeBackup`, `bootVolumeReplica`
* `vpus_per_gb` - (Optional) (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.

	Allowed values:
	* `10`: Represents Balanced option.
	* `20`: Represents Higher Performance option. 
* `boot_volume_replicas_deletion` - (Optional) (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `boot_volume_replicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `boot_volume_replicas` again.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `auto_tuned_vpus_per_gb` - The number of Volume Performance Units per GB that this boot volume is effectively tuned to when it's idle. 
* `availability_domain` - The availability domain of the boot volume.  Example: `Uocm:PHX-AD-1` 
* `boot_volume_replicas` - The list of boot volume replicas of this boot volume
	* `availability_domain` - The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1` 
	* `boot_volume_replica_id` - The boot volume replica's Oracle ID (OCID).
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `compartment_id` - The OCID of the compartment that contains the boot volume.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The boot volume's Oracle ID (OCID).
* `image_id` - The image OCID used to create the boot volume.
* `is_auto_tune_enabled` - Specifies whether the auto-tune performance is enabled for this boot volume. 
* `is_hydrated` - Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup. 
* `kms_key_id` - The OCID of the Key Management master encryption key assigned to the boot volume.
* `size_in_gbs` - The size of the boot volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`. 
* `source_details` - 
	* `id` - The OCID of the boot volume replica.
	* `type` - The type can be one of these values: `bootVolume`, `bootVolumeBackup`, `bootVolumeReplica`
* `state` - The current state of a boot volume.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the boot volume was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `volume_group_id` - The OCID of the source volume group.
* `vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this boot volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.

	Allowed values:
	* `10`: Represents Balanced option.
	* `20`: Represents Higher Performance option. 
* `boot_volume_replicas_deletion` - The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `boot_volume_replicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `boot_volume_replicas` again.


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Boot Volume
	* `update` - (Defaults to 20 minutes), when updating the Boot Volume
	* `delete` - (Defaults to 20 minutes), when destroying the Boot Volume


## Import

BootVolumes can be imported using the `id`, e.g.

```
$ terraform import oci_core_boot_volume.test_boot_volume "id"
```

