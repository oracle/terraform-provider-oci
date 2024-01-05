// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package onesubscription

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_onesubscription_aggregated_computed_usages", OnesubscriptionAggregatedComputedUsagesDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_billing_schedules", OnesubscriptionBillingSchedulesDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_commitment", OnesubscriptionCommitmentDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_commitments", OnesubscriptionCommitmentsDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_computed_usage", OnesubscriptionComputedUsageDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_computed_usages", OnesubscriptionComputedUsagesDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_invoice_line_computed_usages", OnesubscriptionInvoiceLineComputedUsagesDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_invoices", OnesubscriptionInvoicesDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_organization_subscriptions", OnesubscriptionOrganizationSubscriptionsDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_ratecards", OnesubscriptionRatecardsDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_subscribed_service", OnesubscriptionSubscribedServiceDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_subscribed_services", OnesubscriptionSubscribedServicesDataSource())
	tfresource.RegisterDatasource("oci_onesubscription_subscriptions", OnesubscriptionSubscriptionsDataSource())
}
