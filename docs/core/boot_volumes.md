
# oci_core_boot_volumes

## BootVolume DataSource

Gets a list of boot_volumes.

### List Operation
Lists the boot volumes in the specified compartment and Availability Domain.

The following arguments are supported:

* `availability_domain` - (Required) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `volume_group_id` - (Optional) The OCID of the volume group.


The following attributes are exported:

* `boot_volumes` - The list of boot_volumes.

### Example Usage

```hcl
data "oci_core_boot_volumes" "test_boot_volumes" {
	#Required
	availability_domain = "${var.boot_volume_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	volume_group_id = "${oci_core_volume_group.test_volume_group.id}"
}
```
### BootVolume Reference

The following attributes are exported:

* `availability_domain` - The Availability Domain of the boot volume.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the boot volume.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The boot volume's Oracle ID (OCID).
* `image_id` - The image OCID used to create the boot volume.
* `size_in_gbs` - The size of the boot volume in GBs.
* `size_in_mbs` - The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use sizeInGBs. 
* `state` - The current state of a boot volume.
* `time_created` - The date and time the boot volume was created. Format defined by RFC3339.
* `volume_group_id` - The OCID of the source volume group.
