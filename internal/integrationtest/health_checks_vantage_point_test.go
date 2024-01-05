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
	HealthChecksHealthChecksVantagePointDataSourceRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `AWS Asia Pacific South 1`},
		"name":         acctest.Representation{RepType: acctest.Optional, Create: `aws-bom`},
	}

	HealthChecksVantagePointResourceConfig = ""
)

// issue-routing-tag: health_checks/default
func TestHealthChecksVantagePointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksVantagePointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_health_checks_vantage_points.test_vantage_points"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_health_checks_vantage_points", "test_vantage_points", acctest.Optional, acctest.Create, HealthChecksHealthChecksVantagePointDataSourceRepresentation) +
				compartmentIdVariableStr + HealthChecksVantagePointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "AWS Asia Pacific South 1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "aws-bom"),

				resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.0.display_name"),
				resource.TestCheckResourceAttr(datasourceName, "health_checks_vantage_points.0.geo.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.0.provider_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.0.routing.#"),
			),
		},
	})
}
