---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_scheduler_job_counts"
sidebar_current: "docs-oci-datasource-fleet_apps_management-scheduler_job_counts"
description: |-
  Provides the list of Scheduler Job Counts in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_scheduler_job_counts
This data source provides the list of Scheduler Job Counts in Oracle Cloud Infrastructure Fleet Apps Management service.

Retrieve aggregated summary information of Scheduler Jobs within a Tenancy.


## Example Usage

```hcl
data "oci_fleet_apps_management_scheduler_job_counts" "test_scheduler_job_counts" {

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `scheduler_job_aggregation_collection` - The list of scheduler_job_aggregation_collection.

### SchedulerJobCount Reference

The following attributes are exported:

* `items` - List of SchedulerJobAggregation objects.
	* `dimensions` - Aggregated summary information for a SchedulerJob.
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `scheduler_job_count_count` - Count of jobs in a Tenancy.

