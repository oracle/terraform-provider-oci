---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_job"
sidebar_current: "docs-oci-resource-database_migration-job"
description: |-
Provides the Job resource in Oracle Cloud Infrastructure Database Migration service
---

# oci_database_migration_job
This resource provides the Job resource in Oracle Cloud Infrastructure Database Migration service.

Update Migration Job resource details.

Note: If you wish to use the DMS deprecated API version /20210929 it is necessary to pin the Terraform Provider version to v5.47.0. Newer Terraform provider versions will not support the DMS deprecated API version /20210929

## Example Usage

```hcl
resource "oci_database_migration_job" "test_job" {
	#Required
	job_id = oci_database_migration_job.test_job.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.job_display_name
	freeform_tags = var.job_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - (Optional) (Updatable) Name of the job.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"}
* `job_id` - (Required) The OCID of the job


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - Name of the job.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"}
* `id` - The OCID of the Migration Job.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `migration_id` - The OCID of the Migration that this job belongs to.
* `progress` - Progress details of a Migration Job.
	* `current_phase` - Current phase of the job.
	* `current_status` - Current status of the job.
	* `phases` - List of phase status for the job.
		* `action` - The text describing the action required to fix the issue
		* `duration_in_ms` - Duration of the phase in milliseconds
		* `extract` - Summary of phase status results.
			* `message` - Message in entry.
			* `type` - Type of extract.
		* `is_advisor_report_available` - True if a Pre-Migration Advisor report is available for this phase. False or null if no report is available.
		* `issue` - The text describing the root cause of the reported issue
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Job
* `update` - (Defaults to 20 minutes), when updating the Job
* `delete` - (Defaults to 20 minutes), when destroying the Job


## Import

Jobs can be imported using the `id`, e.g.

```
$ terraform import oci_database_migration_job.test_job "id"
```
