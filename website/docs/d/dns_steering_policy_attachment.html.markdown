---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_steering_policy_attachment"
sidebar_current: "docs-oci-datasource-dns-steering_policy_attachment"
description: |-
  Provides details about a specific Steering Policy Attachment in Oracle Cloud Infrastructure Dns service
---

# Data Source: oci_dns_steering_policy_attachment
This data source provides details about a specific Steering Policy Attachment resource in Oracle Cloud Infrastructure Dns service.

Gets information about the specified steering policy attachment.


## Example Usage

```hcl
data "oci_dns_steering_policy_attachment" "test_steering_policy_attachment" {
	#Required
	steering_policy_attachment_id = "${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}"
}
```

## Argument Reference

The following arguments are supported:

* `steering_policy_attachment_id` - (Required) The OCID of the target steering policy attachment.


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

