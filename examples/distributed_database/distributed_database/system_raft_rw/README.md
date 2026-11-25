# ExaScale System-managed Raft Read-write

This example creates an Oracle Globally Distributed Exadata Database on Exascale Infrastructure (ExaScale) with system-managed data distribution and a READ_WRITE Raft cluster. Complete the [prerequisites](../../README.md#prerequisites) before running it.

## Topology

The [template](terraform.tfvars-template) defines one shard catalog database and three shard databases:

```text
DDB
├── Catalog
└── rwcluster — READ_WRITE
    ├── Shard 1
    ├── Shard 2
    └── Shard 3
```

The cluster uses `replication_factor = 3` and `witness_count = 0`. The template sets `system_chunk_count = 120` and `system_raft_replication_unit_count = 6`.

## Inputs

Copy the template to `terraform.tfvars` and supply the OCIDs for your environment.

| Input | What to supply |
|---|---|
| `catalog_details` | The catalog VM cluster, admin password secret, vault, and key OCIDs. |
| `system_raft_clusters` | The cluster name, replication settings, and database details for each shard. |
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

Review the plan before applying.

Configure Sharding runs after the DDB is created and is enabled by default. Check the result with:

```shell
terraform output distributed_database_id
terraform output distributed_database_lifecycle_state
terraform output -json observed_topology
```

Confirm the catalog and all three shard databases are present and configured.

## INSERT example

To add a shard database to the existing cluster, set `actions.insert` in the same `terraform.tfvars`. Use the existing cluster name and replication settings, and supply the new database's OCIDs. The example below matches the template's `rwcluster`.

```hcl
insert = [{
  selection = "systemRaftClusters"
  value = <<-INSERT_JSON
  {
    "name": "rwcluster",
    "replicationFactor": 3,
    "witnessCount": 0,
    "ruMode": "READ_WRITE",
    "databases": [{
      "source": "XS_EXISTING_CLUSTER",
      "vmClusterId": "<new-shard-vm-cluster-ocid>",
      "adminPasswordSecretId": "<new-shard-admin-password-secret-ocid>",
      "vaultId": "<vault-ocid>",
      "kmsKeyId": "<key-ocid>"
    }]
  }
  INSERT_JSON
}]
```

Place this entry inside the existing `actions = { ... }` block. The `databases` list contains only the new shard. Keep the original creation inputs and any completed action inputs unchanged.

Use the [update commands](../../ACTIONS.md#invoke-an-update-separately) to plan and apply the insertion. Configure Sharding runs after INSERT succeeds. Then check `action_work_request_ids` and `observed_topology` to confirm the new shard is configured.

For another insertion after this action completes, follow [additional INSERT resources](../../ACTIONS.md#additional-insert-resources).

## Backup configuration

The template enables automatic backups to `OBJECT_STORE`, uses `SLOT_FIVE` for incremental and full backup windows, schedules full backups on `SUNDAY`, sets a **30-day** recovery window and `DELETE_IMMEDIATELY` deletion policy, and enables immediate full backup. Supply `db_backup_config` in `terraform.tfvars` before the initial plan/apply to include it during creation. To update it after creation, edit the same block and review/apply a separate plan using the existing state; no backup trigger is needed.

## Other operations

- **Remove a shard:** use `systemRaftClusters[?name == 'rwcluster'].databases[?name == '<database-name>']` with [REMOVE](../../ACTIONS.md#remove).
- **Catalog Data Guard:** use [INSERTDG](../../ACTIONS.md#insertdg) or [REMOVEDG](../../ACTIONS.md#removedg). Shard databases use Raft replication.
- **Start or stop:** set `operational_state` as described in [Start and stop](../../ACTIONS.md#start-and-stop).
- **Replace a failed shard:** follow the [recovery steps](../../ACTIONS.md#recovery-from-partial-creation) and insert its replacement into the same cluster.
- **Delete the DDB:** follow [Teardown](../../README.md#teardown) using the same state and input file.
