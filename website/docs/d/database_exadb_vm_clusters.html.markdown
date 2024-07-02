---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadb_vm_clusters"
sidebar_current: "docs-oci-datasource-database-exadb_vm_clusters"
description: |-
  Provides the list of Exadb Vm Clusters in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadb_vm_clusters
This data source provides the list of Exadb Vm Clusters in Oracle Cloud Infrastructure Database service.

Gets a list of the Exadata VM clusters on Exascale Infrastructure in the specified compartment. Applies to Exadata Database Service on Exascale Infrastructure only.


## Example Usage

```hcl
data "oci_database_exadb_vm_clusters" "test_exadb_vm_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.exadb_vm_cluster_display_name
	exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
	state = var.exadb_vm_cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `exascale_db_storage_vault_id` - (Optional) A filter to return only Exadata VM clusters on Exascale Infrastructure that match the given Exascale Database Storage Vault ID. 
* `state` - (Optional) A filter to return only Exadata VM clusters on Exascale Infrastructure that match the given lifecycle state exactly. 


## Attributes Reference

The following attributes are exported:

* `exadb_vm_clusters` - The list of exadb_vm_clusters.

### ExadbVmCluster Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain in which the Exadata VM cluster on Exascale Infrastructure is located. 
* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
* `cluster_name` - The cluster name for Exadata VM cluster on Exascale Infrastructure. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `data_collection_options` - Indicates user preferences for the various diagnostic collection options for the VM cluster/Cloud VM cluster/VMBM DBCS. 
	* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster/Cloud VM cluster/VMBM DBCS. Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues. Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system. You can enable diagnostic collection during VM cluster/Cloud VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` or `updateCloudVmCluster` API. 
	* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel. You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster`, `UpdateCloudVmCluster` or `updateDbsystem` API. 
	* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster`, `updateCloudVmCluster` or `updateDbsystem` API. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Exadata VM cluster on Exascale Infrastructure. The name does not need to be unique.
* `domain` - A domain name used for the Exadata VM cluster on Exascale Infrastructure. If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.  Applies to Exadata Database Service on Exascale Infrastructure only. 
* `exascale_db_storage_vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gi_version` - A valid Oracle Grid Infrastructure (GI) software version.
* `grid_image_id` - Grid Setup will be done using this grid image id
* `grid_image_type` - The type of Grid Image
* `hostname` - The hostname for the Exadata VM cluster on Exascale Infrastructure. The hostname must begin with an alphabetic character, and  can contain alphanumeric characters and hyphens (-). For Exadata systems, the maximum length of the hostname is 12 characters.

	The maximum length of the combined hostname and domain is 63 characters.

	**Note:** The hostname must be unique within the subnet. If it is not unique,  then the Exadata VM cluster on Exascale Infrastructure will fail to provision. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata VM cluster on Exascale Infrastructure.
* `iorm_config_cache` - The IORM settings of the Exadata DB system. 
	* `db_plans` - An array of IORM settings for all the database in the Exadata DB system. 
		* `db_name` - The database name. For the default `DbPlan`, the `dbName` is `default`. 
		* `flash_cache_limit` - The flash cache limit for this database. This value is internally configured based on the share value assigned to the database. 
		* `share` - The relative priority of this database. 
	* `lifecycle_details` - Additional information about the current `lifecycleState`. 
	* `objective` - The current value for the IORM objective. The default is `AUTO`. 
	* `state` - The current state of IORM configuration for the Exadata DB system. 
* `last_update_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance update history entry. This value is updated when a maintenance update starts.
* `license_model` - The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `listener_port` - The port number configured for the listener on the Exadata VM cluster on Exascale Infrastructure.
* `node_config` - The configuration of each node in the Exadata VM cluster on Exascale Infrastructure.
	* `enabled_cpu_count_per_node` - The number of ECPU to enable for each node.
	* `memory_size_in_gbs_per_node` - The memory that you want to be allocated in GBs to each node. Memory is calculated based on 11 GB per VM core reserved.
	* `total_cpu_count_per_node` - The number of Total ECPU for each node.
	* `vm_file_system_storage_size_gbs_per_node` - The file system storage in GBs for each node.
	* `snapshot_file_system_storage_size_gbs_per_node` - The file system storage in GBs for snapshot for each node.
	* `total_file_system_storage_size_gbs_per_node` - Total file system storage in GBs for each node.
* `node_resource` - The list of node in the Exadata VM cluster on Exascale Infrastructure.
	* `node_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the node.
	* `node_hostname` - The host name for the node.
	* `state` - The current state of the node.
* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty. 
* `private_zone_id` - The private zone ID in which you want DNS records to be created. 
* `scan_dns_name` - The FQDN of the DNS record for the SCAN IP addresses that are associated with the Exadata VM cluster on Exascale Infrastructure. 
* `scan_dns_record_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the Exadata VM cluster on Exascale Infrastructure. 
* `scan_ip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the Exadata VM cluster on Exascale Infrastructure. SCAN IP addresses are typically used for load balancing and are not assigned to any interface. Oracle Clusterware directs the requests to the appropriate nodes in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `scan_listener_port_tcp` - The TCP Single Client Access Name (SCAN) port. The default port is 1521.
* `scan_listener_port_tcp_ssl` - The Secured Communication (TCPS) protocol Single Client Access Name (SCAN) port. The default port is 2484. 
* `shape` - The shape of the Exadata VM cluster on Exascale Infrastructure resource 
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
* `state` - The current state of the Exadata VM cluster on Exascale Infrastructure.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `system_version` - Operating system version of the image.
* `time_created` - The date and time that the Exadata VM cluster on Exascale Infrastructure was created.
* `time_zone` - The time zone to use for the Exadata VM cluster on Exascale Infrastructure. For details, see [Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm). 
* `vip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the Exadata VM cluster on Exascale Infrastructure.  The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the Exadata Cloud Service instance to  enable failover. If one node fails, then the VIP is reassigned to another active node in the cluster. 
* `zone_id` - The OCID of the zone with which the Exadata VM cluster on Exascale Infrastructure is associated. 

