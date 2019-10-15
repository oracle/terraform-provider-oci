---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_integration_integration_instance"
sidebar_current: "docs-oci-resource-integration-integration_instance"
description: |-
  Provides the Integration Instance resource in Oracle Cloud Infrastructure Integration service
---

# oci_integration_integration_instance
This resource provides the Integration Instance resource in Oracle Cloud Infrastructure Integration service.

Creates a new Integration Instance.


## Example Usage

```hcl
resource "oci_integration_integration_instance" "test_integration_instance" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.integration_instance_display_name}"
	integration_instance_type = "${var.integration_instance_integration_instance_type}"
	is_byol = "${var.integration_instance_is_byol}"
	message_packs = "${var.integration_instance_message_packs}"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	idcs_at = "${var.integration_instance_idcs_at}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Integration Instance Identifier.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `idcs_at` - (Optional) IDCS Authentication token. This is is required for pre-UCPIS cloud accounts, but not UCPIS, hence not a required parameter
* `integration_instance_type` - (Required) (Updatable) Standard or Enterprise type
* `is_byol` - (Required) (Updatable) Bring your own license.
* `message_packs` - (Required) (Updatable) The number of configured message packs


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

IntegrationInstances can be imported using the `id`, e.g.

```
$ terraform import oci_integration_integration_instance.test_integration_instance "id"
```

