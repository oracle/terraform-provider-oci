---
subcategory: "Gdp"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_gdp_gdp_pipeline"
sidebar_current: "docs-oci-resource-gdp-gdp_pipeline"
description: |-
  Provides the Gdp Pipeline resource in Oracle Cloud Infrastructure Gdp service
---

# oci_gdp_gdp_pipeline
This resource provides the Gdp Pipeline resource in Oracle Cloud Infrastructure Gdp service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/gdp

Creates a new pipeline.


## Example Usage

```hcl
resource "oci_gdp_gdp_pipeline" "test_gdp_pipeline" {
	#Required
	bucket_details {
		#Required
		bucket_type = var.gdp_pipeline_bucket_details_bucket_type
		id = var.gdp_pipeline_bucket_details_id
		name = var.gdp_pipeline_bucket_details_name
		namespace = var.gdp_pipeline_bucket_details_namespace
	}
	compartment_id = var.compartment_id
	display_name = var.gdp_pipeline_display_name
	peering_region = var.gdp_pipeline_peering_region
	pipeline_type = var.gdp_pipeline_pipeline_type

	#Optional
	approval_key_vault_id = oci_kms_vault.test_vault.id
	authorization_details = var.gdp_pipeline_authorization_details
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.gdp_pipeline_description
	env = var.gdp_env
	file_types = var.gdp_pipeline_file_types
	freeform_tags = {"bar-key"= "value"}
	is_approval_needed = var.gdp_pipeline_is_approval_needed
	is_chunking_enabled = var.gdp_pipeline_is_chunking_enabled
	is_file_override_in_destination_enabled = var.gdp_pipeline_is_file_override_in_destination_enabled
	is_scanning_enabled = var.gdp_pipeline_is_scanning_enabled
	service_log_group_id = oci_logging_log_group.test_log_group.id
}
```

## Argument Reference

The following arguments are supported:

* `approval_key_vault_id` - (Optional) (Updatable) The KMS vault OCID used for cryptographic approvals of transfers.
* `authorization_details` - (Optional) (Updatable) Authorization information about the pipeline being configured.
* `bucket_details` - (Required) Configuration information about the buckets used for this pipeline.
	* `bucket_type` - (Required) Type of bucket. SENDER pipelines can be SOURCE, TRANSFER, REJECT, or FAILED. RECEIVER pipelines have a DESTINATION bucket.
	* `id` - (Required) OCID of the bucket.
	* `name` - (Required) Name of the bucket.
	* `namespace` - (Required) Namespace of the bucket.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Short field input by customer for a description of the data pipeline use-case.
* `display_name` - (Required) (Updatable) Pipeline short name.
* `env` - (Optional) The environment where the pipeline resides. Valid values are COMMERCIAL or USGOV. Defaults to COMMERCIAL.
* `file_types` - (Optional) (Updatable) List of file types allowed to be transferred in the pipeline according to the authorization details (e.g. .pdf, .xml, .doc).
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_approval_needed` - (Optional) (Updatable) Determines whether file transfers need to go through an approval workflow.
* `is_chunking_enabled` - (Optional) (Updatable) Determines whether file must be chunked during the transfer. This is only a property of SENDER pipelines.
* `is_file_override_in_destination_enabled` - (Optional) (Updatable) Enable file override feature in destination bucket. If 2 files with same name exist in destination bucket, original file will be overwritten.
* `is_scanning_enabled` - (Optional) (Updatable) Determines whether GDP Scanning should be enabled for the pipeline.
* `peering_region` - (Required) Public region name where the peered pipeline exists.
* `pipeline_type` - (Required) Type of pipeline. Can be SENDER or RECEIVER.
* `service_log_group_id` - (Optional) (Updatable) the OCID of the service log group.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `is_file_override_in_destination_enabled` - Enable file override feature in destination bucket. If 2 files with same name exist in destination bucket, original file will be overwritten
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Gdp Pipeline
	* `update` - (Defaults to 20 minutes), when updating the Gdp Pipeline
	* `delete` - (Defaults to 20 minutes), when destroying the Gdp Pipeline


## Import

GdpPipelines can be imported using the `id`, e.g.

```
$ terraform import oci_gdp_gdp_pipeline.test_gdp_pipeline "id"
```

