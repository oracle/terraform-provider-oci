---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_autonomous_databases_insert_dg"
sidebar_current: "docs-oci-resource-distributed_database-distributed_autonomous_databases_insert_dg"
description: |-
  Executes INSERTDG operations in Oracle Cloud Infrastructure Distributed Database service.
---

# oci_distributed_database_distributed_autonomous_databases_insert_dg

Executes INSERTDG operations against an existing distributed autonomous database. Configure one to three `insert_dg` blocks; all blocks are submitted together in one PATCH request. The provider generates the API operation fields automatically.

## Example Usage

```hcl
resource "oci_distributed_database_distributed_autonomous_databases_insert_dg" "insert_dg" {
  distributed_autonomous_database_id = oci_distributed_database_distributed_autonomous_database.example.id

  insert_dg {
    selection    = "catalogDetails[?originalReplica.database.name == 'catalog-database-a'].originalReplica.database"
    replica_name = "standby"
    database {
      source                         = "ADBD_EXISTING_CLUSTER"
      cloud_autonomous_vm_cluster_id = var.standby_vm_cluster_id
    }
  }
}
```

The example references an existing main database resource and input variables. Replace the topology and database names with names from your database.

## Argument Reference

* `distributed_autonomous_database_id` - (Required) Target database OCID. Changing this value replaces the action resource and executes its configured operations against the new target.
* `insert_dg` - (Required, Updatable) One to three repeatable operation blocks. Changing this list executes all configured operations, including unchanged blocks within that instance.
    * `selection` - (Required, Updatable) API selection that must resolve to exactly one original database. For example, select a catalog original database as shown in the example, or use `systemDataGuardDatabases.originalReplica.databases[?name == 'database-a']` for a system Data Guard shard.
    * `replica_name` - (Required, Updatable) Name of the new Data Guard replica. Starts with a letter, contains only letters, digits, and underscores, and is at most 40 characters.
    * `database` - (Required, Updatable) Exactly one standby database placement block. Fields required by the selected source must be supplied.
        * `cloud_autonomous_vm_cluster_id` - (Required, Updatable) Cloud Autonomous VM cluster OCID for the standby.
        * `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time preference based on data loss tolerance in seconds.
        * `protection_mode` - (Optional, Updatable) Data Guard protection mode. Defaults to `MAXIMUM_PERFORMANCE`.
        * `source` - (Required, Updatable) Placement source. Must be `ADBD_EXISTING_CLUSTER`.
        * `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database.

Each block generates `operation: MERGE` with a JSON-encoded `value` containing `subOperation: INSERTDG` and `dataGuardReplicas: [{name: replica_name, database: ...}]`. The provider maps Terraform field names to API field names.

Autonomous Data Guard placement does not accept admin credentials, compute sizing, or storage sizing in this block.

## Action Lifecycle

Creation executes the configured operations. An update executes them when the operation list changes, without an incremental trigger. Refresh reads the database state and does not execute operations. Unchanged action instances are not replayed when additional instances are added. Recreating an action instance, including after its state is removed, executes its configured operations again.

Resource labels are user-defined. For example, name two instances `insert_dg` and `insert_dg_2`; each submits its own operations. Keep successful instances unchanged and add another instance for additional operations. Changing an existing instance submits its full operation list again. Use `depends_on` to order instances targeting the same database when they must run sequentially.

Removing an operation block does not execute its inverse, but changing a nonempty list submits the remaining blocks again. At least one block is required. Removing the entire action resource only clears its Terraform state; it does not remove topology members or delete the target database. To remove members explicitly, use the corresponding REMOVE or REMOVEDG resource. Action changes do not replace or destroy the main database resource.

The provider polls the asynchronous work request and then refreshes the target until its top-level lifecycle state is `ACTIVE` or `INACTIVE`. Work requests ending in `FAILED`, `CANCELED`, or `NEEDS_ATTENTION` fail the action. Missing work-request IDs and API failures also produce diagnostics. Diagnostics identify the operation, database OCID, and available work-request OCID and service error details.

After adding shards, run [distributed_autonomous_databases configure sharding](distributed_database_distributed_autonomous_databases_configure_sharding.html) after the INSERT resource completes successfully so the new shards participate in the sharding topology. Configure an explicit dependency on the action resource. Each additional INSERT resource needs a new configure-sharding action instance, or an explicit invocation through the main resource's existing configure-sharding trigger.

## Attributes Reference

* `id` - Terraform identifier for this action instance; this is not the target database OCID.
* `work_request_id` - OCID of the work request submitted by the most recent action.
* `state` - Most recently refreshed top-level lifecycle state of the target database.
* `lifecycle_details` - Lifecycle details returned by the target database.
* `retry_token` - Retry token used for the action request.

## Timeouts

The `timeouts` block accepts `create`, `read`, `update`, and `delete`, each defaulting to 12 hours. Destroy only clears action state.

## Import

Import is not supported for this resource.

## Related Resources

* [INSERT](distributed_database_distributed_autonomous_databases_insert.html)
* [REMOVE](distributed_database_distributed_autonomous_databases_remove.html)
* [INSERTDG](distributed_database_distributed_autonomous_databases_insert_dg.html)
* [REMOVEDG](distributed_database_distributed_autonomous_databases_remove_dg.html)
* [Main database](distributed_database_distributed_autonomous_database.html)
