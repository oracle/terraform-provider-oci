---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance"
sidebar_current: "docs-oci-resource-core-instance"
description: |-
  Provides the Instance resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance
This resource provides the Instance resource in Oracle Cloud Infrastructure Core service.

Creates a new instance in the specified compartment and the specified availability domain.
For general information about instances, see
[Overview of the Compute Service](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm).

For information about access control and compartments, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

For information about availability domains, see
[Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the `ListAvailabilityDomains` operation
in the Identity and Access Management Service API.

All Oracle Cloud Infrastructure resources, including instances, get an Oracle-assigned,
unique ID called an Oracle Cloud Identifier (OCID).
When you create a resource, you can find its OCID in the response. You can
also retrieve a resource's OCID by using a List API operation
on that resource type, or by viewing the resource in the Console.

To launch an instance using an image or a boot volume use the `sourceDetails` parameter in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/LaunchInstanceDetails).

When you launch an instance, it is automatically attached to a virtual
network interface card (VNIC), called the *primary VNIC*. The VNIC
has a private IP address from the subnet's CIDR. You can either assign a
private IP address of your choice or let Oracle automatically assign one.
You can choose whether the instance has a public IP address. To retrieve the
addresses, use the [ListVnicAttachments](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/VnicAttachment/ListVnicAttachments)
operation to get the VNIC ID for the instance, and then call
[GetVnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Vnic/GetVnic) with the VNIC ID.

You can later add secondary VNICs to an instance. For more information, see
[Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm).

To launch an instance from a Marketplace image listing, you must provide the image ID of the 
listing resource version that you want, but you also must subscribe to the listing before you try 
to launch the instance. To subscribe to the listing, use the [GetAppCatalogListingAgreements](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements/GetAppCatalogListingAgreements) 
operation to get the signature for the terms of use agreement for the desired listing resource version.  
Then, call [CreateAppCatalogSubscription](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AppCatalogSubscription/CreateAppCatalogSubscription) 
with the signature. To get the image ID for the LaunchInstance operation, call 
[GetAppCatalogListingResourceVersion](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersion/GetAppCatalogListingResourceVersion).


## Example Usage

```hcl
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = "${var.instance_availability_domain}"
	compartment_id = "${var.compartment_id}"
	shape = "${var.instance_shape}"

	#Optional
	agent_config {

		#Optional
		is_management_disabled = "${var.instance_agent_config_is_management_disabled}"
		is_monitoring_disabled = "${var.instance_agent_config_is_monitoring_disabled}"
	}
	create_vnic_details {

		#Optional
		assign_public_ip = "${var.instance_create_vnic_details_assign_public_ip}"
		defined_tags = {"Operations.CostCenter"= "42"}
		display_name = "${var.instance_create_vnic_details_display_name}"
		freeform_tags = {"Department"= "Finance"}
		hostname_label = "${var.instance_create_vnic_details_hostname_label}"
		nsg_ids = "${var.instance_create_vnic_details_nsg_ids}"
		private_ip = "${var.instance_create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.instance_create_vnic_details_skip_source_dest_check}"
		subnet_id = "${oci_core_subnet.test_subnet.id}"
		vlan_id = "${oci_core_vlan.test_vlan.id}"
	}
	dedicated_vm_host_id = "${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}"
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.instance_display_name}"
	extended_metadata = {
		some_string = "stringA"
		nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
	}
	fault_domain = "${var.instance_fault_domain}"
	freeform_tags = {"Department"= "Finance"}
	hostname_label = "${var.instance_hostname_label}"
	ipxe_script = "${var.instance_ipxe_script}"
	is_pv_encryption_in_transit_enabled = "${var.instance_is_pv_encryption_in_transit_enabled}"
	launch_options {

		#Optional
		boot_volume_type = "${var.instance_launch_options_boot_volume_type}"
		firmware = "${var.instance_launch_options_firmware}"
		is_consistent_volume_naming_enabled = "${var.instance_launch_options_is_consistent_volume_naming_enabled}"
		is_pv_encryption_in_transit_enabled = "${var.instance_launch_options_is_pv_encryption_in_transit_enabled}"
		network_type = "${var.instance_launch_options_network_type}"
		remote_data_volume_type = "${var.instance_launch_options_remote_data_volume_type}"
	}
	shape_config {

		#Optional
		ocpus = "${var.instance_shape_config_ocpus}"
    }
	metadata = {
		ssh_authorized_keys = "${var.ssh_public_key}"
		user_data = "${base64encode(file(var.custom_bootstrap_file_name))}"
	}
	source_details {
		#Required
		source_id = "${oci_core_image.test_image.id}"
		source_type = "image"

		#Optional
		boot_volume_size_in_gbs = "${var.instance_source_details_boot_volume_size_in_gbs}"
		kms_key_id = "${oci_kms_key.test_key.id}"
	}
	preserve_boot_volume = false
}
```

## Argument Reference

The following arguments are supported:

* `agent_config` - (Optional) (Updatable) 
	* `is_management_disabled` - (Optional) (Updatable) Whether the agent running on the instance can run all the available management plugins. Default value is false. 
	* `is_monitoring_disabled` - (Optional) (Updatable) Whether the agent running on the instance can gather performance metrics and monitor the instance. Default value is false. 
* `availability_domain` - (Required) The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `create_vnic_details` - (Optional) (Updatable) Details for the primary VNIC, which is automatically created and attached when the instance is launched. 
	* `assign_public_ip` - (Optional) (Updatable) Whether the VNIC should be assigned a public IP address. Defaults to true. If left blank or set to true and `prohibitPublicIpOnVnic` = true, an error is returned.

		**Note:** This public IP address is associated with the primary private IP on the VNIC. For more information, see [IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).

		**Note:** There's a limit to the number of [public IPs](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PublicIp/) a VNIC or instance can have. If you try to create a secondary VNIC with an assigned public IP for an instance that has already reached its public IP limit, an error is returned. For information about the public IP limits, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).

		Example: `false` 
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - (Optional) (Updatable) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
	* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `hostname_label` - (Optional) (Updatable) The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp).

		For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

		When launching an instance, use this `hostnameLabel` instead of the deprecated `hostnameLabel` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/requests/LaunchInstanceDetails). If you provide both, the values must match.

		Example: `bminstance-1` 
	* `nsg_ids` - (Optional) (Updatable) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/).

		If a `vlanId` is specified, the `nsgIds` is ignored. The `vlanId` indicates that the VNIC will belong to a VLAN instead of a subnet. With VLANs, all VNICs in the VLAN belong to the NSGs that are associated with the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Vlan). 
	* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's *primary* private IP address. The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp).

		 If you specify a `vlanId`, the `privateIp` is ignored. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Vlan).

		Example: `10.0.3.3` 
	* `skip_source_dest_check` - (Optional) (Updatable) Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you would skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).

		 If you specify a `vlanId`, the `skipSourceDestCheck` is ignored because the source/destination check is always disabled for VNICs in a VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Vlan).

		Example: `true` 
	* `subnet_id` - (Optional) The OCID of the subnet to create the VNIC in. When launching an instance, use this `subnetId` instead of the deprecated `subnetId` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/requests/LaunchInstanceDetails). At least one of them is required; if you provide both, the values must match.

		If you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN instead of a subnet, provide a `vlanId` instead of a `subnetId`. If you provide both a `vlanId` and `subnetId`, the request fails. 
	* `vlan_id` - (Optional) Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the OCID of the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Vlan).

		Provide a `vlanId` instead of a `subnetId`. If you provide both a `vlanId` and `subnetId`, the request fails. 
* `dedicated_vm_host_id` - (Optional) The OCID of dedicated VM host. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
* `extended_metadata` - (Optional) (Updatable) Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the 'metadata' object.

	They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only). 

	Input in terraform is the same as metadata but allows nested metadata if you pass a valid JSON string as a value. See the example above.
* `fault_domain` - (Optional) A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

	If you do not specify the fault domain, the system selects one for you. To change the fault domain for an instance, terminate it and launch a new instance in the preferred fault domain.

	To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - (Optional) Deprecated. Instead use `hostnameLabel` in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/). If you provide both, the values must match. 
* `image` - (Optional) Deprecated. Use `sourceDetails` with [InstanceSourceViaImageDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/InstanceSourceViaImageDetails) source type instead. If you specify values for both, the values must match. 
* `ipxe_script` - (Optional) This is an advanced option.

	When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

	If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

	The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

	For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

	For more information about iPXE, see http://ipxe.org. 
* `is_pv_encryption_in_transit_enabled` - (Optional) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
* `launch_options` - (Optional) Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
	* `boot_volume_type` - (Optional) Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images. 
	* `firmware` - (Optional) Firmware used to boot VM.  Select the option that matches your operating system.
		* `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
		* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the default for Oracle provided images. 
	* `is_consistent_volume_naming_enabled` - (Optional) Whether to enable consistent volume naming feature. Defaults to false.
	* `is_pv_encryption_in_transit_enabled` - (Optional) Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/datatypes/LaunchInstanceDetails). 
	* `network_type` - (Optional) Emulation type for the physical network interface card (NIC).
		* `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
		* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
		* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using virtio drivers. 
	* `remote_data_volume_type` - (Optional) Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk.This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images. 
* `metadata` - (Optional) (Updatable) Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

	A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
	* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
	* Get information about the instance, including the custom metadata that you provide when you launch the instance.

	**Providing Cloud-Init Metadata**

	You can use the following metadata key names to provide information to Cloud-Init:

	**"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

	**"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

	**Metadata Example**

	```
	"metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227", "user_data" : "<your_public_SSH_key>==" }
	```

	**Getting Metadata on the Instance**

	To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

	```
	curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/
	curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/
	curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>
	```

	You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively.
	
	**Note:** Both the 'user_data' and 'ssh_authorized_keys' fields cannot be changed after an instance has launched. Any request which updates, removes, or adds either of these fields will be rejected. You must provide the same values for 'user_data' and 'ssh_authorized_keys' that already exist on the instance. 
* `preserve_boot_volume` - (Optional) Specifies whether to delete or preserve the boot volume when terminating an instance. The default value is false. Note: This value only applies to destroy operations initiated by Terraform.
* `shape` - (Required) (Updatable) The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

	You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes). 
* `shape_config` - (Optional) (Updatable) 
	* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to the instance. 
* `source_details` - (Optional) Details for creating an instance. Use this parameter to specify whether a boot volume or an image should be used to launch a new instance. 
	* `boot_volume_size_in_gbs` - (Applicable when source_type=image) The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 16384 GB (16TB).
	* `kms_key_id` - (Applicable when source_type=image) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
	* `source_id` - (Required) The OCID of an image or a boot volume to use, depending on the value of `source_type`.
	* `source_type` - (Required) The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
* `state` - (Optional) (Updatable) The target state for the instance. Could be set to RUNNING or STOPPED.
* `subnet_id` - (Optional) Deprecated. Instead use `subnetId` in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/). At least one of them is required; if you provide both, the values must match. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `agent_config` - 
	* `is_management_disabled` - Whether the agent running on the instance can run all the available management plugins. 
	* `is_monitoring_disabled` - Whether the agent running on the instance can gather performance metrics and monitor the instance. 
* `availability_domain` - The availability domain the instance is running in.  Example: `Uocm:PHX-AD-1` 
* `boot_volume_id` - The OCID of the attached boot volume. If the `source_type` is `bootVolume`, this will be the same OCID as the `source_id`.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `dedicated_vm_host_id` - The OCID of dedicated VM host. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My bare metal instance` 
* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the 'metadata' object.

	They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only). 

	Input in terraform is the same as metadata but allows nested metadata if you pass a valid JSON string as a value. See the example below.
* `fault_domain` - The name of the fault domain the instance is running in.

	A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

	If you do not specify the fault domain, the system selects one for you. To change the fault domain for an instance, terminate it and launch a new instance in the preferred fault domain.

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the instance.
* `image` - Deprecated. Use `sourceDetails` instead. 
* `ipxe_script` - When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

	If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

	The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

	For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

	For more information about iPXE, see http://ipxe.org. 
* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
	* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images.
	* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
	* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using virtio drivers.
	* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
* `launch_options` - Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
	* `boot_volume_type` - Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images. 
	* `firmware` - Firmware used to boot VM.  Select the option that matches your operating system.
		* `BIOS` - Boot VM using BIOS style firmware.  This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
		* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems.  This is the default for Oracle provided images. 
	* `is_consistent_volume_naming_enabled` - Whether to enable consistent volume naming feature. Defaults to false.
	* `is_pv_encryption_in_transit_enabled` - Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/datatypes/LaunchInstanceDetails). 
	* `network_type` - Emulation type for the physical network interface card (NIC).
		* `E1000` - Emulated Gigabit ethernet controller.  Compatible with Linux e1000 network driver.
		* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
		* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using virtio drivers. 
	* `remote_data_volume_type` - Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage.  This is the default option for Local data volumes on Oracle provided images.
		* `PARAVIRTUALIZED` - Paravirtualized disk.This is the default for Boot Volumes and Remote Block Storage volumes on Oracle provided images. 
* `metadata` - Custom metadata that you provide.
* `preserve_boot_volume` - Specifies whether to delete or preserve the boot volume when terminating an instance. The default value is false. Note: This value only applies to destroy operations initiated by Terraform.
* `private_ip` - The private IP address of instance VNIC. To set the private IP address, use the `private_ip` argument in create_vnic_details.
* `public_ip` - The public IP address of instance VNIC (if enabled).
* `region` - The region that contains the availability domain the instance is running in.

	For the us-phoenix-1 and us-ashburn-1 regions, `phx` and `iad` are returned, respectively. For all other regions, the full region name is returned.

	Examples: `phx`, `eu-frankfurt-1` 
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance. You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes). 
* `shape_config` - 
	* `gpu_description` - A short description of the instance's graphics processing unit (GPU).

		If the instance does not have any GPUs, this field is `null`. 
	* `gpus` - The number of GPUs available to the instance. 
	* `local_disk_description` - A short description of the local disks available to this instance.

		If the instance does not have any local disks, this field is `null`. 
	* `local_disks` - The number of local disks available to the instance. 
	* `local_disks_total_size_in_gbs` - The aggregate size of all local disks, in gigabytes.

		If the instance does not have any local disks, this field is `null`. 
	* `max_vnic_attachments` - The maximum number of VNIC attachments for the instance. 
	* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
	* `networking_bandwidth_in_gbps` - The networking bandwidth available to the instance, in gigabits per second. 
	* `ocpus` - The total number of OCPUs available to the instance. 
	* `processor_description` - A short description of the instance's processor (CPU). 
* `source_details` - Details for creating an instance
	* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 16384 GB (16TB).
	* `kms_key_id` - The OCID of the Key Management key to assign as the master encryption key for the boot volume.
	* `source_id` - The OCID of an image or a boot volume to use, depending on the value of `source_type`.
	* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
* `state` - The current state of the instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the instance was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_maintenance_reboot_due` - The date and time the instance is expected to be stopped / started,  in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). After that time if instance hasn't been rebooted, Oracle will reboot the instance within 24 hours of the due time. Regardless of how the instance was stopped, the flag will be reset to empty as soon as instance reaches Stopped state. Example: `2018-05-25T21:10:29.600Z` 

## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance.test_instance "id"
```

