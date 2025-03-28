// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LimitsLimitsLimitValueDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"service_name":        acctest.Representation{RepType: acctest.Required, Create: `compute`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `custom-image-count`},
		"scope_type":          acctest.Representation{RepType: acctest.Optional, Create: `AD`},
	}

	LimitsLimitValueResourceConfig = AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_limits_services", "test_services", acctest.Required, acctest.Create, LimitsServiceDataSourceRepresentation)

	LimitsLimitsLimitValueDataSourceRepresentationForSubscriptionTest = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"service_name":    acctest.Representation{RepType: acctest.Required, Create: subscriptionSupportedService},
		"name":            acctest.Representation{RepType: acctest.Required, Create: subscriptionSupportedLimit},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_ocid}`},
	}
)

// issue-routing-tag: limits/default
func TestLimitsLimitValueResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsLimitValueResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_limits_limit_values.test_limit_values"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_limit_values", "test_limit_values", acctest.Required, acctest.Create, LimitsLimitsLimitValueDataSourceRepresentation) +
				compartmentIdVariableStr + LimitsLimitValueResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "custom-image-count"),
				resource.TestCheckResourceAttr(datasourceName, "service_name", "compute"),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_values.#"),
			),
		},
	})
}

func TestLimitsLimitValueResource_subscription(t *testing.T) {
	httpreplay.SetScenario("TestLimitsLimitValueResource_subscription")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	subscriptionOcid := utils.GetEnvSettingWithBlankDefault("subscription_ocid")
	subscriptionOcidVariableStr := fmt.Sprintf("variable \"subscription_ocid\" { default = \"%s\" }\n", subscriptionOcid)

	datasourceName := "data.oci_limits_limit_values.test_limit_values"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_limit_values", "test_limit_values", acctest.Required, acctest.Create, LimitsLimitsLimitValueDataSourceRepresentationForSubscriptionTest) +
				compartmentIdVariableStr + subscriptionOcidVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", subscriptionSupportedLimit),
				resource.TestCheckResourceAttr(datasourceName, "service_name", subscriptionSupportedService),
				resource.TestCheckResourceAttr(datasourceName, "subscription_id", subscriptionOcid),
				resource.TestCheckResourceAttrSet(datasourceName, "limit_values.#"),
			),
		},
	})
}
