---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_snapshot"
sidebar_current: "docs-oci-datasource-database-database_snapshot"
description: |-
  Provides details about a specific Database Snapshot in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_database_snapshot
This data source provides details about a specific Database Snapshot resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Exadata Database Snapshot.


## Example Usage

```hcl
data "oci_database_database_snapshot" "test_database_snapshot" {
	#Required
	database_snapshot_id = oci_database_database_snapshot.test_database_snapshot.id
}
```

## Argument Reference

The following arguments are supported:

* `database_snapshot_id` - (Required) The Exadata Database Snapshot [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

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

