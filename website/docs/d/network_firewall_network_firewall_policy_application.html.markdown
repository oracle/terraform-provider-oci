---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_application"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_application"
description: |-
  Provides details about a specific Network Firewall Policy Application in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_application
This data source provides details about a specific Network Firewall Policy Application resource in Oracle Cloud Infrastructure Network Firewall service.

Get Application by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_application" "test_network_firewall_policy_application" {
	#Required
	application_name = var.network_firewall_policy_application_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `application_name` - (Required) Unique identifier for applications inside a policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `icmp_code` - The value of the ICMP/ICMP_V6 message Code (subtype) field as defined by [RFC 4443](https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
* `icmp_type` - The value of the ICMP/ICMP_V6 message Type field as defined by [RFC 4443](https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
* `name` - Name of the application.
* `parent_resource_id` - OCID of the Network Firewall Policy this application belongs to.
* `type` - Describes the type of application.

