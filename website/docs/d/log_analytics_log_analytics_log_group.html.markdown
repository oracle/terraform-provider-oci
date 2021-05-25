---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_log_group"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_log_group"
description: |-
  Provides details about a specific Log Analytics Log Group in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_log_group
This data source provides details about a specific Log Analytics Log Group resource in Oracle Cloud Infrastructure Log Analytics service.

Gets detailed information about the specified log group such as display name, description, defined tags, and free-form tags.


## Example Usage

```hcl
data "oci_log_analytics_log_analytics_log_group" "test_log_analytics_log_group" {
	#Required
	log_analytics_log_group_id = oci_log_analytics_log_analytics_log_group.test_log_analytics_log_group.id
	namespace = var.log_analytics_log_group_namespace
}
```

## Argument Reference

The following arguments are supported:

* `log_analytics_log_group_id` - (Required) unique logAnalytics log group identifier
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


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

