---
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
	db_node_id = "${var.db_node_id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `backup_vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup VNIC.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `fault_domain` - The name of the fault domain the instance is contained in.
* `hostname` - The host name for the database node.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database node.
* `software_storage_size_in_gb` - The size (in GB) of the block storage volume allocation for the DB system. This attribute applies only for virtual machine DB systems. 
* `state` - The current state of the database node.
* `time_created` - The date and time that the database node was created.
* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC.

