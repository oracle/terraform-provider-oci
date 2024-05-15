---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadata_infrastructures"
sidebar_current: "docs-oci-datasource-database-exadata_infrastructures"
description: |-
  Provides the list of Exadata Infrastructures in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadata_infrastructures
This data source provides the list of Exadata Infrastructures in Oracle Cloud Infrastructure Database service.

Lists the Exadata infrastructure resources in the specified compartment. Applies to Exadata Cloud@Customer instances only.
To list the Exadata Cloud Service infrastructure resources in a compartment, use the  [ListCloudExadataInfrastructures](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudExadataInfrastructure/ListCloudExadataInfrastructures) operation.


## Example Usage

```hcl
data "oci_database_exadata_infrastructures" "test_exadata_infrastructures" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.exadata_infrastructure_display_name
	excluded_fields = var.exadata_infrastructure_excluded_fields
	state = var.exadata_infrastructure_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `excluded_fields` - (Optional) If provided, the specified fields will be excluded in the response.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `exadata_infrastructures` - The list of exadata_infrastructures.

### ExadataInfrastructure Reference

The following attributes are exported:

* `activated_storage_count` - The requested number of additional storage servers activated for the Exadata infrastructure.
* `additional_compute_count` - The requested number of additional compute servers for the Exadata infrastructure.
* `additional_compute_system_model` - Oracle Exadata System Model specification. The system model determines the amount of compute or storage server resources available for use. For more information, please see [System and Shape Configuration Options] (https://docs.oracle.com/en/engineered-systems/exadata-cloud-at-customer/ecccm/ecc-system-config-options.html#GUID-9E090174-5C57-4EB1-9243-B470F9F10D6B) 
* `additional_storage_count` - The requested number of additional storage servers for the Exadata infrastructure.
* `admin_network_cidr` - The CIDR block for the Exadata administration network.
* `availability_domain` - The name of the availability domain that the Exadata infrastructure is located in.
* `cloud_control_plane_server1` - The IP address for the first control plane server.
* `cloud_control_plane_server2` - The IP address for the second control plane server.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_count` - The number of compute servers for the Exadata infrastructure.
* `contacts` - The list of contacts for the Exadata infrastructure.
	* `email` - The email for the Exadata Infrastructure contact.
	* `is_contact_mos_validated` - If `true`, this Exadata Infrastructure contact is a valid My Oracle Support (MOS) contact. If `false`, this Exadata Infrastructure contact is not a valid MOS contact.
	* `is_primary` - If `true`, this Exadata Infrastructure contact is a primary contact. If `false`, this Exadata Infrastructure is a secondary contact.
	* `name` - The name of the Exadata Infrastructure contact.
	* `phone_number` - The phone number for the Exadata Infrastructure contact.
* `corporate_proxy` - The corporate network proxy for access to the control plane network.
* `cpus_enabled` - The number of enabled CPU cores.
* `csi_number` - The CSI Number of the Exadata infrastructure.
* `data_storage_size_in_tbs` - Size, in terabytes, of the DATA disk group. 
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `db_server_version` - The software version of the database servers (dom0) in the Exadata infrastructure.
* `defined_file_system_configurations` - Details of the file system configuration of the Exadata infrastructure.
	* `is_backup_partition` - If true, the file system is used to create a backup prior to Exadata VM OS update.
	* `is_resizable` - If true, the file system resize is allowed for the Exadata Infrastructure cluster. If false, the file system resize is not allowed.
	* `min_size_gb` - The minimum size of file system.
	* `mount_point` - The mount point of file system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Exadata Cloud@Customer infrastructure. The name does not need to be unique.
* `dns_server` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway` - The gateway for the control plane network.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `infini_band_network_cidr` - The CIDR block for the Exadata InfiniBand interconnect.
* `is_cps_offline_report_enabled` - Indicates whether cps offline diagnostic report is enabled for this Exadata infrastructure. This will allow a customer to quickly check status themselves and fix problems on their end, saving time and frustration for both Oracle and the customer when they find the CPS in a disconnected state.You can enable offline diagnostic report during Exadata infrastructure provisioning. You can also disable or enable it at any time using the UpdateExadatainfrastructure API. 
* `is_multi_rack_deployment` - Indicates if deployment is Multi-Rack or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_slo_status` - A field to capture ‘Maintenance SLO Status’ for the Exadata infrastructure with values ‘OK’, ‘DEGRADED’. Default is ‘OK’ when the infrastructure is provisioned.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - If true, enables the monthly patching option.
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `max_cpu_count` - The total number of CPU cores available.
* `max_data_storage_in_tbs` - The total available DATA disk group size.
* `max_db_node_storage_in_gbs` - The total local node storage available in GBs.
* `max_memory_in_gbs` - The total memory available in GBs.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `monthly_db_server_version` - The monthly software version of the database servers (dom0) in the Exadata infrastructure.
* `multi_rack_configuration_file` - The base64 encoded Multi-Rack configuration json file.
* `netmask` - The netmask for the control plane network.
* `network_bonding_mode_details` - Details of bonding mode for Client and Backup and DR networks of an Exadata infrastructure.
	* `backup_network_bonding_mode` - The network bonding mode for the Exadata infrastructure.
	* `client_network_bonding_mode` - The network bonding mode for the Exadata infrastructure.
	* `dr_network_bonding_mode` - The network bonding mode for the Exadata infrastructure.
* `ntp_server` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `rack_serial_number` - The serial number for the Exadata infrastructure.
* `shape` - The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance. 
* `state` - The current lifecycle state of the Exadata infrastructure.
* `storage_count` - The number of Exadata storage servers for the Exadata infrastructure.
* `storage_server_version` - The software version of the storage servers (cells) in the Exadata infrastructure.
* `time_created` - The date and time the Exadata infrastructure was created.
* `time_zone` - The time zone of the Exadata infrastructure. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).

