---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_organization_subscriptions"
sidebar_current: "docs-oci-datasource-onesubscription-organization_subscriptions"
description: |-
  Provides the list of Organization Subscriptions in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_organization_subscriptions
This data source provides the list of Organization Subscriptions in Oracle Cloud Infrastructure Onesubscription service.

API that returns data for the list of subscription ids returned from Organizations API


## Example Usage

```hcl
data "oci_onesubscription_organization_subscriptions" "test_organization_subscriptions" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.


## Attributes Reference

The following attributes are exported:

* `organization_subscriptions` - The list of organization_subscriptions.

### OrganizationSubscription Reference

The following attributes are exported:

* `currency` - Currency details 
	* `iso_code` - Currency Code 
	* `name` - Currency name 
	* `std_precision` - Standard Precision of the Currency 
* `id` - SPM internal Subscription ID 
* `service_name` - Customer friendly service name provided by PRG 
* `status` - Status of the plan 
* `time_end` - Represents the date when the last service of the subscription ends 
* `time_start` - Represents the date when the first service of the subscription was activated 
* `total_value` - Total aggregate TCLV of all lines for the subscription including expired, active, and signed 
* `type` - Subscription Type i.e. IAAS,SAAS,PAAS 

