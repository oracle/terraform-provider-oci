---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_service"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_service"
description: |-
  Provides details about a specific Network Firewall Policy Service in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_service
This data source provides details about a specific Network Firewall Policy Service resource in Oracle Cloud Infrastructure Network Firewall service.

Get Service by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_service" "test_network_firewall_policy_service" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	service_name = var.oci_network_firewall_network_firewall_policy_service_name
}
```

## Argument Reference

The following arguments are supported:

* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `service_name` - (Required) Unique identifier for Services.


## Attributes Reference

The following attributes are exported:

* `name` - Name of the service.
* `parent_resource_id` - OCID of the Network Firewall Policy this service belongs to.
* `port_ranges` - List of port-ranges used.
	* `maximum_port` - The maximum port in the range (inclusive), which may be absent for a single-port range.
	* `minimum_port` - The minimum port in the range (inclusive), or the sole port of a single-port range.
* `type` - Describes the type of Service.

