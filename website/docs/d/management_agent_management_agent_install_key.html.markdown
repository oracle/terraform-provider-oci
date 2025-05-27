---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_install_key"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_install_key"
description: |-
  Provides details about a specific Management Agent Install Key in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_install_key
This data source provides details about a specific Management Agent Install Key resource in Oracle Cloud Infrastructure Management Agent service.

Gets complete details of the Agent install Key for a given key id

## Example Usage

```hcl
data "oci_management_agent_management_agent_install_key" "test_management_agent_install_key" {
	#Required
	management_agent_install_key_id = oci_management_agent_management_agent_install_key.test_management_agent_install_key.id
}
```

## Argument Reference

The following arguments are supported:

* `management_agent_install_key_id` - (Required) Unique Management Agent Install Key identifier


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

