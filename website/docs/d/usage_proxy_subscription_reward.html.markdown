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

Returns the list of rewards for a subscription ID.


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

* `subscription_id` - (Required) The subscription ID for which rewards information is requested for.
* `tenancy_id` - (Required) The OCID of the tenancy.


## Attributes Reference

The following attributes are exported:

* `items` - The monthly summary of rewards.
	* `available_rewards` - The number of rewards available for a specific usage period.
	* `earned_rewards` - The number of rewards earned for the specific usage period.
	* `eligible_usage_amount` - The eligible usage amount for the usage period. 
	* `ineligible_usage_amount` - The ineligible usage amount for the usage period. 
	* `is_manual` - The boolean parameter to indicate whether or not the available rewards are manually posted.
	* `redeemed_rewards` - The number of rewards redeemed for a specific month.
	* `time_rewards_earned` - The date and time when rewards accrue. 
	* `time_rewards_expired` - The date and time when rewards expire.
	* `time_usage_ended` - The end date and time for the usage period. 
	* `time_usage_started` - The start date and time for the usage period. 
	* `usage_amount` - The usage amount for the usage period. 
	* `usage_period_key` - The usage period ID. 
* `summary` - The overall monthly reward summary.
	* `currency` - The currency unit for the reward amount.
	* `redemption_code` - The redemption code used in the Billing Center during the reward redemption process.
	* `rewards_rate` - The current Rewards percentage in decimal format.
	* `subscription_id` - The entitlement ID from MQS, which is the same as the subcription ID.
	* `tenancy_id` - The OCID of the target tenancy.
	* `total_rewards_available` - The total number of available rewards for a given subscription ID.

