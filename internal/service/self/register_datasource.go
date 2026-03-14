// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package self

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_self_partner_subscriptions", SelfPartnerSubscriptionsDataSource())
	tfresource.RegisterDatasource("oci_self_subscription", SelfSubscriptionDataSource())
	tfresource.RegisterDatasource("oci_self_subscription_token", SelfSubscriptionTokenDataSource())
	tfresource.RegisterDatasource("oci_self_subscriptions", SelfSubscriptionsDataSource())
}
