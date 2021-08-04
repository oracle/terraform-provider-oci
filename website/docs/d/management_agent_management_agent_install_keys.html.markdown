---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_install_keys"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_install_keys"
description: |-
  Provides the list of Management Agent Install Keys in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_install_keys
This data source provides the list of Management Agent Install Keys in Oracle Cloud Infrastructure Management Agent service.

Returns a list of Management Agent installed Keys.


## Example Usage

```hcl
data "oci_management_agent_management_agent_install_keys" "test_management_agent_install_keys" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.management_agent_install_key_access_level
	compartment_id_in_subtree = var.management_agent_install_key_compartment_id_in_subtree
	display_name = var.management_agent_install_key_display_name
	state = var.management_agent_install_key_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Value of this is always "ACCESSIBLE" and any other value is not supported.
* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.
* `compartment_id_in_subtree` - (Optional) if set to true then it fetches install key for all compartments where user has access to else only on the compartment specified.
* `display_name` - (Optional) The display name for which the Key needs to be listed.
* `state` - (Optional) Filter to return only Management Agents in the particular lifecycle state.


## Attributes Reference

The following attributes are exported:

* `management_agent_install_keys` - The list of management_agent_install_keys.

### ManagementAgentInstallKey Reference

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

