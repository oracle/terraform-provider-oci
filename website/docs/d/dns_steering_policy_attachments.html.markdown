---
subcategory: "Dns"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_steering_policy_attachments"
sidebar_current: "docs-oci-datasource-dns-steering_policy_attachments"
description: |-
  Provides the list of Steering Policy Attachments in Oracle Cloud Infrastructure Dns service
---

# Data Source: oci_dns_steering_policy_attachments
This data source provides the list of Steering Policy Attachments in Oracle Cloud Infrastructure Dns service.

Lists the steering policy attachments in the specified compartment.


## Example Usage

```hcl
data "oci_dns_steering_policy_attachments" "test_steering_policy_attachments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.steering_policy_attachment_display_name
	domain = var.steering_policy_attachment_domain
	domain_contains = var.steering_policy_attachment_domain_contains
	id = var.steering_policy_attachment_id
	state = var.steering_policy_attachment_state
	steering_policy_id = oci_dns_steering_policy.test_steering_policy.id
	time_created_greater_than_or_equal_to = var.steering_policy_attachment_time_created_greater_than_or_equal_to
	time_created_less_than = var.steering_policy_attachment_time_created_less_than
	zone_id = oci_dns_zone.test_zone.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment the resource belongs to.
* `display_name` - (Optional) The displayName of a resource.
* `domain` - (Optional) Search by domain. Will match any record whose domain (case-insensitive) equals the provided value. 
* `domain_contains` - (Optional) Search by domain. Will match any record whose domain (case-insensitive) contains the provided value. 
* `id` - (Optional) The OCID of a resource.
* `state` - (Optional) The state of a resource.
* `steering_policy_id` - (Optional) Search by steering policy OCID. Will match any resource whose steering policy ID matches the provided value. 
* `time_created_greater_than_or_equal_to` - (Optional) An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created on or after the indicated time. 
* `time_created_less_than` - (Optional) An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created before the indicated time. 
* `zone_id` - (Optional) Search by zone OCID. Will match any resource whose zone ID matches the provided value. 


## Attributes Reference

The following attributes are exported:

* `steering_policy_attachments` - The list of steering_policy_attachments.

### SteeringPolicyAttachment Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the steering policy attachment.
* `display_name` - A user-friendly name for the steering policy attachment. Does not have to be unique and can be changed. Avoid entering confidential information. 
* `domain_name` - The attached domain within the attached zone.
* `id` - The OCID of the resource.
* `rtypes` - The record types covered by the attachment at the domain. The set of record types is determined by aggregating the record types from the answers defined in the steering policy. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `steering_policy_id` - The OCID of the attached steering policy.
* `time_created` - The date and time the resource was created, expressed in RFC 3339 timestamp format.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `zone_id` - The OCID of the attached zone.

