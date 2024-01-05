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
	OnesubscriptionOnesubscriptionBillingScheduleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"subscribed_service_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.subscribed_service_id}`},
	}

	OnesubscriptionBillingScheduleResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_subscribed_services", "test_subscribed_services", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionSubscribedServiceDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_subscriptions", "test_subscriptions", acctest.Required, acctest.Create, OnsOnsSubscriptionDataSourceRepresentation)
)

// issue-routing-tag: onesubscription/default
func TestOnesubscriptionBillingScheduleResource_basic(t *testing.T) {
	t.Skip("Skipping Test:TestOnesubscriptionBillingScheduleResource_basic")
	httpreplay.SetScenario("TestOnesubscriptionBillingScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_onesubscription_billing_schedules.test_billing_schedules"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_billing_schedules", "test_billing_schedules", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionBillingScheduleDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.amount"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.ar_customer_transaction_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.ar_invoice_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.billing_frequency"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.invoice_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.net_unit_price"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.order_number"),
				resource.TestCheckResourceAttr(datasourceName, "billing_schedules.0.product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.quantity"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.subscribed_service_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.time_invoicing"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.time_start"),
			),
		},
	})
}
