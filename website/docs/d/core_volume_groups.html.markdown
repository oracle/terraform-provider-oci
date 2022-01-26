---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_groups"
sidebar_current: "docs-oci-datasource-core-volume_groups"
description: |-
  Provides the list of Volume Groups in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_groups
This data source provides the list of Volume Groups in Oracle Cloud Infrastructure Core service.

Lists the volume groups in the specified compartment and availability domain.
For more information, see [Volume Groups](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).


## Example Usage

```hcl
data "oci_core_volume_groups" "test_volume_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.volume_group_availability_domain
	display_name = var.volume_group_display_name
	state = var.volume_group_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `volume_groups` - The list of volume_groups.

### VolumeGroup Reference

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
	* `type` - The type can be one of these values: `volumeGroupBackupId`, `volumeGroupId`, `volumeIds`
	* `volume_group_backup_id` - The OCID of the volume group backup to restore from, if the type is `volumeGroupBackup` 
	* `volume_group_id` - The OCID of the volume group to clone from, if the type is `volumeGroup`
	* `volume_ids` - OCIDs for the volumes in this volume group, if the type is `volumeIds`
* `state` - The current state of a volume group.
* `time_created` - The date and time the volume group was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `volume_group_replicas` - The list of volume group replicas of this volume group.
	* `availability_domain` - The availability domain of the boot volume replica replica.  Example: `Uocm:PHX-AD-1` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `volume_group_replica_id` - The volume group replica's Oracle ID (OCID).
* `volume_ids` - OCIDs for the volumes in this volume group.

