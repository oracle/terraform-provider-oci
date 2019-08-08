---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_images"
sidebar_current: "docs-oci-datasource-core-images"
description: |-
  Provides the list of Images in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_images
This data source provides the list of Images in Oracle Cloud Infrastructure Core service.

Lists the available images in the specified compartment, including both
[Oracle-provided images](https://docs.cloud.oracle.com/iaas/Content/Compute/References/images.htm) and
[custom images](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/managingcustomimages.htm) that have
been created. The list of images returned is ordered to first show all
Oracle-provided images, then all custom images.

The order of images returned may change when new images are released.


## Example Usage

```hcl
data "oci_core_images" "test_images" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.image_display_name}"
	operating_system = "${var.image_operating_system}"
	operating_system_version = "${var.image_operating_system_version}"
	shape = "${var.image_shape}"
	state = "${var.image_state}"
	sort_by = "${var.image_sort_by}"
	sort_order = "${var.image_sort_order}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `operating_system` - (Optional) The image's operating system.  Example: `Oracle Linux` 
* `operating_system_version` - (Optional) The image's operating system version.  Example: `7.2` 
* `shape` - (Optional) Shape name.
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `sort_by` - (Optional) Sort the resources returned, by creation time or display name. Example `TIMECREATED` or `DISPLAYNAME`.
* `sort_order` - (Optional) The sort order to use, either ascending (`ASC`) or descending (`DESC`).

## Attributes Reference

The following attributes are exported:

* `images` - The list of images.

### Image Reference

The following attributes are exported:

* `agent_features` - 
	* `is_monitoring_supported` - Whether the agent running on the instance can gather performance metrics and monitor the instance. 
* `base_image_id` - The OCID of the image originally used to launch the instance.
* `compartment_id` - The OCID of the compartment containing the instance you want to use as the basis for the image. 
* `create_image_allowed` - Whether instances launched with this image can be used to create new images. For example, you cannot create an image of an Oracle Database instance.  Example: `true` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information.

	You cannot use an Oracle-provided image name as a custom image name.

	Example: `My custom Oracle Linux image` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the image.
* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
	* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images.
	* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
	* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using virtio drivers.
	* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
* `launch_options` - 
	* `boot_volume_type` - Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. 
	* `firmware` - Firmware used to boot VM.  Select the option that matches your operating system.
		* `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
		* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the default for Oracle provided images. 
	* `is_consistent_volume_naming_enabled` - Whether to enable consistent volume naming feature. Defaults to false.
	* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the boot volume's paravirtualized attachment. The default value is false.
	* `network_type` - Emulation type for the physical network interface card (NIC).
		* `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
		* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
		* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using virtio drivers. 
	* `remote_data_volume_type` - Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. 
* `operating_system` - The image's operating system.  Example: `Oracle Linux` 
* `operating_system_version` - The image's operating system version.  Example: `7.2` 
* `size_in_mbs` - The boot volume size for an instance launched from this image, (1 MB = 1048576 bytes). Note this is not the same as the size of the image when it was exported or the actual size of the image.  Example: `47694` 
* `state` - 
* `time_created` - The date and time the image was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

