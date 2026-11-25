---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_autonomous_databases_insert"
sidebar_current: "docs-oci-resource-distributed_database-distributed_autonomous_databases_insert"
description: |-
  Executes INSERT operations in Oracle Cloud Infrastructure Distributed Database service.
---

# oci_distributed_database_distributed_autonomous_databases_insert

Executes INSERT operations against an existing distributed autonomous database. Configure one to three `insert` blocks; all blocks are submitted together in one PATCH request. The provider generates the API operation fields automatically.

## Example Usage

```hcl
resource "oci_distributed_database_distributed_autonomous_databases_insert" "insert" {
  distributed_autonomous_database_id = oci_distributed_database_distributed_autonomous_database.example.id

  insert {
    selection = "systemRaftClusters"
    value = jsonencode({
      name              = "raft-cluster-a"
      replicationFactor = 3
      witnessCount      = 0
      ruMode            = "READ_WRITE"
      databases = [{
        source                     = "ADBD_EXISTING_CLUSTER"
        cloudAutonomousVmClusterId = var.shard_vm_cluster_id
        computeCount               = 2
        dataStorageSizeInGbs       = 32
        adminPassword              = var.shard_admin_password
      }]
    })
  }
}

resource "oci_distributed_database_distributed_autonomous_databases_insert" "insert_2" {
  distributed_autonomous_database_id = oci_distributed_database_distributed_autonomous_database.example.id

  depends_on = [oci_distributed_database_distributed_autonomous_databases_insert.insert]

  insert {
    selection = "systemRaftClusters"
    value = jsonencode({
      name              = "raft-cluster-b"
      replicationFactor = 3
      witnessCount      = 0
      ruMode            = "READ_WRITE"
      databases = [{
        source                     = "ADBD_EXISTING_CLUSTER"
        cloudAutonomousVmClusterId = var.shard_vm_cluster_id_2
        computeCount               = 2
        dataStorageSizeInGbs       = 32
        adminPassword              = var.shard_admin_password
      }]
    })
  }
}
```

The example references an existing main database resource and input variables. Replace the topology and database names with names from your database.

## Argument Reference

* `distributed_autonomous_database_id` - (Required) Target database OCID. Changing this value replaces the action resource and executes its configured operations against the new target.
* `insert` - (Required, Updatable) One to three repeatable operation blocks. Changing this list executes all configured operations, including unchanged blocks within that instance.
    * `selection` - (Required, Updatable) API selection identifying the topology to extend. Supported roots are `catalogDetails`, `userShardSpaces`, `systemRaftClusters`, `compositeRaftShardSpaces`, `systemDataGuardDatabases`, and `compositeDataGuardShardSpaces`; use the root that matches the database sharding and replication methods.
    * `value` - (Required, Sensitive, Updatable) Nonempty JSON object using API field names and the create payload shape for the selected topology. Use `jsonencode` to construct the document. JSON and required payload fields are validated before the request. Equivalent JSON formatting does not cause an update. For a new system Raft cluster, include `replicationFactor`, `witnessCount`, and `ruMode`; adding databases to an existing named cluster does not require those new-cluster fields.

Each block generates `operation: INSERT`. The provider encodes `value` as the JSON string expected by the API.

The three-item limit counts operation blocks, rather than the number of databases inside a valid payload. Sensitive values remain in Terraform state; protect access to the state backend.

## Action Lifecycle

Creation executes the configured operations. An update executes them when the operation list changes, without an incremental trigger. Refresh reads the database state and does not execute operations. Unchanged action instances are not replayed when additional instances are added. Recreating an action instance, including after its state is removed, executes its configured operations again.

Resource labels are user-defined. For example, name two instances `insert` and `insert_2`; each submits its own operations. Keep successful instances unchanged and add another instance for additional operations. Changing an existing instance submits its full operation list again. Use `depends_on` to order instances targeting the same database when they must run sequentially.

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
