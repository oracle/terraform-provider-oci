---
layout: "oci"
page_title: "OCI: oci_core_boot_volume"
sidebar_current: "docs-oci-datasource-core-boot_volume"
description: |-
  Provides details about a specific BootVolume
---

# Data Source: oci_core_boot_volume
The `oci_core_boot_volume` data source provides details about a specific BootVolume

Gets information for the specified boot volume.

## Example Usage

```hcl
data "oci_core_boot_volume" "test_boot_volume" {
	#Required
	boot_volume_id = "${oci_core_boot_volume.test_boot_volume.id}"
}
```

## Argument Reference

The following arguments are supported:

* `boot_volume_id` - (Required) The OCID of the boot volume.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the boot volume.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the boot volume.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The boot volume's Oracle ID (OCID).
* `image_id` - The image OCID used to create the boot volume.
* `is_hydrated` - Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup.
* `size_in_gbs` - The size of the boot volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`. 
* `source_details` - The boot volume source, either an existing boot volume in the same availability domain or a boot volume backup. If null, this means that the boot volume was created from an image. 
	* `id` - The OCID of the boot volume or boot volume backup.
	* `type` - The type of boot volume source. It should be set to either `bootVolumeBackup` or `bootVolume`.
* `state` - The current state of a boot volume.
* `time_created` - The date and time the boot volume was created. Format defined by RFC3339.
* `volume_group_id` - The OCID of the source volume group.

