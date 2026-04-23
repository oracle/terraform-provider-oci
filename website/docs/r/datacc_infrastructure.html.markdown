---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_infrastructure"
sidebar_current: "docs-oci-resource-datacc-infrastructure"
description: |-
  Provides the Infrastructure resource in Oracle Cloud Infrastructure Datacc service
---

# oci_datacc_infrastructure
This resource provides the Infrastructure resource in Oracle Cloud Infrastructure Datacc service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datacc

Create a Database Infrastructure using the specified details.


## Example Usage

```hcl
resource "oci_datacc_infrastructure" "test_infrastructure" {
	#Required
	cloud_control_plane_server1 = var.infrastructure_cloud_control_plane_server1
	cloud_control_plane_server2 = var.infrastructure_cloud_control_plane_server2
	compartment_id = var.compartment_id
	display_name = var.infrastructure_display_name
	dns_servers = var.infrastructure_dns_servers
	gateway = var.infrastructure_gateway
	netmask = var.infrastructure_netmask
	ntp_servers = var.infrastructure_ntp_servers
	shape = var.infrastructure_shape
	system_model = var.infrastructure_system_model

	#Optional
	acfs_file_system_storage_in_gbs = var.infrastructure_acfs_file_system_storage_in_gbs
	admin_networkcidr = var.infrastructure_admin_networkcidr
	backup_network_bonding_interface = var.infrastructure_backup_network_bonding_interface
	backup_network_bonding_mode = var.infrastructure_backup_network_bonding_mode
	client_network_bonding_interface = var.infrastructure_client_network_bonding_interface
	client_network_bonding_mode = var.infrastructure_client_network_bonding_mode
	contacts {
		#Required
		email = var.infrastructure_contacts_email
		is_primary = var.infrastructure_contacts_is_primary
		name = var.infrastructure_contacts_name

		#Optional
		is_contact_mos_validated = var.infrastructure_contacts_is_contact_mos_validated
		phone_number = var.infrastructure_contacts_phone_number
	}
	corporate_proxy = var.infrastructure_corporate_proxy
	cps_network_bonding_interface = var.infrastructure_cps_network_bonding_interface
	cps_network_bonding_mode = var.infrastructure_cps_network_bonding_mode
	data_disk_percentage = var.infrastructure_data_disk_percentage
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.infrastructure_description
	freeform_tags = {"bar-key"= "value"}
	maintenance_window {

		#Optional
		custom_action_timeout_in_mins = var.infrastructure_maintenance_window_custom_action_timeout_in_mins
		days_of_week = var.infrastructure_maintenance_window_days_of_week
		hours_of_day = var.infrastructure_maintenance_window_hours_of_day
		is_custom_action_timeout_enabled = var.infrastructure_maintenance_window_is_custom_action_timeout_enabled
		is_monthly_patching_enabled = var.infrastructure_maintenance_window_is_monthly_patching_enabled
		lead_time_in_weeks = var.infrastructure_maintenance_window_lead_time_in_weeks
		months = var.infrastructure_maintenance_window_months
		patching_mode = var.infrastructure_maintenance_window_patching_mode
		preference = var.infrastructure_maintenance_window_preference
		weeks_of_month = var.infrastructure_maintenance_window_weeks_of_month
	}
	vlan_id = oci_core_vlan.test_vlan.id
}
```

## Argument Reference

The following arguments are supported:

* `acfs_file_system_storage_in_gbs` - (Optional) (Updatable) The amount of storage (in GB) in the DATA disk group that is reserved for creating local storage for VM Clusters and application VMs. 
* `admin_networkcidr` - (Optional) (Updatable) The CIDR block for the system network. The system network is a private network in Database Infrastructure and is not connected to your corporate network. The system network is used for storage (ASM) traffic, high-performance interconnect traffic and administration of infrastructure components.
* `backup_network_bonding_interface` - (Optional) (Updatable) The network bonding interface for backup network for the Database Infrastructure.
* `backup_network_bonding_mode` - (Optional) (Updatable) The network bonding mode for Backup networks for the Database Infrastructure.
* `client_network_bonding_interface` - (Optional) (Updatable) The network bonding interface for client network for the Database Infrastructure.
* `client_network_bonding_mode` - (Optional) (Updatable) The network bonding mode for Client networks for the Database Infrastructure.
* `cloud_control_plane_server1` - (Required) (Updatable) The IP address for the first control plane server.
* `cloud_control_plane_server2` - (Required) (Updatable) The IP address for the second control plane server.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `contacts` - (Optional) (Updatable) The list of contacts for the Database Infrastructure.
	* `email` - (Required) (Updatable) The email for the Database Infrastructure contact.
	* `is_contact_mos_validated` - (Optional) (Updatable) If `true`, this Database Infrastructure contact is a valid My Oracle Support (MOS) contact.  If `false`, this Database Infrastructure contact is not a valid MOS contact. 
	* `is_primary` - (Required) (Updatable) If `true`, this Database Infrastructure contact is a primary contact.  If `false`, this Database Infrastructure is a secondary contact. 
	* `name` - (Required) (Updatable) The name of the Database Infrastructure contact.
	* `phone_number` - (Optional) (Updatable) The phone number for the Database Infrastructure contact.
* `corporate_proxy` - (Optional) (Updatable) The corporate network proxy for access to the control plane network. Oracle recommends using an HTTPS proxy when possible for enhanced security. 
* `cps_network_bonding_interface` - (Optional) (Updatable) The network bonding interface for CPS network for the Database Infrastructure.
* `cps_network_bonding_mode` - (Optional) (Updatable) The network bonding mode for CPS networks for the Database Infrastructure.
* `data_disk_percentage` - (Optional) (Updatable) Percentage of disk space assigned for DATA disk group. Remaining disk space will get assiged to RECO disk group
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Database Infrastructure description.
* `display_name` - (Required) (Updatable) The user-friendly name for the Database Infrastructure. The name does not need to be unique. 
* `dns_servers` - (Required) (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gateway` - (Required) (Updatable) The gateway for the control plane network.
* `maintenance_window` - (Optional) (Updatable) The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - (Optional) (Updatable) Determines the amount of time the system will wait before the start of each Database Infrastructure server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are - 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - (Optional) (Updatable) If true, enables the configuration of a custom action timeout (waiting period) between Database Infrastructure server patching operations.
	* `is_monthly_patching_enabled` - (Optional) (Updatable) If true, enables the monthly patching option.
	* `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time.  The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
	* `patching_mode` - (Optional) (Updatable) Cloud Database Infrastructure node patching method. *IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Database Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle)  for more information. 
	* `preference` - (Optional) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and  22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates,  not days of the week.For example, to allow maintenance during the 2nd week of the month  (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction  with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and  hours that maintenance will be performed. 
* `netmask` - (Required) (Updatable) The netmask for the control plane network.
* `ntp_servers` - (Required) (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
* `shape` - (Required) (Updatable) The shape of the Database Infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance. 
* `system_model` - (Required) (Updatable) Database Infrastructure System Model specification. The system model determines the model of the Database Infrastructure hardware to be used. 
* `vlan_id` - (Optional) (Updatable) The CPS network VLAN ID.
* `scale_storage_trigger` - (Optional) (Updatable) An optional property when incremented triggers Scale Storage. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `acfs_file_system_storage_in_gbs` - The amount of storage (in GB) in the DATA disk group that is reserved for creating local storage for VM Clusters and application VMs. 
* `acfs_file_system_used_storage_in_gbs` - The amount of storage (in GB) in the DATA disk group that is currently utilized for creating local storage for VM Clusters and application VMs. This attribute is deprecated and will be removed in a subsequent release. Please read from systemStorageCapacity instead. 
* `admin_networkcidr` - The CIDR block for the system network. The system network is a private network in Database Infrastructure and is not connected to your corporate network. The system network is used for storage (ASM) traffic, high-performance interconnect traffic and administration of infrastructure components.
* `backup_network_bonding_interface` - The network bonding interface for backup network for the Database Infrastructure.
* `backup_network_bonding_mode` - The network bonding mode for Backup networks for the Database Infrastructure.
* `client_network_bonding_interface` - The network bonding interface for client network for the Database Infrastructure.
* `client_network_bonding_mode` - The network bonding mode for Client networks for the Database Infrastructure.
* `cloud_control_plane_server1` - The IP address for the first control plane server.
* `cloud_control_plane_server2` - The IP address for the second control plane server.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `compute_capacity` - Capacity details of the Database Infrastructure.
	* `allocated_cores` - Total CPU cores count allocated..
	* `available_cores` - Total available CPU cores count.
	* `available_memory_in_gbs` - Available memory, in gigabytes (GB).
	* `reserved_cores` - Total Reserved CPU cores count.
	* `reserved_memory_in_gbs` - Reserved memory, in gigabytes (GB).
	* `total_cores` - Total CPU cores count.
	* `total_memory_in_gbs` - Total memory allocated, in gigabytes (GB).
	* `used_memory_in_gbs` - Memory allocated to Oracle database virtual machine cluster or Instance, in gigabytes (GB).
* `contacts` - The list of contacts for the Database Infrastructure.
	* `email` - The email for the Database Infrastructure contact.
	* `is_contact_mos_validated` - If `true`, this Database Infrastructure contact is a valid My Oracle Support (MOS) contact.  If `false`, this Database Infrastructure contact is not a valid MOS contact. 
	* `is_primary` - If `true`, this Database Infrastructure contact is a primary contact.  If `false`, this Database Infrastructure is a secondary contact. 
	* `name` - The name of the Database Infrastructure contact.
	* `phone_number` - The phone number for the Database Infrastructure contact.
* `corporate_proxy` - The corporate network proxy for access to the control plane network. Oracle recommends using an HTTPS proxy when possible for enhanced security. 
* `cps_network_bonding_interface` - The network bonding interface for CPS network for the Database Infrastructure.
* `cps_network_bonding_mode` - The network bonding mode for CPS networks for the Database Infrastructure.
* `data_disk_percentage` - Percentage of disk space assigned for DATA disk group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Database Infrastructure description.
* `display_name` - The user-friendly name for the Database Infrastructure. The name does not need to be unique. 
* `dns_servers` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gateway` - The gateway for the control plane network.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure. 
* `lifecycle_state_details` - Lifecycle state details of the Database Infrastructure.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each Database Infrastructure server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - Days during the week when maintenance should be performed.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are - 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - If true, enables the configuration of a custom action timeout (waiting period) between Database Infrastructure server patching operations.
	* `is_monthly_patching_enabled` - If true, enables the monthly patching option.
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time.  The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
	* `patching_mode` - Cloud Database Infrastructure node patching method. *IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Database Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle)  for more information. 
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and  22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates,  not days of the week.For example, to allow maintenance during the 2nd week of the month  (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction  with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and  hours that maintenance will be performed. 
* `netmask` - The netmask for the control plane network.
* `network_adapter_configuration` - The network adapter, transceiver and cable configuration for the client and backup networks. 
* `ntp_servers` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `rack_serial_number` - The serial number for the Database Infrastructure.
* `reco_disk_percentage` - Percentage of disk space assigned for RECO disk group.
* `servers` - A list of Database Infrastructure nodes.
	* `base_vm_count` - Number of database virtual machines hosted on the server.
	* `compute_capacity` - Capacity details of the Database Infrastructure.
		* `allocated_cores` - Total CPU cores count allocated..
		* `available_cores` - Total available CPU cores count.
		* `available_memory_in_gbs` - Available memory, in gigabytes (GB).
		* `reserved_cores` - Total Reserved CPU cores count.
		* `reserved_memory_in_gbs` - Reserved memory, in gigabytes (GB).
		* `total_cores` - Total CPU cores count.
		* `total_memory_in_gbs` - Total memory allocated, in gigabytes (GB).
		* `used_memory_in_gbs` - Memory allocated to Oracle database virtual machine cluster or Instance, in gigabytes (GB).
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Data Server of Infrastructure.
	* `ilom_ip_address` - Database Infrastructure Server ILOM IP address.
	* `ilom_name` - Database Infrastructure Server ILOM name.
	* `instance_vm_count` - Number of instances hosted on the server.
	* `server_ip_address` - Database Infrastructure Server IP address.
	* `server_name` - Database Infrastructure Server name.
	* `state` - The current state of the Database Infrastructure server.
* `shape` - The shape of the Database Infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance. 
* `ssd_configuration_requested` - SSD configuration requested for the infrastructure.
* `state` - The current state of the Database Infrastructure.
* `storage_capacity` - Capacity details of the Storage disk group.  This attribute is deprecated and will be removed in a subsequent release. Please use systemStorageCapacity instead. 
	* `disk_group` - Disk group name.
	* `logical_free_space_in_gbs` - The total amount of logical disk space that is currently available for use in a disk group, in gigabytes (GB).
	* `phy_free_space_in_gbs` - The total amount of physical disk space that is currently available for use in a disk group, in gigabytes (GB).
	* `phy_reserved_space_in_gbs` - The total amount of physical disk space that is reserved for system use in a disk group, in gigabytes (GB).
	* `phy_total_space_in_gbs` - The total amount of physical disk space available in a disk group, in gigabytes (GB).
	* `storage_disk_redundancy` - The Disk redundancy for Database Infrastructure storage.
* `subscription_plan_number` - The unique identifier for the subscription plan number.
* `system_model` - Database Infrastructure System Model specification. The system model determines the model of the Database Infrastructure hardware to be used. 
* `system_storage_capacity` - Capacity details of different storage types.
	* `acfs` - The amount of storage (in GB) in the DATA disk group that is currently utilized for creating local storage for VM Clusters and application VMs.
		* `free_space_in_gbs` - The total amount of logical disk space that is currently available for use, in gigabytes (GB).
		* `total_space_in_gbs` - The total amount of logical disk space available, in gigabytes (GB).
	* `disk_groups` - List of storage disk group capacity details.
		* `free_space_in_gbs` - The total amount of logical disk space that is currently available for use, in gigabytes (GB).
		* `reserved_space_in_gbs` - The total amount of logical disk space that is reserved for system use, in gigabytes (GB).
		* `storage_type` - The storage type for the Cloud Database Infrastructure.
		* `total_space_in_gbs` - The total amount of logical disk space available, in gigabytes (GB).
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_activated` - The time, in RFC3339 format, when the Database Infrastructure was activated. 
* `time_created` - The time that the Database Infrastructure cluster was created. An RFC3339 formatted datetime string. 
* `time_last_state_updated` - The time, in RFC3339 format, when the lifecycle state was last updated. 
* `time_updated` - The time that the Database Infrastructure was last updated. An RFC3339 formatted datetime string. 
* `time_validated` - The time, in RFC3339 format, when the Database Infrastructure network was validated. 
* `version` - The version of the system software on the Database Infrastructure. 
* `vlan_id` - The CPS network VLAN ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Infrastructure
	* `update` - (Defaults to 20 minutes), when updating the Infrastructure
	* `delete` - (Defaults to 20 minutes), when destroying the Infrastructure


## Import

Infrastructures can be imported using the `id`, e.g.

```
$ terraform import oci_datacc_infrastructure.test_infrastructure "id"
```

