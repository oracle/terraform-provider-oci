---
layout: "oci"
page_title: "OCI: oci_database_db_nodes"
sidebar_current: "docs-oci-datasource-database-db_nodes"
description: |-
  Provides a list of DbNodes
---

# Data Source: oci_database_db_nodes
The DbNodes data source allows access to the list of OCI db_nodes

Gets a list of database nodes in the specified DB System and compartment. A database node is a server running database software.


## Example Usage

```hcl
data "oci_database_db_nodes" "test_db_nodes" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB System.


## Attributes Reference

The following attributes are exported:

* `db_nodes` - The list of db_nodes.

### DbNode Reference

The following attributes are exported:

* `backup_vnic_id` - The OCID of the backup VNIC.
* `db_system_id` - The OCID of the DB System.
* `hostname` - The host name for the DB Node.
* `id` - The OCID of the DB Node.
* `software_storage_size_in_gb` - Storage size, in GBs, of the software volume that is allocated to the DB system. This is applicable only for VM-based DBs. 
* `state` - The current state of the database node.
* `time_created` - The date and time that the DB Node was created.
* `vnic_id` - The OCID of the VNIC.

