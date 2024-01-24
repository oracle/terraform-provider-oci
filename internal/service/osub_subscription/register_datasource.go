// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_subscription

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_osub_subscription_commitment", OsubSubscriptionCommitmentDataSource())
	tfresource.RegisterDatasource("oci_osub_subscription_commitments", OsubSubscriptionCommitmentsDataSource())
	tfresource.RegisterDatasource("oci_osub_subscription_ratecards", OsubSubscriptionRatecardsDataSource())
	tfresource.RegisterDatasource("oci_osub_subscription_subscriptions", OsubSubscriptionSubscriptionsDataSource())
}
