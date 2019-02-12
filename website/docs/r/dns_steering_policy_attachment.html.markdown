---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_steering_policy_attachment"
sidebar_current: "docs-oci-resource-dns-steering_policy_attachment"
description: |-
  Provides the Steering Policy Attachment resource in Oracle Cloud Infrastructure Dns service
---

# oci_dns_steering_policy_attachment
This resource provides the Steering Policy Attachment resource in Oracle Cloud Infrastructure Dns service.

Creates a new attachment between a steering policy and a domain.
For the purposes of access control, the attachment is automatically placed
into the same compartment as the containing zone of the domain.


## Example Usage

```hcl
resource "oci_dns_steering_policy_attachment" "test_steering_policy_attachment" {
	#Required
	domain_name = "${var.steering_policy_attachment_domain_name}"
	steering_policy_id = "${oci_dns_steering_policy.test_steering_policy.id}"
	zone_id = "${oci_dns_zone.test_zone.id}"

	#Optional
	display_name = "${var.steering_policy_attachment_display_name}"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) (Updatable) A user-friendly name for the steering policy attachment. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `domain_name` - (Required) The attached domain within the attached zone.
* `steering_policy_id` - (Required) The OCID of the attached steering policy.
* `zone_id` - (Required) The OCID of the attached zone.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the steering policy attachment.
* `display_name` - A user-friendly name for the steering policy attachment. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `domain_name` - The attached domain within the attached zone.
* `id` - The OCID of the resource.
* `rtypes` - The record types covered by the attachment at the domain. The set of record types is determined by aggregating the record types from the answers defined in the steering policy. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `steering_policy_id` - The OCID of the attached steering policy.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `zone_id` - The OCID of the attached zone.

## Import

SteeringPolicyAttachments can be imported using the `id`, e.g.

```
$ terraform import oci_dns_steering_policy_attachment.test_steering_policy_attachment "id"
```

