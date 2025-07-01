---
subcategory: "Apm Traces"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_traces_scheduled_query"
sidebar_current: "docs-oci-datasource-apm_traces-scheduled_query"
description: |-
  Provides details about a specific Scheduled Query in Oracle Cloud Infrastructure Apm Traces service
---

# Data Source: oci_apm_traces_scheduled_query
This data source provides details about a specific Scheduled Query resource in Oracle Cloud Infrastructure Apm Traces service.

Retrieve a scheduled query in the APM Domain.


## Example Usage

```hcl
data "oci_apm_traces_scheduled_query" "test_scheduled_query" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	scheduled_query_id = oci_apm_traces_scheduled_query.test_scheduled_query.id
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM Domain ID for the intended request. 
* `scheduled_query_id` - (Required) Id of the scheduled query. 


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

