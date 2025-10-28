---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_resource_pool_members"
sidebar_current: "docs-oci-datasource-database-autonomous_database_resource_pool_members"
description: |-
  Provides the list of Autonomous Database Resource Pool Members in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_resource_pool_members
This data source provides the list of Autonomous Database Resource Pool Members in Oracle Cloud Infrastructure Database service.

Lists the OCIDs of the Autonomous AI Database resource pool members for the specified Autonomous AI Database leader.


## Example Usage

```hcl
data "oci_database_autonomous_database_resource_pool_members" "test_autonomous_database_resource_pool_members" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `resource_pool_member_collection` - The list of resource_pool_member_collection.

### AutonomousDatabaseResourcePoolMember Reference

The following attributes are exported:

* `items` - List of resource pool member summary.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous AI Database.

