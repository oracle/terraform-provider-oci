// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	MeteringComputationAverageCarbonEmissionSingularDataSourceRepresentation = map[string]interface{}{
		"sku_part_number": acctest.Representation{RepType: acctest.Required, Create: `B88317`},
	}

	MeteringComputationAverageCarbonEmissionResourceConfig = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationAverageCarbonEmissionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationAverageCarbonEmissionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_metering_computation_average_carbon_emission.test_average_carbon_emission"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_average_carbon_emission", "test_average_carbon_emission", acctest.Required, acctest.Create, MeteringComputationAverageCarbonEmissionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MeteringComputationAverageCarbonEmissionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "sku_part_number", "B88317"),
			),
		},
	})
}
