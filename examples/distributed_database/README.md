# Terraform for Oracle Globally Distributed AI Database

Use these Terraform examples with Oracle Globally Distributed Autonomous AI Database (Autonomous) or Oracle Globally Distributed Exadata Database on Exascale Infrastructure (ExaScale). Choose a DDB type, supply the catalog and shard inputs, and run Terraform to create the distributed database (DDB) and configure sharding.

## Prerequisites

- Terraform **1.16.1** and an OCI provider build that supports the DDB resources and actions used here.
- Configure your OCI tenancy and required permissions using the setup guide for [Autonomous](https://docs.oracle.com/en/cloud/paas/globally-distributed-autonomous-database/user/set-environment.html) or [ExaScale](https://docs.oracle.com/en/cloud/paas/globally-distributed-exascale-database/user/set-environment.html).
- A target region, compartment OCID, and OCI authentication configured for your execution environment.
- Existing VM cluster OCIDs for the catalog, shards, and any standbys, with enough capacity for the requested databases.
- Admin password secret OCIDs for each database and the vault/key OCIDs required by the DDB type.
- A **unique three-character prefix**, starting with a lowercase letter and followed by lowercase letters or digits, for example `a01`.
- Listener and ONS ports in the **1024–65535** range. Select unused ports for the DDB.
- For Autonomous Configure Sharding, a certificate OCID and CA bundle OCID.

### Environment configuration

Set `auth` to `APIKey`, `SecurityToken`, `InstancePrincipal`, or `ResourcePrincipal`, and supply `region` and `compartment_id`. For API key or security token authentication, select your OCI profile with `config_file_profile`. Instance principal uses the OCI compute instance identity; resource principal uses the supported OCI resource identity. The templates default to `SecurityToken` and profile `DEFAULT`. See [OCI provider authentication](https://docs.oracle.com/en-us/iaas/Content/dev/terraform/configuring.htm#authentication) for setup.

### Infrastructure integration

For Autonomous databases, supply `cloud_autonomous_vm_cluster_id`. For ExaScale databases, supply `vm_cluster_id`, or `db_home_id` when using an existing database home. Enter the relevant OCIDs in each catalog, shard, and standby database block.

### Private endpoints

Supply existing DDB private-endpoint OCIDs through `additional_private_endpoint_ids`. The examples also support creating a private endpoint; the selected DDB type's template contains those inputs.

### Database sources

For Autonomous, set `source` to `ADBD_EXISTING_CLUSTER`. For ExaScale, use `XS_EXISTING_CLUSTER` to deploy on an existing cluster. ExaScale also supports `XD_EXISTING_CLUSTER`, `EXISTING_DB_HOME`, `XS_NEW_CLUSTER`, and `XS_NEW_VAULT_AND_CLUSTER`.

### ExaScale SCAN listener port

`scan_listener_port` is optional and applies only to `XS_NEW_CLUSTER` and `XS_NEW_VAULT_AND_CLUSTER`. Its range is **1024–8999**, with an API default of **1521** when omitted. Omit it for `XS_EXISTING_CLUSTER`. The separate `listener_port` is required.

## Catalog and shard topology

The **shard catalog** stores the distributed database metadata. **Shards** store the application data. Define the catalog in `catalog_details` and the shards in the topology input for your DDB type.

System-managed examples use Raft clusters or Data Guard groups. Composite examples organize these groups into named shardspaces. User-defined examples define a primary database and optional standbys for each shardspace. Catalog Data Guard replicas are configured separately from shard replication.

## Supported DDB types

| DDB type | Autonomous | ExaScale |
|---|---|---|
| Composite with Data Guard | [Example](distributed_autonomous_database/composite_dg/README.md) | [Example](distributed_database/composite_dg/README.md) |
| Composite Raft | [Example](distributed_autonomous_database/composite_raft/README.md) | [Example](distributed_database/composite_raft/README.md) |
| System-managed with Data Guard | [Example](distributed_autonomous_database/system_dg/README.md) | [Example](distributed_database/system_dg/README.md) |
| System-managed Raft RW | [Example](distributed_autonomous_database/system_raft_rw/README.md) | [Example](distributed_database/system_raft_rw/README.md) |
| System-managed Raft / RW-RO | [Example](distributed_autonomous_database/system_raft/README.md) | [Example](distributed_database/systemraft_rwro/README.md) |
| User-defined data distribution | [Example](distributed_autonomous_database/user_managed/README.md) | [Example](distributed_database/user_managed/README.md) |

Each DDB type README describes its topology and includes an INSERT example. Use INSERT/REMOVE for shard changes and INSERTDG/REMOVEDG for Data Guard replicas; see [ACTIONS.md](ACTIONS.md) for execution steps.

## Quickstart

Open the selected DDB type directory and prepare your inputs:

```shell
cd examples/distributed_database/distributed_database/system_raft_rw
cp terraform.tfvars-template terraform.tfvars
```

Replace every placeholder. Inputs such as `prefix`, `catalog_details`, and the shard topology are top-level variables. Fill in the template with your environment values and save it as `terraform.tfvars`.

Leave action lists empty and omit `operational_state` for the initial create. Then run:

```shell
terraform init
terraform validate
terraform plan -var-file=terraform.tfvars -out=create.tfplan
terraform show create.tfplan
terraform apply create.tfplan
```

For a local provider development override, use `terraform get` instead of `terraform init`.

## Configure Sharding

Configure Sharding deploys the sharding configuration after the distributed database (DDB) is created. It is enabled by default through `configure_sharding_enabled`. Set `configure_sharding_is_rebalance_required` to control rebalancing.

INSERT actions have a paired Configure Sharding step. Catalog INSERTDG actions do not require another Configure Sharding step. After completion, check the work requests and `terraform output -json observed_topology` to confirm the members are configured. See [Oracle's deployment guidance](https://docs.oracle.com/en/cloud/paas/globally-distributed-autonomous-database/user/deploying-configuration-distributed-adb.html) for Configure Sharding concepts.

## Optional GDS control node

A GDS control node hosts the GDSCTL administration utility. Set `gds_control_nodes = []` to skip it, or supply a public SSH key and subnet OCID in the list to request a node.

## Start and stop

Set `operational_state = "STOPPED"` to stop the existing DDB, or `operational_state = "AVAILABLE"` to start it. Review and apply each change separately, using the same Terraform state and input file. Wait for Stop to succeed before starting the database.

## Backup and recovery

All Autonomous and ExaScale flavour templates include `db_backup_config`. Supply the respective block in `terraform.tfvars` **before the initial plan/apply** to include backup settings during creation. The templates use Terraform HCL syntax:

### Autonomous

```hcl
db_backup_config = {
  recovery_window_in_days = 30
  backup_destination_details = [{
    type      = "OBJECT_STORE"
    is_remote = false
  }]
}
```

### ExaScale

```hcl
db_backup_config = {
  is_auto_backup_enabled        = true
  auto_backup_window            = "SLOT_FIVE"
  auto_full_backup_day          = "SUNDAY"
  auto_full_backup_window       = "SLOT_FIVE"
  backup_deletion_policy        = "DELETE_IMMEDIATELY"
  recovery_window_in_days       = 30
  can_run_immediate_full_backup = true
  backup_destination_details = [{
    type = "OBJECT_STORE"
  }]
}
```

For `terraform.tfvars.json`, place the equivalent JSON object under the top-level `"db_backup_config"` key. The standalone nightly backup payloads contain only the inner object because their runners add that key.

To test a post-create backup configuration change, edit the same block, review the plan, and apply it using the existing Terraform state. For example, change the recovery window from the default 30 days to another supported value. No `change_db_backup_config_trigger` increment is needed. Review retention, backup scheduling and deletion policy values for your test environment before applying.

If an operation fails, check its work request and the current database state before retrying. Retain the Terraform state and successful resources, and avoid replaying completed actions. See [recovery from partial creation](ACTIONS.md#recovery-from-partial-creation) for the recovery procedure.

## Teardown

From the same DDB type directory, use the existing state and input file:

```shell
terraform plan -destroy -var-file=terraform.tfvars -out=destroy.tfplan
terraform show destroy.tfplan
terraform apply destroy.tfplan
```

Review the resources to be deleted before applying. Existing clusters and endpoints supplied only as OCIDs are not deleted by this example.
