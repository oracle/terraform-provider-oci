---
subcategory: "Jms Utils"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_utils_subscription_acknowledgment_configuration"
sidebar_current: "docs-oci-datasource-jms_utils-subscription_acknowledgment_configuration"
description: |-
  Provides details about a specific Subscription Acknowledgment Configuration in Oracle Cloud Infrastructure Jms Utils service
---

# Data Source: oci_jms_utils_subscription_acknowledgment_configuration
This data source provides details about a specific Subscription Acknowledgment Configuration resource in Oracle Cloud Infrastructure Jms Utils service.

Returns the configuration for subscription acknowledgment.


## Example Usage

```hcl
data "oci_jms_utils_subscription_acknowledgment_configuration" "test_subscription_acknowledgment_configuration" {

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `acknowledged_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal that ackwnoledged the subscription.
* `is_acknowledged` - Flag to determine whether the subscription was acknowledged or not.
* `time_acknowledged` - The date and time the subscription was acknowledged (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)). 

