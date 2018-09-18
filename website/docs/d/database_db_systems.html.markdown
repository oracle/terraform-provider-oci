---
layout: "oci"
page_title: "OCI: oci_database_db_systems"
sidebar_current: "docs-oci-datasource-database-db_systems"
description: |-
  Provides a list of DbSystems
---

# Data Source: oci_database_db_systems
The `oci_database_db_systems` data source allows access to the list of OCI db_systems

Gets a list of the DB systems in the specified compartment. You can specify a backupId to list only the DB systems that support creating a database using this backup in this compartment.
    


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

* `backup_id` - (Optional) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the backup. Specify a backupId to list only the DB systems that support creating a database using this backup in this compartment.
* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `db_systems` - The list of db_systems.

### DbSystem Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the DB system is located in.
* `backup_subnet_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.

	**Subnet Restriction:** See the subnet restrictions information for **subnetId**. 
* `cluster_name` - The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - The number of CPU cores enabled on the DB system.
* `data_storage_percentage` - The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 40 and 80. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems. 
* `data_storage_size_in_gb` - The data storage size, in gigabytes, that is currently available to the DB system. Applies only for virtual machine DB systems. 
* `database_edition` - The Oracle Database edition that applies to all the databases on the DB system. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `disk_redundancy` - The type of redundancy configured for the DB system. NORMAL is 2-way redundancy. HIGH is 3-way redundancy. 
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `domain` - The domain name for the DB system.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The hostname for the DB system.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.
* `last_patch_history_entry_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
* `license_model` - The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED. Allowed values are: LICENSE_INCLUDED, BRING_YOUR_OWN_LICENSE.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `listener_port` - The port number configured for the listener on the DB system.
* `node_count` - The number of nodes in the DB system. For RAC DB systems, the value is greater than 1. 
* `reco_storage_size_in_gb` - The RECO/REDO storage size, in gigabytes, that is currently allocated to the DB system. Applies only for virtual machine DB systems. 
* `scan_dns_record_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the DB system. 
* `scan_ip_ids` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the DB system. SCAN IP addresses are typically used for load balancing and are not assigned to any interface. Oracle Clusterware directs the requests to the appropriate nodes in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `shape` - The shape of the DB system. The shape determines resources to allocate to the DB system.
	* For virtual machine shapes, the number of CPU cores and memory
	* For bare metal and Exadata shapes, the number of CPU cores, storage, and memory 
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the DB system.
* `state` - The current state of the DB system.
* `subnet_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time the DB system was created.
* `version` - The Oracle Database version of the DB system.
* `vip_ids` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the DB system. The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB system to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

	**Note:** For a single-node DB system, this list is empty. 

