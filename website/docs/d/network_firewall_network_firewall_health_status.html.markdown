---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_health_status"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_health_status"
description: |-
  Provides details about a specific Network Firewall Health Status in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_health_status
This data source provides details about a specific Network Firewall Health Status resource in Oracle Cloud Infrastructure Network Firewall service.

Get Overall health status of Network Firewall


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_health_status" "test_network_firewall_health_status" {
	#Required
	network_firewall_id = oci_network_firewall_network_firewall.test_network_firewall.id
}
```

## Argument Reference

The following arguments are supported:

* `network_firewall_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Network Firewall resource.


## Attributes Reference

The following attributes are exported:

* `status` - Overall health status of Network firewall 

