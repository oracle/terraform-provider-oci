---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_group"
sidebar_current: "docs-oci-resource-core-volume_group"
description: |-
  Provides the Volume Group resource in Oracle Cloud Infrastructure Core service
---

# oci_core_volume_group
This resource provides the Volume Group resource in Oracle Cloud Infrastructure Core service.

Creates a new volume group in the specified compartment.
A volume group is a collection of volumes and may be created from a list of volumes, cloning an existing
volume group, or by restoring a volume group backup.
You may optionally specify a *display name* for the volume group, which is simply a friendly name or
description. It does not have to be unique, and you can change it. Avoid entering confidential information.

For more information, see [Volume Groups](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).

Note: If the volume group is created from another volume group or from a volume group backup, a copy of the volumes from the source is made in your compartment. However, this is not automatically deleted by Terraform when this volume group is deleted. To track these volumes, you can import them into the terraform statefile and run terraform destroy. Alternatively, you can also use another interface like CLI, SDK, or Console to remove them manually. 

## Example Usage

```hcl
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = var.volume_group_availability_domain
	compartment_id = var.compartment_id
	source_details {
		#Required
		type = "volumeIds"
		volume_ids = [var.volume_group_source_id]
	}

	#Optional
	backup_policy_id = data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.volume_group_display_name
	freeform_tags = {"Department"= "Finance"}
	volume_group_replicas {
		#Required
		availability_domain = var.volume_group_volume_group_replicas_availability_domain

		#Optional
		display_name = var.volume_group_volume_group_replicas_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of the volume group.
* `backup_policy_id` - (Optional) If provided, specifies the ID of the volume backup policy to assign to the newly created volume group. If omitted, no policy will be assigned. 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the volume group.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `source_details` - (Required) Specifies the source for a volume group.
	* `type` - (Required) The type can be one of these values: `volumeGroupBackupId`, `volumeGroupId`, `volumeGroupReplicaId`, `volumeIds`
	* `volume_group_backup_id` - (Required when type=volumeGroupBackupId) The OCID of the volume group backup to restore from.
	* `volume_group_id` - (Required when type=volumeGroupId) The OCID of the volume group to clone from.
	* `volume_group_replica_id` - (Required when type=volumeGroupReplicaId) The OCID of the volume group replica.
	* `volume_ids` - (Required when type=volumeIds) OCIDs for the volumes in this volume group.
* `volume_group_replicas` - (Optional) (Updatable) The list of volume group replicas that this volume group will be enabled to have in the specified destination availability domains. 
	* `availability_domain` - (Required) (Updatable) The availability domain of the volume group replica.  Example: `Uocm:PHX-AD-1` 
	* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the volume group.
* `compartment_id` - The OCID of the compartment that contains the volume group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID for the volume group.
* `is_hydrated` - Specifies whether the newly created cloned volume group's data has finished copying from the source volume group or backup. 
* `size_in_gbs` - The aggregate size of the volume group in GBs.
* `size_in_mbs` - The aggregate size of the volume group in MBs.
* `source_details` - Specifies the source for a volume group.
	* `type` - The type can be one of these values: `volumeGroupBackupId`, `volumeGroupId`, `volumeGroupReplicaId`, `volumeIds`
	* `volume_group_backup_id` - The OCID of the volume group backup to restore from.
	* `volume_group_id` - The OCID of the volume group to clone from.
	* `volume_group_replica_id` - The OCID of the volume group replica.
	* `volume_ids` - OCIDs for the volumes in this volume group.
* `state` - The current state of a volume group.
* `time_created` - The date and time the volume group was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `volume_group_replicas` - The list of volume group replicas of this volume group.
	* `availability_domain` - The availability domain of the boot volume replica replica.  Example: `Uocm:PHX-AD-1` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `volume_group_replica_id` - The volume group replica's Oracle ID (OCID).
* `volume_ids` - OCIDs for the volumes in this volume group.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Volume Group
	* `update` - (Defaults to 20 minutes), when updating the Volume Group
	* `delete` - (Defaults to 20 minutes), when destroying the Volume Group


## Import

VolumeGroups can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume_group.test_volume_group "id"
```

