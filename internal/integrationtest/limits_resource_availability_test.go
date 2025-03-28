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
	LimitsLimitsResourceAvailabilitySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"limit_name":          acctest.Representation{RepType: acctest.Required, Create: `adb-free-count`},
		"service_name":        acctest.Representation{RepType: acctest.Required, Create: `database`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	LimitsResourceAvailabilityResourceConfig = AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_limits_services", "test_services", acctest.Required, acctest.Create, LimitsServiceDataSourceRepresentation)
	LimitsLimitsResourceAvailabilitySingularDataSourceRepresentationForSubscriptionTest = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"service_name":    acctest.Representation{RepType: acctest.Required, Create: subscriptionSupportedService},
		"limit_name":      acctest.Representation{RepType: acctest.Required, Create: subscriptionSupportedLimit},
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_ocid}`},
	}
)

// issue-routing-tag: limits/default
func TestLimitsResourceAvailabilityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsResourceAvailabilityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	singularDatasourceName := "data.oci_limits_resource_availability.test_resource_availability"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_resource_availability", "test_resource_availability", acctest.Required, acctest.Create, LimitsLimitsResourceAvailabilitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + LimitsResourceAvailabilityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "limit_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fractional_availability"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fractional_usage"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "used"),
			),
		},
	})
}

func TestLimitsResourceAvailabilityResource_subscription(t *testing.T) {
	httpreplay.SetScenario("TestLimitsResourceAvailabilityResource_subscription")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	subscriptionOcid := utils.GetEnvSettingWithBlankDefault("subscription_ocid")
	subscriptionOcidVariableStr := fmt.Sprintf("variable \"subscription_ocid\" { default = \"%s\" }\n", subscriptionOcid)

	singularDatasourceName := "data.oci_limits_resource_availability.test_resource_availability"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_resource_availability", "test_resource_availability", acctest.Required, acctest.Create, LimitsLimitsResourceAvailabilitySingularDataSourceRepresentationForSubscriptionTest) +
				compartmentIdVariableStr + LimitsResourceAvailabilityResourceConfig + subscriptionOcidVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "limit_name", subscriptionSupportedLimit),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_name", subscriptionSupportedService),
				resource.TestCheckResourceAttr(singularDatasourceName, "subscription_id", subscriptionOcid),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fractional_availability"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fractional_usage"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "used"),
			),
		},
	})
}
