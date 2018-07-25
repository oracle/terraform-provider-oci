---
layout: "oci"
page_title: "OCI: oci_core_instances"
sidebar_current: "docs-oci-datasource-core-instances"
description: |-
  Provides a list of Instances
---

# Data Source: oci_core_instances
The Instances data source allows access to the list of OCI instances

Lists the instances in the specified compartment and the specified Availability Domain.
You can filter the results by specifying an instance name (the list will include all the identically-named
instances in the compartment).


## Example Usage

```hcl
data "oci_core_instances" "test_instances" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.instance_availability_domain}"
	display_name = "${var.instance_display_name}"
	state = "${var.instance_state}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `instances` - The list of instances.

### Instance Reference

The following attributes are exported:

* `availability_domain` - The Availability Domain the instance is running in.  Example: `Uocm:PHX-AD-1` 
* `boot_volume_id` - The OCID of the attached boot volume. If the `source_type` is `bootVolume`, this will be the same OCID as the `source_id`.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
* `extended_metadata` - Additional metadata key/value pairs that you provide.  They serve a similar purpose and functionality from fields in the 'metadata' object.  They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only).  If you don't need nested metadata values, it is strongly advised to avoid using this object and use the Metadata object instead. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the instance.
* `image` - Deprecated. Use `sourceDetails` instead. 
* `ipxe_script` - When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.  If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.  The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.  For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/References/bringyourownimage.htm).  For more information about iPXE, see http://ipxe.org. 
* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are: 
	* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images. 
	* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller. 
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
	* `network_type` - Emulation type for NIC. 
		* `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver. 
		* `VFIO` - Direct attached Virtual Function network controller.  Default for Oracle provided images. 
	* `remote_data_volume_type` - Emulation type for volume. 
		* `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images. 
		* `SCSI` - Emulated SCSI disk. * `IDE` - Emulated IDE disk. 
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images. 
		* `PARAVIRTUALIZED` - Paravirtualized disk. 
* `metadata` - Custom metadata that you provide.
* `preserve_boot_volume` - Specifies whether to delete or preserve the boot volume when terminating an instance. The default value is false. Note: This value only applies to destroy operations initiated by Terraform.
* `private_ip` - The private IP address of instance VNIC. To set the private IP address, use the `private_ip` argument in create_vnic_details.
* `public_ip` - The public IP address of instance VNIC (if enabled).
* `region` - The region that contains the Availability Domain the instance is running in.  Example: `phx` 
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance. You can enumerate all available shapes by calling [ListShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Shape/ListShapes). 
* `source_details` - Details for creating an instance
	* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
	* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 16384 GB (16TB). This should only be specified when `source_type` is `image`.
	* `source_id` - The OCID of an image or a boot volume to use, depending on the value of `source_type`. 
* `state` - The current state of the instance.
* `time_created` - The date and time the instance was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

