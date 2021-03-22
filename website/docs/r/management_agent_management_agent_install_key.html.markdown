---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_install_key"
sidebar_current: "docs-oci-resource-management_agent-management_agent_install_key"
description: |-
  Provides the Management Agent Install Key resource in Oracle Cloud Infrastructure Management Agent service
---

# oci_management_agent_management_agent_install_key
This resource provides the Management Agent Install Key resource in Oracle Cloud Infrastructure Management Agent service.

User creates a new install key as part of this API.


## Example Usage

```hcl
resource "oci_management_agent_management_agent_install_key" "test_management_agent_install_key" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.management_agent_install_key_display_name

	#Optional
	allowed_key_install_count = var.management_agent_install_key_allowed_key_install_count
	time_expires = var.management_agent_install_key_time_expires
}
```

## Argument Reference

The following arguments are supported:

* `allowed_key_install_count` - (Optional) Total number of install for this keys
* `compartment_id` - (Required) Compartment Identifier
* `display_name` - (Required) (Updatable) Management Agent install Key Name
* `time_expires` - (Optional) date after which key would expire after creation


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `allowed_key_install_count` - Total number of install for this keys
* `compartment_id` - Compartment Identifier
* `created_by_principal_id` - Principal id of user who created the Agent Install key
* `current_key_install_count` - Total number of install for this keys
* `display_name` - Management Agent Install Key Name
* `id` - Agent install Key identifier
* `key` - Management Agent Install Key
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - Status of Key
* `time_created` - The time when Management Agent install Key was created. An RFC3339 formatted date time string
* `time_expires` - date after which key would expire after creation
* `time_updated` - The time when Management Agent install Key was updated. An RFC3339 formatted date time string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Agent Install Key
	* `update` - (Defaults to 20 minutes), when updating the Management Agent Install Key
	* `delete` - (Defaults to 20 minutes), when destroying the Management Agent Install Key


## Import

ManagementAgentInstallKeys can be imported using the `id`, e.g.

```
$ terraform import oci_management_agent_management_agent_install_key.test_management_agent_install_key "id"
```

