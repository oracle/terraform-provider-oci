// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
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
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"deployment_type": acctest.Representation{RepType: acctest.Optional, Create: `DATABASE_ORACLE`},
	}

	GoldenGateDeploymentTypeResourceConfig = ""
)

/*
 Note:
 Set the following environmentVariables in order to make it work:
  TF_VAR_compartment_id=com
*/
// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentTypeResource_basic")
	defer httpreplay.SaveScenario()

	const CompartmentId = "compartment_id"

	config := acctest.ProviderTestConfig() + makeVariableStr(CompartmentId, t)

	var compartmentId = utils.GetEnvSettingWithBlankDefault(CompartmentId)

	datasourceName := "data.oci_golden_gate_deployment_types.test_deployment_types"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_types", "test_deployment_types", acctest.Optional, acctest.Create, GoldenGateGoldenGateDeploymentTypeDataSourceRepresentation) +
				GoldenGateDeploymentTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_type_collection.0.items.0.ogg_version"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_type_collection.0.items.0.deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_type_collection.0.items.0.category", "DATA_REPLICATION"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_type_collection.#"),
			),
		},
	})
}
