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
	OnesubscriptionOnesubscriptionSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"buyer_email":             acctest.Representation{RepType: acctest.Optional, Create: `buyerEmail`},
		"is_commit_info_required": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"plan_number":             acctest.Representation{RepType: acctest.Optional, Create: `planNumber`},
		"subscription_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
	}
)

// issue-routing-tag: onesubscription/default
func TestOnesubscriptionSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnesubscriptionSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_onesubscription_subscriptions.test_subscriptions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_subscriptions", "test_subscriptions", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr, /*+ oneRegionVariableStr*/
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.#"),
				resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.currency.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.service_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.status"),
				resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.subscribed_services.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.time_start"),
			),
		},
	})
}
