// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OnesubscriptionOnesubscriptionInvoiceLineComputedUsageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"invoice_line_id": acctest.Representation{RepType: acctest.Required, Create: `${var.invoice_line_id}`},
		"fields":          acctest.Representation{RepType: acctest.Optional, Create: []string{`fields`}},
	}
)

// issue-routing-tag: onesubscription/default
func TestOnesubscriptionInvoiceLineComputedUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnesubscriptionInvoiceLineComputedUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	invoiceLineId := utils.GetEnvSettingWithBlankDefault("invoice_line_id")
	invoiceLineIdVariableStr := fmt.Sprintf("variable \"invoice_line_id\" { default = \"%s\" }\n", invoiceLineId)

	datasourceName := "data.oci_onesubscription_invoice_line_computed_usages.test_invoice_line_computed_usages"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_invoice_line_computed_usages", "test_invoice_line_computed_usages", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionInvoiceLineComputedUsageDataSourceRepresentation) +
				compartmentIdVariableStr + invoiceLineIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(datasourceName, "fields.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoice_line_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "invoiceline_computed_usages.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoiceline_computed_usages.0.cost"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoiceline_computed_usages.0.cost_rounded"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoiceline_computed_usages.0.net_unit_price"),
				resource.TestCheckResourceAttr(datasourceName, "invoiceline_computed_usages.0.parent_product.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "invoiceline_computed_usages.0.product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoiceline_computed_usages.0.quantity"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoiceline_computed_usages.0.time_metered_on"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoiceline_computed_usages.0.type"),
			),
		},
	})
}
