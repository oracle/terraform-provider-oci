// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_osp_gateway_address", OspGatewayAddressDataSource())
	tfresource.RegisterDatasource("oci_osp_gateway_address_rule", OspGatewayAddressRuleDataSource())
	tfresource.RegisterDatasource("oci_osp_gateway_invoice", OspGatewayInvoiceDataSource())
	tfresource.RegisterDatasource("oci_osp_gateway_invoices", OspGatewayInvoicesDataSource())
	tfresource.RegisterDatasource("oci_osp_gateway_invoices_invoice_line", OspGatewayInvoicesInvoiceLineDataSource())
	tfresource.RegisterDatasource("oci_osp_gateway_invoices_invoice_lines", OspGatewayInvoicesInvoiceLinesDataSource())
	tfresource.RegisterDatasource("oci_osp_gateway_subscription", OspGatewaySubscriptionDataSource())
	tfresource.RegisterDatasource("oci_osp_gateway_subscriptions", OspGatewaySubscriptionsDataSource())
}
