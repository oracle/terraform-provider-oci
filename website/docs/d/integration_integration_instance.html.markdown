---
subcategory: "Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_integration_integration_instance"
sidebar_current: "docs-oci-datasource-integration-integration_instance"
description: |-
  Provides details about a specific Integration Instance in Oracle Cloud Infrastructure Integration service
---

# Data Source: oci_integration_integration_instance
This data source provides details about a specific Integration Instance resource in Oracle Cloud Infrastructure Integration service.

Gets a IntegrationInstance by identifier

## Example Usage

```hcl
data "oci_integration_integration_instance" "test_integration_instance" {
	#Required
	integration_instance_id = "${oci_integration_integration_instance.test_integration_instance.id}"
}
```

## Argument Reference

The following arguments are supported:

* `integration_instance_id` - (Required) Unique Integration Instance identifier.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Integration Instance Identifier, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `instance_url` - The Integration Instance URL.
* `integration_instance_type` - Standard or Enterprise type
* `is_byol` - Bring your own license.
* `message_packs` - The number of configured message packs (if any)
* `state` - The current state of the integration instance.
* `state_message` - An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `time_created` - The time the the Integration Instance was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the IntegrationInstance was updated. An RFC3339 formatted datetime string.

