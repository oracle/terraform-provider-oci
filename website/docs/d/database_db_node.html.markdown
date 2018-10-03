---
layout: "oci"
page_title: "OCI: oci_database_db_node"
sidebar_current: "docs-oci-datasource-database-db_node"
description: |-
  Provides details about a specific DbNode
---

# Data Source: oci_database_db_node
The `oci_database_db_node` data source provides details about a specific DbNode

Gets information about the specified database node.

## Example Usage

```hcl
data "oci_database_db_node" "test_db_node" {
	#Required
	db_node_id = "${oci_database_db_node.test_db_node.id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `backup_vnic_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the backup VNIC.
* `db_system_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.
* `hostname` - The host name for the database node.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the database node.
* `software_storage_size_in_gb` - The size (in GB) of the block storage volume allocation for the DB system. This attribute applies only for virtual machine DB systems. 
* `state` - The current state of the database node.
* `time_created` - The date and time that the database node was created.
* `vnic_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VNIC.

