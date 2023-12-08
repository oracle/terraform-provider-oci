---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_logs"
sidebar_current: "docs-oci-datasource-logging-logs"
description: |-
  Provides the list of Logs in Oracle Cloud Infrastructure Logging service
---

# Data Source: oci_logging_logs
This data source provides the list of Logs in Oracle Cloud Infrastructure Logging service.

Lists the specified log group's log objects.

## Example Usage

```hcl
data "oci_logging_logs" "test_logs" {
	#Required
	log_group_id = oci_logging_log_group.test_log_group.id

	#Optional
	display_name = var.log_display_name
	log_type = var.log_log_type
	source_resource = var.log_source_resource
	source_service = var.log_source_service
	state = var.log_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) Resource name.
* `log_group_id` - (Required) OCID of a log group to work with.
* `log_type` - (Optional) The logType that the log object is for, whether custom or service.
* `source_resource` - (Optional) Log object resource, which is a field of LogSummary.Configuration.Source.
* `source_service` - (Optional) Service that created the log object, which is a field of LogSummary.Configuration.Source.
* `state` - (Optional) Lifecycle state of the log object


## Attributes Reference

The following attributes are exported:

* `logs` - The list of logs.

### Log Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that the resource belongs to.
* `configuration` - Log object configuration.
	* `compartment_id` - The OCID of the compartment that the resource belongs to.
	* `source` - The source the log object comes from.
		* `category` - Log object category.
		* `parameters` - Log category parameters are stored here.
		* `resource` - The unique identifier of the resource emitting the log.
		* `service` - Service generating log.
		* `source_type` - The log source.
			* **OCISERVICE:** Oracle Service. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource.
* `is_enabled` - Whether or not this resource is currently enabled.
* `log_group_id` - Log group OCID.
* `log_type` - The logType that the log object is for, whether custom or service.
* `retention_duration` - Log retention duration in 30-day increments (30, 60, 90 and so on until 180).
* `state` - The pipeline state.
* `tenancy_id` - The OCID of the tenancy.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

