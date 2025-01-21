---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_snapshot_management"
sidebar_current: "docs-oci-resource-database-db_node_snapshot_management"
description: |-
  Provides the Db Node Snapshot Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_db_node_snapshot_management
This resource provides the Db Node Snapshot Management resource in Oracle Cloud Infrastructure Database service.

Create Exadata Database Node Snapshots in the Exadb VM cluster.


## Example Usage

```hcl
resource "oci_database_db_node_snapshot_management" "test_db_node_snapshot_management" {
	#Required
	exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
	source_dbnode_ids = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.node_resource[*].node_id
	name = var.db_node_snapshot_suffix

	#Optional
	defined_tags = var.db_node_snapshot_defined_tags
	freeform_tags = var.db_node_snapshot_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) Defined tags for the Exadata Database Node Snapshots. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `exadb_vm_cluster_id` - (Required) The Exadata VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on Exascale Infrastructure.
* `freeform_tags` - (Optional) Free-form tags for the Exadata Database Node Snapshots. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The suffix of the Exadata Database Node Snapshot names (Snpashot name = Node hostname + "-" + suffix). The Exadata Database Node Snapshot name should be unique.
* `source_dbnode_ids` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Nodes for which snapshots will be created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `name` - (Required) The suffix of the Exadata Database Node Snapshot names (Snpashot name = Node hostname + "-" + suffix). The Exadata Database Node Snapshot name should be unique.
* `snapshots` - The list of created Exadata Database Node Snapshots. 
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
* `create` - (Defaults to 20 minutes), when creating the Db Node Snapshot Management


## Import

Import is not supported for this resource.
