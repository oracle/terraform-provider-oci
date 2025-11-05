---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_snapshots"
sidebar_current: "docs-oci-datasource-database-db_node_snapshots"
description: |-
  Provides the list of Db Node Snapshots in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_node_snapshots
This data source provides the list of Db Node Snapshots in Oracle Cloud Infrastructure Database service.

Gets a list of the Exadata Database Node Snapshots in the specified compartment.


## Example Usage

```hcl
data "oci_database_db_node_snapshots" "test_db_node_snapshots" {
	#Required
	compartment_id = var.compartment_id
  
	#Optional
	cluster_id       = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.exadb_vm_cluster_id
	name             = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].name
	source_dbnode_id = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].source_dbnode_id
	state            = oci_database_db_node_snapshot_management.test_db_node_snapshot_management.snapshots[0].state

	## Example: filter db_node_snapshots by name
	#filter {
	#  name  = "name"
	#  regex = true
	#  values = ["^\\w+-${oci_database_db_node_snapshot_management.test_db_node_snapshot_management.name}$"]
	#}
	
	## Example: Get all but Terminated db_node_snapshots
	#filter {
	#  name = "state"
	#  values = ["CREATING", "AVAILABLE", "FAILED", "MOUNTED", "MOUNTING", "UNMOUNTING"]
	#}
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) A filter to return only Exadata Database Node Snapshots that match the given VM cluster.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - (Optional) A filter to return only resources that match the entire name given. The match is not case sensitive.
* `source_dbnode_id` - (Optional) A filter to return only Exadata Database Snapshots that match the given database node.
* `state` - (Optional) A filter to return only Exadata Database Snapshots that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `dbnode_snapshots` - The list of dbnode_snapshots.

### DbNodeSnapshot Reference

The following attributes are exported:

* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `dbnode_snapshot_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Node Snapshot.
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

