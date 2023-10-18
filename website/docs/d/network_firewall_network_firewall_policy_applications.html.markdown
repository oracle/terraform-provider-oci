---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_applications"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_applications"
description: |-
  Provides the list of Network Firewall Policy Applications in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_applications
This data source provides the list of Network Firewall Policy Applications in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of Applications for the policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_applications" "test_network_firewall_policy_applications" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_application_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `application_summary_collection` - The list of application_summary_collection.

### NetworkFirewallPolicyApplication Reference

The following attributes are exported:

* `icmp_code` - The value of the ICMP6 message Code (subtype) field as defined by [RFC 4443](https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
* `icmp_type` - The value of the ICMP6 message Type field as defined by [RFC 4443](https://www.rfc-editor.org/rfc/rfc4443.html#section-2.1).
* `name` - Name of the application.
* `parent_resource_id` - OCID of the Network Firewall Policy this application belongs to.
* `type` - Describes the type of Application.

