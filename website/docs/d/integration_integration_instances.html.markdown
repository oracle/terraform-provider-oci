---
subcategory: "Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_integration_integration_instances"
sidebar_current: "docs-oci-datasource-integration-integration_instances"
description: |-
  Provides the list of Integration Instances in Oracle Cloud Infrastructure Integration service
---

# Data Source: oci_integration_integration_instances
This data source provides the list of Integration Instances in Oracle Cloud Infrastructure Integration service.

Returns a list of Integration Instances.


## Example Usage

```hcl
data "oci_integration_integration_instances" "test_integration_instances" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.integration_instance_display_name}"
	state = "${var.integration_instance_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) Life cycle state to query on.


## Attributes Reference

The following attributes are exported:

* `integration_instances` - The list of integration_instances.

### IntegrationInstance Reference

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

