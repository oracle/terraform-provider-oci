// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package usage_proxy

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_usage_proxy_resource_quotas", UsageProxyResourceQuotasDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_resources", UsageProxyResourcesDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_subscription_product", UsageProxySubscriptionProductDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_subscription_products", UsageProxySubscriptionProductsDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_subscription_redeemable_user", UsageProxySubscriptionRedeemableUserDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_subscription_redeemable_users", UsageProxySubscriptionRedeemableUsersDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_subscription_redemption", UsageProxySubscriptionRedemptionDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_subscription_redemptions", UsageProxySubscriptionRedemptionsDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_subscription_reward", UsageProxySubscriptionRewardDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_subscription_rewards", UsageProxySubscriptionRewardsDataSource())
	tfresource.RegisterDatasource("oci_usage_proxy_usagelimits", UsageProxyUsagelimitsDataSource())
}
