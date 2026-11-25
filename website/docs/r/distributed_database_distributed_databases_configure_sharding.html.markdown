---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_databases_configure_sharding"
sidebar_current: "docs-oci-resource-distributed_database-distributed_databases_configure_sharding"
description: |-
  Provides the Distributed Databases Configure Sharding resource in Oracle Cloud Infrastructure Distributed Database service
---

# oci_distributed_database_distributed_databases_configure_sharding
This resource provides the Distributed Databases Configure Sharding resource in Oracle Cloud Infrastructure Distributed Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/


Once all components of Globally distributed database are provisioned, and signed GSM certificates are successfully uploaded, this
api shall be invoked to configure sharding on the Globally distributed database. Note that this 'ConfigureSharding' API also needs to be
invoked after successfully adding a new shard to the Globally distributed database using PATCH api. If this API is not
invoked after successfully adding a new shard, then that new shard will not be a participant in sharding topology of
the Globally distributed database.


When shards are added through [distributed_databases INSERT](distributed_database_distributed_databases_insert.html), make this action depend on the successful INSERT resource with `depends_on`. A database ID reference alone does not order the two actions. For each additional INSERT resource, create another configure-sharding action instance or invoke configure sharding through the main database resource's existing trigger; an unchanged configure-sharding instance does not rerun automatically.

## Example Usage

```hcl
resource "oci_distributed_database_distributed_databases_configure_sharding" "test_distributed_databases_configure_sharding" {
	#Required
	distributed_database_id = oci_distributed_database_distributed_database.test_distributed_database.id

	#Optional
	is_rebalance_required = var.distributed_databases_configure_sharding_is_rebalance_required
}
```

## Argument Reference

The following arguments are supported:

* `distributed_database_id` - (Required) Globally distributed database identifier.
* `is_rebalance_required` - (Optional) The flag to indicate whether chunks need to be re-balanced. This flag is not applicable for USER-defined sharding type. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 12 hours), when creating the Distributed Databases Configure Sharding
	* `update` - (Defaults to 12 hours), reserved for updates to the Distributed Databases Configure Sharding
	* `delete` - (Defaults to 12 hours), when destroying the Distributed Databases Configure Sharding


## Import

Import is not supported for this resource.

