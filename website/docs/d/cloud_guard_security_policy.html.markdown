---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_security_policy"
sidebar_current: "docs-oci-datasource-cloud_guard-security_policy"
description: |-
  Provides details about a specific Security Policy in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_security_policy
This data source provides details about a specific Security Policy resource in Oracle Cloud Infrastructure Cloud Guard service.

Gets a security zone policy using its identifier. When a policy is enabled in a security zone, then any action in the zone that attempts to violate that policy is denied.

## Example Usage

```hcl
data "oci_cloud_guard_security_policy" "test_security_policy" {
	#Required
	security_policy_id = oci_cloud_guard_security_policy.test_security_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `security_policy_id` - (Required) The unique identifier of the security zone policy (`SecurityPolicy`)


## Attributes Reference

The following attributes are exported:

* `category` - The category of security policy
* `compartment_id` - The id of the security policy's compartment
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The security policy's description
* `display_name` - The security policy's full name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `friendly_name` - A shorter version of the security policy's name
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, this can be used to provide actionable information for a resource in a `Failed` state.
* `owner` - The owner of the security policy
* `services` - The list of services that the security policy protects
* `state` - The current state of the security policy
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the security policy was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the security policy was last updated. An RFC3339 formatted datetime string.

