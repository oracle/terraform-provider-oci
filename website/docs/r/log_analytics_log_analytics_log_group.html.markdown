---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_log_group"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_log_group"
description: |-
  Provides the Log Analytics Log Group resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_log_group
This resource provides the Log Analytics Log Group resource in Oracle Cloud Infrastructure Log Analytics service.

Creates a new log group in the specified compartment with the input display name. You may also specify optional information such as description, defined tags, and free-form tags.


## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_log_group" "test_log_analytics_log_group" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.log_analytics_log_group_display_name
	namespace = var.log_analytics_log_group_namespace

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.log_analytics_log_group_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description for this resource. 
* `display_name` - (Required) (Updatable) A user-friendly name that is changeable and that does not have to be unique. Format: a leading alphanumeric, followed by zero or more alphanumerics, underscores, spaces, backslashes, or hyphens in any order). No trailing spaces allowed. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description for this resource. 
* `display_name` - A user-friendly name that is changeable and that does not have to be unique. Format: a leading alphanumeric, followed by zero or more alphanumerics, underscores, spaces, backslashes, or hyphens in any order). No trailing spaces allowed. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
* `time_created` - The date and time the resource was created, in the format defined by RFC3339. 
* `time_updated` - The date and time the resource was last updated, in the format defined by RFC3339. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log Analytics Log Group
	* `update` - (Defaults to 20 minutes), when updating the Log Analytics Log Group
	* `delete` - (Defaults to 20 minutes), when destroying the Log Analytics Log Group


## Import

LogAnalyticsLogGroups can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_log_analytics_log_group.test_log_analytics_log_group "namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}" 
```

