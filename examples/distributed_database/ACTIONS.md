# DDB operations and recovery

Use this guide to update an existing distributed database (DDB) or replace failed database members using action API resources. To add or remove members, use INSERT, REMOVE, INSERTDG, or REMOVEDG; editing the original catalog or shard lists does not perform these operations. Run Terraform from the DDB type directory with the same state and `terraform.tfvars` used to create it.

## Recovery from partial creation

A database creation failure can leave some catalog or shard resources provisioned. Review the Terraform error and work-request details to identify the failure. For a DDB retained in Terraform state, inspect a refresh-only plan before choosing a recovery action:

```shell
terraform plan -refresh-only -var-file=terraform.tfvars -out=inspect.tfplan
terraform show inspect.tfplan
```

This reads the current resource details without changing the DDB.

| What you find | Next step |
|---|---|
| The work request is still running, even though the client timed out | Wait for the request to finish before retrying. |
| The DDB and catalog exist, but a shard database failed or is missing | Correct the cause and use supported REMOVE/INSERT operations to replace the affected member. |
| The DDB or catalog is unavailable, or the service rejects member recovery | Resolve the service failure before proceeding with member operations. Cleanup and recreation may be required. |

For a retained DDB with a failed member:

1. Record the DDB OCID, work-request OCID, failed member name, and the healthy member OCIDs. Check the Terraform state with `terraform state list`.
2. Fix the cause reported by the work request, such as insufficient capacity, an invalid secret, or a resource access issue.
3. If the failed member is still present and the service allows its removal, use [REMOVE](#remove) to remove only that member. Confirm the work request succeeds before continuing. A missing member does not need removal.
4. Use [INSERT](#insert) to add the replacement to the intended cluster or shardspace. Keep the original creation inputs and completed actions unchanged. The DDB type README contains the INSERT example.
5. Complete Configure Sharding, then verify the expected members, their readiness, and the work-request results. Confirm the DDB and healthy member OCIDs have not changed.

Review the plan before every apply. If it proposes replacing the whole DDB, resolve that condition before continuing with member recovery.

A **“No changes”** plan does not confirm that every member is healthy. Terraform does not automatically replace failed members; use the recovery steps above.

## Invoke an update separately

For INSERT, REMOVE, INSERTDG, or REMOVEDG, set the required operation under `actions` in your local `terraform.tfvars`. Leave unused lists empty and preserve completed action inputs. For backup settings, follow [Update backup configuration](#update-backup-configuration).

```shell
terraform plan -var-file=terraform.tfvars -out=update.tfplan
terraform show update.tfplan
terraform apply update.tfplan
```

Check the plan before applying. Terraform loads `actions.tf` automatically; you do not run that file separately.

Each action submits a request. Changing a completed action's input or recreating its resource can submit the operation again. Removing the action from Terraform does not undo its changes to the DDB.

## INSERT

Use INSERT to add shard databases. Each entry has a `selection` identifying the topology and a `value` containing the new database details.

| DDB type | Selection root |
|---|---|
| System-managed Raft | `systemRaftClusters` |
| Composite Raft | `compositeRaftShardSpaces` |
| System-managed with Data Guard | `systemDataGuardDatabases` |
| Composite with Data Guard | `compositeDataGuardShardSpaces` |
| User-defined | `userShardSpaces` |

`value` is a JSON string with API camelCase field names. Use a JSON heredoc in HCL tfvars or an escaped JSON string in JSON tfvars. Supply admin password secret OCIDs for the new databases. Each [DDB type README](README.md#supported-ddb-types) has a complete INSERT example.

INSERT runs its paired Configure Sharding step after the insertion succeeds.

For Raft DDBs, insert the replacement into the same cluster and, for Composite Raft, the same shardspace as the failed member.

## REMOVE

Use REMOVE to delete a selected shard member. For example, this entry inside `actions` selects one database in a System-managed Raft cluster:

```hcl
remove = [{
  selection = "systemRaftClusters[?name == 'rwcluster'].databases[?name == 'database-a']"
}]
```

Replace the names with those returned by the service. Use the selection shown in your DDB type README and confirm removal is supported by its current topology.

## INSERTDG

Use INSERTDG to add a Data Guard replica to an existing primary database. Select the primary and supply the replica name and standby database details.

For a catalog standby on ExaScale, set this entry inside `actions`:

```hcl
insert_dg = [{
  selection    = "catalogDetails[?originalReplica.database.name == 'catalog-database-a'].originalReplica.database"
  replica_name = "standby2"
  database = {
    source                   = "XS_EXISTING_CLUSTER"
    vm_cluster_id            = "<standby-vm-cluster-ocid>"
    admin_password_secret_id = "<standby-password-secret-ocid>"
  }
}]
```

For Autonomous, use these database fields instead:

```hcl
database = {
  source                         = "ADBD_EXISTING_CLUSTER"
  cloud_autonomous_vm_cluster_id = "<standby-cloud-autonomous-vm-cluster-ocid>"
  protection_mode                = "MAXIMUM_PERFORMANCE"
}
```

Autonomous INSERTDG does not take admin credentials or compute/storage sizing. For shard Data Guard replicas, use the primary-database selection in the DDB type README. Raft shard members use Raft replication; INSERTDG applies to the catalog for Raft DDBs.

Catalog INSERTDG runs against an already configured DDB and does not require another Configure Sharding step.

## REMOVEDG

Use REMOVEDG to remove a named Data Guard replica. Select its primary database and identify the replica:

```hcl
remove_dg = [{
  selection    = "catalogDetails[?originalReplica.database.name == 'catalog-database-a'].originalReplica.database"
  replica_name = "standby2"
}]
```

Replace the names with the current service-returned values.

## Additional INSERT resources

After an INSERT completes, keep that resource and its inputs unchanged. For another insertion:

1. Add a separate INSERT resource in `actions.tf`, for example `insert_2`.
2. Add its input list to the typed `actions` variable in `variables.tf` and to `terraform.tfvars`. Wire the new resource to that list.
3. Add a corresponding Configure Sharding resource that runs after the new INSERT succeeds.
4. Update the dependencies and readback to wait for the new operation. Review the plan to confirm completed operations will not run again.

## Limits and Configure Sharding

Each action resource supports up to **three operations per request**. Empty lists skip that action. Use separate resources and dependencies when operations need to run in sequence.

| Operation | Configure Sharding in these examples |
|---|---|
| INSERT | Runs after a successful INSERT. |
| Catalog INSERTDG | Not required for an already configured DDB. |
| REMOVE / REMOVEDG | No automatic Configure Sharding step. |

## Readback and recovery

After each operation, inspect the work request and refreshed topology:

```shell
terraform output action_work_request_ids
terraform output -json observed_topology
terraform output database_operational_state
terraform plan -var-file=terraform.tfvars
```

`action_work_request_ids` covers action resources. For an initial CREATE failure, use the work-request OCID reported in the Terraform error or provider logs.

Confirm the expected members are present and configured. `ACTIVE` describes the DDB lifecycle; `AVAILABLE` or `STOPPED` describes its operational state. Check each member’s status to confirm it is configured.

Keep the state and diagnostics for failed operations. Check an uncertain request's outcome before retrying, and retain successful action inputs when resuming.

## Start and stop

Omit `operational_state` during creation. For an existing DDB, set it to `STOPPED` to stop or `AVAILABLE` to start, then plan and apply the change. Verify Stop succeeds before starting the DDB. Starting it does not repair failed members.

## Update backup configuration

Templates include `db_backup_config` with a **30-day** recovery window during creation. To update an existing DDB, edit the same block in `terraform.tfvars`, review the in-place plan, and apply using the existing state. No action entry, backup trigger increment, or Configure Sharding is needed.
