---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_protection_capabilities"
sidebar_current: "docs-oci-datasource-waf-protection_capabilities"
description: |-
  Provides the list of Protection Capabilities in Oracle Cloud Infrastructure Waf service
---

# Data Source: oci_waf_protection_capabilities
This data source provides the list of Protection Capabilities in Oracle Cloud Infrastructure Waf service.

Lists of protection capabilities filtered by query parameters.


## Example Usage

```hcl
data "oci_waf_protection_capabilities" "test_protection_capabilities" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.protection_capability_display_name
	group_tag = var.protection_capability_group_tag
	is_latest_version = var.protection_capability_is_latest_version
	key = var.protection_capability_key
	type = var.protection_capability_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `group_tag` - (Optional) A filter to return only resources that are accociated given group tag.
* `is_latest_version` - (Optional) A filter to return only resources that matches given isLatestVersion.
* `key` - (Optional) The unique key of protection capability to filter by.
* `type` - (Optional) A filter to return only resources that matches given type.


## Attributes Reference

The following attributes are exported:

* `protection_capability_collection` - The list of protection_capability_collection.

### ProtectionCapability Reference

The following attributes are exported:

* `items` - List of protection capabilities.
	* `collaborative_action_threshold` - The default collaborative action threshold for OCI-managed collaborative protection capability. Collaborative protection capabilities are made of several simple, non-collaborative protection capabilities (referred to as `contributing capabilities` later on) which have weights assigned to them. These weights can be found in the `collaborativeWeights` array. For incoming/outgoing HTTP messages, all contributing capabilities are executed and the sum of all triggered contributing capabilities weights is calculated. Only if this sum is greater than or equal to `collaborativeActionThreshold` is the incoming/outgoing HTTP message marked as malicious.

		This field is ignored for non-collaborative capabilities. 
	* `collaborative_weights` - The weights of contributing capabilities. Defines how much each contributing capability contributes towards the action threshold of a collaborative protection capability.

		This field is ignored for non-collaborative capabilities. 
		* `display_name` - The display name of contributing protection capability.
		* `key` - Unique key of contributing protection capability.
		* `weight` - The weight of contributing protection capability.
	* `description` - The description of protection capability.
	* `display_name` - The display name of protection capability.
	* `group_tags` - The list of unique names protection capability group tags that are associated with this capability. Example: ["PCI", "Recommended"] 
	* `is_latest_version` - The field that shows if this is the latest version of protection capability.
	* `key` - Unique key of protection capability.
	* `type` - The type of protection capability.
		* **REQUEST_PROTECTION_CAPABILITY** can only be used in `requestProtection` module of WebAppFirewallPolicy.
		* **RESPONSE_PROTECTION_CAPABILITY** can only be used in `responseProtection` module of WebAppFirewallPolicy. 
	* `version` - The version of protection capability.

