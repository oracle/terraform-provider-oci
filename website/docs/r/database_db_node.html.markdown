---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node"
sidebar_current: "docs-oci-resource-database-db_node"
description: |-
  Provides the Db Node resource in Oracle Cloud Infrastructure Database service
---

# oci_database_db_node
This resource provides the Db Node resource in Oracle Cloud Infrastructure Database service.

Updates the specified database node.

## Example Usage

```hcl
resource "oci_database_db_node" "test_db_node" {
	#Required
	db_node_id = oci_database_db_node.test_db_node.id

	#Optional
	defined_tags = var.db_node_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - Additional information about the planned maintenance.
* `backup_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup IP address associated with the database node. Use this OCID with either the [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp) or the [GetPublicIpByPrivateIpId](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PublicIp/GetPublicIpByPrivateIpId) API to get the IP address  needed to make a database connection.

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Db Node
	* `update` - (Defaults to 20 minutes), when updating the Db Node
	* `delete` - (Defaults to 20 minutes), when destroying the Db Node


## Import

DbNodes can be imported using the `id`, e.g.

```
$ terraform import oci_database_db_node.test_db_node "id"
```

