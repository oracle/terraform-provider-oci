---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_pluggable_database_snapshots"
sidebar_current: "docs-oci-datasource-database-pluggable_database_snapshots"
description: |-
  Provides the list of Pluggable Database Snapshots in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_pluggable_database_snapshots
This data source provides the list of Pluggable Database Snapshots in Oracle Cloud Infrastructure Database service.

Gets a list of the Exadata Pluggable Database Snapshots in the specified compartment.


## Example Usage

```hcl
data "oci_database_pluggable_database_snapshots" "test_pluggable_database_snapshots" {

	#Optional
	cluster_id = oci_containerengine_cluster.test_cluster.id
	compartment_id = var.compartment_id
	name = var.pluggable_database_snapshot_name
	pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id
	state = var.pluggable_database_snapshot_state
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) A filter to return only Exadata Database Node Snapshots that match the given VM cluster.
* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - (Optional) A filter to return only resources that match the entire name given. The match is not case sensitive.
* `pluggable_database_id` - (Optional) A filter to return only Exadata Pluggable Database Snapshots that match the given database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only Exadata Pluggable Database Snapshots that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `pluggable_database_snapshots` - The list of pluggable_database_snapshots.

### PluggableDatabaseSnapshot Reference

The following attributes are exported:

* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Pluggable Database Snapshot.
* `lifecycle_details` - Additional information about the current lifecycle state of the Exadata Pluggable Database Snapshot.
* `name` - The user-friendly name for the Database Snapshot. The name should be unique.
* `pluggable_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Pluggable Database.
* `state` - The current state of the Exadata Pluggable Database Snapshot.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time that the Exadata Pluggable Database Snapshot was created, as expressed in RFC 3339 format. For example: 2023-06-27T21:10:29Z 

