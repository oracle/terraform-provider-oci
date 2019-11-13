---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_boot_volume"
sidebar_current: "docs-oci-resource-core-boot_volume"
description: |-
  Provides the Boot Volume resource in Oracle Cloud Infrastructure Core service
---

# oci_core_boot_volume
This resource provides the Boot Volume resource in Oracle Cloud Infrastructure Core service.

Creates a new boot volume in the specified compartment from an existing boot volume or a boot volume backup.
For general information about boot volumes, see [Boot Volumes](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumes.htm).
You may optionally specify a *display name* for the volume, which is simply a friendly name or
description. It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_boot_volume" "test_boot_volume" {
	#Required
	availability_domain = "${var.boot_volume_availability_domain}"
	compartment_id = "${var.compartment_id}"
	source_details {
		#Required
		id = "${var.boot_volume_source_details_id}"
		type = "${var.boot_volume_source_details_type}"
	}

	#Optional
	backup_policy_id = "${data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id}"
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.boot_volume_display_name}"
	freeform_tags = {"Department"= "Finance"}
	kms_key_id = "${oci_kms_key.test_key.id}"
	size_in_gbs = "${var.boot_volume_size_in_gbs}"
	vpus_per_gb = "${var.boot_volume_vpus_per_gb}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain of the boot volume.  Example: `Uocm:PHX-AD-1` 
* `backup_policy_id` - (Optional) If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned. 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the boot volume.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `kms_key_id` - (Optional) (Updatable) The OCID of the KMS key to be used as the master encryption key for the boot volume.
* `size_in_gbs` - (Optional) (Updatable) The size of the volume in GBs.
* `source_details` - (Required) Specifies the boot volume source details for a new boot volume. The volume source is either another boot volume in the same availability domain or a boot volume backup. This is a mandatory field for a boot volume. 
	* `id` - (Required) The OCID of the boot volume or boot volume backup.
	* `type` - (Required) The type can be one of these values: `bootVolume`, `bootVolumeBackup`
* `vpus_per_gb` - (Optional) (Updatable) The number of Volume Performance Units that will be applied to this boot volume per GB.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the boot volume.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the boot volume.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The boot volume's Oracle ID (OCID).
* `image_id` - The image OCID used to create the boot volume.
* `is_hydrated` - Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup.
* `kms_key_id` - The OCID of the KMS key which is the master encryption key for the boot volume.
* `size_in_gbs` - The size of the boot volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`. 
* `source_details` - The boot volume source, either an existing boot volume in the same availability domain or a boot volume backup. If null, this means that the boot volume was created from an image. 
	* `id` - The OCID of the boot volume or boot volume backup.
	* `type` - The type can be one of these values: `bootVolume`, `bootVolumeBackup`
* `state` - The current state of a boot volume.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the boot volume was created. Format defined by RFC3339.
* `volume_group_id` - The OCID of the source volume group.
* `vpus_per_gb` - The number of Volume Performance Units that will be applied to this boot volume per GB.

## Import

BootVolumes can be imported using the `id`, e.g.

```
$ terraform import oci_core_boot_volume.test_boot_volume "id"
```

