---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migration_object_type"
sidebar_current: "docs-oci-datasource-database_migration-migration_object_type"
description: |-
  Provides details about a specific Migration Object Type in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_migration_object_type
This data source provides details about a specific Migration Object Type resource in Oracle Cloud Infrastructure Database Migration service.

Display sample object types to exclude or include for a Migration.


## Example Usage

```hcl
data "oci_database_migration_migration_object_type" "test_migration_object_type" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `items` - Items in collection. 
	* `name` - Object type name 

