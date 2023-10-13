---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_application_group"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_application_group"
description: |-
  Provides details about a specific Network Firewall Policy Application Group in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_application_group
This data source provides details about a specific Network Firewall Policy Application Group resource in Oracle Cloud Infrastructure Network Firewall service.

Get ApplicationGroup by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_application_group" "test_network_firewall_policy_application_group" {
	#Required
	application_group_name = var.network_firewall_policy_application_group_display_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `application_group_name` - (Required) Unique name identifier for Application Lists in the scope of Network Firewall Policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `apps` - List of apps in the group.
* `name` - Name of the application Group.
* `parent_resource_id` - OCID of the Network Firewall Policy this application group belongs to.
* `total_apps` - Count of total applications in the given application group.

