---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_vm_cluster"
sidebar_current: "docs-oci-resource-database-cloud_vm_cluster"
description: |-
  Provides the Cloud Vm Cluster resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_vm_cluster
This resource provides the Cloud Vm Cluster resource in Oracle Cloud Infrastructure Database service.

Creates a cloud VM cluster.


## Example Usage

```hcl
resource "oci_database_cloud_vm_cluster" "test_cloud_vm_cluster" {
	#Required
	backup_subnet_id = oci_core_subnet.test_subnet.id
	cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
	compartment_id = var.compartment_id
	cpu_core_count = var.cloud_vm_cluster_cpu_core_count
	display_name = var.cloud_vm_cluster_display_name
	gi_version = var.cloud_vm_cluster_gi_version
	hostname = var.cloud_vm_cluster_hostname
	ssh_public_keys = var.cloud_vm_cluster_ssh_public_keys
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	backup_network_nsg_ids = var.cloud_vm_cluster_backup_network_nsg_ids
	cloud_automation_update_details {

		#Optional
		apply_update_time_preference {

			#Optional
			apply_update_preferred_end_time = var.cloud_vm_cluster_cloud_automation_update_details_apply_update_time_preference_apply_update_preferred_end_time
			apply_update_preferred_start_time = var.cloud_vm_cluster_cloud_automation_update_details_apply_update_time_preference_apply_update_preferred_start_time
		}
		freeze_period {

			#Optional
			freeze_period_end_time = var.cloud_vm_cluster_cloud_automation_update_details_freeze_period_freeze_period_end_time
			freeze_period_start_time = var.cloud_vm_cluster_cloud_automation_update_details_freeze_period_freeze_period_start_time
		}
		is_early_adoption_enabled = var.cloud_vm_cluster_cloud_automation_update_details_is_early_adoption_enabled
		is_freeze_period_enabled = var.cloud_vm_cluster_cloud_automation_update_details_is_freeze_period_enabled
	}
	cluster_name = var.cloud_vm_cluster_cluster_name
	data_collection_options {

		#Optional
		is_diagnostics_events_enabled = var.cloud_vm_cluster_data_collection_options_is_diagnostics_events_enabled
        is_health_monitoring_enabled = var.cloud_vm_cluster_data_collection_options_is_health_monitoring_enabled
        is_incident_logs_enabled = var.cloud_vm_cluster_data_collection_options_is_incident_logs_enabled
	}
	data_storage_percentage = var.cloud_vm_cluster_data_storage_percentage
	data_storage_size_in_tbs = var.cloud_vm_cluster_data_storage_size_in_tbs
	db_node_storage_size_in_gbs = var.cloud_vm_cluster_db_node_storage_size_in_gbs
	db_servers = var.cloud_vm_cluster_db_servers
	defined_tags = var.cloud_vm_cluster_defined_tags
	domain = var.cloud_vm_cluster_domain
	file_system_configuration_details {

		#Optional
		file_system_size_gb = var.cloud_vm_cluster_file_system_configuration_details_file_system_size_gb
		mount_point = var.cloud_vm_cluster_file_system_configuration_details_mount_point
	}
	freeform_tags = {"Department"= "Finance"}
	is_local_backup_enabled = var.cloud_vm_cluster_is_local_backup_enabled
	is_sparse_diskgroup_enabled = var.cloud_vm_cluster_is_sparse_diskgroup_enabled
	license_model = var.cloud_vm_cluster_license_model
	memory_size_in_gbs = var.cloud_vm_cluster_memory_size_in_gbs
	nsg_ids = var.cloud_vm_cluster_nsg_ids
	ocpu_count = var.cloud_vm_cluster_ocpu_count
	private_zone_id = oci_dns_zone.test_zone.id
	scan_listener_port_tcp = var.cloud_vm_cluster_scan_listener_port_tcp
	scan_listener_port_tcp_ssl = var.cloud_vm_cluster_scan_listener_port_tcp_ssl
	security_attributes = var.cloud_vm_cluster_security_attributes
	subscription_id = var.tenant_subscription_id
	system_version = var.cloud_vm_cluster_system_version
	time_zone = var.cloud_vm_cluster_time_zone
	vm_cluster_type = var.cloud_vm_cluster_vm_cluster_type
}
```

## Argument Reference

The following arguments are supported:

* `backup_network_nsg_ids` - (Optional) (Updatable) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
* `backup_subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the cloud VM cluster. 
* `cloud_automation_update_details` - (Optional) (Updatable) Specifies the properties necessary for cloud automation updates. This includes modifying the apply update time preference, enabling or disabling early adoption, and enabling, modifying, or disabling the update freeze period. 
	* `apply_update_time_preference` - (Optional) (Updatable) Configure the time slot for applying VM cloud automation software updates to the cluster. When nothing is selected, the default time slot is 12 AM to 2 AM UTC. Any 2-hour slot is available starting at 12 AM. 
		* `apply_update_preferred_end_time` - (Optional) (Updatable) End time for polling VM cloud automation software updates for the cluster. If the endTime is not specified, 2 AM UTC is used by default. 
		* `apply_update_preferred_start_time` - (Optional) (Updatable) Start time for polling VM cloud automation software updates for the cluster. If the startTime is not specified, 12 AM UTC is used by default. 
	* `freeze_period` - (Optional) (Updatable) Enables a freeze period for the VM cluster prohibiting the VMs from getting cloud automation software updates during critical business cycles. Freeze period start date. Starts at 12:00 AM UTC on the selected date and ends at 11:59:59 PM UTC on the selected date. Validates to ensure the freeze period does not exceed 45 days. 
		* `freeze_period_end_time` - (Optional) (Updatable) End time of the freeze period cycle. 
		* `freeze_period_start_time` - (Optional) (Updatable) Start time of the freeze period cycle. 
	* `is_early_adoption_enabled` - (Optional) (Updatable) Annotates whether the cluster should be part of early access to apply VM cloud automation software updates. Those clusters annotated as early access will download the software bits for cloud automation in the first week after the update is available, while other clusters will have to wait until the following week. 
	* `is_freeze_period_enabled` - (Optional) (Updatable) Specifies if the freeze period is enabled for the VM cluster to prevent the VMs from receiving cloud automation software updates during critical business cycles. Freeze period starts at 12:00 AM UTC and ends at 11:59:59 PM UTC on the selected date. Ensure that the freezing period does not exceed 45 days. 
* `cloud_exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure resource.
* `cluster_name` - (Optional) The cluster name for cloud VM cluster. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - (Required) (Updatable) The number of CPU cores to enable for a cloud VM cluster. Valid values depend on the specified shape:
	* Exadata.Base.48 - Specify a multiple of 2, from 0 to 48.
	* Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84.
	* Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168.
	* Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.
	* Exadata.Quarter2.92 - Specify a multiple of 2, from 0 to 92.
	* Exadata.Half2.184 - Specify a multiple of 4, from 0 to 184.
	* Exadata.Full2.368 - Specify a multiple of 8, from 0 to 368. 
* `data_collection_options` - (Optional) (Updatable) Indicates user preferences for the various diagnostic collection options for the VM cluster/Cloud VM cluster/VMBM DBCS. 
	* `is_diagnostics_events_enabled` - (Optional) (Updatable) Indicates whether diagnostic collection is enabled for the VM cluster/Cloud VM cluster/VMBM DBCS. Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues. Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system. You can enable diagnostic collection during VM cluster/Cloud VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` or `updateCloudVmCluster` API.
	* `is_health_monitoring_enabled` - (Optional) (Updatable) Indicates whether health monitoring is enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel. You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster`, `UpdateCloudVmCluster` or `updateDbsystem` API. 
    * `is_incident_logs_enabled` - (Optional) (Updatable) Indicates whether incident logs and trace collection are enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster`, `updateCloudVmCluster` or `updateDbsystem` API.  
* `data_storage_percentage` - (Optional) The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 35, 40, 60 and 80. The default is 80 percent assigned to DATA storage. See [Storage Configuration](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaoverview.htm#Exadata) in the Exadata documentation for details on the impact of the configuration settings on storage. 
* `data_storage_size_in_tbs` - (Optional) (Updatable) The data disk group size to be allocated in TBs.
* `db_node_storage_size_in_gbs` - (Optional) (Updatable) The local node storage to be allocated in GBs.
* `db_servers` - (Optional) The list of DB servers.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The user-friendly name for the cloud VM cluster. The name does not need to be unique.
* `domain` - (Optional) A domain name used for the cloud VM cluster. If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted. Applies to Exadata Cloud Service instances only. 
* `file_system_configuration_details` - (Optional) (Updatable) Details of the file system configuration of the VM cluster.
	* `file_system_size_gb` - (Optional) (Updatable) The file system size to be allocated in GBs.
	* `mount_point` - (Optional) (Updatable) The mount point of file system.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gi_version` - (Required) A valid Oracle Grid Infrastructure (GI) software version.
* `hostname` - (Required) The hostname for the cloud VM cluster. The hostname must begin with an alphabetic character, and can contain alphanumeric characters and hyphens (-). The maximum length of the hostname is 16 characters for bare metal and virtual machine DB systems, and 12 characters for Exadata systems.

	The maximum length of the combined hostname and domain is 63 characters.

	**Note:** The hostname must be unique within the subnet. If it is not unique, the cloud VM Cluster will fail to provision. 
* `is_local_backup_enabled` - (Optional) If true, database backup on local Exadata storage is configured for the cloud VM cluster. If false, database backup on local Exadata storage is not available in the cloud VM cluster. 
* `is_sparse_diskgroup_enabled` - (Optional) If true, the sparse disk group is configured for the cloud VM cluster. If false, the sparse disk group is not created. 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the cloud VM cluster. The default is BRING_YOUR_OWN_LICENSE. 
* `memory_size_in_gbs` - (Optional) (Updatable) The memory to be allocated in GBs.
* `nsg_ids` - (Optional) (Updatable) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty. 
* `ocpu_count` - (Optional) (Updatable) The number of OCPU cores to enable for a cloud VM cluster. Only 1 decimal place is allowed for the fractional part.
* `private_zone_id` - (Optional) The private zone id in which DNS records need to be created.
* `scan_listener_port_tcp` - (Optional) The TCP Single Client Access Name (SCAN) port. The default port is 1521.
* `scan_listener_port_tcp_ssl` - (Optional) The TCPS Single Client Access Name (SCAN) port. The default port is 2484.
* `security_attributes` - (Optional) (Updatable) Security Attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}` 
* `ssh_public_keys` - (Required) (Updatable) The public key portion of one or more key pairs used for SSH access to the cloud VM cluster.
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the cloud VM cluster. 
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `system_version` - (Optional) Operating system version of the image.
* `time_zone` - (Optional) The time zone to use for the cloud VM cluster. For details, see [Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm). 
* `vm_cluster_type` - (Optional) The vmcluster type for the VM cluster/Cloud VM cluster.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the cloud Exadata infrastructure resource is located in.
* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the cloud VM cluster.

	**Subnet Restriction:** See the subnet restrictions information for **subnetId**. 
* `cloud_automation_update_details` - Specifies the properties necessary for cloud automation updates. This includes modifying the apply update time preference, enabling or disabling early adoption, and enabling, modifying, or disabling the update freeze period. 
	* `apply_update_time_preference` - Configure the time slot for applying VM cloud automation software updates to the cluster. When nothing is selected, the default time slot is 12 AM to 2 AM UTC. Any 2-hour slot is available starting at 12 AM. 
		* `apply_update_preferred_end_time` - End time for polling VM cloud automation software updates for the cluster. If the endTime is not specified, 2 AM UTC is used by default. 
		* `apply_update_preferred_start_time` - Start time for polling VM cloud automation software updates for the cluster. If the startTime is not specified, 12 AM UTC is used by default. 
	* `freeze_period` - Enables a freeze period for the VM cluster prohibiting the VMs from getting cloud automation software updates during critical business cycles. Freeze period start date. Starts at 12:00 AM UTC on the selected date and ends at 11:59:59 PM UTC on the selected date. Validates to ensure the freeze period does not exceed 45 days. 
		* `freeze_period_end_time` - End time of the freeze period cycle. 
		* `freeze_period_start_time` - Start time of the freeze period cycle. 
	* `is_early_adoption_enabled` - Annotates whether the cluster should be part of early access to apply VM cloud automation software updates. Those clusters annotated as early access will download the software bits for cloud automation in the first week after the update is available, while other clusters will have to wait until the following week. 
	* `is_freeze_period_enabled` - Specifies if the freeze period is enabled for the VM cluster to prevent the VMs from receiving cloud automation software updates during critical business cycles. Freeze period starts at 12:00 AM UTC and ends at 11:59:59 PM UTC on the selected date. Ensure that the freezing period does not exceed 45 days. 
* `cloud_exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
* `cluster_name` - The cluster name for cloud VM cluster. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_model` - The compute model of the cloud VM cluster.
* `cpu_core_count` - The number of CPU cores enabled on the cloud VM cluster.
* `data_collection_options` - Indicates user preferences for the various diagnostic collection options for the VM cluster/Cloud VM cluster/VMBM DBCS. 
	* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster/Cloud VM cluster/VMBM DBCS. Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues. Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system. You can enable diagnostic collection during VM cluster/Cloud VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` or `updateCloudVmCluster` API. 
	* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel. You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster`, `UpdateCloudVmCluster` or `updateDbsystem` API. 
	* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster`, `updateCloudVmCluster` or `updateDbsystem` API. 
* `data_storage_percentage` - The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 35, 40, 60 and 80. The default is 80 percent assigned to DATA storage. See [Storage Configuration](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaoverview.htm#Exadata) in the Exadata documentation for details on the impact of the configuration settings on storage. 
* `data_storage_size_in_tbs` - The data disk group size to be allocated in TBs.
* `db_node_storage_size_in_gbs` - The local node storage to be allocated in GBs.
* `db_servers` - The list of DB servers.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `disk_redundancy` - The type of redundancy configured for the cloud Vm cluster. NORMAL is 2-way redundancy. HIGH is 3-way redundancy. 
* `display_name` - The user-friendly name for the cloud VM cluster. The name does not need to be unique.
* `domain` - The domain name for the cloud VM cluster.
* `file_system_configuration_details` - Details of the file system configuration of the VM cluster.
	* `file_system_size_gb` - The file system size to be allocated in GBs.
	* `mount_point` - The mount point of file system.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gi_version` - A valid Oracle Grid Infrastructure (GI) software version.
* `hostname` - The hostname for the cloud VM cluster.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud VM cluster.
* `iorm_config_cache` - The IORM settings of the Exadata DB system. 
	* `db_plans` - An array of IORM settings for all the database in the Exadata DB system. 
		* `db_name` - The database name. For the default `DbPlan`, the `dbName` is `default`. 
		* `flash_cache_limit` - The flash cache limit for this database. This value is internally configured based on the share value assigned to the database. 
		* `share` - The relative priority of this database. 
	* `lifecycle_details` - Additional information about the current `lifecycleState`. 
	* `objective` - The current value for the IORM objective. The default is `AUTO`. 
	* `state` - The current state of IORM configuration for the Exadata DB system. 
* `is_local_backup_enabled` - If true, database backup on local Exadata storage is configured for the cloud VM cluster. If false, database backup on local Exadata storage is not available in the cloud VM cluster. 
* `is_sparse_diskgroup_enabled` - If true, sparse disk group is configured for the cloud VM cluster. If false, sparse disk group is not created. 
* `last_update_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance update history entry. This value is updated when a maintenance update starts.
* `license_model` - The Oracle license model that applies to the cloud VM cluster. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `listener_port` - The port number configured for the listener on the cloud VM cluster.
* `memory_size_in_gbs` - The memory to be allocated in GBs.
* `node_count` - The number of nodes in the cloud VM cluster. 
* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
* `ocpu_count` - The number of OCPU cores to enable on the cloud VM cluster. Only 1 decimal place is allowed for the fractional part.
* `scan_dns_name` - The FQDN of the DNS record for the SCAN IP addresses that are associated with the cloud VM cluster. 
* `scan_dns_record_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the cloud VM cluster. 
* `scan_ip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IPv4 addresses associated with the cloud VM cluster. SCAN IPv4 addresses are typically used for load balancing and are not assigned to any interface. Oracle Clusterware directs the requests to the appropriate nodes in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `scan_ipv6ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IPv6 addresses associated with the cloud VM cluster. SCAN IPv6 addresses are typically used for load balancing and are not assigned to any interface. Oracle Clusterware directs the requests to the appropriate nodes in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `scan_listener_port_tcp` - The TCP Single Client Access Name (SCAN) port. The default port is 1521.
* `scan_listener_port_tcp_ssl` - The TCPS Single Client Access Name (SCAN) port. The default port is 2484.
* `security_attributes` - Security Attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}` 
* `shape` - The model name of the Exadata hardware running the cloud VM cluster. 
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the cloud VM cluster.
* `state` - The current state of the cloud VM cluster.
* `storage_size_in_gbs` - The storage allocation for the disk group, in gigabytes (GB).
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the cloud VM cluster.

	**Subnet Restrictions:**
	* For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `system_version` - Operating system version of the image.
* `time_created` - The date and time that the cloud VM cluster was created.
* `time_zone` - The time zone of the cloud VM cluster. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) IPv4 addresses associated with the cloud VM cluster. The Cluster Ready Services (CRS) creates and maintains one VIP IPv4 address for each node in the Exadata Cloud Service instance to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `vipv6ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) IPv6 addresses associated with the cloud VM cluster. The Cluster Ready Services (CRS) creates and maintains one VIP IPv6 address for each node in the Exadata Cloud Service instance to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `vm_cluster_type` - The vmcluster type for the VM cluster/Cloud VM cluster.
* `zone_id` - The OCID of the zone the cloud VM cluster is associated with. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 12 hours), when creating the Cloud Vm Cluster
	* `update` - (Defaults to 12 hours), when updating the Cloud Vm Cluster
	* `delete` - (Defaults to 12 hours), when destroying the Cloud Vm Cluster


## Import

CloudVmClusters can be imported using the `id`, e.g.

```
$ terraform import oci_database_cloud_vm_cluster.test_cloud_vm_cluster "id"
```

