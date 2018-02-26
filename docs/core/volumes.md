# oci\_core\_volume

## Volume Resource

### Volume Reference

The following attributes are exported:

* `availability_domain` - The Availability Domain of the volume.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the volume.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The OCID of the volume.
* `is_hydrated` - Specifies whether the cloned volume's data has finished copying from the source volume or backup.
* `size_in_gbs` - The size of the volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. This field is deprecated. Use sizeInGBs instead.
* `source_details` - The volume source, either an existing volume in the same Availability Domain or a volume backup. If null, an empty volume is created. 
	* `type` - 
* `state` - The current state of a volume.
* `time_created` - The date and time the volume was created. Format defined by RFC3339.



### Create Operation
Creates a new volume in the specified compartment. Volumes can be created in sizes ranging from
50 GB (51200 MB) to 16 TB (16777216 MB), in 1 GB (1024 MB) increments. By default, volumes are 1 TB (1048576 MB).
For general information about block volumes, see
[Overview of Block Volume Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Block/Concepts/overview.htm).

A volume and instance can be in separate compartments but must be in the same Availability Domain.
For information about access control and compartments, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about
Availability Domains, see [Regions and Availability Domains](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).
To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
in the Identity and Access Management Service API.

You may optionally specify a *display name* for the volume, which is simply a friendly name or
description. It does not have to be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `availability_domain` - (Required) The Availability Domain of the volume.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment that contains the volume.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `size_in_gbs` - (Optional) The size of the volume in GBs.
* `size_in_mbs` - (Optional) The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Use sizeInGBs instead. 
* `source_details` - (Optional) Specifies the volume source details for a new Block volume. The volume source is either another Block volume in the same Availability Domain or a Block volume backup. This is an optional field. If not specified or set to null, the new Block volume will be empty. When specified, the new Block volume will contain data from the source volume or backup. 
	* `type` - (Required) 
* `volume_backup_id` - (Optional) The OCID of the volume backup from which the data should be restored on the newly created volume. This field is deprecated. Use the sourceDetails field instead to specify the backup for the volume. 


### Update Operation
Updates the specified volume's display name.
Avoid entering confidential information.


The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_core_volume" "test_volume" {
	#Required
	availability_domain = "${var.volume_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.volume_display_name}"
	size_in_gbs = "${var.volume_size_in_gbs}"
	size_in_mbs = "${var.volume_size_in_mbs}"
	source_details {
		#Required
		type = "${var.volume_source_details_type}"
	}
	volume_backup_id = "${oci_core_volume_backup.test_volume_backup.id}"
}
```

# oci\_core\_volumes

## Volume DataSource

Gets a list of volumes.

### List Operation
Lists the volumes in the specified compartment and Availability Domain.

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


The following attributes are exported:

* `volumes` - The list of volumes.

### Example Usage

```
data "oci_core_volumes" "test_volumes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.volume_availability_domain}"
	display_name = "${var.volume_display_name}"
	state = "${var.volume_state}"
}
```