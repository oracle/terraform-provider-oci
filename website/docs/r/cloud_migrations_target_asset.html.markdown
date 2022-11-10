---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_target_asset"
sidebar_current: "docs-oci-resource-cloud_migrations-target_asset"
description: |-
  Provides the Target Asset resource in Oracle Cloud Infrastructure Cloud Migrations service
---

# oci_cloud_migrations_target_asset
This resource provides the Target Asset resource in Oracle Cloud Infrastructure Cloud Migrations service.

Creates a target asset.


## Example Usage

```hcl
resource "oci_cloud_migrations_target_asset" "test_target_asset" {
	#Required
	is_excluded_from_execution = var.target_asset_is_excluded_from_execution
	migration_plan_id = oci_cloud_migrations_migration_plan.test_migration_plan.id
	preferred_shape_type = var.target_asset_preferred_shape_type
	type = var.target_asset_type
	user_spec {

		#Optional
		agent_config {

			#Optional
			are_all_plugins_disabled = var.target_asset_user_spec_agent_config_are_all_plugins_disabled
			is_management_disabled = var.target_asset_user_spec_agent_config_is_management_disabled
			is_monitoring_disabled = var.target_asset_user_spec_agent_config_is_monitoring_disabled
			plugins_config {
				#Required
				desired_state = var.target_asset_user_spec_agent_config_plugins_config_desired_state
				name = var.target_asset_user_spec_agent_config_plugins_config_name
			}
		}
		availability_domain = var.target_asset_user_spec_availability_domain
		capacity_reservation_id = oci_cloud_migrations_capacity_reservation.test_capacity_reservation.id
		compartment_id = var.compartment_id
		create_vnic_details {

			#Optional
			assign_private_dns_record = var.target_asset_user_spec_create_vnic_details_assign_private_dns_record
			assign_public_ip = var.target_asset_user_spec_create_vnic_details_assign_public_ip
			defined_tags = {"foo-namespace.bar-key"= "value"}
			display_name = var.target_asset_user_spec_create_vnic_details_display_name
			freeform_tags = {"bar-key"= "value"}
			hostname_label = var.target_asset_user_spec_create_vnic_details_hostname_label
			nsg_ids = var.target_asset_user_spec_create_vnic_details_nsg_ids
			private_ip = var.target_asset_user_spec_create_vnic_details_private_ip
			skip_source_dest_check = var.target_asset_user_spec_create_vnic_details_skip_source_dest_check
			subnet_id = oci_core_subnet.test_subnet.id
			vlan_id = oci_core_vlan.test_vlan.id
		}
		dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
		defined_tags = {"foo-namespace.bar-key"= "value"}
		display_name = var.target_asset_user_spec_display_name
		fault_domain = var.target_asset_user_spec_fault_domain
		freeform_tags = {"bar-key"= "value"}
		hostname_label = var.target_asset_user_spec_hostname_label
		instance_options {

			#Optional
			are_legacy_imds_endpoints_disabled = var.target_asset_user_spec_instance_options_are_legacy_imds_endpoints_disabled
		}
		ipxe_script = var.target_asset_user_spec_ipxe_script
		is_pv_encryption_in_transit_enabled = var.target_asset_user_spec_is_pv_encryption_in_transit_enabled
		preemptible_instance_config {
			#Required
			preemption_action {
				#Required
				type = var.target_asset_user_spec_preemptible_instance_config_preemption_action_type

				#Optional
				preserve_boot_volume = var.target_asset_user_spec_preemptible_instance_config_preemption_action_preserve_boot_volume
			}
		}
		shape = var.target_asset_user_spec_shape
		shape_config {

			#Optional
			baseline_ocpu_utilization = var.target_asset_user_spec_shape_config_baseline_ocpu_utilization
			memory_in_gbs = var.target_asset_user_spec_shape_config_memory_in_gbs
			ocpus = var.target_asset_user_spec_shape_config_ocpus
		}
		source_details {
			#Required
			source_type = var.target_asset_user_spec_source_details_source_type

			#Optional
			boot_volume_id = oci_core_boot_volume.test_boot_volume.id
			boot_volume_size_in_gbs = var.target_asset_user_spec_source_details_boot_volume_size_in_gbs
			boot_volume_vpus_per_gb = var.target_asset_user_spec_source_details_boot_volume_vpus_per_gb
			image_id = oci_core_image.test_image.id
			kms_key_id = oci_kms_key.test_key.id
		}
	}

	#Optional
	block_volumes_performance = var.target_asset_block_volumes_performance
	ms_license = var.target_asset_ms_license
}
```

## Argument Reference

The following arguments are supported:

* `block_volumes_performance` - (Optional) (Updatable) Performance of the block volumes.
* `is_excluded_from_execution` - (Required) (Updatable) A boolean indicating whether the asset should be migrated.
* `migration_plan_id` - (Required) OCID of the associated migration plan.
* `ms_license` - (Optional) (Updatable) Microsoft license for the VM configuration.
* `preferred_shape_type` - (Required) (Updatable) Preferred VM shape type that you provide.
* `type` - (Required) (Updatable) The type of target asset.
* `user_spec` - (Required) (Updatable) Instance launch details. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance. 
	* `agent_config` - (Optional) (Updatable) Configuration options for the Oracle Cloud Agent software running on the instance.
		* `are_all_plugins_disabled` - (Optional) (Updatable) Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

			To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
		* `is_management_disabled` - (Optional) (Updatable) Whether Oracle Cloud Agent can run all the available management plugins. By default, the value is false (management plugins are enabled).

			These are the management plugins: OS Management Service Agent and Compute instance run command.

			The management plugins are controlled by this parameter and the per-plugin configuration in the `pluginsConfig` object.
			* If `isManagementDisabled` is true, all the management plugins are disabled, regardless of the per-plugin configuration.
			* If `isManagementDisabled` is false, all the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
		* `is_monitoring_disabled` - (Optional) (Updatable) Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. By default, the value is false (monitoring plugins are enabled).

			These are the monitoring plugins: Compute instance monitoring and Custom logs monitoring.

			The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
			* If `isMonitoringDisabled` is true, all the monitoring plugins are disabled, regardless of the per-plugin configuration.
			* If `isMonitoringDisabled` is false, all the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
		* `plugins_config` - (Optional) (Updatable) The configuration of plugins associated with this instance.
			* `desired_state` - (Required) (Updatable) Whether the plugin should be enabled or disabled.

				To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
			* `name` - (Required) (Updatable) The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
	* `availability_domain` - (Optional) (Updatable) The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
	* `capacity_reservation_id` - (Optional) (Updatable) The OCID of the compute capacity reservation under which this instance is launched. You can opt out of all default reservations by specifying an empty string as input for this field. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default). 
	* `compartment_id` - (Optional) (Updatable) The OCID of the compartment.
	* `create_vnic_details` - (Optional) (Updatable) Contains properties for a VNIC. You use this object when creating the primary VNIC during instance launch or when creating a secondary VNIC. For more information about VNICs, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
		* `assign_private_dns_record` - (Optional) (Updatable) Whether the VNIC should be assigned a DNS record. If set to false, there will be no DNS record registration for the VNIC. If set to true, the DNS record will be registered. By default, the value is true.

			If you specify a `hostnameLabel`, then `assignPrivateDnsRecord` must be set to true. 
		* `assign_public_ip` - (Optional) (Updatable) Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the [Subnet](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Subnet/)), then no public IP address is assigned. If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then a public IP address is assigned. If set to true and `prohibitPublicIpOnVnic` = true, an error is returned.

			**Note:** This public IP address is associated with the primary private IP on the VNIC. For more information, see [IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).

			**Note:** There's a limit to the number of [public IPs](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/) a VNIC or instance can have. If you try to create a secondary VNIC with an assigned public IP for an instance that has already reached its public IP limit, an error is returned. For information about the public IP limits, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).

			Example: `false`

			If you specify a `vlanId`, then `assignPublicIp` must be set to false. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
		* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
		* `hostname_label` - (Optional) (Updatable) The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

			For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

			When launching an instance, use this `hostnameLabel` instead of the deprecated `hostnameLabel` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). If you provide both, the values must match.

			Example: `bminstance-1`

			If you specify a `vlanId`, the `hostnameLabel` cannot be specified. VNICs on a VLAN can not be assigned a hostname. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `nsg_ids` - (Optional) (Updatable) List of OCIDs of the network security groups (NSGs) that are added to the VNIC. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).

			If a `vlanId` is specified, the `nsgIds` cannot be specified. The `vlanId` indicates that the VNIC will belong to a VLAN instead of a subnet. With VLANs, all VNICs in the VLAN belong to the NSGs that are associated with the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `private_ip` - (Optional) (Updatable) A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's *primary* private IP address. The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

			 If you specify a `vlanId`, the `privateIp` cannot be specified. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Example: `10.0.3.3` 
		* `skip_source_dest_check` - (Optional) (Updatable) Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you should skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).

			 If you specify a `vlanId`, the `skipSourceDestCheck` cannot be specified because the source/destination check is always disabled for VNICs in a VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Example: `true` 
		* `subnet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create the VNIC. When launching an instance, use this `subnetId` instead of the deprecated `subnetId` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). At least one of them is required; if you provide both, the values must match.

			If you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN instead of a subnet, provide a `vlanId` instead of a `subnetId`. If you provide both `vlanId` and `subnetId`, the request fails. 
		* `vlan_id` - (Optional) (Updatable) Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Provide a `vlanId` instead of a `subnetId`. If you provide both `vlanId` and `subnetId`, the request fails. 
	* `dedicated_vm_host_id` - (Optional) (Updatable) The OCID of the dedicated VM host. 
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `fault_domain` - (Optional) (Updatable) A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains lets you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

		If you do not specify the fault domain, the system selects one for you.

		 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

		Example: `FAULT-DOMAIN-1` 
	* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
	* `hostname_label` - (Optional) (Updatable) Deprecated. Instead use `hostnameLabel` in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/). If you provide both, the values must match. 
	* `instance_options` - (Optional) (Updatable) Optional mutable instance options
		* `are_legacy_imds_endpoints_disabled` - (Optional) (Updatable) Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
	* `ipxe_script` - (Optional) (Updatable) This is an advanced option.

		When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

		If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots. Be aware that the same iPXE script will run every time an instance boots, not only after the initial LaunchInstance call.

		By default, the iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI in the same way as the default iPXE script, use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

		If your instance boot volume type is paravirtualized, the boot volume is attached to the instance through virtio-scsi and no iPXE script is used. If your instance boot volume type is paravirtualized and you use custom iPXE to perform network-boot into your instance, the primary boot volume is attached as a data volume through the virtio-scsi drive.

		For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

		For more information about iPXE, see http://ipxe.org. 
	* `is_pv_encryption_in_transit_enabled` - (Optional) (Updatable) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. By default, the value is false.
	* `preemptible_instance_config` - (Optional) (Updatable) Configuration options for preemptible instances. 
		* `preemption_action` - (Required) (Updatable) The action to run when the preemptible instance is interrupted for eviction. 
			* `preserve_boot_volume` - (Optional) (Updatable) Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. By default, it is false if not specified. 
			* `type` - (Required) (Updatable) The type of action to run when the instance is interrupted for eviction.
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
		* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory in gigabytes that is available to the instance. 
		* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to the instance. 
	* `source_details` - (Optional) (Updatable) 
		* `boot_volume_id` - (Required when source_type=bootVolume) (Updatable) The OCID of the boot volume used to boot the instance.
		* `boot_volume_size_in_gbs` - (Applicable when source_type=image) (Updatable) The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB). 
		* `boot_volume_vpus_per_gb` - (Applicable when source_type=image) (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB that represents the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

			Allowed values:
			* `10`: Represents Balanced option.
			* `20`: Represents Higher Performance option.
			* `30`-`120`: Represents the Ultra High Performance option.

			For volumes with the auto-tuned performance feature enabled, this is set to the default (minimum) VPUs/GB. 
		* `image_id` - (Required when source_type=image) (Updatable) The OCID of the image used to boot the instance.
		* `kms_key_id` - (Applicable when source_type=image) (Updatable) The OCID of the key management key to assign as the master encryption key for the boot volume.
		* `source_type` - (Required) (Updatable) The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `block_volumes_performance` - Performance of the block volumes.
* `compartment_id` - Compartment identifier
* `compatibility_messages` - Messages about the compatibility issues.
	* `message` - Detailed description of the compatibility issue.
	* `name` - Name of the compatibility issue.
	* `severity` - Severity level of the compatibility issue.
* `created_resource_id` - Created resource identifier
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `estimated_cost` - Cost estimation description
	* `compute` - Cost estimation for compute
		* `gpu_count` - Total number of GPU
		* `gpu_per_hour` - GPU per hour
		* `gpu_per_hour_by_subscription` - GPU per hour by subscription
		* `memory_amount_gb` - Total usage of memory
		* `memory_gb_per_hour` - Gigabyte per hour
		* `memory_gb_per_hour_by_subscription` - Gigabyte per hour by subscription
		* `ocpu_count` - Total number of OCPUs
		* `ocpu_per_hour` - OCPU per hour
		* `ocpu_per_hour_by_subscription` - OCPU per hour by subscription
		* `total_per_hour` - Total per hour
		* `total_per_hour_by_subscription` - Total usage per hour by subscription
	* `currency_code` - Currency code in the ISO format.
	* `os_image` - Cost estimation for the OS image.
		* `total_per_hour` - Total price per hour
		* `total_per_hour_by_subscription` - Total price per hour by subscription
	* `storage` - Cost estimation for storage
		* `total_gb_per_month` - Gigabyte storage capacity per month.
		* `total_gb_per_month_by_subscription` - Gigabyte storage capacity per month by subscription.
		* `volumes` - Volume estimation
			* `capacity_gb` - Gigabyte storage capacity
			* `description` - Volume description
			* `total_gb_per_month` - Gigabyte storage capacity per month.
			* `total_gb_per_month_by_subscription` - Gigabyte storage capacity per month by subscription
	* `subscription_id` - Subscription ID
	* `total_estimation_per_month` - Total estimation per month
	* `total_estimation_per_month_by_subscription` - Total estimation per month by subscription.
* `id` - Unique identifier that is immutable on creation.
* `is_excluded_from_execution` - A boolean indicating whether the asset should be migrated.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `migration_asset` - Description of the migration asset.
	* `availability_domain` - Availability domain
	* `compartment_id` - Compartment Identifier
	* `depended_on_by` - List of migration assets that depend on the asset.
	* `depends_on` - List of migration assets that depends on the asset.
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	* `id` - Asset ID generated by mirgration service. It is used in the mirgration service pipeline.
	* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	* `migration_id` - OCID of the associated migration.
	* `notifications` - List of notifications
	* `parent_snapshot` - The parent snapshot of the migration asset to be used by the replication task.
	* `replication_compartment_id` - Replication compartment identifier
	* `replication_schedule_id` - Replication schedule identifier
	* `snap_shot_bucket_name` - Name of snapshot bucket
	* `snapshots` - Key-value pair representing disks ID mapped to the OCIDs of replicated or hydration server volume snapshots. Example: `{"bar-key": "value"}` 
		* `unmodified_volume_id` - ID of the unmodified volume
		* `uuid` - ID of the vCenter disk obtained from Inventory.
		* `volume_id` - ID of the hydration server volume
		* `volume_type` - The hydration server volume type
	* `source_asset_data` - Key-value pair representing asset metadata keys and values scoped to a namespace. Example: `{"bar-key": "value"}` 
	* `source_asset_id` - OCID that is referenced to an asset for an inventory.
	* `state` - The current state of the migration asset.
	* `tenancy_id` - Tenancy identifier
	* `time_created` - The time when the migration asset was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time when the migration asset was updated. An RFC3339 formatted datetime string.
	* `type` - The type of asset referenced for inventory.
* `migration_plan_id` - OCID of the associated migration plan.
* `ms_license` - Microsoft license for VM configuration.
* `preferred_shape_type` - Preferred VM shape type that you provide.
* `recommended_spec` - Instance launch details. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance. 
	* `agent_config` - Configuration options for the Oracle Cloud Agent software running on the instance.
		* `are_all_plugins_disabled` - Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

			To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
		* `is_management_disabled` - Whether Oracle Cloud Agent can run all the available management plugins. By default, the value is false (management plugins are enabled).

			These are the management plugins: OS Management Service Agent and Compute instance run command.

			The management plugins are controlled by this parameter and the per-plugin configuration in the `pluginsConfig` object.
			* If `isManagementDisabled` is true, all the management plugins are disabled, regardless of the per-plugin configuration.
			* If `isManagementDisabled` is false, all the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
		* `is_monitoring_disabled` - Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. By default, the value is false (monitoring plugins are enabled).

			These are the monitoring plugins: Compute instance monitoring and Custom logs monitoring.

			The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
			* If `isMonitoringDisabled` is true, all the monitoring plugins are disabled, regardless of the per-plugin configuration.
			* If `isMonitoringDisabled` is false, all the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
		* `plugins_config` - The configuration of plugins associated with this instance.
			* `desired_state` - Whether the plugin should be enabled or disabled.

				To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
			* `name` - The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
	* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
	* `capacity_reservation_id` - The OCID of the compute capacity reservation under which this instance is launched. You can opt out of all default reservations by specifying an empty string as input for this field. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default). 
	* `compartment_id` - The OCID of the compartment.
	* `create_vnic_details` - Contains properties for a VNIC. You use this object when creating the primary VNIC during instance launch or when creating a secondary VNIC. For more information about VNICs, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
		* `assign_private_dns_record` - Whether the VNIC should be assigned a DNS record. If set to false, there will be no DNS record registration for the VNIC. If set to true, the DNS record will be registered. By default, the value is true.

			If you specify a `hostnameLabel`, then `assignPrivateDnsRecord` must be set to true. 
		* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the [Subnet](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Subnet/)), then no public IP address is assigned. If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then a public IP address is assigned. If set to true and `prohibitPublicIpOnVnic` = true, an error is returned.

			**Note:** This public IP address is associated with the primary private IP on the VNIC. For more information, see [IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).

			**Note:** There's a limit to the number of [public IPs](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/) a VNIC or instance can have. If you try to create a secondary VNIC with an assigned public IP for an instance that has already reached its public IP limit, an error is returned. For information about the public IP limits, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).

			Example: `false`

			If you specify a `vlanId`, then `assignPublicIp` must be set to false. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
		* `hostname_label` - The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

			For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

			When launching an instance, use this `hostnameLabel` instead of the deprecated `hostnameLabel` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). If you provide both, the values must match.

			Example: `bminstance-1`

			If you specify a `vlanId`, the `hostnameLabel` cannot be specified. VNICs on a VLAN can not be assigned a hostname. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `nsg_ids` - List of OCIDs of the network security groups (NSGs) that are added to the VNIC. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).

			If a `vlanId` is specified, the `nsgIds` cannot be specified. The `vlanId` indicates that the VNIC will belong to a VLAN instead of a subnet. With VLANs, all VNICs in the VLAN belong to the NSGs that are associated with the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `private_ip` - A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's *primary* private IP address. The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

			 If you specify a `vlanId`, the `privateIp` cannot be specified. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Example: `10.0.3.3` 
		* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you should skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).

			 If you specify a `vlanId`, the `skipSourceDestCheck` cannot be specified because the source/destination check is always disabled for VNICs in a VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Example: `true` 
		* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create the VNIC. When launching an instance, use this `subnetId` instead of the deprecated `subnetId` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). At least one of them is required; if you provide both, the values must match.

			If you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN instead of a subnet, provide a `vlanId` instead of a `subnetId`. If you provide both `vlanId` and `subnetId`, the request fails. 
		* `vlan_id` - Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Provide a `vlanId` instead of a `subnetId`. If you provide both `vlanId` and `subnetId`, the request fails. 
	* `dedicated_vm_host_id` - The OCID of the dedicated VM host. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `fault_domain` - A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains lets you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

		If you do not specify the fault domain, the system selects one for you.

		 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

		Example: `FAULT-DOMAIN-1` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
	* `hostname_label` - Deprecated. Instead use `hostnameLabel` in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/). If you provide both, the values must match. 
	* `instance_options` - Optional mutable instance options
		* `are_legacy_imds_endpoints_disabled` - Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
	* `ipxe_script` - This is an advanced option.

		When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

		If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots. Be aware that the same iPXE script will run every time an instance boots, not only after the initial LaunchInstance call.

		By default, the iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI in the same way as the default iPXE script, use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

		If your instance boot volume type is paravirtualized, the boot volume is attached to the instance through virtio-scsi and no iPXE script is used. If your instance boot volume type is paravirtualized and you use custom iPXE to perform network-boot into your instance, the primary boot volume is attached as a data volume through the virtio-scsi drive.

		For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

		For more information about iPXE, see http://ipxe.org. 
	* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. By default, the value is false.
	* `preemptible_instance_config` - Configuration options for preemptible instances. 
		* `preemption_action` - The action to run when the preemptible instance is interrupted for eviction. 
			* `preserve_boot_volume` - Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. By default, it is false if not specified. 
			* `type` - The type of action to run when the instance is interrupted for eviction.
	* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

		You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
	* `shape_config` - The shape configuration requested for the instance.

		If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

		Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
		* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

			The following values are supported:
			* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
			* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
			* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - The total amount of memory in gigabytes that is available to the instance. 
		* `ocpus` - The total number of OCPUs available to the instance. 
	* `source_details` - 
		* `boot_volume_id` - The OCID of the boot volume used to boot the instance.
		* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB). 
		* `boot_volume_vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB that represents the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

			Allowed values:
			* `10`: Represents Balanced option.
			* `20`: Represents Higher Performance option.
			* `30`-`120`: Represents the Ultra High Performance option.

			For volumes with the auto-tuned performance feature enabled, this is set to the default (minimum) VPUs/GB. 
		* `image_id` - The OCID of the image used to boot the instance.
		* `kms_key_id` - The OCID of the key management key to assign as the master encryption key for the boot volume.
		* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
* `state` - The current state of the target asset.
* `test_spec` - Instance launch details. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance. 
	* `agent_config` - Configuration options for the Oracle Cloud Agent software running on the instance.
		* `are_all_plugins_disabled` - Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

			To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
		* `is_management_disabled` - Whether Oracle Cloud Agent can run all the available management plugins. By default, the value is false (management plugins are enabled).

			These are the management plugins: OS Management Service Agent and Compute instance run command.

			The management plugins are controlled by this parameter and the per-plugin configuration in the `pluginsConfig` object.
			* If `isManagementDisabled` is true, all the management plugins are disabled, regardless of the per-plugin configuration.
			* If `isManagementDisabled` is false, all the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
		* `is_monitoring_disabled` - Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. By default, the value is false (monitoring plugins are enabled).

			These are the monitoring plugins: Compute instance monitoring and Custom logs monitoring.

			The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
			* If `isMonitoringDisabled` is true, all the monitoring plugins are disabled, regardless of the per-plugin configuration.
			* If `isMonitoringDisabled` is false, all the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
		* `plugins_config` - The configuration of plugins associated with this instance.
			* `desired_state` - Whether the plugin should be enabled or disabled.

				To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
			* `name` - The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
	* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
	* `capacity_reservation_id` - The OCID of the compute capacity reservation under which this instance is launched. You can opt out of all default reservations by specifying an empty string as input for this field. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default). 
	* `compartment_id` - The OCID of the compartment.
	* `create_vnic_details` - Contains properties for a VNIC. You use this object when creating the primary VNIC during instance launch or when creating a secondary VNIC. For more information about VNICs, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
		* `assign_private_dns_record` - Whether the VNIC should be assigned a DNS record. If set to false, there will be no DNS record registration for the VNIC. If set to true, the DNS record will be registered. By default, the value is true.

			If you specify a `hostnameLabel`, then `assignPrivateDnsRecord` must be set to true. 
		* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the [Subnet](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Subnet/)), then no public IP address is assigned. If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then a public IP address is assigned. If set to true and `prohibitPublicIpOnVnic` = true, an error is returned.

			**Note:** This public IP address is associated with the primary private IP on the VNIC. For more information, see [IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).

			**Note:** There's a limit to the number of [public IPs](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/) a VNIC or instance can have. If you try to create a secondary VNIC with an assigned public IP for an instance that has already reached its public IP limit, an error is returned. For information about the public IP limits, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).

			Example: `false`

			If you specify a `vlanId`, then `assignPublicIp` must be set to false. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
		* `hostname_label` - The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

			For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

			When launching an instance, use this `hostnameLabel` instead of the deprecated `hostnameLabel` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). If you provide both, the values must match.

			Example: `bminstance-1`

			If you specify a `vlanId`, the `hostnameLabel` cannot be specified. VNICs on a VLAN can not be assigned a hostname. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `nsg_ids` - List of OCIDs of the network security groups (NSGs) that are added to the VNIC. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).

			If a `vlanId` is specified, the `nsgIds` cannot be specified. The `vlanId` indicates that the VNIC will belong to a VLAN instead of a subnet. With VLANs, all VNICs in the VLAN belong to the NSGs that are associated with the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `private_ip` - A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's *primary* private IP address. The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

			 If you specify a `vlanId`, the `privateIp` cannot be specified. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Example: `10.0.3.3` 
		* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you should skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).

			 If you specify a `vlanId`, the `skipSourceDestCheck` cannot be specified because the source/destination check is always disabled for VNICs in a VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Example: `true` 
		* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create the VNIC. When launching an instance, use this `subnetId` instead of the deprecated `subnetId` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). At least one of them is required; if you provide both, the values must match.

			If you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN instead of a subnet, provide a `vlanId` instead of a `subnetId`. If you provide both `vlanId` and `subnetId`, the request fails. 
		* `vlan_id` - Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Provide a `vlanId` instead of a `subnetId`. If you provide both `vlanId` and `subnetId`, the request fails. 
	* `dedicated_vm_host_id` - The OCID of the dedicated VM host. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `fault_domain` - A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains lets you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

		If you do not specify the fault domain, the system selects one for you.

		 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

		Example: `FAULT-DOMAIN-1` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
	* `hostname_label` - Deprecated. Instead use `hostnameLabel` in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/). If you provide both, the values must match. 
	* `instance_options` - Optional mutable instance options
		* `are_legacy_imds_endpoints_disabled` - Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
	* `ipxe_script` - This is an advanced option.

		When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

		If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots. Be aware that the same iPXE script will run every time an instance boots, not only after the initial LaunchInstance call.

		By default, the iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI in the same way as the default iPXE script, use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

		If your instance boot volume type is paravirtualized, the boot volume is attached to the instance through virtio-scsi and no iPXE script is used. If your instance boot volume type is paravirtualized and you use custom iPXE to perform network-boot into your instance, the primary boot volume is attached as a data volume through the virtio-scsi drive.

		For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

		For more information about iPXE, see http://ipxe.org. 
	* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. By default, the value is false.
	* `preemptible_instance_config` - Configuration options for preemptible instances. 
		* `preemption_action` - The action to run when the preemptible instance is interrupted for eviction. 
			* `preserve_boot_volume` - Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. By default, it is false if not specified. 
			* `type` - The type of action to run when the instance is interrupted for eviction.
	* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

		You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
	* `shape_config` - The shape configuration requested for the instance.

		If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

		Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
		* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

			The following values are supported:
			* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
			* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
			* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - The total amount of memory in gigabytes that is available to the instance. 
		* `ocpus` - The total number of OCPUs available to the instance. 
	* `source_details` - 
		* `boot_volume_id` - The OCID of the boot volume used to boot the instance.
		* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB). 
		* `boot_volume_vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB that represents the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

			Allowed values:
			* `10`: Represents Balanced option.
			* `20`: Represents Higher Performance option.
			* `30`-`120`: Represents the Ultra High Performance option.

			For volumes with the auto-tuned performance feature enabled, this is set to the default (minimum) VPUs/GB. 
		* `image_id` - The OCID of the image used to boot the instance.
		* `kms_key_id` - The OCID of the key management key to assign as the master encryption key for the boot volume.
		* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
* `time_assessed` - The time when the assessment was done. An RFC3339 formatted datetime string.
* `time_created` - The time when the target asset was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the target asset was updated. An RFC3339 formatted datetime string.
* `type` - The type of target asset.
* `user_spec` - Instance launch details. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance. 
	* `agent_config` - Configuration options for the Oracle Cloud Agent software running on the instance.
		* `are_all_plugins_disabled` - Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

			To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
		* `is_management_disabled` - Whether Oracle Cloud Agent can run all the available management plugins. By default, the value is false (management plugins are enabled).

			These are the management plugins: OS Management Service Agent and Compute instance run command.

			The management plugins are controlled by this parameter and the per-plugin configuration in the `pluginsConfig` object.
			* If `isManagementDisabled` is true, all the management plugins are disabled, regardless of the per-plugin configuration.
			* If `isManagementDisabled` is false, all the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
		* `is_monitoring_disabled` - Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. By default, the value is false (monitoring plugins are enabled).

			These are the monitoring plugins: Compute instance monitoring and Custom logs monitoring.

			The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
			* If `isMonitoringDisabled` is true, all the monitoring plugins are disabled, regardless of the per-plugin configuration.
			* If `isMonitoringDisabled` is false, all the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
		* `plugins_config` - The configuration of plugins associated with this instance.
			* `desired_state` - Whether the plugin should be enabled or disabled.

				To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
			* `name` - The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
	* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
	* `capacity_reservation_id` - The OCID of the compute capacity reservation under which this instance is launched. You can opt out of all default reservations by specifying an empty string as input for this field. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default). 
	* `compartment_id` - The OCID of the compartment.
	* `create_vnic_details` - Contains properties for a VNIC. You use this object when creating the primary VNIC during instance launch or when creating a secondary VNIC. For more information about VNICs, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
		* `assign_private_dns_record` - Whether the VNIC should be assigned a DNS record. If set to false, there will be no DNS record registration for the VNIC. If set to true, the DNS record will be registered. By default, the value is true.

			If you specify a `hostnameLabel`, then `assignPrivateDnsRecord` must be set to true. 
		* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the [Subnet](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Subnet/)), then no public IP address is assigned. If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then a public IP address is assigned. If set to true and `prohibitPublicIpOnVnic` = true, an error is returned.

			**Note:** This public IP address is associated with the primary private IP on the VNIC. For more information, see [IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).

			**Note:** There's a limit to the number of [public IPs](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/) a VNIC or instance can have. If you try to create a secondary VNIC with an assigned public IP for an instance that has already reached its public IP limit, an error is returned. For information about the public IP limits, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).

			Example: `false`

			If you specify a `vlanId`, then `assignPublicIp` must be set to false. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
		* `hostname_label` - The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

			For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

			When launching an instance, use this `hostnameLabel` instead of the deprecated `hostnameLabel` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). If you provide both, the values must match.

			Example: `bminstance-1`

			If you specify a `vlanId`, the `hostnameLabel` cannot be specified. VNICs on a VLAN can not be assigned a hostname. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `nsg_ids` - List of OCIDs of the network security groups (NSGs) that are added to the VNIC. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).

			If a `vlanId` is specified, the `nsgIds` cannot be specified. The `vlanId` indicates that the VNIC will belong to a VLAN instead of a subnet. With VLANs, all VNICs in the VLAN belong to the NSGs that are associated with the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
		* `private_ip` - A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's *primary* private IP address. The value appears in the [Vnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/) object and also the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) object returned by [ListPrivateIps](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp).

			 If you specify a `vlanId`, the `privateIp` cannot be specified. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Example: `10.0.3.3` 
		* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you should skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).

			 If you specify a `vlanId`, the `skipSourceDestCheck` cannot be specified because the source/destination check is always disabled for VNICs in a VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Example: `true` 
		* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create the VNIC. When launching an instance, use this `subnetId` instead of the deprecated `subnetId` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/requests/LaunchInstanceDetails). At least one of them is required; if you provide both, the values must match.

			If you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN instead of a subnet, provide a `vlanId` instead of a `subnetId`. If you provide both `vlanId` and `subnetId`, the request fails. 
		* `vlan_id` - Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

			Provide a `vlanId` instead of a `subnetId`. If you provide both `vlanId` and `subnetId`, the request fails. 
	* `dedicated_vm_host_id` - The OCID of the dedicated VM host. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
	* `fault_domain` - A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains lets you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

		If you do not specify the fault domain, the system selects one for you.

		 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

		Example: `FAULT-DOMAIN-1` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
	* `hostname_label` - Deprecated. Instead use `hostnameLabel` in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/). If you provide both, the values must match. 
	* `instance_options` - Optional mutable instance options
		* `are_legacy_imds_endpoints_disabled` - Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
	* `ipxe_script` - This is an advanced option.

		When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

		If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots. Be aware that the same iPXE script will run every time an instance boots, not only after the initial LaunchInstance call.

		By default, the iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI in the same way as the default iPXE script, use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

		If your instance boot volume type is paravirtualized, the boot volume is attached to the instance through virtio-scsi and no iPXE script is used. If your instance boot volume type is paravirtualized and you use custom iPXE to perform network-boot into your instance, the primary boot volume is attached as a data volume through the virtio-scsi drive.

		For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

		For more information about iPXE, see http://ipxe.org. 
	* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. By default, the value is false.
	* `preemptible_instance_config` - Configuration options for preemptible instances. 
		* `preemption_action` - The action to run when the preemptible instance is interrupted for eviction. 
			* `preserve_boot_volume` - Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. By default, it is false if not specified. 
			* `type` - The type of action to run when the instance is interrupted for eviction.
	* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

		You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
	* `shape_config` - The shape configuration requested for the instance.

		If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

		Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
		* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

			The following values are supported:
			* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
			* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
			* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - The total amount of memory in gigabytes that is available to the instance. 
		* `ocpus` - The total number of OCPUs available to the instance. 
	* `source_details` - 
		* `boot_volume_id` - The OCID of the boot volume used to boot the instance.
		* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB). 
		* `boot_volume_vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB that represents the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

			Allowed values:
			* `10`: Represents Balanced option.
			* `20`: Represents Higher Performance option.
			* `30`-`120`: Represents the Ultra High Performance option.

			For volumes with the auto-tuned performance feature enabled, this is set to the default (minimum) VPUs/GB. 
		* `image_id` - The OCID of the image used to boot the instance.
		* `kms_key_id` - The OCID of the key management key to assign as the master encryption key for the boot volume.
		* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Target Asset
	* `update` - (Defaults to 20 minutes), when updating the Target Asset
	* `delete` - (Defaults to 20 minutes), when destroying the Target Asset


## Import

TargetAssets can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_migrations_target_asset.test_target_asset "id"
```

