---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_job_advisor_report"
sidebar_current: "docs-oci-datasource-database_migration-job_advisor_report"
description: |-
  Provides details about a specific Job Advisor Report in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_job_advisor_report
This data source provides details about a specific Job Advisor Report resource in Oracle Cloud Infrastructure Database Migration service.

Get the Pre-Migration Advisor report details


## Example Usage

```hcl
data "oci_database_migration_job_advisor_report" "test_job_advisor_report" {
	#Required
	job_id = oci_database_migration_job.test_job.id
}
```

## Argument Reference

The following arguments are supported:

* `job_id` - (Required) The OCID of the job 


## Attributes Reference

The following attributes are exported:

* `number_of_fatal` - Number of Fatal results in the advisor report. 
* `number_of_fatal_blockers` - Number of Fatal Blocker results in the advisor report. 
* `number_of_informational_results` - Number of Informational results in the advisor report. 
* `number_of_warnings` - Number of Warning results in the advisor report. 
* `report_location_details` - Details to access Pre-Migration Advisor report. 
	* `location_in_source` - Path in the Source Registered Connection where the Pre-Migration advisor report can be accessed. 
	* `object_storage_details` - Details to access Pre-Migration Advisor report in the specified Object Storage bucket, if any. 
		* `bucket` - Name of the bucket containing the Pre-Migration Advisor report. 
		* `namespace` - Object Storage namespace. 
		* `object` - Pre-Migration Advisor report object name. 
* `result` - Pre-Migration advisor result. 

