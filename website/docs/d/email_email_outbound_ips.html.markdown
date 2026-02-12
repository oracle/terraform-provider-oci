---
subcategory: "Email"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_email_email_outbound_ips"
sidebar_current: "docs-oci-datasource-email-email_outbound_ips"
description: |-
  Provides the list of Email Outbound Ips in Oracle Cloud Infrastructure Email service
---

# Data Source: oci_email_email_outbound_ips
This data source provides the list of Email Outbound Ips in Oracle Cloud Infrastructure Email service.

Returns a list of all Outbound Public IPs assigned for a given tenant.

## Example Usage

```hcl
data "oci_email_email_outbound_ips" "test_email_outbound_ips" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	assignment_state = var.email_outbound_ip_assignment_state
	outbound_ip = var.email_outbound_ip_outbound_ip
	state = var.email_outbound_ip_state
}
```

## Argument Reference

The following arguments are supported:

* `assignment_state` - (Optional) Filter returned list by specified assignment state. 
* `compartment_id` - (Required) The OCID for the compartment.
* `outbound_ip` - (Optional) The outbound IP address assigned to the tenancy.
* `state` - (Optional) Filter returned list by specified lifecycle state. This parameter is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `email_outbound_ip_collection` - The list of email_outbound_ip_collection.

### EmailOutboundIp Reference

The following attributes are exported:

* `items` - List of public IPs.
	* `assignment_state` - The assignment state of the public IP address.
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'DRAINING' state. 
	* `outbound_ip` - The public IP address assigned to the tenancy.
	* `state` - The current state of the Email Outbound Public IP.

