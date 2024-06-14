---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migration_object_types"
sidebar_current: "docs-oci-datasource-database_migration-migration_object_types"
description: |-
  Provides the list of Migration Object Types in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_migration_object_types
This data source provides the list of Migration Object Types in Oracle Cloud Infrastructure Database Migration service.

Display sample object types to exclude or include for a Migration.


## Example Usage

```hcl
data "oci_database_migration_migration_object_types" "test_migration_object_types" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `migration_object_type_summary_collection` - The list of migration_object_type_summary_collection.

### MigrationObjectType Reference

The following attributes are exported:

* `items` - Items in collection. 
	* `name` - Object type name 

