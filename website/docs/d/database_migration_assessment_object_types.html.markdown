---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_assessment_object_types"
sidebar_current: "docs-oci-datasource-database_migration-assessment_object_types"
description: |-
  Provides the list of Assessment Object Types in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_assessment_object_types
This data source provides the list of Assessment Object Types in Oracle Cloud Infrastructure Database Migration service.

Display sample object types to exclude or include for an Assessment.


## Example Usage

```hcl
data "oci_database_migration_assessment_object_types" "test_assessment_object_types" {
	#Required
	connection_type = var.assessment_object_type_connection_type
}
```

## Argument Reference

The following arguments are supported:

* `connection_type` - (Required) The connection type for assessment objects.


## Attributes Reference

The following attributes are exported:

* `assessment_object_type_summary_collection` - The list of assessment_object_type_summary_collection.

### AssessmentObjectType Reference

The following attributes are exported:

* `items` - Items in collection.
	* `name` - Object type name

