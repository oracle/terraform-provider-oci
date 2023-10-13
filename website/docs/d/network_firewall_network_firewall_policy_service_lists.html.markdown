---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_service_lists"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_service_lists"
description: |-
  Provides the list of Network Firewall Policy Service Lists in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_service_lists
This data source provides the list of Network Firewall Policy Service Lists in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of ServiceLists for the policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_service_lists" "test_network_firewall_policy_service_lists" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_service_list_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `service_list_summary_collection` - The list of service_list_summary_collection.

### NetworkFirewallPolicyServiceList Reference

The following attributes are exported:

* `name` - Name of the service Group.
* `parent_resource_id` - OCID of the Network Firewall Policy this serviceList belongs to.
* `services` - List of services in the group.
* `total_services` - Count of total services in the given service List.

