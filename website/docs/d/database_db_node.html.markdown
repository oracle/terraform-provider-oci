---
layout: "oci"
page_title: "OCI: oci_database_db_node"
sidebar_current: "docs-oci-datasource-database-db_node"
description: |-
Provides details about a specific DbNode
---

# Data Source: oci_database_db_node
The DbNode data source provides details about a specific DbNode

Gets information about the specified database node.

## Example Usage

```hcl
data "oci_database_db_node" "test_db_node" {
	#Required
	db_node_id = "${var.db_node_db_node_id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `backup_vnic_id` - The OCID of the backup VNIC.
* `db_system_id` - The OCID of the DB System.
* `hostname` - The host name for the DB Node.
* `id` - The OCID of the DB Node.
* `software_storage_size_in_gb` - Storage size, in GBs, of the software volume that is allocated to the DB system. This is applicable only for VM-based DBs. 
* `state` - The current state of the database node.
* `time_created` - The date and time that the DB Node was created.
* `vnic_id` - The OCID of the VNIC.

