---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_job_output"
sidebar_current: "docs-oci-datasource-database_migration-job_output"
description: |-
  Provides details about a specific Job Output in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_job_output
This data source provides details about a specific Job Output resource in Oracle Cloud Infrastructure Database Migration service.

List the Job Outputs


## Example Usage

```hcl
data "oci_database_migration_job_output" "test_job_output" {
	#Required
	job_id = oci_database_migration_job.test_job.id
}
```

## Argument Reference

The following arguments are supported:

* `job_id` - (Required) The OCID of the job 


## Attributes Reference

The following attributes are exported:

* `items` - Items in collection. 
	* `message` - Job output line. 

