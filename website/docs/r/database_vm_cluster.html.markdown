---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster"
sidebar_current: "docs-oci-resource-database-vm_cluster"
description: |-
  Provides the Vm Cluster resource in Oracle Cloud Infrastructure Database service
---

# oci_database_vm_cluster
This resource provides the Vm Cluster resource in Oracle Cloud Infrastructure Database service.

Creates an Exadata Cloud@Customer VM cluster.


## Example Usage

```hcl
resource "oci_database_vm_cluster" "test_vm_cluster" {
	#Required
	compartment_id = var.compartment_id
	cpu_core_count = var.vm_cluster_cpu_core_count
	display_name = var.vm_cluster_display_name
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	gi_version = var.vm_cluster_gi_version
	ssh_public_keys = var.vm_cluster_ssh_public_keys
	vm_cluster_network_id = oci_database_vm_cluster_network.test_vm_cluster_network.id

	#Optional
	cloud_automation_update_details {

		#Optional
		apply_update_time_preference {

			#Optional
			apply_update_preferred_end_time = var.vm_cluster_cloud_automation_update_details_apply_update_time_preference_apply_update_preferred_end_time
			apply_update_preferred_start_time = var.vm_cluster_cloud_automation_update_details_apply_update_time_preference_apply_update_preferred_start_time
		}
		freeze_period {

			#Optional
			freeze_period_end_time = var.vm_cluster_cloud_automation_update_details_freeze_period_freeze_period_end_time
			freeze_period_start_time = var.vm_cluster_cloud_automation_update_details_freeze_period_freeze_period_start_time
		}
		is_early_adoption_enabled = var.vm_cluster_cloud_automation_update_details_is_early_adoption_enabled
		is_freeze_period_enabled = var.vm_cluster_cloud_automation_update_details_is_freeze_period_enabled
	}
	data_collection_options {

		#Optional
		is_diagnostics_events_enabled = var.vm_cluster_data_collection_options_is_diagnostics_events_enabled
		is_health_monitoring_enabled = var.vm_cluster_data_collection_options_is_health_monitoring_enabled
		is_incident_logs_enabled = var.vm_cluster_data_collection_options_is_incident_logs_enabled
	}
	data_storage_size_in_tbs = var.vm_cluster_data_storage_size_in_tbs
	db_node_storage_size_in_gbs = var.vm_cluster_db_node_storage_size_in_gbs
	db_servers = var.vm_cluster_db_servers
	defined_tags = var.vm_cluster_defined_tags
	file_system_configuration_details {

		#Optional
		file_system_size_gb = var.vm_cluster_file_system_configuration_details_file_system_size_gb
		mount_point = var.vm_cluster_file_system_configuration_details_mount_point
	}
	freeform_tags = {"Department"= "Finance"}
	is_local_backup_enabled = var.vm_cluster_is_local_backup_enabled
	is_sparse_diskgroup_enabled = var.vm_cluster_is_sparse_diskgroup_enabled
	license_model = var.vm_cluster_license_model
	memory_size_in_gbs = var.vm_cluster_memory_size_in_gbs
	system_version = var.vm_cluster_system_version
	time_zone = var.vm_cluster_time_zone
}
```

## Argument Reference

The following arguments are supported:

* `cloud_automation_update_details` - (Optional) (Updatable) Specifies the properties necessary for cloud automation updates. This includes modifying the apply update time preference, enabling or disabling early adoption, and enabling, modifying, or disabling the update freeze period. 
	* `apply_update_time_preference` - (Optional) (Updatable) Configure the time slot for applying VM cloud automation software updates to the cluster. When nothing is selected, the default time slot is 12 AM to 2 AM UTC. Any 2-hour slot is available starting at 12 AM. 
		* `apply_update_preferred_end_time` - (Optional) (Updatable) End time for polling VM cloud automation software updates for the cluster. If the endTime is not specified, 2 AM UTC is used by default. 
		* `apply_update_preferred_start_time` - (Optional) (Updatable) Start time for polling VM cloud automation software updates for the cluster. If the startTime is not specified, 12 AM UTC is used by default. 
	* `freeze_period` - (Optional) (Updatable) Enables a freeze period for the VM cluster prohibiting the VMs from getting cloud automation software updates during critical business cycles. Freeze period start date. Starts at 12:00 AM UTC on the selected date and ends at 11:59:59 PM UTC on the selected date. Validates to ensure the freeze period does not exceed 45 days. 
		* `freeze_period_end_time` - (Optional) (Updatable) End time of the freeze period cycle. 
		* `freeze_period_start_time` - (Optional) (Updatable) Start time of the freeze period cycle. 
	* `is_early_adoption_enabled` - (Optional) (Updatable) Annotates whether the cluster should be part of early access to apply VM cloud automation software updates. Those clusters annotated as early access will download the software bits for cloud automation in the first week after the update is available, while other clusters will have to wait until the following week. 
	* `is_freeze_period_enabled` - (Optional) (Updatable) Specifies if the freeze period is enabled for the VM cluster to prevent the VMs from receiving cloud automation software updates during critical business cycles. Freeze period starts at 12:00 AM UTC and ends at 11:59:59 PM UTC on the selected date. Ensure that the freezing period does not exceed 45 days. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - (Required) (Updatable) The number of CPU cores to enable for the VM cluster. *Note:* If `cpu_core_count` is modified in `DISCONNECTED` state, the provider could experience a drift in Terraform state. To remediate this, refresh your Terraform state and update the configuration file when the Oracle Cloud Infrastructure connection is established.
* `data_collection_options` - (Optional) (Updatable) Indicates user preferences for the various diagnostic collection options for the VM cluster/Cloud VM cluster/VMBM DBCS. 
	* `is_diagnostics_events_enabled` - (Optional) (Updatable) Indicates whether diagnostic collection is enabled for the VM cluster/Cloud VM cluster/VMBM DBCS. Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues. Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system. You can enable diagnostic collection during VM cluster/Cloud VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` or `updateCloudVmCluster` API. 
	* `is_health_monitoring_enabled` - (Optional) (Updatable) Indicates whether health monitoring is enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel. You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster`, `UpdateCloudVmCluster` or `updateDbsystem` API. 
	* `is_incident_logs_enabled` - (Optional) (Updatable) Indicates whether incident logs and trace collection are enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster`, `updateCloudVmCluster` or `updateDbsystem` API. 
* `data_storage_size_in_gb` - (Optional) (Updatable) The data disk group size to be allocated in GBs.
* `data_storage_size_in_tbs` - (Optional) (Updatable) The data disk group size to be allocated in TBs.
* `db_node_storage_size_in_gbs` - (Optional) (Updatable) The local node storage to be allocated in GBs.
* `db_servers` - (Optional) The list of Db server.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `file_system_configuration_details` - (Optional) (Updatable) Details of the file system configuration of the VM cluster.
	* `file_system_size_gb` - (Optional) (Updatable) The file system size to be allocated in GBs.
	* `mount_point` - (Optional) (Updatable) The mount point of file system.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gi_version` - (Required) The Oracle Grid Infrastructure software version for the VM cluster.
* `is_local_backup_enabled` - (Optional) If true, database backup on local Exadata storage is configured for the VM cluster. If false, database backup on local Exadata storage is not available in the VM cluster. 
* `is_sparse_diskgroup_enabled` - (Optional) If true, the sparse disk group is configured for the VM cluster. If false, the sparse disk group is not created. 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the VM cluster. The default is BRING_YOUR_OWN_LICENSE. 
* `memory_size_in_gbs` - (Optional) (Updatable) The memory to be allocated in GBs.
* `ssh_public_keys` - (Required) (Updatable) The public key portion of one or more key pairs used for SSH access to the VM cluster.
* `system_version` - (Optional) Operating system version of the image.
* `time_zone` - (Optional) The time zone to use for the VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vm_cluster_network_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the VM cluster is located in.
* `cloud_automation_update_details` - Specifies the properties necessary for cloud automation updates. This includes modifying the apply update time preference, enabling or disabling early adoption, and enabling, modifying, or disabling the update freeze period. 
	* `apply_update_time_preference` - Configure the time slot for applying VM cloud automation software updates to the cluster. When nothing is selected, the default time slot is 12 AM to 2 AM UTC. Any 2-hour slot is available starting at 12 AM. 
		* `apply_update_preferred_end_time` - End time for polling VM cloud automation software updates for the cluster. If the endTime is not specified, 2 AM UTC is used by default. 
		* `apply_update_preferred_start_time` - Start time for polling VM cloud automation software updates for the cluster. If the startTime is not specified, 12 AM UTC is used by default. 
	* `freeze_period` - Enables a freeze period for the VM cluster prohibiting the VMs from getting cloud automation software updates during critical business cycles. Freeze period start date. Starts at 12:00 AM UTC on the selected date and ends at 11:59:59 PM UTC on the selected date. Validates to ensure the freeze period does not exceed 45 days. 
		* `freeze_period_end_time` - End time of the freeze period cycle. 
		* `freeze_period_start_time` - Start time of the freeze period cycle. 
	* `is_early_adoption_enabled` - Annotates whether the cluster should be part of early access to apply VM cloud automation software updates. Those clusters annotated as early access will download the software bits for cloud automation in the first week after the update is available, while other clusters will have to wait until the following week. 
	* `is_freeze_period_enabled` - Specifies if the freeze period is enabled for the VM cluster to prevent the VMs from receiving cloud automation software updates during critical business cycles. Freeze period starts at 12:00 AM UTC and ends at 11:59:59 PM UTC on the selected date. Ensure that the freezing period does not exceed 45 days. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpus_enabled` - The number of enabled CPU cores.
* `data_collection_options` - Indicates user preferences for the various diagnostic collection options for the VM cluster/Cloud VM cluster/VMBM DBCS. 
	* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster/Cloud VM cluster/VMBM DBCS. Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues. Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system. You can enable diagnostic collection during VM cluster/Cloud VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` or `updateCloudVmCluster` API. 
	* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel. You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster`, `UpdateCloudVmCluster` or `updateDbsystem` API. 
	* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster`, `updateCloudVmCluster` or `updateDbsystem` API. 
* `data_storage_size_in_gb` - Size of the DATA disk group in GBs.
* `data_storage_size_in_tbs` - Size, in terabytes, of the DATA disk group.
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `db_servers` - The list of Db server.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Exadata Cloud@Customer VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `file_system_configuration_details` - Details of the file system configuration of the VM cluster.
	* `file_system_size_gb` - The file system size to be allocated in GBs.
	* `mount_point` - The mount point of file system.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gi_version` - The Oracle Grid Infrastructure software version for the VM cluster.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
* `is_local_backup_enabled` - If true, database backup on local Exadata storage is configured for the VM cluster. If false, database backup on local Exadata storage is not available in the VM cluster. 
* `is_sparse_diskgroup_enabled` - If true, sparse disk group is configured for the VM cluster. If false, sparse disk group is not created. 
* `last_patch_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
* `license_model` - The Oracle license model that applies to the VM cluster. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `shape` - The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance. 
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the VM cluster.
* `state` - The current state of the VM cluster.
* `system_version` - Operating system version of the image.
* `time_created` - The date and time that the VM cluster was created.
* `time_zone` - The time zone of the Exadata infrastructure. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vm_cluster_network_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vm Cluster
	* `update` - (Defaults to 20 minutes), when updating the Vm Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Vm Cluster


## Import

VmClusters can be imported using the `id`, e.g.

```
$ terraform import oci_database_vm_cluster.test_vm_cluster "id"
```

