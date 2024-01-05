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
	OsubBillingScheduleOsubBillingScheduleBillingScheduleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subscription_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"x_one_origin_region": acctest.Representation{RepType: acctest.Required, Create: `${var.x_one_origin_region}`},
	}
)

// issue-routing-tag: osub_billing_schedule/default
func TestOsubBillingScheduleBillingScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsubBillingScheduleBillingScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	oneRegion := utils.GetEnvSettingWithBlankDefault("region")
	oneRegionVariableStr := fmt.Sprintf("variable \"x_one_origin_region\" { default = \"%s\" }\n", oneRegion)

	datasourceName := "data.oci_osub_billing_schedule_billing_schedules.test_billing_schedules"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osub_billing_schedule_billing_schedules", "test_billing_schedules", acctest.Required, acctest.Create, OsubBillingScheduleOsubBillingScheduleBillingScheduleDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + oneRegionVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "x_one_origin_region"),

				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.amount"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.billing_frequency"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.invoice_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.net_unit_price"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.order_number"),
				resource.TestCheckResourceAttr(datasourceName, "billing_schedules.0.product.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.quantity"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.time_invoicing"),
				resource.TestCheckResourceAttrSet(datasourceName, "billing_schedules.0.time_start"),
			),
		},
	})
}
