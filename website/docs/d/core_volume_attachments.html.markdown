---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_attachments"
sidebar_current: "docs-oci-datasource-core-volume_attachments"
description: |-
  Provides the list of Volume Attachments in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_attachments
This data source provides the list of Volume Attachments in Oracle Cloud Infrastructure Core service.

Lists the volume attachments in the specified compartment. You can filter the
list by specifying an instance OCID, volume OCID, or both.

Currently, the only supported volume attachment type are [IScsiVolumeAttachment](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IScsiVolumeAttachment/) and
[ParavirtualizedVolumeAttachment](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ParavirtualizedVolumeAttachment/).


## Example Usage

```hcl
data "oci_core_volume_attachments" "test_volume_attachments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.volume_attachment_availability_domain
	instance_id = oci_core_instance.test_instance.id
	volume_id = oci_core_volume.test_volume.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1`
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `instance_id` - (Optional) The OCID of the instance.
* `volume_id` - (Optional) The OCID of the volume.


## Attributes Reference

The following attributes are exported:

* `volume_attachments` - The list of volume_attachments.

### VolumeAttachment Reference

The following attributes are exported:

* `attachment_type` - The type of volume attachment.
* `availability_domain` - The availability domain of an instance.  Example: `Uocm:PHX-AD-1`
* `chap_secret` - The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name. (Also called the "CHAP password".)
* `chap_username` - The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name. See [RFC 1994](https://tools.ietf.org/html/rfc1994) for more on CHAP.  Example: `ocid1.volume.oc1.phx.<unique_ID>`
* `compartment_id` - The OCID of the compartment.
* `device` - The device name.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `encryption_in_transit_type` - Refer the top-level definition of encryptionInTransitType. The default value is NONE.
* `id` - The OCID of the volume attachment.
* `instance_id` - The OCID of the instance the volume is attached to.
* `ipv4` - The volume's iSCSI IP address.  Example: `169.254.0.2` 
* `iqn` - The target volume's iSCSI Qualified Name in the format defined by [RFC 3720](https://tools.ietf.org/html/rfc3720#page-32).  Example: `iqn.2015-12.us.oracle.com:<CHAP_username>` 
* `is_agent_auto_iscsi_login_enabled` - Whether Oracle Cloud Agent is enabled perform the iSCSI login and logout commands after the volume attach or detach operations for non multipath-enabled iSCSI attachments. 
* `is_multipath` - Whether the Iscsi or Paravirtualized attachment is multipath or not, it is not applicable to NVMe attachment.
* `is_pv_encryption_in_transit_enabled` - Whether in-transit encryption for the data volume's paravirtualized attachment is enabled or not.
* `is_read_only` - Whether the attachment was created in read-only mode.
* `is_volume_created_during_launch` - Flag indicating if this volume was created for the customer as part of a simplified launch. Used to determine whether the volume requires deletion on instance termination. 
* `iscsi_login_state` - The iscsi login state of the volume attachment. For a Iscsi volume attachment, all iscsi sessions need to be all logged-in or logged-out to be in logged-in or logged-out state.
* `multipath_devices` - A list of secondary multipath devices
	* `ipv4` - The volume's iSCSI IP address.  Example: `169.254.2.2`
	* `iqn` - The target volume's iSCSI Qualified Name in the format defined by [RFC 3720](https://tools.ietf.org/html/rfc3720#page-32).  Example: `iqn.2015-12.com.oracleiaas:40b7ee03-883f-46c6-a951-63d2841d2195`
	* `port` - The volume's iSCSI port, usually port 860 or 3260.  Example: `3260`
* `port` - The volume's iSCSI port, usually port 860 or 3260.  Example: `3260`
* `state` - The current state of the volume attachment.
* `time_created` - The date and time the volume was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
* `volume_id` - The OCID of the volume.