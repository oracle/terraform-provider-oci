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
	OnesubscriptionOnesubscriptionInvoiceDataSourceRepresentation = map[string]interface{}{
		"ar_customer_transaction_id": acctest.Representation{RepType: acctest.Required, Create: `${var.ar_customer_transaction_id}`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fields":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`fields`}},
		"time_from":                  acctest.Representation{RepType: acctest.Optional, Create: `timeFrom`},
		"time_to":                    acctest.Representation{RepType: acctest.Optional, Create: `timeTo`},
	}
)

// issue-routing-tag: onesubscription/default
func TestOnesubscriptionInvoiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnesubscriptionInvoiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	arCustomerTransactionId := utils.GetEnvSettingWithBlankDefault("ar_customer_transaction_id")
	arCustomerTransactionIdVariableStr := fmt.Sprintf("variable \"ar_customer_transaction_id\" { default = \"%s\" }\n", arCustomerTransactionId)

	datasourceName := "data.oci_onesubscription_invoices.test_invoices"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_invoices", "test_invoices", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionInvoiceDataSourceRepresentation) +
				compartmentIdVariableStr + arCustomerTransactionIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "ar_customer_transaction_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "invoices.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.ar_invoices"),
				resource.TestCheckResourceAttr(datasourceName, "invoices.0.bill_to_address.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "invoices.0.bill_to_contact.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "invoices.0.bill_to_customer.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "invoices.0.currency.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "invoices.0.invoice_lines.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "invoices.0.organization.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.payment_method"),
				resource.TestCheckResourceAttr(datasourceName, "invoices.0.payment_term.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.receipt_method"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.spm_invoice_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.subscription_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.time_invoice_date"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "invoices.0.updated_by"),
			),
		},
	})
}
