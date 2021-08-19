---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node"
sidebar_current: "docs-oci-datasource-database-db_node"
description: |-
  Provides details about a specific Db Node in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_node
This data source provides details about a specific Db Node resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified database node.

## Example Usage

```hcl
data "oci_database_db_node" "test_db_node" {
	#Required
	db_node_id = var.db_node_id
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `additional_details` - Additional information about the planned maintenance.
* `backup_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup IP address associated with the database node.

	**Note:** Applies only to Exadata Cloud Service. 
* `backup_vnic2id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second backup VNIC.

	**Note:** Applies only to Exadata Cloud Service. 
* `backup_vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup VNIC.
* `cpu_core_count` - The number of CPU cores enabled on the Db node.
* `db_node_storage_size_in_gbs` - The allocated local node storage in GBs on the Db node.
* `db_server_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exacc Db server associated with the database node.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `fault_domain` - The name of the Fault Domain the instance is contained in.
* `host_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host IP address associated with the database node.

	**Note:** Applies only to Exadata Cloud Service. 
* `hostname` - The host name for the database node.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database node.
* `maintenance_type` - The type of database node maintenance.
* `memory_size_in_gbs` - The allocated memory in GBs on the Db node.
* `software_storage_size_in_gb` - The size (in GB) of the block storage volume allocation for the DB system. This attribute applies only for virtual machine DB systems. 
* `state` - The current state of the database node.
* `time_created` - The date and time that the database node was created.
* `time_maintenance_window_end` - End date and time of maintenance window.
* `time_maintenance_window_start` - Start date and time of maintenance window.
* `vnic2id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second VNIC.

	**Note:** Applies only to Exadata Cloud Service. 
* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC.

