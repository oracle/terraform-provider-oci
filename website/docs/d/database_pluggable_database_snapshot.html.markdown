---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_pluggable_database_snapshot"
sidebar_current: "docs-oci-datasource-database-pluggable_database_snapshot"
description: |-
  Provides details about a specific Pluggable Database Snapshot in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_pluggable_database_snapshot
This data source provides details about a specific Pluggable Database Snapshot resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Exadata Pluggable Database Snapshot in the specified compartment.


## Example Usage

```hcl
data "oci_database_pluggable_database_snapshot" "test_pluggable_database_snapshot" {
	#Required
	pluggable_database_snapshot_id = oci_database_pluggable_database_snapshot.test_pluggable_database_snapshot.id
}
```

## Argument Reference

The following arguments are supported:

* `pluggable_database_snapshot_id` - (Required) The Exadata Pluggable Database Snapshot [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

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

