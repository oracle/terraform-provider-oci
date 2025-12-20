---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_job_advisor_report_checks"
sidebar_current: "docs-oci-datasource-database_migration-job_advisor_report_checks"
description: |-
  Provides the list of Job Advisor Report Checks in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_job_advisor_report_checks
This data source provides the list of Job Advisor Report Checks in Oracle Cloud Infrastructure Database Migration service.

List of Pre-Migration checks from the advisor.


## Example Usage

```hcl
data "oci_database_migration_job_advisor_report_checks" "test_job_advisor_report_checks" {
	#Required
	job_id = oci_database_migration_job.test_job.id
}
```

## Argument Reference

The following arguments are supported:

* `job_id` - (Required) The OCID of the job 


## Attributes Reference

The following attributes are exported:

* `advisor_report_check_collection` - The list of advisor_report_check_collection.

### JobAdvisorReportCheck Reference

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

