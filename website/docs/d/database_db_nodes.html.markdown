---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_nodes"
sidebar_current: "docs-oci-datasource-database-db_nodes"
description: |-
  Provides the list of Db Nodes in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_nodes
This data source provides the list of Db Nodes in Oracle Cloud Infrastructure Database service.

Lists the database nodes in the specified compartment. A database node is a server running database software. In addition to the other required parameters, either '--db-system-id' or '--vm-cluster-id' also must be provided, depending on the service being accessed.


## Example Usage

```hcl
data "oci_database_db_nodes" "test_db_nodes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	db_server_id = oci_database_db_server.test_db_server.id
	db_system_id = oci_database_db_system.test_db_system.id
	state = var.db_node_state
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_server_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exacc Db server.
* `db_system_id` - (Optional) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). If provided, filters the results to the set of database versions which are supported for the DB system.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `vm_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.


## Attributes Reference

The following attributes are exported:

* `db_nodes` - The list of db_nodes.

### DbNode Reference

The following attributes are exported:

* `additional_details` - Additional information about the planned maintenance.
* `backup_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup IP address associated with the database node. Use this OCID with either the [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp) or the [GetPublicIpByPrivateIpId](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PublicIp/GetPublicIpByPrivateIpId) API to get the IP address needed to make a database connection.

	**Note:** Applies only to Exadata Cloud Service. 
* `backup_vnic2id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second backup VNIC.

	**Note:** Applies only to Exadata Cloud Service. 
* `backup_vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup VNIC.
* `cpu_core_count` - The number of CPU cores enabled on the Db node.
* `db_node_storage_size_in_gbs` - The allocated local node storage in GBs on the Db node.
* `db_server_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exacc Db server associated with the database node.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `fault_domain` - The name of the Fault Domain the instance is contained in.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `host_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host IP address associated with the database node. Use this OCID with either the  [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp) or the [GetPublicIpByPrivateIpId](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PublicIp/GetPublicIpByPrivateIpId) API to get the IP address  needed to make a database connection.
	**Note:** Applies only to Exadata Cloud Service. 
* `hostname` - The host name for the database node.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database node.
* `lifecycle_details` - Information about the current lifecycle state.
* `maintenance_type` - The type of database node maintenance.
* `memory_size_in_gbs` - The allocated memory in GBs on the Db node.
* `software_storage_size_in_gb` - The size (in GB) of the block storage volume allocation for the DB system. This attribute applies only for virtual machine DB systems. 
* `state` - The current state of the database node.
* `time_created` - The date and time that the database node was created.
* `time_maintenance_window_end` - End date and time of maintenance window.
* `time_maintenance_window_start` - Start date and time of maintenance window.
* `total_cpu_core_count` - The total number of CPU cores reserved on the Db node.
* `vnic2id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the second VNIC.

	**Note:** Applies only to Exadata Cloud Service. 
* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC.

