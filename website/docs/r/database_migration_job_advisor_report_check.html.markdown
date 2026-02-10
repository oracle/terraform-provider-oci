---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_job_advisor_report_check"
sidebar_current: "docs-oci-resource-database_migration-job_advisor_report_check"
description: |-
  Provides the Job Advisor Report Check resource in Oracle Cloud Infrastructure Database Migration service
---

# oci_database_migration_job_advisor_report_check
This resource provides the Job Advisor Report Check resource in Oracle Cloud Infrastructure Database Migration service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-migration/latest/AdvisorReportCheck

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemigration

Update the premigration extended Advisor report check.


## Example Usage

```hcl
resource "oci_database_migration_job_advisor_report_check" "test_job_advisor_report_check" {
	#Required
	advisor_report_check_id = oci_database_migration_advisor_report_check.test_advisor_report_check.id
	is_reviewed = var.job_advisor_report_check_is_reviewed
	job_id = oci_database_migration_job.test_job.id
}
```

## Argument Reference

The following arguments are supported:

* `advisor_report_check_id` - (Required) The ID of the advisor check 
* `is_reviewed` - (Required) (Updatable) User flag for advisor report check. 
* `job_id` - (Required) The OCID of the job 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - Array of advisor check items. 
	* `action` - Fixing the issue. 
	* `columns` - Array of the column of the objects table. 
		* `display_name` - Display name of column. 
		* `key` - Id of column. 
	* `display_name` - Pre-Migration сheck display name. 
	* `fixup_script_location` - The path to the fixup script for this check. 
	* `impact` - Impact of the issue on data migration. 
	* `is_exclusion_allowed` - If false, objects cannot be excluded from migration. 
	* `is_reviewed` - User flag for advisor report check. 
	* `issue` - Description of the issue. 
	* `key` - Pre-Migration сheck id. 
	* `metadata` - Metadata of object. 
		* `object_name_column` - The field that stores the name of the object. 
		* `object_type_column` - The field that stores the type of the object. 
		* `object_type_fixed` - The field that stores the fixed type of the object. 
		* `schema_owner_column` - The field that stores the owner of the object. 
	* `object_count` - Number of database objects to migrate. 
	* `result_type` - Pre-Migration advisor result. 
* `summary` - Pre-Migration extended advisor report summary. 
	* `blocker_results_total_count` - Number of BLOCKER results in the extended advisor report. 
	* `fatal_results_total_count` - Number of FATAL results in the extended advisor report. 
	* `informational_results_total_count` - Number of INFORMATIONAL results in the extended advisor report. 
	* `pass_results_total_count` - Number of PASS results in the extended advisor report. 
	* `warning_results_total_count` - Number of WARNING results in the extended advisor report. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Job Advisor Report Check
	* `update` - (Defaults to 20 minutes), when updating the Job Advisor Report Check
	* `delete` - (Defaults to 20 minutes), when destroying the Job Advisor Report Check


## Import

JobAdvisorReportChecks can be imported using the `id`, e.g.

```
$ terraform import oci_database_migration_job_advisor_report_check.test_job_advisor_report_check "jobs/{jobId}/advisorReportChecks/{advisorReportCheckId}" 
```

