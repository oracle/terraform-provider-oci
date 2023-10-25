---
subcategory: "NoSQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_nosql_table_replica"
sidebar_current: "docs-oci-resource-nosql-table_replica"
description: |-
  Provides the Table Replica resource in Oracle Cloud Infrastructure NoSQL Database service
---

# oci_nosql_table_replica
This resource provides the Table Replica resource in Oracle Cloud Infrastructure NoSQL Database service.

Add a replica for this table

## Example Usage

```hcl
resource "oci_nosql_table_replica" "test_table_replica" {
	#Required
	region = var.table_replica_region
	table_name_or_id = oci_nosql_table_name_or.test_table_name_or.id

	#Optional
	compartment_id = var.compartment_id
	max_read_units = var.table_replica_max_read_units
	max_write_units = var.table_replica_max_write_units
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the table's compartment.  Required if the tableNameOrId path parameter is a table name. Optional if tableNameOrId is an OCID.  If tableNameOrId is an OCID, and compartmentId is supplied, the latter must match the identified table's compartmentId. 
* `max_read_units` - (Optional) Maximum sustained read throughput limit for the new replica table. If not specified, the local table's read limit is used. 
* `max_write_units` - (Optional) Maximum sustained write throughput limit for the new replica table. If not specified, the local table's write limit is used. 
* `region` - (Required) Name of the remote region in standard Oracle Cloud Infrastructure format, i.e. us-ashburn-1 
* `table_name_or_id` - (Required) A table name within the compartment, or a table OCID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Table Replica
	* `update` - (Defaults to 20 minutes), when updating the Table Replica
	* `delete` - (Defaults to 20 minutes), when destroying the Table Replica


## Import

TableReplicas can be imported using the `id`, e.g.

```
$ terraform import oci_nosql_table_replica.test_table_replica "tables/{tableNameOrId}/replicas/{region}"
```

