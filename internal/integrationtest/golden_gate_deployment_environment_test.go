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
	GoldenGateDeploymentEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	GoldenGateDeploymentEnvironmentResourceConfig = ""
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentEnvironmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentEnvironmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_deployment_environments.test_deployment_environments"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_environments", "test_deployment_environments", acctest.Required, acctest.Create, GoldenGateDeploymentEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGateDeploymentEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "deployment_environment_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_environment_collection.0.items.#", "6"),
			),
		},
	})
}
