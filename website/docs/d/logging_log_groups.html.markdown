---
subcategory: "Logging"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_logging_log_groups"
sidebar_current: "docs-oci-datasource-logging-log_groups"
description: |-
  Provides the list of Log Groups in Oracle Cloud Infrastructure Logging service
---

# Data Source: oci_logging_log_groups
This data source provides the list of Log Groups in Oracle Cloud Infrastructure Logging service.

Lists all log groups for the specified compartment or tenancy.

## Example Usage

```hcl
data "oci_logging_log_groups" "test_log_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.log_group_display_name
	is_compartment_id_in_subtree = var.log_group_is_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Compartment OCID to list resources in. Please see compartmentIdInSubtree for nested compartments traversal. 
* `display_name` - (Optional) Resource name
* `is_compartment_id_in_subtree` - (Optional) Specifies whether or not nested compartments should be traversed. Defaults to false.


## Attributes Reference

The following attributes are exported:

* `log_groups` - The list of log_groups.

### LogGroup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that the resource belongs to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description for this resource.
* `display_name` - The display name of a user-friendly name. It has to be unique within enclosing resource, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource.
* `state` - The state of the log group object.
* `time_created` - Time the resource was created.
* `time_last_modified` - Time the resource was last modified.

