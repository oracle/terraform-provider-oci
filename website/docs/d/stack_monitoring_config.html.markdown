---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_config"
sidebar_current: "docs-oci-datasource-stack_monitoring-config"
description: |-
  Provides details about a specific Config in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_config
This data source provides details about a specific Config resource in Oracle Cloud Infrastructure Stack Monitoring service.

Gets the details of a configuration.

## Example Usage

```hcl
data "oci_stack_monitoring_config" "test_config" {
	#Required
	config_id = oci_stack_monitoring_config.test_config.id
}
```

## Argument Reference

The following arguments are supported:

* `config_id` - (Required) Unique Config identifier.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the configuration.
* `config_type` - The type of configuration.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Unique Oracle ID (OCID) that is immutable on creation.
* `is_enabled` - True if automatic promotion or enterprise extensibility is enabled, false if it is not enabled.
* `license` - License edition.
* `resource_type` - The type of resource to configure for automatic promotion.
* `state` - The current state of the configuration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the configuration was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Config was updated.

