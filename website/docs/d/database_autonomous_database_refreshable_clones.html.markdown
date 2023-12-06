---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_refreshable_clones"
sidebar_current: "docs-oci-datasource-database-autonomous_database_refreshable_clones"
description: |-
  Provides the list of Autonomous Database Refreshable Clones in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_refreshable_clones
This data source provides the list of Autonomous Database Refreshable Clones in Oracle Cloud Infrastructure Database service.

Lists the OCIDs of the Autonomous Database local and connected remote refreshable clones with the region where they exist for the specified source database.


## Example Usage

```hcl
data "oci_database_autonomous_database_refreshable_clones" "test_autonomous_database_refreshable_clones" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `refreshable_clone_collection` - The list of refreshable_clone_collection.

### AutonomousDatabaseRefreshableClone Reference

The following attributes are exported:

* `items` - 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	* `region` - The name of the region where the refreshable clone exists.

