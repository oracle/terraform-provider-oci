---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflow_job_facts"
sidebar_current: "docs-oci-datasource-media_services-media_workflow_job_facts"
description: |-
  Provides the list of Media Workflow Job Facts in Oracle Cloud Infrastructure Media Services service
---
> **_NOTE:_** This data source has been deprecated and is no longer supported.
# Data Source: oci_media_services_media_workflow_job_facts
This data source provides the list of Media Workflow Job Facts in Oracle Cloud Infrastructure Media Services service.

Internal API to get a point-in-time snapshot of a MediaWorkflowJob.

## Example Usage

```hcl
data "oci_media_services_media_workflow_job_facts" "test_media_workflow_job_facts" {
	#Required
	media_workflow_job_id = oci_media_services_media_workflow_job.test_media_workflow_job.id

	#Optional
	key = var.media_workflow_job_fact_key
	type = var.media_workflow_job_fact_type
}
```

## Argument Reference

The following arguments are supported:

* `key` - (Optional) Filter by MediaWorkflowJob ID and MediaWorkflowJobFact key. 
* `media_workflow_job_id` - (Required) Unique MediaWorkflowJob identifier.
* `type` - (Optional) Types of details to include.


## Attributes Reference

The following attributes are exported:

* `media_workflow_job_fact_collection` - The list of media_workflow_job_fact_collection.

### MediaWorkflowJobFact Reference

The following attributes are exported:

* `detail` - The body of the detail captured as JSON.
* `key` - System generated serial number to uniquely identify a detail in order within a MediaWorkflowJob.
* `media_workflow_job_id` - Reference to the parent job.
* `name` - Unique name. It is read-only and generated for the fact.
* `type` - The type of information contained in this detail.

