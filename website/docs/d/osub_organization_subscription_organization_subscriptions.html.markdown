---
subcategory: "Osub Organization Subscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osub_organization_subscription_organization_subscriptions"
sidebar_current: "docs-oci-datasource-osub_organization_subscription-organization_subscriptions"
description: |-
  Provides the list of Organization Subscriptions in Oracle Cloud Infrastructure Osub Organization Subscription service
---

# Data Source: oci_osub_organization_subscription_organization_subscriptions
This data source provides the list of Organization Subscriptions in Oracle Cloud Infrastructure Osub Organization Subscription service.

API that returns data for the list of subscription ids returned from Organizations API


## Example Usage

```hcl
data "oci_osub_organization_subscription_organization_subscriptions" "test_organization_subscriptions" {
	#Required
	compartment_id = var.compartment_id
	subscription_ids = var.organization_subscription_subscription_ids

	#Optional
	x_one_origin_region = var.organization_subscription_x_one_origin_region
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `subscription_ids` - (Required) Comma separated list of subscription ids, pass "DUMMY" as value
* `x_one_origin_region` - (Optional) The Oracle Cloud Infrastructure home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc. 


## Attributes Reference

The following attributes are exported:

* `subscriptions` - The list of subscriptions.

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

