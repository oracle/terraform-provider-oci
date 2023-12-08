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
	MeteringComputationCleanEnergyUsageSingularDataSourceRepresentation = map[string]interface{}{
		"region": acctest.Representation{RepType: acctest.Required, Create: `us-seattle-1`},
		"ad":     acctest.Representation{RepType: acctest.Optional, Create: ``},
	}

	MeteringComputationCleanEnergyUsageResourceConfig = ""
)

// issue-routing-tag: metering_computation/default
func TestMeteringComputationCleanEnergyUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMeteringComputationCleanEnergyUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_metering_computation_clean_energy_usage.test_clean_energy_usage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_metering_computation_clean_energy_usage", "test_clean_energy_usage", acctest.Required, acctest.Create, MeteringComputationCleanEnergyUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MeteringComputationCleanEnergyUsageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ad"),
				resource.TestCheckResourceAttr(singularDatasourceName, "region", "us-seattle-1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "usage", "80"),
			),
		},
	})
}
