---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_security_policies"
sidebar_current: "docs-oci-datasource-cloud_guard-security_policies"
description: |-
  Provides the list of Security Policies in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_security_policies
This data source provides the list of Security Policies in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of security zone policies (SecurityPolicySummary resources),
identified by compartmentId.


## Example Usage

```hcl
data "oci_cloud_guard_security_policies" "test_security_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.security_policy_display_name
	id = var.security_policy_id
	state = var.security_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) The unique identifier of the security zone policy. (`SecurityPolicy`)
* `state` - (Optional) The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.


## Attributes Reference

The following attributes are exported:

* `security_policy_collection` - The list of security_policy_collection.

### SecurityPolicy Reference

The following attributes are exported:

* `category` - The category of the security policy
* `compartment_id` - The OCID of the security policy's compartment
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The security policy's description
* `display_name` - The security policy's display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `friendly_name` - A shorter version of the security policy's name
* `id` - Unique identifier that can’t be changed after creation
* `lifecycle_details` - A message describing the current state in more detail. For example, this can be used to provide actionable information for a resource in a `Failed` state.
* `owner` - The owner of the security policy
* `services` - The list of services that the security policy protects
* `state` - The current lifecycle state of the security policy
* `time_created` - The time the security policy was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the security policy was last updated. An RFC3339 formatted datetime string.

