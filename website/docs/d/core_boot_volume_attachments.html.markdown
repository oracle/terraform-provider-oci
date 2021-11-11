---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_boot_volume_attachments"
sidebar_current: "docs-oci-datasource-core-boot_volume_attachments"
description: |-
  Provides the list of Boot Volume Attachments in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_boot_volume_attachments
This data source provides the list of Boot Volume Attachments in Oracle Cloud Infrastructure Core service.

Lists the boot volume attachments in the specified compartment. You can filter the
list by specifying an instance OCID, boot volume OCID, or both.


## Example Usage

```hcl
data "oci_core_boot_volume_attachments" "test_boot_volume_attachments" {
	#Required
	availability_domain = var.boot_volume_attachment_availability_domain
	compartment_id = var.compartment_id

	#Optional
	boot_volume_id = oci_core_boot_volume.test_boot_volume.id
	instance_id = oci_core_instance.test_instance.id
}
```
For more detailed implementation refer the [instance example](https://github.com/oracle/terraform-provider-oci/tree/master/examples/compute/instance)

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `boot_volume_id` - (Optional) The OCID of the boot volume.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `instance_id` - (Optional) The OCID of the instance.


## Attributes Reference

The following attributes are exported:

* `boot_volume_attachments` - The list of boot_volume_attachments.

### BootVolumeAttachment Reference

The following attributes are exported:

* `availability_domain` - The availability domain of an instance.  Example: `Uocm:PHX-AD-1` 
* `boot_volume_id` - The OCID of the boot volume.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `encryption_in_transit_type` - Refer the top-level definition of encryptionInTransitType. The default value is NONE. 
* `id` - The OCID of the boot volume attachment.
* `instance_id` - The OCID of the instance the boot volume is attached to.
* `is_pv_encryption_in_transit_enabled` - Whether in-transit encryption for the boot volume's paravirtualized attachment is enabled or not.
* `state` - The current state of the boot volume attachment.
* `time_created` - The date and time the boot volume was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

