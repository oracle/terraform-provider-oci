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
	MeteringComputationUsageCarbonEmissionsConfigSingularDataSourceRepresentation = map[string]interface{}{
		"tenant_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
	}

	MeteringComputationUsageCarbonEmissionsConfigResourceConfig = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationUsageCarbonEmissionsConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationUsageCarbonEmissionsConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)

	singularDatasourceName := "data.oci_metering_computation_usage_carbon_emissions_config.test_usage_carbon_emissions_config"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + tenancyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_usage_carbon_emissions_config", "test_usage_carbon_emissions_config", acctest.Required, acctest.Create, MeteringComputationUsageCarbonEmissionsConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MeteringComputationUsageCarbonEmissionsConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "22"),
			),
		},
	})
}
