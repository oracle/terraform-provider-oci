---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_log"
sidebar_current: "docs-oci-resource-logging-log"
description: |-
  Provides the Log resource in Oracle Cloud Infrastructure Logging service
---

# oci_logging_log
This resource provides the Log resource in Oracle Cloud Infrastructure Logging service.

Creates a log within the specified log group. This call fails if a log group has already been created
with the same displayName or (service, resource, category) triplet.


## Example Usage

```hcl
resource "oci_logging_log" "test_log" {
	#Required
	display_name = var.log_display_name
	log_group_id = oci_logging_log_group.test_log_group.id
	log_type = var.log_log_type

	#Optional
	configuration {
		#Required
		source {
			#Required
			category = var.log_configuration_source_category
			resource = var.log_configuration_source_resource
			service = var.log_configuration_source_service
			source_type = var.log_configuration_source_source_type
		}

		#Optional
		compartment_id = var.compartment_id
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_enabled = var.log_is_enabled
	retention_duration = var.log_retention_duration
}
```

## Argument Reference

The following arguments are supported:

* `configuration` - (Optional) Log object configuration.
	* `compartment_id` - (Optional) The OCID of the compartment that the resource belongs to.
	* `source` - (Required) The source the log object comes from.
		* `category` - (Required) Log object category.
		* `resource` - (Required) The unique identifier of the resource emitting the log.
		* `service` - (Required) Service generating log.
		* `source_type` - (Required) The log source.
			* **OCISERVICE:** Oracle Service. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_enabled` - (Optional) (Updatable) Whether or not this resource is currently enabled.
* `log_group_id` - (Required) (Updatable) OCID of a log group to work with.
* `log_type` - (Required) The logType that the log object is for, whether custom or service.
* `retention_duration` - (Optional) (Updatable) Log retention duration in 30-day increments (30, 60, 90 and so on).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that the resource belongs to.
* `configuration` - Log object configuration.
	* `compartment_id` - The OCID of the compartment that the resource belongs to.
	* `source` - The source the log object comes from.
		* `category` - Log object category.
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
* `retention_duration` - Log retention duration in 30-day increments (30, 60, 90 and so on).
* `state` - The pipeline state.
* `tenancy_id` - The OCID of the tenancy.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log
	* `update` - (Defaults to 20 minutes), when updating the Log
	* `delete` - (Defaults to 20 minutes), when destroying the Log


## Import

Logs can be imported using the `id`, e.g.

```
$ terraform import oci_logging_log.test_log "logGroupId/{logGroupId}/logId/{logId}"
```

