---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_databases_remove"
sidebar_current: "docs-oci-resource-distributed_database-distributed_databases_remove"
description: |-
  Executes REMOVE operations in Oracle Cloud Infrastructure Distributed Database service.
---

# oci_distributed_database_distributed_databases_remove

Executes REMOVE operations against an existing distributed database. Configure one to three `remove` blocks; all blocks are submitted together in one PATCH request. The provider generates the API operation fields automatically.

## Example Usage

```hcl
resource "oci_distributed_database_distributed_databases_remove" "remove" {
  distributed_database_id = oci_distributed_database_distributed_database.example.id

  remove {
    selection = "systemRaftClusters[?name == 'raft-cluster-a'].databases[?name == 'database-a']"
  }
}
```

The example references an existing main database resource and input variables. Replace the topology and database names with names from your database.

## Argument Reference

* `distributed_database_id` - (Required) Target database OCID. Changing this value replaces the action resource and executes its configured operations against the new target.
* `remove` - (Required, Updatable) One to three repeatable operation blocks. Changing this list executes all configured operations, including unchanged blocks within that instance.
    * `selection` - (Required, Updatable) API selection identifying the topology members to remove. Use the actual database and topology names returned by the service.

Each block generates only `operation: REMOVE` and `selection`. No `value` is sent.

## Action Lifecycle

Creation executes the configured operations. An update executes them when the operation list changes, without an incremental trigger. Refresh reads the database state and does not execute operations. Unchanged action instances are not replayed when additional instances are added. Recreating an action instance, including after its state is removed, executes its configured operations again.

Resource labels are user-defined. For example, name two instances `remove` and `remove_2`; each submits its own operations. Keep successful instances unchanged and add another instance for additional operations. Changing an existing instance submits its full operation list again. Use `depends_on` to order instances targeting the same database when they must run sequentially.

Removing an operation block does not execute its inverse, but changing a nonempty list submits the remaining blocks again. At least one block is required. Removing the entire action resource only clears its Terraform state; it does not remove topology members or delete the target database. To remove members explicitly, use the corresponding REMOVE or REMOVEDG resource. Action changes do not replace or destroy the main database resource.

The provider polls the asynchronous work request and then refreshes the target until its top-level lifecycle state is `ACTIVE` or `INACTIVE`. Work requests ending in `FAILED`, `CANCELED`, or `NEEDS_ATTENTION` fail the action. Missing work-request IDs and API failures also produce diagnostics. Diagnostics identify the operation, database OCID, and available work-request OCID and service error details.

After adding shards, run [distributed_databases configure sharding](distributed_database_distributed_databases_configure_sharding.html) after the INSERT resource completes successfully so the new shards participate in the sharding topology. Configure an explicit dependency on the action resource. Each additional INSERT resource needs a new configure-sharding action instance, or an explicit invocation through the main resource's existing configure-sharding trigger.

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

* [INSERT](distributed_database_distributed_databases_insert.html)
* [REMOVE](distributed_database_distributed_databases_remove.html)
* [INSERTDG](distributed_database_distributed_databases_insert_dg.html)
* [REMOVEDG](distributed_database_distributed_databases_remove_dg.html)
* [Main database](distributed_database_distributed_database.html)
