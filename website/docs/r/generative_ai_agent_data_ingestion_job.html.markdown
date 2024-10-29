---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_data_ingestion_job"
sidebar_current: "docs-oci-resource-generative_ai_agent-data_ingestion_job"
description: |-
  Provides the Data Ingestion Job resource in Oracle Cloud Infrastructure Generative Ai Agent service
---

# oci_generative_ai_agent_data_ingestion_job
This resource provides the Data Ingestion Job resource in Oracle Cloud Infrastructure Generative Ai Agent service.

**CreateDataIngestionJob**

Creates a data ingestion job.


## Example Usage

```hcl
resource "oci_generative_ai_agent_data_ingestion_job" "test_data_ingestion_job" {
	#Required
	compartment_id = var.compartment_id
	data_source_id = oci_generative_ai_agent_data_source.test_data_source.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.data_ingestion_job_description
	display_name = var.data_ingestion_job_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the data ingestion job in. 
* `data_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent DataSource.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) A user-friendly description of the data ingestion job.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Ingestion Job
	* `update` - (Defaults to 20 minutes), when updating the Data Ingestion Job
	* `delete` - (Defaults to 20 minutes), when destroying the Data Ingestion Job


## Import

DataIngestionJobs can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_agent_data_ingestion_job.test_data_ingestion_job "id"
```

