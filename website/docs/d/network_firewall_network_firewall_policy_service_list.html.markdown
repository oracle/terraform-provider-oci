---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_service_list"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_service_list"
description: |-
  Provides details about a specific Network Firewall Policy Service List in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_service_list
This data source provides details about a specific Network Firewall Policy Service List resource in Oracle Cloud Infrastructure Network Firewall service.

Get ServiceList by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_service_list" "test_network_firewall_policy_service_list" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	service_list_name = var.network_firewall_policy_service_list_service_list_name
}
```

## Argument Reference

The following arguments are supported:

* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `service_list_name` - (Required) Unique name identifier for Service Lists in the scope of Network Firewall Policy.


## Attributes Reference

The following attributes are exported:

* `name` - Name of the service Group.
* `parent_resource_id` - OCID of the Network Firewall Policy this serviceList belongs to.
* `services` - List of services in the group.
* `total_services` - Count of total services in the given service List.

