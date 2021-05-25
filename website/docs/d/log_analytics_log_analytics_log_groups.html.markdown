---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_log_groups"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_log_groups"
description: |-
  Provides the list of Log Analytics Log Groups in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_log_groups
This data source provides the list of Log Analytics Log Groups in Oracle Cloud Infrastructure Log Analytics service.

Returns a list of log groups in a compartment. You may limit the number of log groups, provide sorting options, and filter the results by specifying a display name.


## Example Usage

```hcl
data "oci_log_analytics_log_analytics_log_groups" "test_log_analytics_log_groups" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.log_analytics_log_group_namespace

	#Optional
	display_name = var.log_analytics_log_group_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only log analytics log groups whose displayName matches the entire display name given. The match is case-insensitive. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `log_analytics_log_group_summary_collection` - The list of log_analytics_log_group_summary_collection.

### LogAnalyticsLogGroup Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description for this resource. 
* `display_name` - A user-friendly name that is changeable and that does not have to be unique. Format: a leading alphanumeric, followed by zero or more alphanumerics, underscores, spaces, backslashes, or hyphens in any order). No trailing spaces allowed. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
* `time_created` - The date and time the resource was created, in the format defined by RFC3339. 
* `time_updated` - The date and time the resource was last updated, in the format defined by RFC3339. 

