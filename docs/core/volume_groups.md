# oci_core_volume_group

## VolumeGroup Resource

### VolumeGroup Reference

The following attributes are exported:

* `availability_domain` - The Availability Domain of the volume group.
* `compartment_id` - The OCID of the compartment that contains the volume group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume group. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The Oracle Cloud ID (OCID) that uniquely identifies the volume group.
* `size_in_mbs` - The aggregate size of the volume group in MBs.
* `source_details` - The volume group source. The volume source is either another a list of volume ids in the same Availability Domain, another volume group or a volume group backup. 
	* `type` - The type of the volume group source. It should be set to either `volumeIds`, `volumeGroup`, or `volumeBackup`
	* `volume_ids` - OCIDs for the volumes in this volume group, if the type is `volumeIds`
	* `volume_group_id` - The OCID of the volume group to clone from, if the type is `volumeGroup`
	* `volume_group_backup_id` - The OCID of the volume group backup to restore from, if the type is `volumeGroupBackup` 
* `state` - The current state of a volume group.
* `time_created` - The date and time the volume group was created. Format defined by RFC3339.
* `volume_ids` - OCIDs for the volumes in this volume group.



### Create Operation
Creates a new volume group in the specified compartment. A volume group can have at most 20 block volumes.
A volume group is a collection of volumes and may be created from a list of volumes, cloning an existing
volume group or by restoring a volume group backup.
You may optionally specify a *display name* for the volume group, which is simply a friendly name or
description. It does not have to be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `availability_domain` - (Required) The Availability Domain of the volume group.
* `compartment_id` - (Required) The OCID of the compartment that contains the volume group.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name for the volume group. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `source_details` - (Required) Specifies the volume group source details for a new volume group. The volume source is either another a list of volume ids in the same Availability Domain, another volume group or a volume group backup. 
	* `type` - (Required) The type of the volume group source. It should be set to either `volumeIds`, `volumeGroup`, or `volumeBackup`
	* `volume_ids` (Optional) - OCIDs for the volumes in this volume group, if the type is `volumeIds`
	* `volume_group_id` (Optional) - The OCID of the volume group to clone from, if the type is `volumeGroup`
	* `volume_group_backup_id` (Optional) - The OCID of the volume group backup to restore from, if the type is `volumeGroupBackup` 


### Update Operation
Updates the set of volumes in a volume group along with (optionally) it's display name. This call can be used
to add or remove volumes in a volume group. The entire list of volume ids must be specified.
Avoid entering confidential information.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume group. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_volume_group" "test_volume_group" {
	#Required
	availability_domain = "${var.volume_group_availability_domain}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		type = "volumeIds"
		volume_ids = ["${var.volume_group_source_id}"]
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.volume_group_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

# oci_core_volume_groups

## VolumeGroup DataSource

Gets a list of volume_groups.

### List Operation
Lists the volume groups in the specified compartment and Availability Domain.

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.


The following attributes are exported:

* `volume_groups` - The list of volume_groups.

### Example Usage

```hcl
data "oci_core_volume_groups" "test_volume_groups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.volume_group_availability_domain}"
	display_name = "${var.volume_group_display_name}"
	state = "${var.volume_group_state}"
}
```