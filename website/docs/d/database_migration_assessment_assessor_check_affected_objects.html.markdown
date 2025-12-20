---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_assessment_assessor_check_affected_objects"
sidebar_current: "docs-oci-datasource-database_migration-assessment_assessor_check_affected_objects"
description: |-
  Provides the list of Assessment Assessor Check Affected Objects in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_assessment_assessor_check_affected_objects
This data source provides the list of Assessment Assessor Check Affected Objects in Oracle Cloud Infrastructure Database Migration service.

Display Check Affected objects.


## Example Usage

```hcl
data "oci_database_migration_assessment_assessor_check_affected_objects" "test_assessment_assessor_check_affected_objects" {
	#Required
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessor_name = var.assessment_assessor_check_affected_object_assessor_name
	check_name = var.assessment_assessor_check_affected_object_check_name
}
```

## Argument Reference

The following arguments are supported:

* `assessment_id` - (Required) The OCID of the Assessment 
* `assessor_name` - (Required) The name of the Assessor
* `check_name` - (Required) The Name of the assessor check


## Attributes Reference

The following attributes are exported:

* `affected_objects_collection` - The list of affected_objects_collection.

### AssessmentAssessorCheckAffectedObject Reference

The following attributes are exported:

* `items` - Items in collection.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `fields` - 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"} 
	* `is_excluded` - If the object was excluded from migration, then it is true. 
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

