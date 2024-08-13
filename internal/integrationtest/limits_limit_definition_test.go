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
	LimitsLimitDefinitionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `custom-image-count`},
		"service_name":   acctest.Representation{RepType: acctest.Required, Create: `compute`},
	}

	LimitsLimitDefinitionDataSourceRepresentationForSubscription = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":            acctest.Representation{RepType: acctest.Required, Create: subscriptionSupportedLimit},
		"service_name":    acctest.Representation{RepType: acctest.Required, Create: subscriptionSupportedService},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_ocid}`},
	}
)

// issue-routing-tag: limits/default
func TestLimitsLimitDefinitionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsLimitDefinitionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	//subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_ocid")
	//subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_ocid\" { default = \"%s\" }\n", subscriptionId)

	datasourceName := "data.oci_limits_limit_definitions.test_limit_definitions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_limit_definitions", "test_limit_definitions", acctest.Required, acctest.Create, LimitsLimitDefinitionDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "custom-image-count"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.are_quotas_supported"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_deprecated"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_dynamic"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_eligible_for_limit_increase"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_resource_availability_supported"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.scope_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.service_name"),
			),
		},
	})
}

func TestLimitsLimitDefinitionResource_subscription_support(t *testing.T) {
	httpreplay.SetScenario("TestLimitsLimitDefinitionResource_subscription_support")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	subscriptionOcid := utils.GetEnvSettingWithBlankDefault("subscription_ocid")
	subscriptionOcidVariableStr := fmt.Sprintf("variable \"subscription_ocid\" { default = \"%s\" }\n", subscriptionOcid)

	datasourceName := "data.oci_limits_limit_definitions.test_limit_definitions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_limit_definitions", "test_limit_definitions", acctest.Required, acctest.Create, LimitsLimitDefinitionDataSourceRepresentationForSubscription) +
				compartmentIdVariableStr + subscriptionOcidVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", subscriptionSupportedLimit),
				resource.TestCheckResourceAttr(datasourceName, "service_name", subscriptionSupportedService),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionOcid),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.are_quotas_supported"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_deprecated"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_dynamic"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_eligible_for_limit_increase"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.is_resource_availability_supported"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.scope_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_definitions.0.service_name"),
				resource.TestCheckResourceAttr(datasourceName, "limit_definitions.0.supported_subscriptions.#", "1"),
			),
		},
	})
}
