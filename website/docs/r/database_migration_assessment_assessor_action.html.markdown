---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_assessment_assessor_action"
sidebar_current: "docs-oci-resource-database_migration-assessment_assessor_action"
description: |-
  Provides the Assessment Assessor Action resource in Oracle Cloud Infrastructure Database Migration service
---

# oci_database_migration_assessment_assessor_action
This resource provides the Assessment Assessor Action resource in Oracle Cloud Infrastructure Database Migration service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-migration/latest/AssessorAction

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemigration

Assessor Action.

## Example Usage

```hcl
resource "oci_database_migration_assessment_assessor_action" "test_assessment_assessor_action" {
	#Required
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessor_action = var.assessment_assessor_action_assessor_action
	assessor_name = var.assessment_assessor_action_assessor_name
	items {
		#Required
		name = var.assessment_assessor_action_items_name
		value = var.assessment_assessor_action_items_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `assessment_id` - (Required) The OCID of the Assessment 
* `assessor_action` - (Required) The Accessor Action
* `assessor_name` - (Required) The name of the Assessor
* `items` - (Required) Array of name-value details for assessor action.
	* `name` - (Required) The property name.
	* `value` - (Required) The property value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Assessment Assessor Action
	* `update` - (Defaults to 20 minutes), when updating the Assessment Assessor Action
	* `delete` - (Defaults to 20 minutes), when destroying the Assessment Assessor Action


## Import

AssessmentAssessorActions can be imported using the `id`, e.g.

```
$ terraform import oci_database_migration_assessment_assessor_action.test_assessment_assessor_action "id"
```

