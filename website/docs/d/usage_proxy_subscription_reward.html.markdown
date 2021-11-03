---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_subscription_reward"
sidebar_current: "docs-oci-datasource-usage_proxy-subscription_reward"
description: |-
  Provides details about a specific Subscription Reward in Oracle Cloud Infrastructure Usage Proxy service
---

# Data Source: oci_usage_proxy_subscription_reward
This data source provides details about a specific Subscription Reward resource in Oracle Cloud Infrastructure Usage Proxy service.

This API returns list of rewards for a subscription Id


## Example Usage

```hcl
data "oci_usage_proxy_subscription_reward" "test_subscription_reward" {
	#Required
	subscription_id = oci_ons_subscription.test_subscription.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (Required) The subscriptionId for which rewards information is requested for.
* `tenancy_id` - (Required) The OCID of the tenancy.


## Attributes Reference

The following attributes are exported:

* `items` - The monthly summary of rewards.
	* `available_rewards` - The number of rewards available for a specific usage period.
	* `earned_rewards` - The number of rewards earned for the specific usage period.
	* `eligible_usage_amount` - The eligible usage amount for the usage period. 
	* `ineligible_usage_amount` - The in eligible usage amount for the usage period. 
	* `is_manual` - The boolean flag to tell if the available rewards are posted manually or not.
	* `redeemed_rewards` - The number of rewards redeemed for a specific month.
	* `time_rewards_earned` - The date and time on which rewards are accrued. 
	* `time_rewards_expired` - The date and time on which rewards are expired.
	* `time_usage_ended` - The end date and time for the usage period. 
	* `time_usage_started` - The start date and time for the usage period. 
	* `usage_amount` - The usage amount for the usage period. 
	* `usage_period_key` - The id for the usage period. 
* `summary` - The overrall reward summary of the monthly summary rewards.
	* `currency` - The currency unit for the reward amount.
	* `rewards_rate` - The current Rewards percentage in decimal format.
	* `subscription_id` - The entitlement id from MQS and it is same as subcription id.
	* `tenancy_id` - The OCID of the target tenancy.
	* `total_rewards_available` - The total number of available rewards for a given subscription Id.

