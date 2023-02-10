---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflow_job_fact"
sidebar_current: "docs-oci-datasource-media_services-media_workflow_job_fact"
description: |-
  Provides details about a specific Media Workflow Job Fact in Oracle Cloud Infrastructure Media Services service
---
> **_NOTE:_** This data source has been deprecated and is no longer supported.
# Data Source: oci_media_services_media_workflow_job_fact
This data source provides details about a specific Media Workflow Job Fact resource in Oracle Cloud Infrastructure Media Services service.

Get the MediaWorkflowJobFact identified by the mediaWorkflowJobId and Fact ID.

## Example Usage

```hcl
data "oci_media_services_media_workflow_job_fact" "test_media_workflow_job_fact" {
	#Required
	key = var.media_workflow_job_fact_key
	media_workflow_job_id = oci_media_services_media_workflow_job.test_media_workflow_job.id
}
```

## Argument Reference

The following arguments are supported:

* `key` - (Required) Identifier of the MediaWorkflowJobFact within a MediaWorkflowJob.
* `media_workflow_job_id` - (Required) Unique MediaWorkflowJob identifier.


## Attributes Reference

The following attributes are exported:

* `detail` - The body of the detail captured as JSON.
* `key` - System generated serial number to uniquely identify a detail in order within a MediaWorkflowJob.
* `media_workflow_job_id` - Reference to the parent job.
* `name` - Unique name. It is read-only and generated for the fact.
* `type` - The type of information contained in this detail.

