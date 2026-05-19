---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_snapshots"
sidebar_current: "docs-oci-datasource-database-database_snapshots"
description: |-
  Provides the list of Database Snapshots in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_database_snapshots
This data source provides the list of Database Snapshots in Oracle Cloud Infrastructure Database service.

Gets a list of the Exadata Database Snapshots in the specified compartment.


## Example Usage

```hcl
data "oci_database_database_snapshots" "test_database_snapshots" {

	#Optional
	cluster_id = oci_containerengine_cluster.test_cluster.id
	compartment_id = var.compartment_id
	database_id = oci_database_database.test_database.id
	name = var.database_snapshot_name
	state = var.database_snapshot_state
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) A filter to return only Exadata Database Snapshots that match the given VM cluster.
* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `database_id` - (Optional) A filter to return only Exadata Database Snapshots that match the given Database Id.
* `name` - (Optional) A filter to return only resources that match the entire name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only Exadata Database Snapshots that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `database_snapshot_collection` - The list of database_snapshot_collection.

### DatabaseSnapshot Reference

The following attributes are exported:

* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Snapshot.
* `lifecycle_details` - Additional information about the current lifecycle state of the Exadata Database Snapshot.
* `name` - The user-friendly name for the Database Snapshot. The name should be unique.
* `state` - The current state of the Exadata Database Snapshot.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time that the Exadata Database Snapshot was created, as expressed in RFC 3339 format. For example: 2023-06-27T21:10:29Z 

