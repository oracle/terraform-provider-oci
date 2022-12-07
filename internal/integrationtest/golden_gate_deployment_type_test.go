// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	GoldenGateGoldenGateDeploymentTypeSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`},
	}

	GoldenGateGoldenGateDeploymentTypeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`},
	}

	GoldenGateDeploymentTypeResourceConfig = ""
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_deployment_types.test_deployment_types"
	singularDatasourceName := "data.oci_golden_gate_deployment_type.test_deployment_type"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_types", "test_deployment_types", acctest.Required, acctest.Create, GoldenGateGoldenGateDeploymentTypeDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGateDeploymentTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),

				resource.TestCheckResourceAttrSet(datasourceName, "deployment_type_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_type_collection.0.items.#", "3"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_type", "test_deployment_type", acctest.Required, acctest.Create, GoldenGateGoldenGateDeploymentTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGateDeploymentTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "3"),
			),
		},
	})
}
