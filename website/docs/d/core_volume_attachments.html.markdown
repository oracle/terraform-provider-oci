---
layout: "oci"
page_title: "OCI: oci_core_volume_attachments"
sidebar_current: "docs-oci-datasource-core-volume_attachments"
description: |-
  Provides a list of VolumeAttachments
---

# Data Source: oci_core_volume_attachments
The `oci_core_volume_attachments` data source allows access to the list of OCI volume_attachments

Lists the volume attachments in the specified compartment. You can filter the
list by specifying an instance OCID, volume OCID, or both.

Currently, the only supported volume attachment type are [IScsiVolumeAttachment](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IScsiVolumeAttachment/) and
[ParavirtualizedVolumeAttachment](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/ParavirtualizedVolumeAttachment/).


## Example Usage

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

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `instance_id` - (Optional) The OCID of the instance.
* `volume_id` - (Optional) The OCID of the volume.


## Attributes Reference

The following attributes are exported:

* `volume_attachments` - The list of volume_attachments.

### VolumeAttachment Reference

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

