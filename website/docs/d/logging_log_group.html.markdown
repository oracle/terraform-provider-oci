---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_log_group"
sidebar_current: "docs-oci-datasource-logging-log_group"
description: |-
  Provides details about a specific Log Group in Oracle Cloud Infrastructure Logging service
---

# Data Source: oci_logging_log_group
This data source provides details about a specific Log Group resource in Oracle Cloud Infrastructure Logging service.

Get the specified log group's information.

## Example Usage

```hcl
data "oci_logging_log_group" "test_log_group" {
	#Required
	log_group_id = oci_logging_log_group.test_log_group.id
}
```

## Argument Reference

The following arguments are supported:

* `log_group_id` - (Required) OCID of a log group to work with.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that the resource belongs to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description for this resource.
* `display_name` - The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource.
* `state` - The log group object state.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

