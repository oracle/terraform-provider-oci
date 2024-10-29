---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_data_ingestion_job_log_content"
sidebar_current: "docs-oci-datasource-generative_ai_agent-data_ingestion_job_log_content"
description: |-
  Provides details about a specific Data Ingestion Job Log Content in Oracle Cloud Infrastructure Generative Ai Agent service
---

# Data Source: oci_generative_ai_agent_data_ingestion_job_log_content
This data source provides details about a specific Data Ingestion Job Log Content resource in Oracle Cloud Infrastructure Generative Ai Agent service.

**GetDataIngestionJobLogContent**

Returns the raw log file for the specified data ingestion job in text format.


## Example Usage

```hcl
data "oci_generative_ai_agent_data_ingestion_job_log_content" "test_data_ingestion_job_log_content" {
	#Required
	data_ingestion_job_id = oci_generative_ai_agent_data_ingestion_job.test_data_ingestion_job.id
}
```

## Argument Reference

The following arguments are supported:

* `data_ingestion_job_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data ingestion job.


## Attributes Reference

The following attributes are exported:


