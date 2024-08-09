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
	LimitsServiceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	subscriptionSupportedService                             = "limits-testing"
	subscriptionSupportedLimit                               = "limits-multisub-dev-count"
	LimitsServiceDataSourceRepresentationForSubscriptionTest = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_ocid}`},
	}
)

// issue-routing-tag: limits/default
func TestLimitsServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_limits_services.test_services"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_services", "test_services", acctest.Required, acctest.Create, LimitsServiceDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "services.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.name"),
			),
		},
	})
}

func TestLimitsServiceResource_with_subscription(t *testing.T) {
	httpreplay.SetScenario("TestLimitsServiceResource_with_subscription")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	subscriptionOcid := utils.GetEnvSettingWithBlankDefault("subscription_ocid")
	subscriptionOcidVariableStr := fmt.Sprintf("variable \"subscription_ocid\" { default = \"%s\" }\n", subscriptionOcid)

	datasourceName := "data.oci_limits_services.test_services"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_services", "test_services", acctest.Required, acctest.Create, LimitsServiceDataSourceRepresentationForSubscriptionTest) +
				compartmentIdVariableStr + subscriptionOcidVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionOcid),
				resource.TestCheckResourceAttrSet(datasourceName, "services.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "services.0.name"),
				resource.TestCheckResourceAttr(datasourceName, "services.0.supported_subscriptions.#", "1"),
			),
		},
	})
}
