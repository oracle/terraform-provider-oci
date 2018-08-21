---
layout: "oci"
page_title: "OCI: oci_database_db_systems"
sidebar_current: "docs-oci-datasource-database-db_systems"
description: |-
  Provides a list of DbSystems
---

# Data Source: oci_database_db_systems
The `oci_database_db_systems` data source allows access to the list of OCI db_systems

Gets a list of the DB Systems in the specified compartment. You can specify a backupId to list only the DB Systems that support creating a database using this backup in this compartment.
    


## Example Usage

```hcl
data "oci_database_db_systems" "test_db_systems" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	backup_id = "${oci_database_backup.test_backup.id}"
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Optional) The OCID of the backup. Specify a backupId to list only the DB Systems that support creating a database using this backup in this compartment.
* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `db_systems` - The list of db_systems.

### DbSystem Reference

The following attributes are exported:

* `availability_domain` - The name of the Availability Domain that the DB System is located in.
* `backup_subnet_id` - The OCID of the backup network subnet the DB System is associated with. Applicable only to Exadata.

	**Subnet Restriction:** See above subnetId's 'Subnet Restriction'. to malfunction. 
* `cluster_name` - Cluster name for Exadata and 2-node RAC DB Systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - The OCID of the compartment.
* `cpu_core_count` - The number of CPU cores enabled on the DB System.
* `data_storage_percentage` - The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 40 and 80. 
* `data_storage_size_in_gb` - Data storage size, in GBs, that is currently available to the DB system. This is applicable only for VM-based DBs. 
* `database_edition` - The Oracle Database Edition that applies to all the databases on the DB System. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `disk_redundancy` - The type of redundancy configured for the DB System. Normal is 2-way redundancy. High is 3-way redundancy. 
* `display_name` - The user-friendly name for the DB System. It does not have to be unique.
* `domain` - The domain name for the DB System.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The host name for the DB Node.
* `id` - The OCID of the DB System.
* `last_patch_history_entry_id` - The OCID of the last patch history. This is updated as soon as a patch operation is started.
* `license_model` - The Oracle license model that applies to all the databases on the DB System. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycleState.
* `listener_port` - The port number configured for the listener on the DB System.
* `node_count` - Number of nodes in this DB system. For RAC DBs, this will be greater than 1. 
* `reco_storage_size_in_gb` - RECO/REDO storage size, in GBs, that is currently allocated to the DB system. This is applicable only for VM-based DBs. 
* `scan_dns_record_id` - The OCID of the DNS record for the SCAN IP addresses that are associated with the DB System. 
* `scan_ip_ids` - The OCID of the Single Client Access Name (SCAN) IP addresses associated with the DB System. SCAN IP addresses are typically used for load balancing and are not assigned to any interface. Clusterware directs the requests to the appropriate nodes in the cluster.

	* For a single-node DB System, this list is empty. 
* `shape` - The shape of the DB System. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the DB System.
* `state` - The current state of the DB System.
* `subnet_id` - The OCID of the subnet the DB System is associated with.

	**Subnet Restrictions:**
	* For single node and 2-node (RAC) DB Systems, do not use a subnet that overlaps with 192.168.16.16/28
	* For Exadata and VM-based RAC DB Systems, do not use a subnet that overlaps with 192.168.128.0/20

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time the DB System was created.
* `version` - The version of the DB System.
* `vip_ids` - The OCID of the virtual IP (VIP) addresses associated with the DB System. The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB System to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

	* For a single-node DB System, this list is empty. 

