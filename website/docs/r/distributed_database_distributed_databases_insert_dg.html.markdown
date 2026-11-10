---
subcategory: "Distributed Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_distributed_database_distributed_databases_insert_dg"
sidebar_current: "docs-oci-resource-distributed_database-distributed_databases_insert_dg"
description: |-
  Executes INSERTDG operations in Oracle Cloud Infrastructure Distributed Database service.
---

# oci_distributed_database_distributed_databases_insert_dg

Executes INSERTDG operations against an existing distributed database. Configure one to three `insert_dg` blocks; all blocks are submitted together in one PATCH request. The provider generates the API operation fields automatically.

## Example Usage

```hcl
resource "oci_distributed_database_distributed_databases_insert_dg" "insert_dg" {
  distributed_database_id = oci_distributed_database_distributed_database.example.id

  insert_dg {
    selection    = "catalogDetails[?originalReplica.database.name == 'catalog-database-a'].originalReplica.database"
    replica_name = "standby"
    database {
      source                   = "XS_EXISTING_CLUSTER"
      vm_cluster_id            = var.standby_vm_cluster_id
      admin_password_secret_id = var.standby_password_secret_id
    }
  }
}
```

The example references an existing main database resource and input variables. Replace the topology and database names with names from your database.

## Argument Reference

* `distributed_database_id` - (Required) Target database OCID. Changing this value replaces the action resource and executes its configured operations against the new target.
* `insert_dg` - (Required, Updatable) One to three repeatable operation blocks. Changing this list executes all configured operations, including unchanged blocks within that instance.
    * `selection` - (Required, Updatable) API selection that must resolve to exactly one original database. For example, select a catalog original database as shown in the example, or use `systemDataGuardDatabases.originalReplica.databases[?name == 'database-a']` for a system Data Guard shard.
    * `replica_name` - (Required, Updatable) Name of the new Data Guard replica. Starts with a letter, contains only letters, digits, and underscores, and is at most 40 characters.
    * `database` - (Required, Updatable) Exactly one standby database placement block. Fields required by the selected source must be supplied.
        * `admin_password` - (Optional, Sensitive, Updatable) Admin password. Supply exactly one of this field or `admin_password_secret_id`.
        * `admin_password_secret_id` - (Optional, Sensitive, Updatable) Vault secret OCID for the admin password. Supply exactly one of this field or `admin_password`.
        * `admin_password_secret_version_number` - (Optional, Updatable) Secret version from 1 through 60; requires `admin_password_secret_id`. Omit to use the latest version.
        * `availability_domain` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The name of the availability domain that the distributed database catalog will be located in.
        * `db_home_id` - (Required for `EXISTING_DB_HOME`, Updatable) Database home OCID.
        * `db_storage_vault_details` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exascale db vault storage for shard or catalog of the distributed database.
            * `additional_flash_cache_in_percent` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The size of additional Flash Cache in percentage of High Capacity database storage.
            * `autoscale_limit_in_gbs` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
            * `compartment_id` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
            * `high_capacity_database_storage` - (Required when source=XS_NEW_VAULT_AND_CLUSTER) Total storage capacity in GB for vault storage.
            * `is_autoscale_enabled` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
            * `subscription_id` - (Applicable when source=XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
        * `exascale_db_storage_vault_id` - (Required when source=XS_NEW_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
        * `protection_mode` - (Optional, Updatable) Data Guard protection mode. Defaults to `MAXIMUM_PERFORMANCE`.
        * `source` - (Required, Updatable) Placement source. One of `XS_EXISTING_CLUSTER`, `XS_NEW_CLUSTER`, `XS_NEW_VAULT_AND_CLUSTER`, `XD_EXISTING_CLUSTER`, or `EXISTING_DB_HOME`.
        * `transport_type` - (Optional, Updatable) Data Guard transport type. Defaults to `ASYNC`.
        * `vm_cluster_details` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Details of the request to create exadb vm cluster for shard or catalog of the distributed database.
            * `backup_network_nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to.  Setting this to an empty array after the list is created removes the resource from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
            * `backup_subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet associated with the Exadata VM cluster on Exascale Infrastructure.
            * `compartment_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for VM Cluster.
            * `domain` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) A domain name used for the Exadata VM cluster on Exascale Infrastructure.  If the Oracle-provided internet and VCN resolver is enabled for the specified subnet, then the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name.  Hyphens (-) are not permitted. Applies to Exadata Database Service on Exascale Infrastructure only.
            * `enabled_ecpu_count` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of ECPUs to enable for an Exadata VM cluster on Exascale Infrastructure.
            * `is_diagnostics_events_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether diagnostic collection is enabled for the VM cluster.  Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues.  Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system.  You can enable diagnostic collection during VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` API.
            * `is_health_monitoring_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether health monitoring is enabled for the VM cluster.  Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel.  You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster` API.
            * `is_incident_logs_enabled` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) Indicates whether incident logs and trace collection are enabled for the VM cluster.  Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster` API.
            * `license_model` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The Oracle license model that applies to the Exadata VM cluster on Exascale Infrastructure. The default is BRING_YOUR_OWN_LICENSE.
            * `nsg_ids` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs.  Setting this to an empty list removes all resources from all NSGs.  For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
            * `private_zone_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The private zone ID in which you want DNS records to be created.
            * `ssh_public_keys` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The public key portion of one or more key pairs used for SSH access to the Exadata VM cluster on Exascale Infrastructure.
            * `subnet_id` - (Required when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the Exadata VM cluster on Exascale Infrastructure.
            * `subscription_id` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
            * `total_ecpu_count` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) The number of Total ECPUs for an Exadata VM cluster on Exascale Infrastructure.
            * `vm_file_system_storage_size` - (Applicable when source=XS_NEW_CLUSTER | XS_NEW_VAULT_AND_CLUSTER) File System Storage Size in GBs for Exadata VM cluster.
        * `vm_cluster_id` - (Required for `XS_EXISTING_CLUSTER` or `XD_EXISTING_CLUSTER`, Updatable) VM cluster OCID.

Each block generates `operation: MERGE` with a JSON-encoded `value` containing `subOperation: INSERTDG` and `dataGuardReplicas: [{name: replica_name, database: ...}]`. The provider maps Terraform field names to API field names.

For `XS_NEW_CLUSTER`, supply `availability_domain`, `exascale_db_storage_vault_id`, and `vm_cluster_details`. For `XS_NEW_VAULT_AND_CLUSTER`, supply `availability_domain`, `db_storage_vault_details`, and `vm_cluster_details`. Supply exactly one admin password or password secret for every source. Sensitive inputs remain in Terraform state; protect access to the state backend.

## Action Lifecycle

Creation executes the configured operations. An update executes them when the operation list changes, without an incremental trigger. Refresh reads the database state and does not execute operations. Unchanged action instances are not replayed when additional instances are added. Recreating an action instance, including after its state is removed, executes its configured operations again.

Resource labels are user-defined. For example, name two instances `insert_dg` and `insert_dg_2`; each submits its own operations. Keep successful instances unchanged and add another instance for additional operations. Changing an existing instance submits its full operation list again. Use `depends_on` to order instances targeting the same database when they must run sequentially.

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
