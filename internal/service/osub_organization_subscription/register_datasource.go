// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_organization_subscription

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_osub_organization_subscription_organization_subscriptions", OsubOrganizationSubscriptionOrganizationSubscriptionsDataSource())
}
