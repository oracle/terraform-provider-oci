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
	OspGatewayOspGatewayInvoicesInvoiceLineDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"internal_invoice_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_osp_gateway_invoices.test_invoices.invoice_collection.0.items[3], "internal_invoice_id")}`},
		"osp_home_region":     acctest.Representation{RepType: acctest.Required, Create: `${var.home_region}`},
	}

	OspGatewayInvoicesInvoiceLineResourceConfig = ""
)

// issue-routing-tag: osp_gateway/default
func TestOspGatewayInvoicesInvoiceLineResource_basic(t *testing.T) {
	t.Skip("Invoice tests are not supported due to test resource unavailability.")

	httpreplay.SetScenario("TestOspGatewayInvoicesInvoiceLineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	homeRegion := utils.GetEnvSettingWithBlankDefault("region")
	regionVariableStr := fmt.Sprintf("variable \"home_region\" { default = \"%s\" }\n", homeRegion)

	datasourceName := "data.oci_osp_gateway_invoices_invoice_lines.test_invoices_invoice_lines"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_invoices", "test_invoices", acctest.Required, acctest.Create, OspGatewayOspGatewayInvoiceDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osp_gateway_invoices_invoice_lines", "test_invoices_invoice_lines", acctest.Required, acctest.Create, OspGatewayOspGatewayInvoicesInvoiceLineDataSourceRepresentation) +
				compartmentIdVariableStr + regionVariableStr + OspGatewayInvoicesInvoiceLineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "internal_invoice_id"),
				resource.TestCheckResourceAttr(datasourceName, "osp_home_region", homeRegion),

				resource.TestCheckResourceAttrSet(datasourceName, "invoice_line_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "invoice_line_collection.0.items.#", "1"),
			),
		},
	})
}
