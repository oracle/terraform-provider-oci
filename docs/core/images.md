# oci_core_image

## Image Resource

### Image Reference

The following attributes are exported:

* `base_image_id` - The OCID of the image originally used to launch the instance.
* `compartment_id` - The OCID of the compartment containing the instance you want to use as the basis for the image. 
* `create_image_allowed` - Whether instances launched with this image can be used to create new images. For example, you cannot create an image of an Oracle Database instance. Example: `true` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information. You cannot use an Oracle-provided image name as a custom image name.  Example: `My custom Oracle Linux image` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the image.
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
	* `firmware` - Firmware used to boot VM.  Select the option that matches your operating system. 
	    * `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders. 
	    * `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the default for Oracle provided images. 
	* `network_type` - Emulation type for NIC. 
	    * `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver. 
	    * `VFIO` - Direct attached Virtual Function network controller.  Default for Oracle provided images. 
	* `remote_data_volume_type` - Emulation type for volume. 
	    * `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images. 
	    * `SCSI` - Emulated SCSI disk. * `IDE` - Emulated IDE disk. * `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images. 
* `operating_system` - The image's operating system.  Example: `Oracle Linux` 
* `operating_system_version` - The image's operating system version.  Example: `7.2` 
* `size_in_mbs` - Image size (1 MB = 1048576 bytes)  Example: `47694` 
* `state` - The current state of the image.
* `time_created` - The date and time the image was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a boot disk image for the specified instance or imports an exported image from the Oracle Cloud Infrastructure Object Storage service.

When creating a new image, you must provide the OCID of the instance you want to use as the basis for the image, and
the OCID of the compartment containing that instance. For more information about images,
see [Managing Custom Images](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingcustomimages.htm).

You may optionally specify a *display name* for the image, which is simply a friendly name or description.
It does not have to be unique, and you can change it. See [UpdateImage](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Image/UpdateImage).
Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the instance you want to use as the basis for the image.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information.  You cannot use an Oracle-provided image name as a custom image name.  Example: `My Oracle Linux image`
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `image_source_details` - (Optional) Details for creating an image through import
    * `source_type` - (Required) The source type for the image. Use objectStorageTuple when specifying the namespace, bucket name, and object name. Use objectStorageUri when specifying the Object Storage URL. Allowed values are: - `objectStorageTuple` - `objectStorageUri` 
    * `source_image_type` - (Optional) The format of the image to be imported. Exported Oracle images are QCOW2. Only monolithic images are supported. Allowed values are: - `QCOW2` - `VMDK`
    * `source_uri` - (Required for objectStorageUri source_type) The Object Storage URL for the image.
    * `bucket_name` - (Required for objectStorageTuple source_type) The Object Storage bucket for the image.
    * `namespace_name` - (Required for objectStorageTuple source_type) The Object Storage namespace for the image.
    * `object_name` - (Required for objectStorageTuple source_type) The Object Storage name for the image.
* `instance_id` - (Optional -- required when not specifying `image_source_details`) The OCID of the instance you want to use as the basis for the image.
* `launch_mode` - (Optional) Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are: 
    * `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images. 
    * `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller. 
    * `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 


### Update Operation
Updates the display name of the image. Avoid entering confidential information.

The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the image. It does not have to be unique, and it's changeable. Avoid entering confidential information.  You cannot use an Oracle-provided image name as a custom image name.  Example: `My Oracle Linux image` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

#### Create image from instance in tenancy
```hcl
resource "oci_core_image" "test_image" {
	#Required
	compartment_id = "${var.compartment_id}"
	instance_id = "${oci_core_instance.test_instance.id}"

	#Optional
	display_name = "${var.image_display_name}"
	defined_tags = {"Operations.CostCenter"= "42"}
	launch_mode = "${var.image_launch_mode}"
	freeform_tags = {"Department"= "Finance"}
}
```

#### Create image from exported image via direct access to object store 
```hcl
resource "oci_core_image" "test_image" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.image_display_name}"
	launch_mode = "${var.image_launch_mode}"
	
	image_source_details {
		source_type = "objectStorageTuple"
		bucket_name = "${var.bucket_name"}
		namespace_name = "${var.namespace}"
		object_name = "${var.object_name}" # exported image name
        
		#Optional
		source_image_type = "${var.source_image_type}"
	}
}
```

#### Create image from exported image at publicly accessible uri  
```hcl
resource "oci_core_image" "test_image" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.image_display_name}"
	launch_mode = "${var.image_launch_mode}"
	
	image_source_details {
		source_type = "objectStorageUri"
		source_uri = "${var.source_uri}" 

		#Optional
		source_image_type = "${var.source_image_type}"
    }
}
```

# oci_core_images

## Image DataSource

Gets a list of images.

### List Operation
Lists the available images in the specified compartment.
If you specify a value for the `sortBy` parameter, Oracle-provided images appear first in the list, followed by custom images.
For more
information about images, see
[Managing Custom Images](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingcustomimages.htm).

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `operating_system` - (Optional) The image's operating system.  Example: `Oracle Linux` 
* `operating_system_version` - (Optional) The image's operating system version.  Example: `7.2` 
* `shape` - (Optional) Shape name.
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


The following attributes are exported:

* `images` - The list of images.

### Example Usage

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
}
```