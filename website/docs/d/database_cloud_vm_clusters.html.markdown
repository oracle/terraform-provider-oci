---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_vm_clusters"
sidebar_current: "docs-oci-datasource-database-cloud_vm_clusters"
description: |-
  Provides the list of Cloud Vm Clusters in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_cloud_vm_clusters
This data source provides the list of Cloud Vm Clusters in Oracle Cloud Infrastructure Database service.

Gets a list of the cloud VM clusters in the specified compartment. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.


## Example Usage

```hcl
data "oci_database_cloud_vm_clusters" "test_cloud_vm_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
	display_name = var.cloud_vm_cluster_display_name
	state = var.cloud_vm_cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `cloud_exadata_infrastructure_id` - (Optional) If provided, filters the results for the specified cloud Exadata infrastructure.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only cloud VM clusters that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `cloud_vm_clusters` - The list of cloud_vm_clusters.

### CloudVmCluster Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the cloud Exadata infrastructure resource is located in.
* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the cloud VM cluster.

	**Subnet Restriction:** See the subnet restrictions information for **subnetId**. 
* `cloud_exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
* `cluster_name` - The cluster name for cloud VM cluster. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - The number of CPU cores enabled on the cloud VM cluster.
* `data_storage_percentage` - The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 35, 40, 60 and 80. The default is 80 percent assigned to DATA storage. See [Storage Configuration](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaoverview.htm#Exadata) in the Exadata documentation for details on the impact of the configuration settings on storage. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `disk_redundancy` - The type of redundancy configured for the cloud Vm cluster. NORMAL is 2-way redundancy. HIGH is 3-way redundancy. 
* `display_name` - The user-friendly name for the cloud VM cluster. The name does not need to be unique.
* `domain` - The domain name for the cloud VM cluster.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gi_version` - A valid Oracle Grid Infrastructure (GI) software version.
* `hostname` - The hostname for the cloud VM cluster.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud VM cluster.
* `is_local_backup_enabled` - If true, database backup on local Exadata storage is configured for the cloud VM cluster. If false, database backup on local Exadata storage is not available in the cloud VM cluster. 
* `is_sparse_diskgroup_enabled` - If true, sparse disk group is configured for the cloud VM cluster. If false, sparse disk group is not created. 
* `last_update_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance update history entry. This value is updated when a maintenance update starts.
* `license_model` - The Oracle license model that applies to the cloud VM cluster. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `listener_port` - The port number configured for the listener on the cloud VM cluster.
* `node_count` - The number of nodes in the cloud VM cluster. 
* `nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty. 
* `ocpu_count` - The number of OCPU cores to enable on the cloud VM cluster. Only 1 decimal place is allowed for the fractional part.
* `scan_dns_name` - The FQDN of the DNS record for the SCAN IP addresses that are associated with the cloud VM cluster. 
* `scan_dns_record_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the cloud VM cluster. 
* `scan_ip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the cloud VM cluster. SCAN IP addresses are typically used for load balancing and are not assigned to any interface. Oracle Clusterware directs the requests to the appropriate nodes in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `scan_listener_port_tcp` - The TCP Single Client Access Name (SCAN) port. The default port is 1521.
* `scan_listener_port_tcp_ssl` - The TCPS Single Client Access Name (SCAN) port. The default port is 2484.
* `shape` - The model name of the Exadata hardware running the cloud VM cluster. 
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the cloud VM cluster.
* `state` - The current state of the cloud VM cluster.
* `storage_size_in_gbs` - The storage allocation for the disk group, in gigabytes (GB).
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the cloud VM cluster.

	**Subnet Restrictions:**
	* For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `system_version` - Operating system version of the image.
* `time_created` - The date and time that the cloud VM cluster was created.
* `time_zone` - The time zone of the cloud VM cluster. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the cloud VM cluster. The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the Exadata Cloud Service instance to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `zone_id` - The OCID of the zone the cloud VM cluster is associated with. 

