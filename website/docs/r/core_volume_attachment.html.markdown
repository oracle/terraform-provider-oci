---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_attachment"
sidebar_current: "docs-oci-resource-core-volume_attachment"
description: |-
  Provides the Volume Attachment resource in Oracle Cloud Infrastructure Core service
---

# oci_core_volume_attachment
This resource provides the Volume Attachment resource in Oracle Cloud Infrastructure Core service.

Attaches the specified storage volume to the specified instance.


## Example Usage

```hcl
resource "oci_core_volume_attachment" "test_volume_attachment" {
	#Required
	attachment_type = "${var.volume_attachment_attachment_type}"
	instance_id = "${oci_core_instance.test_instance.id}"
	volume_id = "${oci_core_volume.test_volume.id}"

	#Optional
	display_name = "${var.volume_attachment_display_name}"
	is_read_only = "${var.volume_attachment_is_read_only}"
	use_chap = "${var.volume_attachment_use_chap}"
}
```

## Argument Reference

The following arguments are supported:

* `attachment_type` - (Required) The type of volume. The only supported value are "iscsi" and "paravirtualized".
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information. 
* `instance_id` - (Required) The OCID of the instance.
* `is_read_only` - (Optional) Whether the attachment was created in read-only mode.
* `use_chap` - (Applicable when attachment_type=iscsi) Whether to use CHAP authentication for the volume attachment. Defaults to false.
* `volume_id` - (Required) The OCID of the volume.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `attachment_type` - The type of volume attachment.
* `availability_domain` - The Availability Domain of an instance.  Example: `Uocm:PHX-AD-1` 
* `chap_secret` - The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name. (Also called the "CHAP password".)  Example: `d6866c0d-298b-48ba-95af-309b4faux45e` 
* `chap_username` - The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name.  Example: `ocid1.volume.oc1.phx.abyhqljrgvttnlx73nmrwfaux7kcvzfs3s66izvxf2h4lgvyndsdsnoiwr5q` 
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.  Example: `My volume attachment` 
* `id` - The OCID of the volume attachment.
* `instance_id` - The OCID of the instance the volume is attached to.
* `ipv4` - The volume's iSCSI IP address.  Example: `169.254.0.2` 
* `iqn` - The target volume's iSCSI Qualified Name in the format defined by RFC 3720.  Example: `iqn.2015-12.us.oracle.com:456b0391-17b8-4122-bbf1-f85fc0bb97d9` 
* `is_read_only` - Whether the attachment was created in read-only mode.
* `port` - The volume's iSCSI port.  Example: `3260` 
* `state` - The current state of the volume attachment.
* `time_created` - The date and time the volume was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `volume_id` - The OCID of the volume.

## Import

VolumeAttachments can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume_attachment.test_volume_attachment "id"
```

