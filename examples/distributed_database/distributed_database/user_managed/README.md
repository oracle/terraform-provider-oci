# ExaScale User-defined Sharding

This example creates an Oracle Globally Distributed Exadata Database on Exascale Infrastructure (ExaScale) with user-defined data distribution and optional Data Guard replicas for each shardspace. Complete the [prerequisites](../../README.md#prerequisites) before running it.

## Topology

The [template](terraform.tfvars-template) defines one shard catalog database, 2 primary shard databases, and 1 standby shard database:

```text
DDB
  Catalog
  shardspace1
    primary1 — 1 primary shard
    standby1 — 1 standby shard
  shardspace2
    primary2 — 1 primary shard
```

## Inputs

Copy the template to `terraform.tfvars` and supply the OCIDs for your environment.

| Input | What to supply |
|---|---|
| `catalog_details` | The catalog VM cluster, admin password secret, vault, and key OCIDs. |
| `user_shard_spaces` | Shardspace names, a primary database, and any standby databases. |
| `source` | `XS_EXISTING_CLUSTER`, `XS_NEW_CLUSTER`, `XS_NEW_VAULT_AND_CLUSTER`, `XD_EXISTING_CLUSTER`, or `EXISTING_DB_HOME`. The template uses `XS_EXISTING_CLUSTER`. |
| `vm_cluster_id` | The VM cluster OCID for `XS_EXISTING_CLUSTER` or `XD_EXISTING_CLUSTER`. |
| `db_home_id` | The database home OCID for `EXISTING_DB_HOME`. |
| `additional_private_endpoint_ids` | Existing DDB private-endpoint OCIDs. |
| `prefix` | A unique three-character prefix, for example `a01`. |

Supply `admin_password_secret_id`, `vault_id`, and `kms_key_id` in each catalog and shard database block. Leave `admin_password_secret_version_number` as `null` to use the latest secret version, or set the required version.

Choose unused listener and ONS ports in the **1024–65535** range. Omit `scan_listener_port` for `XS_EXISTING_CLUSTER`; see the [SCAN port guidance](../../README.md#exascale-scan-listener-port) for other sources.

## Create the DDB

Run from this directory:

```shell
cp terraform.tfvars-template terraform.tfvars
```

Replace the placeholders and set your prefix and ports. Leave the `actions` lists empty and omit `operational_state`. Set `gds_control_nodes = []` unless you want a GDS control node; to add one, supply its subnet OCID and public SSH key.

```shell
terraform init
terraform validate
terraform plan -var-file=terraform.tfvars -out=create.tfplan
terraform show create.tfplan
terraform apply create.tfplan
```

Review the plan before applying. Configure Sharding runs after the DDB is created and is enabled by default. Check the result with:

```shell
terraform output distributed_database_id
terraform output distributed_database_lifecycle_state
terraform output -json observed_topology
```

Confirm the catalog and all expected primary and standby databases are present and configured.

## INSERT example

To add a new shardspace with a primary shard, set `actions.insert` in the same `terraform.tfvars`. This example creates `shardspace3` without a standby; use a name that does not already exist.

```hcl
insert = [{
  selection = "userShardSpaces"
  value = <<-INSERT_JSON
  {
    "name": "shardspace3",
    "originalReplica": {
      "name": "primary3",
      "database": {
        "source": "XS_EXISTING_CLUSTER",
        "vmClusterId": "<new-shard-vm-cluster-ocid>",
        "adminPasswordSecretId": "<new-shard-admin-password-secret-ocid>",
        "vaultId": "<vault-ocid>",
        "kmsKeyId": "<key-ocid>"
      }
    },
    "dataGuardReplicas": []
  }
  INSERT_JSON
}]
```

Place this entry inside the existing `actions = { ... }` block. Each replica uses one `database` object. Use `dataGuardReplicas` when adding standby databases with the new shardspace. Keep the original creation inputs and any completed action inputs unchanged.

Use the [update commands](../../ACTIONS.md#invoke-an-update-separately) to plan and apply the insertion. Configure Sharding runs after INSERT succeeds. Then check `action_work_request_ids` and `observed_topology` to confirm the new databases are configured.

For another insertion after this action completes, follow [additional INSERT resources](../../ACTIONS.md#additional-insert-resources).

## Backup configuration

The template enables automatic backups to `OBJECT_STORE`, uses `SLOT_FIVE` for incremental and full backup windows, schedules full backups on `SUNDAY`, sets a **30-day** recovery window and `DELETE_IMMEDIATELY` deletion policy, and enables immediate full backup. Supply `db_backup_config` in `terraform.tfvars` before the initial plan/apply to include it during creation. To update it after creation, edit the same block and review/apply a separate plan using the existing state; no backup trigger is needed.

## Other operations

- **Remove a shardspace:** use `userShardSpaces[?name == 'shardspace1']` with [REMOVE](../../ACTIONS.md#remove).
- **Data Guard replicas:** use [INSERTDG](../../ACTIONS.md#insertdg) or [REMOVEDG](../../ACTIONS.md#removedg). Select the primary database whose standby you want to add or remove.
- **Start or stop:** set `operational_state` as described in [Start and stop](../../ACTIONS.md#start-and-stop).
- **Replace a failed shard:** follow the [recovery steps](../../ACTIONS.md#recovery-from-partial-creation).
- **Delete the DDB:** follow [Teardown](../../README.md#teardown) using the same state and input file.
