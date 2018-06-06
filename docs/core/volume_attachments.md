# oci_core_volume_attachment

## VolumeAttachment Resource

### VolumeAttachment Reference

The following attributes are exported:

* `attachment_type` - The type of volume attachment.
* `availability_domain` - The Availability Domain of an instance.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.  Example: `My volume attachment` 
* `id` - The OCID of the volume attachment.
* `instance_id` - The OCID of the instance the volume is attached to.
* `is_read_only` - Whether the attachment was created in read-only mode.
* `state` - The current state of the volume attachment.
* `time_created` - The date and time the volume was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `volume_id` - The OCID of the volume.

The following additional attributes are exported for the iSCSI volume attachment type:
* `use_chap` - Whether to use CHAP authentication for the volume attachment.
* `chap_username` - The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name.
* `chap_secret` - The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name. (Also called the "CHAP password".)
* `ipv4` - The volume's iSCSI IP address.
* `port` - The volume's iSCSI port.
* `iqn` - The target volume's iSCSI Qualified Name in the format defined by RFC 3720.

### Create Operation
Attaches the specified storage volume to the specified instance.


The following arguments are supported:

* `attachment_type` - (Required) The type of volume. The only supported value are "iscsi" and "paravirtualized".
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information. 
* `instance_id` - (Required) The OCID of the instance.
* `is_read_only` - (Optional) Whether the attachment was created in read-only mode.
* `use_chap` - (Optional) Whether to use CHAP authentication for the volume attachment. This only applies if the attachment type is "iscsi".
* `volume_id` - (Required) The OCID of the volume.


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_volume_attachment" "test_volume_attachment" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"
	attachment_type = "iscsi"
	volume_id = "${oci_core_volume.test_volume.id}"

	#Optional
	display_name = "${var.volume_attachment_display_name}"
	is_read_only = "${var.volume_attachment_is_read_only}"
}
```

# oci_core_volume_attachments

## VolumeAttachment DataSource

Gets a list of volume_attachments.

### List Operation
Lists the volume attachments in the specified compartment. You can filter the
list by specifying an instance OCID, volume OCID, or both.

Currently, the only supported volume attachment type are [IScsiVolumeAttachment](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IScsiVolumeAttachment/) and
[ParavirtualizedVolumeAttachment](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ParavirtualizedVolumeAttachment/).

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `instance_id` - (Optional) The OCID of the instance.
* `volume_id` - (Optional) The OCID of the volume.


The following attributes are exported:

* `volume_attachments` - The list of volume_attachments.

### Example Usage

```hcl
data "oci_core_volume_attachments" "test_volume_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.volume_attachment_availability_domain}"
	instance_id = "${oci_core_instance.test_instance.id}"
	volume_id = "${oci_core_volume.test_volume.id}"
}
```