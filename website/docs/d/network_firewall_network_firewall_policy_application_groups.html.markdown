---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_application_groups"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_application_groups"
description: |-
  Provides the list of Network Firewall Policy Application Groups in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_application_groups
This data source provides the list of Network Firewall Policy Application Groups in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of ApplicationGroups for the policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_application_groups" "test_network_firewall_policy_application_groups" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_application_group_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `application_group_summary_collection` - The list of application_group_summary_collection.

### NetworkFirewallPolicyApplicationGroup Reference

The following attributes are exported:

* `apps` - List of apps in the group.
* `name` - Name of the application Group.
* `parent_resource_id` - OCID of the Network Firewall Policy this application group belongs to.
* `total_apps` - Count of total applications in the given application group.

