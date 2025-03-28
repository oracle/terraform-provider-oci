---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_data_ingestion_jobs"
sidebar_current: "docs-oci-datasource-generative_ai_agent-data_ingestion_jobs"
description: |-
  Provides the list of Data Ingestion Jobs in Oracle Cloud Infrastructure Generative Ai Agent service
---

# Data Source: oci_generative_ai_agent_data_ingestion_jobs
This data source provides the list of Data Ingestion Jobs in Oracle Cloud Infrastructure Generative Ai Agent service.

**ListDataIngestionJobs**

Gets a list of data ingestion jobs.


## Example Usage

```hcl
data "oci_generative_ai_agent_data_ingestion_jobs" "test_data_ingestion_jobs" {

	#Optional
	compartment_id = var.compartment_id
	data_source_id = oci_generative_ai_agent_data_source.test_data_source.id
	display_name = var.data_ingestion_job_display_name
	state = var.data_ingestion_job_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `data_source_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data source.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `data_ingestion_job_collection` - The list of data_ingestion_job_collection.

### DataIngestionJob Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `data_ingestion_job_statistics` - **DataIngestionJobStatistics**

	The statistics of data ingestion job. 
	* `duration_in_seconds` - The duration of this ingestion job.
	* `number_of_failed_files` - The number of files that have failed during the ingestion.
	* `number_of_ingested_files` - The number of files that have been successfully ingested during the ingestion.
* `data_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent DataSource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user-friendly name. Does not have to be unique, and it's changeable.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DataIngestionJob.
* `lifecycle_details` - A message that describes the current state of the data ingestion job in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of the data ingestion job.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the data ingestion job was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the data ingestion job was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

