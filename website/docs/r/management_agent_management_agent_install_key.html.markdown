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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/management-agent/latest/ManagementAgentInstallKey

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/management_agent

User creates a new install key as part of this API.


## Example Usage

```hcl
resource "oci_management_agent_management_agent_install_key" "test_management_agent_install_key" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.management_agent_install_key_display_name

	#Optional
	allowed_key_install_count = var.management_agent_install_key_allowed_key_install_count
	is_unlimited = var.management_agent_install_key_is_unlimited
	time_expires = var.management_agent_install_key_time_expires
}
```

## Argument Reference

The following arguments are supported:

* `allowed_key_install_count` - (Optional) Total number of install for this keys
* `compartment_id` - (Required) Compartment Identifier
* `display_name` - (Required) (Updatable) Management Agent install Key Name
* `is_unlimited` - (Optional) If set to true, the install key has no expiration date or usage limit. Defaults to false
* `time_expires` - (Optional) date after which key would expire after creation


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `allowed_key_install_count` - Total number of install for this keys
* `compartment_id` - Compartment Identifier
* `created_by_principal_id` - Principal id of user who created the Agent Install key
* `current_key_install_count` - Total number of install for this keys
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Management Agent Install Key Name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Agent install Key identifier
* `is_unlimited` - If set to true, the install key has no expiration date or usage limit. Properties allowedKeyInstallCount and timeExpires are ignored if set to true. Defaults to false.
* `key` - Management Agent Install Key
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - Status of Key
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when Management Agent install Key was created. An RFC3339 formatted date time string
* `time_expires` - date after which key would expire after creation
* `time_updated` - The time when Management Agent install Key was updated. An RFC3339 formatted date time string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Agent Install Key
	* `update` - (Defaults to 20 minutes), when updating the Management Agent Install Key
	* `delete` - (Defaults to 20 minutes), when destroying the Management Agent Install Key


## Import

ManagementAgentInstallKeys can be imported using the `id`, e.g.

```
$ terraform import oci_management_agent_management_agent_install_key.test_management_agent_install_key "id"
```

