---
subcategory: "Gdp"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_gdp_gdp_pipeline"
sidebar_current: "docs-oci-datasource-gdp-gdp_pipeline"
description: |-
  Provides details about a specific Gdp Pipeline in Oracle Cloud Infrastructure Gdp service
---

# Data Source: oci_gdp_gdp_pipeline
This data source provides details about a specific Gdp Pipeline resource in Oracle Cloud Infrastructure Gdp service.

Retrieves a pipeline by identifier.

## Example Usage

```hcl
data "oci_gdp_gdp_pipeline" "test_gdp_pipeline" {
	#Required
	gdp_pipeline_id = oci_gdp_gdp_pipeline.test_gdp_pipeline.id
	#Optional
	env = var.gdp_env
}
```

## Argument Reference

The following arguments are supported:

* `gdp_pipeline_id` - (Required) Unique pipeline identifier.
* `env` - (Optional) The environment where the pipeline resides. Valid values are COMMERCIAL or USGOV. Defaults to COMMERCIAL.

## Attributes Reference

The following attributes are exported:

* `approval_key_vault_id` - The KMS vault OCID used for cryptographic approvals of transfers.
* `authorization_details` - Authorization information about the pipeline being configured.
* `bucket_details` - Configuration information about the buckets used for this pipeline.
	* `bucket_type` - Type of bucket. SENDER pipelines can be SOURCE, TRANSFER, REJECT, or FAILED. RECEIVER pipelines have a DESTINATION bucket.
	* `id` - OCID of the bucket.
	* `name` - Name of the bucket.
	* `namespace` - Namespace of the bucket.
* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Short field input by customer for a description of the data pipeline use-case.
* `display_name` - Pipeline short name.
* `file_types` - List of file types allowed to be transferred in the pipeline according to the authorization details (e.g. .pdf, .xml, .doc).
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the pipeline.
* `is_approval_needed` - Determines whether file transfers need to go through an approval workflow.
* `is_chunking_enabled` - Determines whether file must be chunked during the transfer. This is only a property of SENDER pipelines.
* `is_file_override_in_destination_enabled` - Enable file override feature in destination bucket. If 2 files with same name exist in destination bucket, original file will be overwritten.
* `is_scanning_enabled` - Determines whether GDP Scanning should be enabled for the pipeline.
* `lifecycle_details` - Additional details about the current state of the pipeline.
* `peered_gdp_pipeline_id` - OCID of the peered pipeline. This null for SENDER pipeline.
* `peering_region` - Public region name where the peered pipeline exists.
* `pipeline_type` - Type of pipeline. Can be SENDER or RECEIVER.
* `service_log_group_id` - the OCID of the service log group.
* `state` - The current state of the pipeline.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the pipeline was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the pipeline was updated. An RFC3339 formatted datetime string.

