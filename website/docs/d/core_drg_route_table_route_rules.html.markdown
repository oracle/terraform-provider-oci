---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_route_table_route_rules"
sidebar_current: "docs-oci-datasource-core-drg_route_table_route_rules"
description: |-
  Provides the list of Drg Route Table Route Rules in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_drg_route_table_route_rules
This data source provides the list of Drg Route Table Route Rules in Oracle Cloud Infrastructure Core service.

Lists the route rules in the specified DRG route table.

## Example Usage

```hcl
data "oci_core_drg_route_table_route_rules" "test_drg_route_table_route_rules" {
	#Required
	drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id

	#Optional
	route_type = var.drg_route_table_route_rule_route_type
}
```

## Argument Reference

The following arguments are supported:

* `drg_route_table_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.
* `route_type` - (Optional) Static routes are specified through the DRG route table API. Dynamic routes are learned by the DRG from the DRG attachments through various routing protocols. 


## Attributes Reference

The following attributes are exported:

* `drg_route_rules` - The list of drg_route_rules.

### DrgRouteTableRouteRule Reference

The following attributes are exported:

* `attributes` - Additional properties for the route, computed by the service. 
* `destination` - Represents the range of IP addresses to match against when routing traffic.

	Potential values:
	* An IP address range (IPv4 or IPv6) in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`.
	* When you're setting up a security rule for traffic destined for a particular `Service` through a service gateway, this is the `cidrBlock` value associated with that [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Service/). For example: `oci-phx-objectstorage`. 
* `destination_type` - The type of destination for the rule. the type is required if `direction` = `EGRESS`.

	Allowed values:
	* `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
	* `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic destined for a particular `Service` through a service gateway). 
* `id` - The Oracle-assigned ID of the DRG route rule. 
* `is_blackhole` - Indicates that if the next hop attachment does not exist, so traffic for this route is discarded without notification. 
* `is_conflict` - Indicates that the route was not imported due to a conflict between route rules. 
* `next_hop_drg_attachment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next hop DRG attachment responsible for reaching the network destination.

	A value of `BLACKHOLE` means traffic for this route is discarded without notification. 
* `route_provenance` - The earliest origin of a route. If a route is advertised to a DRG through an IPsec tunnel attachment, and is propagated to peered DRGs via RPC attachments, the route's provenance in the peered DRGs remains `IPSEC_TUNNEL`, because that is the earliest origin.

	No routes with a provenance `IPSEC_TUNNEL` or `VIRTUAL_CIRCUIT` will be exported to IPsec tunnel or virtual circuit attachments, regardless of the attachment's export distribution. 
* `route_type` - You can specify static routes for the DRG route table using the API. The DRG learns dynamic routes from the DRG attachments using various routing protocols. 

