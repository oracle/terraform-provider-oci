---
subcategory: "Apm Traces"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_traces_scheduled_query"
sidebar_current: "docs-oci-resource-apm_traces-scheduled_query"
description: |-
  Provides the Scheduled Query resource in Oracle Cloud Infrastructure Apm Traces service
---

# oci_apm_traces_scheduled_query
This resource provides the Scheduled Query resource in Oracle Cloud Infrastructure Apm Traces service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/latest/ScheduledQuery

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/apm/apm_traces

Create a scheduled query in the APM Domain.


## Example Usage

```hcl
resource "oci_apm_traces_scheduled_query" "test_scheduled_query" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	opc_dry_run = var.scheduled_query_opc_dry_run
	scheduled_query_description = var.scheduled_query_scheduled_query_description
	scheduled_query_maximum_runtime_in_seconds = var.scheduled_query_scheduled_query_maximum_runtime_in_seconds
	scheduled_query_name = oci_apm_traces_scheduled_query.test_scheduled_query.name
	scheduled_query_processing_configuration {

		#Optional
		custom_metric {
			#Required
			name = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_name

			#Optional
			compartment = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_compartment
			description = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_description
			is_anomaly_detection_enabled = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_is_anomaly_detection_enabled
			is_metric_published = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_is_metric_published
			namespace = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_namespace
			resource_group = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_resource_group
			unit = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_unit
		}
		object_storage {

			#Optional
			bucket = var.scheduled_query_scheduled_query_processing_configuration_object_storage_bucket
			name_space = var.scheduled_query_scheduled_query_processing_configuration_object_storage_name_space
			object_name_prefix = var.scheduled_query_scheduled_query_processing_configuration_object_storage_object_name_prefix
		}
		streaming {

			#Optional
			stream_id = oci_streaming_stream.test_stream.id
		}
	}
	scheduled_query_processing_sub_type = var.scheduled_query_scheduled_query_processing_sub_type
	scheduled_query_processing_type = var.scheduled_query_scheduled_query_processing_type
	scheduled_query_retention_criteria = var.scheduled_query_scheduled_query_retention_criteria
	scheduled_query_retention_period_in_ms = var.scheduled_query_scheduled_query_retention_period_in_ms
	scheduled_query_schedule = var.scheduled_query_scheduled_query_schedule
	scheduled_query_text = var.scheduled_query_scheduled_query_text
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) (Updatable) The APM Domain ID for the intended request. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `opc_dry_run` - (Optional) (Updatable) Indicates that the request is a dry run, if set to "true". A dry run request does not create or modify the resource  and is used only to perform validation on the submitted data. 
* `scheduled_query_description` - (Optional) (Updatable) Description for the scheduled query. 
* `scheduled_query_maximum_runtime_in_seconds` - (Optional) (Updatable) Maximum runtime for the scheduled query in seconds. 
* `scheduled_query_name` - (Optional) (Updatable) Name of the scheduled query. 
* `scheduled_query_processing_configuration` - (Optional) (Updatable) Definition of the scheduled query processing configuration. 
	* `custom_metric` - (Optional) (Updatable) Definition of the Custom Metric. 
		* `compartment` - (Optional) (Updatable) Compartment of the Monitoring Service. It defaults to the APM domain's compartment if not specified.  If specified, the necessary Oracle Cloud Infrastructure policies should be set to allow APM to write to that compartment. 
		* `description` - (Optional) (Updatable) Description of the Custom Metric. 
		* `is_anomaly_detection_enabled` - (Optional) (Updatable) Indicates whether anomaly Detection should be performed on the generated metric. 
		* `is_metric_published` - (Optional) (Updatable) Used in conjunction with the dry run header.  When the dry run header is set and the isPublishMetric flag is set to true, the  scheduled query is not created, but validations happen to check if the right Oracle Cloud Infrastructure policies have been set to write to the specified namespace/compartment. 
		* `name` - (Required) (Updatable) Name of the Custom Metric. 
		* `namespace` - (Optional) (Updatable) Namespace in the Custom Metric. It defaults to `oracle_apm_custom` if not specified.  If specified, the necessary Oracle Cloud Infrastructure policies should be set to allow APM to write to that namespace. 
		* `resource_group` - (Optional) (Updatable) Resource Group of the Custom Metric. 
		* `unit` - (Optional) (Updatable) Unit in which the metric value is reported. For example 'ms'. 
	* `object_storage` - (Optional) (Updatable) Definition of the object storage. 
		* `bucket` - (Optional) (Updatable) Bucket name in the object store. 
		* `name_space` - (Optional) (Updatable) Namespace in the object store. 
		* `object_name_prefix` - (Optional) (Updatable) Object name prefix in the object store. 
	* `streaming` - (Optional) (Updatable) Definition of the Stream. 
		* `stream_id` - (Optional) (Updatable) Stream Id. 
* `scheduled_query_processing_sub_type` - (Optional) (Updatable) Processing sub type of the scheduled query. 
* `scheduled_query_processing_type` - (Optional) (Updatable) Type of the scheduled query. 
* `scheduled_query_retention_criteria` - (Optional) (Updatable) Retention criteria for the scheduled query. 
* `scheduled_query_retention_period_in_ms` - (Optional) (Updatable) Retention period for the scheduled query in milliseconds. 
* `scheduled_query_schedule` - (Optional) (Updatable) Schedule for the scheduled query. 
* `scheduled_query_text` - (Optional) (Updatable) Scheduled query to be run. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the scheduled query . An OCID is generated when the scheduled query is created. 
* `scheduled_query_description` - Description for the scheduled query. 
* `scheduled_query_instances` - Scheduled query instances. 
* `scheduled_query_maximum_runtime_in_seconds` - Maximum runtime for the scheduled query in seconds. 
* `scheduled_query_name` - Name of the scheduled query. 
* `scheduled_query_next_run_in_ms` - Next run for the scheduled query. 
* `scheduled_query_processing_configuration` - Definition of the scheduled query processing configuration. 
	* `custom_metric` - Definition of the Custom Metric. 
		* `compartment` - Compartment of the Monitoring Service. It defaults to the APM domain's compartment if not specified.  If specified, the necessary Oracle Cloud Infrastructure policies should be set to allow APM to write to that compartment. 
		* `description` - Description of the Custom Metric. 
		* `is_anomaly_detection_enabled` - Indicates whether anomaly Detection should be performed on the generated metric. 
		* `is_metric_published` - Used in conjunction with the dry run header.  When the dry run header is set and the isPublishMetric flag is set to true, the  scheduled query is not created, but validations happen to check if the right Oracle Cloud Infrastructure policies have been set to write to the specified namespace/compartment. 
		* `name` - Name of the Custom Metric. 
		* `namespace` - Namespace in the Custom Metric. It defaults to `oracle_apm_custom` if not specified.  If specified, the necessary Oracle Cloud Infrastructure policies should be set to allow APM to write to that namespace. 
		* `resource_group` - Resource Group of the Custom Metric. 
		* `unit` - Unit in which the metric value is reported. For example 'ms'. 
	* `object_storage` - Definition of the object storage. 
		* `bucket` - Bucket name in the object store. 
		* `name_space` - Namespace in the object store. 
		* `object_name_prefix` - Object name prefix in the object store. 
	* `streaming` - Definition of the Stream. 
		* `stream_id` - Stream Id. 
* `scheduled_query_processing_sub_type` - Processing sub type of the scheduled query. 
* `scheduled_query_processing_type` - Processing type of the scheduled query. 
* `scheduled_query_retention_criteria` - Retention criteria for the scheduled query. 
* `scheduled_query_retention_period_in_ms` - Retention period for the scheduled query in milliseconds. 
* `scheduled_query_schedule` - Schedule for the scheduled query. 
* `scheduled_query_text` - Scheduled query to be run. 
* `state` - The current lifecycle state of the Scheduled Query.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Scheduled Query
	* `update` - (Defaults to 20 minutes), when updating the Scheduled Query
	* `delete` - (Defaults to 20 minutes), when destroying the Scheduled Query


## Import

ScheduledQueries can be imported using the `id`, e.g.

```
$ terraform import oci_apm_traces_scheduled_query.test_scheduled_query "scheduledQueries/{scheduledQueryId}/apmDomainId/{apmDomainId}" 
```

