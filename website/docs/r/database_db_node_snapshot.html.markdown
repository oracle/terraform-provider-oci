---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_snapshot"
sidebar_current: "docs-oci-resource-database-db_node_snapshot"
description: |-
  Provides the Db Node Snapshot resource in Oracle Cloud Infrastructure Database service
---

# oci_database_db_node_snapshot
This resource provides the Db Node Snapshot resource in Oracle Cloud Infrastructure Database service.

Manage the specified Db Node Snapshot in the Exadb VM cluster.


## Example Usage

Unmount the specified Db Node Snapshot
```hcl
resource "oci_database_db_node_snapshot" "test_db_node_snapshot" {
    #Required
    dbnode_snapshot_id = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].id
    mount_dbnode_id    = "null"
}
```

Mount the specified Db Node Snapshot to the specified Db Node
```hcl
resource "oci_database_db_node_snapshot" "test_db_node_snapshot" {
    #Required
    dbnode_snapshot_id = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].id
    mount_dbnode_id    = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].source_dbnode_id
}
```

## Argument Reference

The following arguments are supported:

* `dbnode_snapshot_id` - (Required) The Exadata Database Node Snapshot [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `mount_dbnode_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Node to which the snapshot will be mounted. Set `mount_dbnode_id` to `"null"` (string, not `null`) to unmount the Db Node Snapshot.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Node Snapshot.
* `lifecycle_details` - Additional information about the current lifecycle state of the Exadata Database Node Snapshot.
* `mount_dbnode_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Node to which the snapshot is mounted. If the snapshot is not mounted to any node, then the value of `mount_dbnode_id` will be `"null"`.
* `mount_points` - Details of the mount points
	* `db_node_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Node where snapshot was mounted.
	* `name` - Mount Point Name
* `name` - The user-friendly name for the Database Node Snapshot. The name should be unique.
* `source_dbnode_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Node.
* `state` - The current state of the Exadata Database Node Snapshot.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `time_created` - The date and time that the Exadata Database Node Snapshot was created.
* `volumes` - Details of the volumes
	* `name` - Volume Name
	* `size` - Volume Size

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Db Node Snapshot
* `update` - (Defaults to 20 minutes), when updating the Db Node Snapshot
* `delete` - (Defaults to 20 minutes), when deleting the Db Node Snapshot	


## Import

DbNodeSnapshots can be imported using the `id`, e.g.

```
$ terraform import oci_database_db_node_snapshot.test_db_node_snapshot "id"
```
