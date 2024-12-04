---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_pluggable_database_snapshot"
sidebar_current: "docs-oci-resource-database-pluggable_database_snapshot"
description: |-
  Provides the Pluggable Database Snapshot resource in Oracle Cloud Infrastructure Database service
---

# oci_database_pluggable_database_snapshot
This resource provides the Pluggable Database Snapshot resource in Oracle Cloud Infrastructure Database service.

Creates a Pluggable Database Snapshot


## Example Usage

```hcl
resource "oci_database_pluggable_database_snapshot" "test_pluggable_database_snapshot" {
	#Required
	name = var.pluggable_database_snapshot_name
	pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id

	#Optional
	defined_tags = var.pluggable_database_snapshot_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The user-friendly name for the Database Snapshot. The name should be unique.
* `pluggable_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Pluggable Database.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pluggable Database Snapshot
	* `update` - (Defaults to 20 minutes), when updating the Pluggable Database Snapshot
	* `delete` - (Defaults to 20 minutes), when destroying the Pluggable Database Snapshot


## Import

PluggableDatabaseSnapshots can be imported using the `id`, e.g.

```
$ terraform import oci_database_pluggable_database_snapshot.test_pluggable_database_snapshot "id"
```

