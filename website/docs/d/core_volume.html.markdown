---
layout: "oci"
page_title: "OCI: oci_core_volume"
sidebar_current: "docs-oci-datasource-core-volume"
description: |-
  Provides details about a specific Volume
---

# Data Source: oci_core_volume
The `oci_core_volume` data source provides details about a specific Volume

Gets information for the specified volume.

## Example Usage

```hcl
data "oci_core_volume" "test_volume" {
	#Required
	volume_id = "${oci_core_volume.test_volume.id}"
}
```

## Argument Reference

The following arguments are supported:

* `volume_id` - (Required) The OCID of the volume.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the volume.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume.
* `is_hydrated` - Specifies whether the cloned volume's data has finished copying from the source volume or backup.
* `size_in_gbs` - The size of the volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. This field is deprecated. Use sizeInGBs instead.
* `source_details` - The volume source, either an existing volume in the same availability domain or a volume backup. If null, an empty volume is created. 
	* `id` - The OCID of the volume or volume backup.
	* `type` - The type of volume source. It should be set to either `volumeBackup` or `volume`.
* `state` - The current state of a volume.
* `time_created` - The date and time the volume was created. Format defined by RFC3339.
* `volume_group_id` - The OCID of the source volume group.

