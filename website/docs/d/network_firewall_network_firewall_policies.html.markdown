---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policies"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policies"
description: |-
  Provides the list of Network Firewall Policies in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policies
This data source provides the list of Network Firewall Policies in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of Network Firewall Policies.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policies" "test_network_firewall_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.network_firewall_policy_display_name
	id = var.network_firewall_policy_id
	state = var.network_firewall_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique Network Firewall Policy identifier
* `state` - (Optional) A filter to return only resources with a lifecycleState matching the given value.


## Attributes Reference

The following attributes are exported:

* `network_firewall_policy_summary_collection` - The list of network_firewall_policy_summary_collection.

### NetworkFirewallPolicy Reference

The following attributes are exported:

* `attached_network_firewall_count` - Count of number of Network Firewall attached to the Policy. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the NetworkFirewall Policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly optional name for the firewall policy. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource - Network Firewall Policy.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the Network Firewall Policy.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time instant at which the Network Firewall Policy was created in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The time instant at which the Network Firewall Policy was updated in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

