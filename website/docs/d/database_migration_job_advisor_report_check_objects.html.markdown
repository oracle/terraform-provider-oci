---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_job_advisor_report_check_objects"
sidebar_current: "docs-oci-datasource-database_migration-job_advisor_report_check_objects"
description: |-
  Provides the list of Job Advisor Report Check Objects in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_job_advisor_report_check_objects
This data source provides the list of Job Advisor Report Check Objects in Oracle Cloud Infrastructure Database Migration service.

Get the Pre-Migration extended Advisor report object list.


## Example Usage

```hcl
data "oci_database_migration_job_advisor_report_check_objects" "test_job_advisor_report_check_objects" {
	#Required
	advisor_report_check_id = oci_database_migration_advisor_report_check.test_advisor_report_check.id
	job_id = oci_database_migration_job.test_job.id
}
```

## Argument Reference

The following arguments are supported:

* `advisor_report_check_id` - (Required) The ID of the advisor check 
* `job_id` - (Required) The OCID of the job 


## Attributes Reference

The following attributes are exported:

* `advisor_report_check_objects_collection` - The list of advisor_report_check_objects_collection.

### JobAdvisorReportCheckObject Reference

The following attributes are exported:

* `items` - Array of check objects. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `fields` - 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"} 
	* `is_excluded` - If the object was excluded from migration, then it is true. 
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

