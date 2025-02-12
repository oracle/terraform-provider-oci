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

To launch an instance using an image or a boot volume use the `sourceDetails` parameter in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/LaunchInstanceDetails).

When you launch an instance, it is automatically attached to a virtual
network interface card (VNIC), called the *primary VNIC*. The VNIC
has a private IP address from the subnet's CIDR. You can either assign a
private IP address of your choice or let Oracle automatically assign one.
You can choose whether the instance has a public IP address. To retrieve the
addresses, use the [ListVnicAttachments](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/VnicAttachment/ListVnicAttachments)
operation to get the VNIC ID for the instance, and then call
[GetVnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/GetVnic) with the VNIC ID.

You can later add secondary VNICs to an instance. For more information, see
[Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm).

To launch an instance from a Marketplace image listing, you must provide the image ID of the
listing resource version that you want, but you also must subscribe to the listing before you try
to launch the instance. To subscribe to the listing, use the [GetAppCatalogListingAgreements](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements/GetAppCatalogListingAgreements)
operation to get the signature for the terms of use agreement for the desired listing resource version.
Then, call [CreateAppCatalogSubscription](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AppCatalogSubscription/CreateAppCatalogSubscription)
with the signature. To get the image ID for the LaunchInstance operation, call
[GetAppCatalogListingResourceVersion](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersion/GetAppCatalogListingResourceVersion).

When launching an instance, you may provide the `securityAttributes` parameter in
[LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/LaunchInstanceDetails) to manage security attributes via the instance, 
or in the embedded [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) to manage security attributes
via the VNIC directly, but not both.  Providing `securityAttributes` in both locations will return a
400 Bad Request response.

To determine whether capacity is available for a specific shape before you create an instance,
use the [CreateComputeCapacityReport](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ComputeCapacityReport/CreateComputeCapacityReport)
operation.


## Example Usage

```hcl
resource "oci_core_instance" "test_instance" {
	#Required
	availability_domain = var.instance_availability_domain
	compartment_id = var.compartment_id
	shape = var.instance_shape

	#Optional
	agent_config {

		#Optional
		are_all_plugins_disabled = var.instance_agent_config_are_all_plugins_disabled
		is_management_disabled = var.instance_agent_config_is_management_disabled
		is_monitoring_disabled = var.instance_agent_config_is_monitoring_disabled
		plugins_config {
			#Required
			desired_state = var.instance_agent_config_plugins_config_desired_state
			name = var.instance_agent_config_plugins_config_name
		}
	}
	availability_config {

		#Optional
		is_live_migration_preferred = var.instance_availability_config_is_live_migration_preferred
		recovery_action = var.instance_availability_config_recovery_action
	}
	cluster_placement_group_id = oci_identity_group.test_group.id
	compute_cluster_id = oci_core_compute_cluster.test_compute_cluster.id
	create_vnic_details {

		#Optional
		assign_ipv6ip = var.instance_create_vnic_details_assign_ipv6ip
		assign_private_dns_record = var.instance_create_vnic_details_assign_private_dns_record
		assign_public_ip = var.instance_create_vnic_details_assign_public_ip
		defined_tags = {"Operations.CostCenter"= "42"}
		display_name = var.instance_create_vnic_details_display_name
		freeform_tags = {"Department"= "Finance"}
		hostname_label = var.instance_create_vnic_details_hostname_label
		ipv6address_ipv6subnet_cidr_pair_details = var.instance_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details
		nsg_ids = var.instance_create_vnic_details_nsg_ids
		private_ip = var.instance_create_vnic_details_private_ip
		security_attributes = var.instance_create_vnic_details_security_attributes
		skip_source_dest_check = var.instance_create_vnic_details_skip_source_dest_check
		subnet_id = oci_core_subnet.test_subnet.id
		vlan_id = oci_core_vlan.test_vlan.id
	}
	dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.instance_display_name
	extended_metadata = {
		some_string = "stringA"
		nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
	}
	fault_domain = var.instance_fault_domain
	freeform_tags = {"Department"= "Finance"}
	hostname_label = var.instance_hostname_label
	instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id
	instance_options {

		#Optional
		are_legacy_imds_endpoints_disabled = var.instance_instance_options_are_legacy_imds_endpoints_disabled
	}
	ipxe_script = var.instance_ipxe_script
	is_pv_encryption_in_transit_enabled = var.instance_is_pv_encryption_in_transit_enabled
	launch_options {

		#Optional
		boot_volume_type = var.instance_launch_options_boot_volume_type
		firmware = var.instance_launch_options_firmware
		is_consistent_volume_naming_enabled = var.instance_launch_options_is_consistent_volume_naming_enabled
		is_pv_encryption_in_transit_enabled = var.instance_launch_options_is_pv_encryption_in_transit_enabled
		network_type = var.instance_launch_options_network_type
		remote_data_volume_type = var.instance_launch_options_remote_data_volume_type
	}
	launch_volume_attachments {
		#Required
		type = var.instance_launch_volume_attachments_type

		#Optional
		device = var.instance_launch_volume_attachments_device
		display_name = var.instance_launch_volume_attachments_display_name
		encryption_in_transit_type = var.instance_launch_volume_attachments_encryption_in_transit_type
		is_agent_auto_iscsi_login_enabled = var.instance_launch_volume_attachments_is_agent_auto_iscsi_login_enabled
		is_pv_encryption_in_transit_enabled = var.instance_launch_volume_attachments_is_pv_encryption_in_transit_enabled
		is_read_only = var.instance_launch_volume_attachments_is_read_only
		is_shareable = var.instance_launch_volume_attachments_is_shareable
		launch_create_volume_details {
			#Required
			size_in_gbs = var.instance_launch_volume_attachments_launch_create_volume_details_size_in_gbs
			volume_creation_type = var.instance_launch_volume_attachments_launch_create_volume_details_volume_creation_type

			#Optional
			compartment_id = var.compartment_id
			display_name = var.instance_launch_volume_attachments_launch_create_volume_details_display_name
			kms_key_id = oci_kms_key.test_key.id
			vpus_per_gb = var.instance_launch_volume_attachments_launch_create_volume_details_vpus_per_gb
		}
		use_chap = var.instance_launch_volume_attachments_use_chap
		volume_id = oci_core_volume.test_volume.id
	}
	licensing_configs {
		#Required
		type = var.instance_licensing_configs_type

		#Optional
		license_type = var.instance_licensing_configs_license_type
	}
	metadata = var.instance_metadata
	platform_config {
		#Required
		type = var.instance_platform_config_type

		#Optional
		are_virtual_instructions_enabled = var.instance_platform_config_are_virtual_instructions_enabled
		config_map = var.instance_platform_config_config_map
		is_access_control_service_enabled = var.instance_platform_config_is_access_control_service_enabled
		is_input_output_memory_management_unit_enabled = var.instance_platform_config_is_input_output_memory_management_unit_enabled
		is_measured_boot_enabled = var.instance_platform_config_is_measured_boot_enabled
		is_memory_encryption_enabled = var.instance_platform_config_is_memory_encryption_enabled
		is_secure_boot_enabled = var.instance_platform_config_is_secure_boot_enabled
		is_symmetric_multi_threading_enabled = var.instance_platform_config_is_symmetric_multi_threading_enabled
		is_trusted_platform_module_enabled = var.instance_platform_config_is_trusted_platform_module_enabled
		numa_nodes_per_socket = var.instance_platform_config_numa_nodes_per_socket
		percentage_of_cores_enabled = var.instance_platform_config_percentage_of_cores_enabled
	}
	preemptible_instance_config {
		#Required
		preemption_action {
			#Required
			type = var.instance_preemptible_instance_config_preemption_action_type

			#Optional
			preserve_boot_volume = var.instance_preemptible_instance_config_preemption_action_preserve_boot_volume
		}
	}
	security_attributes = var.instance_security_attributes
	shape = var.instance_shape
	shape_config {

		#Optional
		baseline_ocpu_utilization = var.instance_shape_config_baseline_ocpu_utilization
		memory_in_gbs = var.instance_shape_config_memory_in_gbs
		nvmes = var.instance_shape_config_nvmes
		ocpus = var.instance_shape_config_ocpus
		vcpus = var.instance_shape_config_vcpus
	}
	source_details {
		#Required
		source_id = oci_core_image.test_image.id
		source_type = "image"

		#Optional
		boot_volume_size_in_gbs = var.instance_source_details_boot_volume_size_in_gbs
		boot_volume_vpus_per_gb = var.instance_source_details_boot_volume_vpus_per_gb
		instance_source_image_filter_details {
			#Required
			compartment_id = var.compartment_id

			#Optional
			defined_tags_filter = var.instance_source_details_instance_source_image_filter_details_defined_tags_filter
			operating_system = var.instance_source_details_instance_source_image_filter_details_operating_system
			operating_system_version = var.instance_source_details_instance_source_image_filter_details_operating_system_version
		}
		kms_key_id = oci_kms_key.test_key.id
	}
	preserve_boot_volume = false
}
```

## Argument Reference

The following arguments are supported:

* `async` - (Optional) Whether Terraform creates and destroys the resource asynchronously. The default value is false.
	* If `async` is true, all the creation and deletion of instances are asynchronous
	* If `async` is false, all the creation and deletion of instances are synchronous as normal behavior

-> Please follow this guideline [Terraform support asynchronous operation](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/terraform-async.htm) for more detail of this advanced option.

* `agent_config` - (Optional) (Updatable) Configuration options for the Oracle Cloud Agent software running on the instance.
	* `are_all_plugins_disabled` - (Optional) (Updatable) Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

		To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
	* `is_management_disabled` - (Optional) (Updatable) Whether Oracle Cloud Agent can run all the available management plugins. Default value is false (management plugins are enabled).

		These are the management plugins: OS Management Service Agent and Compute Instance Run Command.

		The management plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
		* If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of the per-plugin configuration.
		* If `isManagementDisabled` is false, all of the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
	* `is_monitoring_disabled` - (Optional) (Updatable) Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. Default value is false (monitoring plugins are enabled).

		These are the monitoring plugins: Compute Instance Monitoring and Custom Logs Monitoring.

		The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
		* If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of the per-plugin configuration.
		* If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
	* `plugins_config` - (Optional) (Updatable) The configuration of plugins associated with this instance.
		* `desired_state` - (Required) (Updatable) Whether the plugin should be enabled or disabled.

			To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
		* `name` - (Required) (Updatable) The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
* `availability_config` - (Optional) (Updatable) Options for VM migration during infrastructure maintenance events and for defining the availability of a VM instance after a maintenance event that impacts the underlying hardware. 
	* `is_live_migration_preferred` - (Optional) (Updatable) Whether to live migrate supported VM instances to a healthy physical VM host without disrupting running instances during infrastructure maintenance events. If null, Oracle chooses the best option for migrating the VM during infrastructure maintenance events. 
	* `recovery_action` - (Optional) (Updatable) The lifecycle state for an instance when it is recovered after infrastructure maintenance.
		* `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event. If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
		* `STOP_INSTANCE` - The instance is recovered in the stopped state. 
* `availability_domain` - (Required) The availability domain of the instance.  Example: `Uocm:PHX-AD-1`
* `capacity_reservation_id` - (Optional) (Updatable) The OCID of the compute capacity reservation this instance is launched under. You can opt out of all default reservations by specifying an empty string as input for this field. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default). 
* `cluster_placement_group_id` - (Optional) The OCID of the cluster placement group of the instance.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `compute_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm) that the instance will be created in. 
* `create_vnic_details` - (Optional) (Updatable) Contains properties for a VNIC. You use this object when creating the primary VNIC during instance launch or when creating a secondary VNIC. For more information about VNICs, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
	* `assign_ipv6ip` - (Optional) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
	* `assign_private_dns_record` - (Optional) Whether the VNIC should be assigned a DNS record. If set to false, there will be no DNS record registration for the VNIC. If set to true, the DNS record will be registered. The default value is true.
		If you specify a `hostnameLabel`, the `assignPrivateDnsRecord` is require to be set to true. 
	* `assign_public_ip` - (Optional) (Updatable) Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the [Subnet](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Subnet/)), then no public IP address is assigned. If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then a public IP address is assigned. If set to true and `prohibitPublicIpOnVnic` = true, an error is returned.

		**Note:** This public IP address is associated with the primary private IP on the VNIC. For more information, see [IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).

		**Note:** There's a limit to the number of [public IPs](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/) a VNIC or instance can have. If you try to create a secondary VNIC with an assigned public IP for an instance that has already reached its public IP limit, an error is returned. For information about the public IP limits, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).

		Example: `false`

		If you specify a `vlanId`, then `assignPublicIp` must be set to false. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `hostname_label` - (Optional) (Updatable) The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance1` in FQDN `bminstance1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

		For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

		When launching an instance, use this `hostnameLabel` instead of the deprecated `hostnameLabel` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). If you provide both, the values must match.

		Example: `bminstance1`

		If you specify a `vlanId`, the `hostnameLabel` cannot be specified. VNICs on a VLAN can not be assigned a hostname. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
	* `ipv6address_ipv6subnet_cidr_pair_details` - (Optional) A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges from which Oracle Cloud Infrastructure will select an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range. 
	* `nsg_ids` - (Optional) (Updatable) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).

		If a `vlanId` is specified, the `nsgIds` cannot be specified. The `vlanId` indicates that the VNIC will belong to a VLAN instead of a subnet. With VLANs, all VNICs in the VLAN belong to the NSGs that are associated with the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
	* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's *primary* private IP address. The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

		 If you specify a `vlanId`, the `privateIp` cannot be specified. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

		Example: `10.0.3.3` 

	* `security_attributes` - (Optional) Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.  Example: `{"Oracle-DataSecurity-ZPR.MaxEgressCount.value": "42", "Oracle-DataSecurity-ZPR.MaxEgressCount.mode": "audit"}` 

	* `skip_source_dest_check` - (Optional) (Updatable) Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you would skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).

		 If you specify a `vlanId`, the `skipSourceDestCheck` cannot be specified because the source/destination check is always disabled for VNICs in a VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

		Example: `true` 
	* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create the VNIC in. When launching an instance, use this `subnetId` instead of the deprecated `subnetId` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). At least one of them is required; if you provide both, the values must match.

		If you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN instead of a subnet, provide a `vlanId` instead of a `subnetId`. If you provide both a `vlanId` and `subnetId`, the request fails. 
	* `vlan_id` - (Optional) Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

		Provide a `vlanId` instead of a `subnetId`. If you provide both a `vlanId` and `subnetId`, the request fails. 
* `dedicated_vm_host_id` - (Optional) (Updatable) The OCID of the dedicated virtual machine host to place the instance on. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `extended_metadata` - (Optional) (Updatable) Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

	They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

	The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 

	Input in terraform is the same as metadata but allows nested metadata if you pass a valid JSON string as a value. See the example above.
* `fault_domain` - (Optional) (Updatable) A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

	If you do not specify the fault domain, the system selects one for you.

	 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - (Optional) Deprecated. Instead use `hostnameLabel` in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/). If you provide both, the values must match. 
* `image` - (Optional) Deprecated. Use `sourceDetails` with [InstanceSourceViaImageDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/InstanceSourceViaImageDetails) source type instead. If you specify values for both, the values must match. 
* `instance_configuration_id` - (Optional) The OCID of the Instance Configuration containing instance launch details. Any other fields supplied in this instance launch request will override the details stored in the Instance Configuration for this instance launch.
* `instance_options` - (Optional) (Updatable) Optional mutable instance options
	* `are_legacy_imds_endpoints_disabled` - (Optional) (Updatable) Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
* `ipxe_script` - (Optional) This is an advanced option.

	When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

	If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots. Be aware that the same iPXE script will run every time an instance boots, not only after the initial LaunchInstance call.

	The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

	If your instance boot volume attachment type is paravirtualized, the boot volume is attached to the instance through virtio-scsi and no iPXE script is used. If your instance boot volume attachment type is paravirtualized and you use custom iPXE to network boot into your instance, the primary boot volume is attached as a data volume through virtio-scsi drive.

	For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

	For more information about iPXE, see http://ipxe.org. 
* `is_pv_encryption_in_transit_enabled` - (Optional) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false. Use this field only during create. To update use `is_pv_encryption_in_transit_enabled` under `launch_options` instead.
* `launch_options` - (Optional) (Updatable) Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
	* `boot_volume_type` - (Optional) (Updatable) Emulation type for the boot volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
	* `firmware` - (Optional) Firmware used to boot VM. Select the option that matches your operating system.
		* `BIOS` - Boot VM using BIOS style firmware. This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
		* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems. This is the default for platform images. 
	* `is_consistent_volume_naming_enabled` - (Optional) Whether to enable consistent volume naming feature. Defaults to false.
	* `is_pv_encryption_in_transit_enabled` - (Optional) (Updatable) Use this for update operation only. This field is  Deprecated during create. For create use `isPvEncryptionInTransitEnabled` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/datatypes/LaunchInstanceDetails). 
	* `network_type` - (Optional) (Updatable) Emulation type for the physical network interface card (NIC).
		* `E1000` - Emulated Gigabit ethernet controller. Compatible with Linux e1000 network driver.
		* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
		* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers. 
	* `remote_data_volume_type` - (Optional) Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
* `launch_volume_attachments` - (Optional) Volume attachments to create as part of the launch instance operation.

     **Note:** This property is used for initial instance provisioning only. Updates to this property will not be supported. To update volume attachments, user should use `oci_core_volume_attachment`. To update volume details, user should use `oci_core_volume`
	
    * `device` - (Optional) The device name. To retrieve a list of devices for a given instance, see [ListInstanceDevices](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Device/ListInstanceDevices).
	* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `encryption_in_transit_type` - (Applicable when type=iscsi) Refer the top-level definition of encryptionInTransitType. The default value is NONE. 
	* `is_agent_auto_iscsi_login_enabled` - (Applicable when type=iscsi) Whether to enable Oracle Cloud Agent to perform the iSCSI login and logout commands after the volume attach or detach operations for non multipath-enabled iSCSI attachments. 
	* `is_pv_encryption_in_transit_enabled` - (Applicable when type=paravirtualized) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
	* `is_read_only` - (Optional) Whether the attachment was created in read-only mode.
	* `is_shareable` - (Optional) Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified. 
	* `launch_create_volume_details` - (Optional) Define a volume that will be created and attached or attached to an instance on creation.
		* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the volume. If not provided,  it will be inherited from the instance. 
		* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `kms_key_id` - (Optional) The OCID of the Vault service key to assign as the master encryption key for the volume. 
		* `size_in_gbs` - (Required) The size of the volume in GBs.
		* `volume_creation_type` - (Required) Specifies the method for volume creation.
		* `vpus_per_gb` - (Optional) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

			Allowed values:
			* `0`: Represents Lower Cost option.
			* `10`: Represents Balanced option.
			* `20`: Represents Higher Performance option.
			* `30`-`120`: Represents the Ultra High Performance option. 
	* `type` - (Required) The type of volume attachment. Currently, the only supported values are "iscsi" and "paravirtualized".
	* `use_chap` - (Applicable when type=iscsi) Whether to use CHAP authentication for the volume attachment. Defaults to false. 
	* `volume_id` - (Optional) The OCID of the volume. If CreateVolumeDetails is specified, this field must be omitted from the request. 
* `licensing_configs` - (Optional) (Updatable) List of licensing configurations associated with target launch values.
	* `license_type` - (Optional) (Updatable) License Type for the OS license.
		* `OCI_PROVIDED` - Oracle Cloud Infrastructure provided license (e.g. metered $/OCPU-hour).
		* `BRING_YOUR_OWN_LICENSE` - Bring your own license. 
	* `type` - (Required) (Updatable) Operating System type of the Configuration.
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

	The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes.
	
	**Note:** Both the 'user_data' and 'ssh_authorized_keys' fields cannot be changed after an instance has launched. Any request which updates, removes, or adds either of these fields will be rejected. You must provide the same values for 'user_data' and 'ssh_authorized_keys' that already exist on the instance.
* `platform_config` - (Optional) (Updatable only for VM's) The platform configuration requested for the instance.

	If you provide the parameter, the instance is created with the platform configuration that you specify. For any values that you omit, the instance uses the default configuration values for the `shape` that you specify. If you don't provide the parameter, the default values for the `shape` are used.

	Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.

	For more information about shielded instances, see [Shielded Instances](https://docs.cloud.oracle.com/iaas/Content/Compute/References/shielded-instances.htm).

	For more information about BIOS settings for bare metal instances, see [BIOS Settings for Bare Metal Instances](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bios-settings.htm). 
	* `are_virtual_instructions_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM) Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes or VT-x for Intel shapes. 
	* `config_map` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) Instance Platform Configuration Configuration Map for flexible setting input. 
	* `is_access_control_service_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM) Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device pass-through. 
	* `is_input_output_memory_management_unit_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) Whether the input-output memory management unit is enabled. 
	* `is_measured_boot_enabled` - (Optional) Whether the Measured Boot feature is enabled on the instance. 
	* `is_memory_encryption_enabled` - (Optional) Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`. 
	* `is_secure_boot_enabled` - (Optional) Whether Secure Boot is enabled on the instance. 
	* `is_symmetric_multi_threading_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | AMD_VM | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM | INTEL_VM) (Updatable only for INTEL_VM and AMD_VM) Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also called simultaneous multithreading (SMT) or Intel Hyper-Threading.

		Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple independent threads of execution, to better use the resources and increase the efficiency of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which can provide higher or more predictable performance for some workloads. 
	* `is_trusted_platform_module_enabled` - (Optional) Whether the Trusted Platform Module (TPM) is enabled on the instance. 
	* `numa_nodes_per_socket` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) The number of NUMA nodes per socket (NPS). 
	* `percentage_of_cores_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) The percentage of cores enabled. Value must be a multiple of 25%. If the requested percentage results in a fractional number of cores, the system rounds up the number of cores across processors and provisions an instance with a whole number of cores.

		If the applications that you run on the instance use a core-based licensing model and need fewer cores than the full size of the shape, you can disable cores to reduce your licensing costs. The instance itself is billed for the full shape, regardless of whether all cores are enabled. 
	* `type` - (Required) The type of platform being configured. 
* `preemptible_instance_config` - (Optional) Configuration options for preemptible instances. 
	* `preemption_action` - (Required) The action to run when the preemptible instance is interrupted for eviction. 
		* `preserve_boot_volume` - (Optional) Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
		* `type` - (Required) The type of action to run when the instance is interrupted for eviction.
* `security_attributes` - (Optional) (Updatable) [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
* `shape` - (Optional) (Updatable) The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

	You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
* `shape_config` - (Optional) (Updatable) The shape configuration requested for the instance.

	If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

	Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
	* `baseline_ocpu_utilization` - (Optional) (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

		The following values are supported:
		* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
		* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
		* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
	* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory available to the instance, in gigabytes. 
	* `nvmes` - (Optional) (Updatable) The number of NVMe drives to be used for storage. A single drive has 6.8 TB available. 
	* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to the instance. 
	* `vcpus` - (Optional) (Updatable) The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2.
* `source_details` - (Optional) (Updatable) 
	* `boot_volume_size_in_gbs` - (Applicable when source_type=image) (Updatable) The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 32,768 GB (32 TB). 
	* `boot_volume_vpus_per_gb` - (Applicable when source_type=image) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

		Allowed values:
		* `10`: Represents Balanced option.
		* `20`: Represents Higher Performance option.
		* `30`-`120`: Represents the Ultra High Performance option.

		For volumes with the auto-tuned performance feature enabled, this is set to the default (minimum) VPUs/GB. 
	* `instance_source_image_filter_details` - (Applicable when source_type=image) These are the criteria for selecting an image. This is required if imageId is not specified.
		* `compartment_id` - (Required when source_type=image) (Updatable) The OCID of the compartment containing images to search
		* `defined_tags_filter` - (Applicable when source_type=image) Filter based on these defined tags. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
		* `operating_system` - (Applicable when source_type=image) The image's operating system.  Example: `Oracle Linux` 
		* `operating_system_version` - (Applicable when source_type=image) The image's operating system version.  Example: `7.2`
	* `kms_key_id` - (Applicable when source_type=image) (Updatable) The OCID of the Vault service key to assign as the master encryption key for the boot volume.
	* `source_id` - (Required) (Updatable) The OCID of the boot volume used to boot the instance. Updates are supported only for linux Images. The user will need to manually destroy and re-create the resource for other image types.
	* `source_type` - (Required) (Updatable) The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID.
    * `is_preserve_boot_volume_enabled` - (Optional) (Updatable) Whether to preserve the boot volume that was previously attached to the instance after a successful replacement of that boot volume.
* `subnet_id` - (Optional) Deprecated. Instead use `subnetId` in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/). At least one of them is required; if you provide both, the values must match. 
* `state` - (Optional) (Updatable) The target state for the instance. Could be set to RUNNING or STOPPED.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `agent_config` - Configuration options for the Oracle Cloud Agent software running on the instance.
	* `are_all_plugins_disabled` - Whether Oracle Cloud Agent can run all of the available plugins. This includes the management and monitoring plugins.

		For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
	* `is_management_disabled` - Whether Oracle Cloud Agent can run all the available management plugins.

		These are the management plugins: OS Management Service Agent and Compute Instance Run Command.

		The management plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
		* If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of the per-plugin configuration.
		* If `isManagementDisabled` is false, all of the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
	* `is_monitoring_disabled` - Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins.

		These are the monitoring plugins: Compute Instance Monitoring and Custom Logs Monitoring.

		The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
		* If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of the per-plugin configuration.
		* If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
	* `plugins_config` - The configuration of plugins associated with this instance.
		* `desired_state` - Whether the plugin should be enabled or disabled.

			To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
		* `name` - The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
* `availability_config` - Options for defining the availabiity of a VM instance after a maintenance event that impacts the underlying hardware. 
	* `is_live_migration_preferred` - Whether to live migrate supported VM instances to a healthy physical VM host without disrupting running instances during infrastructure maintenance events. If null, Oracle chooses the best option for migrating the VM during infrastructure maintenance events. 
	* `recovery_action` - The lifecycle state for an instance when it is recovered after infrastructure maintenance.
		* `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event. If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
		* `STOP_INSTANCE` - The instance is recovered in the stopped state. 
* `availability_domain` - The availability domain the instance is running in.  Example: `Uocm:PHX-AD-1`
* `capacity_reservation_id` - The OCID of the compute capacity reservation this instance is launched under. When this field contains an empty string or is null, the instance is not currently in a capacity reservation. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
* `boot_volume_id` - The OCID of the attached boot volume. If the `source_type` is `bootVolume`, this will be the same OCID as the `source_id`.
* `capacity_reservation_id` - The OCID of the compute capacity reservation this instance is launched under. When this field contains an empty string or is null, the instance is not currently in a capacity reservation. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
* `compartment_id` - The OCID of the compartment that contains the instance.
* `dedicated_vm_host_id` - The OCID of the dedicated virtual machine host that the instance is placed on. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

	They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only). 

	Input in terraform is the same as metadata but allows nested metadata if you pass a valid JSON string as a value. See the example below.
* `fault_domain` - The name of the fault domain the instance is running in.

	A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

	If you do not specify the fault domain, the system selects one for you.

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the instance.
* `image` - Deprecated. Use `sourceDetails` instead. 
* `instance_configuration_id` - The OCID of the Instance Configuration used to source launch details for this instance. Any other fields supplied in the instance launch request override the details stored in the Instance Configuration for this instance launch.
* `instance_options` - Optional mutable instance options
	* `are_legacy_imds_endpoints_disabled` - Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
* `ipxe_script` - When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

	If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots. Be aware that the same iPXE script will run every time an instance boots, not only after the initial LaunchInstance call.

	The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

	If your instance boot volume attachment type is paravirtualized, the boot volume is attached to the instance through virtio-scsi and no iPXE script is used. If your instance boot volume attachment type is paravirtualized and you use custom iPXE to network boot into your instance, the primary boot volume is attached as a data volume through virtio-scsi drive.

	For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

	For more information about iPXE, see http://ipxe.org. 
* `is_cross_numa_node` - Whether the instance’s OCPUs and memory are distributed across multiple NUMA nodes. 
* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
	* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for platform images.
	* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
	* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
	* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
* `launch_options` - Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
	* `boot_volume_type` - Emulation type for the boot volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
	* `firmware` - Firmware used to boot VM. Select the option that matches your operating system.
		* `BIOS` - Boot VM using BIOS style firmware. This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
		* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems. This is the default for platform images. 
	* `is_consistent_volume_naming_enabled` - Whether to enable consistent volume naming feature. Defaults to false.
	* `is_pv_encryption_in_transit_enabled` - Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/datatypes/LaunchInstanceDetails). 
	* `network_type` - Emulation type for the physical network interface card (NIC).
		* `E1000` - Emulated Gigabit ethernet controller. Compatible with Linux e1000 network driver.
		* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
		* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers. 
	* `remote_data_volume_type` - Emulation type for volume.
		* `ISCSI` - ISCSI attached block storage device.
		* `SCSI` - Emulated SCSI disk.
		* `IDE` - Emulated IDE disk.
		* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
		* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
* `licensing_configs` - List of licensing configurations associated with the instance.
	* `license_type` - License Type for the OS license.
		* `OCI_PROVIDED` - Oracle Cloud Infrastructure provided license (e.g. metered $/OCPU-hour).
		* `BRING_YOUR_OWN_LICENSE` - Bring your own license. 
	* `os_version` - The Operating System version of the license config.
	* `type` - Operating System type of the Configuration.
* `metadata` - Custom metadata that you provide.

* `platform_config` - The platform configuration for the instance. 
	* `are_virtual_instructions_enabled` - Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes or VT-x for Intel shapes. 
	* `config_map` - Instance Platform Configuration Configuration Map for flexible setting input. 
	* `is_access_control_service_enabled` - Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device pass-through. 
	* `is_input_output_memory_management_unit_enabled` - Whether the input-output memory management unit is enabled. 
	* `is_measured_boot_enabled` - Whether the Measured Boot feature is enabled on the instance. 
	* `is_memory_encryption_enabled` - Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`. 
	* `is_secure_boot_enabled` - Whether Secure Boot is enabled on the instance. 
	* `is_symmetric_multi_threading_enabled` - Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also called simultaneous multithreading (SMT) or Intel Hyper-Threading.

		Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple independent threads of execution, to better use the resources and increase the efficiency of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which can provide higher or more predictable performance for some workloads. 
	* `is_trusted_platform_module_enabled` - Whether the Trusted Platform Module (TPM) is enabled on the instance. 

* `platform_config` - The platform configuration for the instance.
	* `are_virtual_instructions_enabled` - Whether virtualization instructions are available.
	* `is_access_control_service_enabled` - Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device passthrough.
	* `is_input_output_memory_management_unit_enabled` - Whether the input-output memory management unit is enabled.
	* `is_measured_boot_enabled` - Whether the Measured Boot feature is enabled on the instance.
	* `is_secure_boot_enabled` - Whether Secure Boot is enabled on the instance.
	* `is_symmetric_multi_threading_enabled` - Whether symmetric multi-threading is enabled on the instance.
	* `is_trusted_platform_module_enabled` - Whether the Trusted Platform Module (TPM) is enabled on the instance.
	* `numa_nodes_per_socket` - The number of NUMA nodes per socket (NPS). 
	* `percentage_of_cores_enabled` - The percentage of cores enabled.
	* `type` - The type of platform being configured. (Supported types=[INTEL_VM, AMD_MILAN_BM, AMD_ROME_BM, AMD_ROME_BM_GPU, INTEL_ICELAKE_BM, INTEL_SKYLAKE_BM])
* `preemptible_instance_config` - (Optional) Configuration options for preemptible instances. 
	* `preemption_action` - (Required) The action to run when the preemptible instance is interrupted for eviction. 
		* `preserve_boot_volume` - (Optional) Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
		* `type` - (Required) The type of action to run when the instance is interrupted for eviction.
* `private_ip` - The private IP address of instance VNIC. To set the private IP address, use the `private_ip` argument in create_vnic_details.
* `public_ip` - The public IP address of instance VNIC (if enabled).
* `region` - The region that contains the availability domain the instance is running in.

	For the us-phoenix-1 and us-ashburn-1 regions, `phx` and `iad` are returned, respectively. For all other regions, the full region name is returned.

	Examples: `phx`, `eu-frankfurt-1`
* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
* `security_attributes_state` - The lifecycle state of the `securityAttributes`
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance. You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
* `shape_config` - The shape configuration for an instance. The shape configuration determines the resources allocated to an instance. 
	* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

		The following values are supported:
		* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
		* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
		* `BASELINE_1_1` - baseline usage is the entire OCPU. This represents a non-burstable instance. 
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
	* `vcpus` - The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2.
* `source_details` - 
	* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 32,768 GB (32 TB). 
	* `boot_volume_vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

		Allowed values:
		* `10`: Represents Balanced option.
		* `20`: Represents Higher Performance option.
		* `30`-`120`: Represents the Ultra High Performance option.

		For volumes with the auto-tuned performance feature enabled, this is set to the default (minimum) VPUs/GB. 
	* `instance_source_image_filter_details` - These are the criteria for selecting an image. This is required if imageId is not specified.
		* `compartment_id` - The OCID of the compartment containing images to search
		* `defined_tags_filter` - Filter based on these defined tags. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
		* `operating_system` - The image's operating system.  Example: `Oracle Linux` 
		* `operating_system_version` - The image's operating system version.  Example: `7.2` 
	* `kms_key_id` - The OCID of the Vault service key to assign as the master encryption key for the boot volume.
	* `source_id` - The OCID of an image or a boot volume to use, depending on the value of `source_type`.
	* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
* `state` - The current state of the instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the instance was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_maintenance_reboot_due` - The date and time the instance is expected to be stopped / started,  in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). After that time if instance hasn't been rebooted, Oracle will reboot the instance within 24 hours of the due time. Regardless of how the instance was stopped, the flag will be reset to empty as soon as instance reaches Stopped state. Example: `2018-05-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 45 minutes), when creating the Instance
	* `update` - (Defaults to 45 minutes), when updating the Instance
	* `delete` - (Defaults to 75 minutes), when destroying the Instance


## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance.test_instance "id"
```
