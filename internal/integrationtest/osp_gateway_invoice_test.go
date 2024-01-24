// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OspGatewayOspGatewayInvoiceSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"internal_invoice_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_osp_gateway_invoices.test_invoices.invoice_collection.0.items[3], "internal_invoice_id")}`},
		"osp_home_region":     acctest.Representation{RepType: acctest.Required, Create: `${var.home_region}`},
	}

	OspGatewayOspGatewayInvoiceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"osp_home_region": acctest.Representation{RepType: acctest.Required, Create: `${var.home_region}`},
	}

	OspGatewayInvoiceResourceConfig = ""
)

// issue-routing-tag: osp_gateway/default
func TestOspGatewayInvoiceResource_basic(t *testing.T) {
	t.Skip("Invoice tests are not supported due to test resource unavailability.")

	httpreplay.SetScenario("TestOspGatewayInvoiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	homeRegion := utils.GetEnvSettingWithBlankDefault("region")
	regionVariableStr := fmt.Sprintf("variable \"home_region\" { default = \"%s\" }\n", homeRegion)

	datasourceName := "data.oci_osp_gateway_invoices.test_invoices"
	singularDatasourceName := "data.oci_osp_gateway_invoice.test_invoice"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_invoices", "test_invoices", acctest.Required, acctest.Create, OspGatewayOspGatewayInvoiceDataSourceRepresentation) +
				compartmentIdVariableStr + regionVariableStr + OspGatewayInvoiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "osp_home_region", homeRegion),

				resource.TestCheckResourceAttrSet(datasourceName, "invoice_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_invoices", "test_invoices", acctest.Required, acctest.Create, OspGatewayOspGatewayInvoiceDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_invoice", "test_invoice", acctest.Required, acctest.Create, OspGatewayOspGatewayInvoiceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + regionVariableStr + OspGatewayInvoiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "internal_invoice_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "osp_home_region", homeRegion),

				resource.TestCheckResourceAttr(singularDatasourceName, "bill_to_address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "currency.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "internal_invoice_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_amount"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_amount_adjusted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_amount_applied"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_amount_credited"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_amount_due"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_ref_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "invoice_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_credit_card_payable"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_display_download_pdf"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_payable"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_pdf_email_available"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_payment_detail.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "payment_terms"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "preferred_email"),
				resource.TestCheckResourceAttr(singularDatasourceName, "subscription_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tax"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_invoice"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_invoice_due"),
			),
		},
	})
}
