# oci\_core\_volume\_attachment

Provides a volume attachment resource

## Example Usage

```
resource "oci_core_volume_attachment" "t" {
    attachment_type = "attachment_type"
    compartment_id = "compartment_id"
    instance_id = "instance_id"
    volume_id = "volume_id"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) A user-friendly name. Does not have to be unique, and it cannot be changed.
* `instance_id` - (Required) The OCID of the instance.
* `volume_id` - (Required) The OCID of the volume.
* `type` - (Required) The type of volume. The only supported value is "iscsi".


## Attributes Reference
* `attachment_type` - The type of volume attachment.
* `availability_domain` - The Availability Domain of an instance.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it cannot be changed.
* `id` - The OCID of the volume attachment.
* `instance_id` - The OCID of the instance the volume is attached to.
* `state` - The current state of the volume attachment: [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `time_created` - The date and time the volume was created
* `volume_id` - The OCID of the volume.
* `chap_username` - The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name.
* `chap_secret` - The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name. (Also called the "CHAP password".)
* `ipv4` - The volume's iSCSI IP address.
* `port` - The volume's iSCSI port.
* `iqn` - The target volume's iSCSI Qualified Name in the format defined by RFC 3720.
