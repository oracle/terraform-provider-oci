---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_job"
sidebar_current: "docs-oci-datasource-database_migration-job"
description: |-
  Provides details about a specific Job in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_job
This data source provides details about a specific Job resource in Oracle Cloud Infrastructure Database Migration service.

Get a migration job.

## Example Usage

```hcl
data "oci_database_migration_job" "test_job" {
	#Required
	job_id = oci_database_migration_job.test_job.id
}
```

## Argument Reference

The following arguments are supported:

* `job_id` - (Required) The OCID of the job 


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Name of the job. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Migration Job. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `migration_id` - The OCID of the Migration that this job belongs to. 
* `progress` - Progress details of a Migration Job. 
	* `current_phase` - Current phase of the job. 
	* `current_status` - Current status of the job. 
	* `phases` - List of phase status for the job. 
		* `duration_in_ms` - Duration of the phase in milliseconds 
		* `extract` - Summary of phase status results. 
			* `message` - Message in entry. 
			* `type` - Type of extract. 
		* `is_advisor_report_available` - True if a Pre-Migration Advisor report is available for this phase. False or null if no report is available. 
		* `log_location` - Details to access log file in the specified Object Storage bucket, if any. 
			* `bucket` - Name of the bucket containing the log file. 
			* `namespace` - Object Storage namespace. 
			* `object` - Log object name. 
		* `name` - Phase name 
		* `progress` - Percent progress of job phase. 
		* `status` - Phase status 
* `state` - The current state of the migration job. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Migration Job was created. An RFC3339 formatted datetime string 
* `time_updated` - The time the Migration Job was last updated. An RFC3339 formatted datetime string 
* `type` - The job type. 
* `unsupported_objects` - Database objects not supported. 
	* `object` - Name of the object (regular expression is allowed) 
	* `owner` - Owner of the object (regular expression is allowed) 
	* `type` - Type of unsupported object 

