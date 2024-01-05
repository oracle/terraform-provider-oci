// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	GoldenGateDeploymentVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"deployment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.test_deployment.id}`},
		"deployment_type": acctest.Representation{RepType: acctest.Required, Create: `DATABASE_ORACLE`},
	}
)

/*
 Note:
 Set the following environmentVariables in order to make it work:
  TF_VAR_compartment_id=compartment - use any compartmentId what you have rights for
*/
// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentVersionResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID = "compartment_id"
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t)

	var (
		compartmentId = utils.GetEnvSettingWithBlankDefault(COMPARTMENT_ID)
	)

	datasourceName := "data.oci_golden_gate_deployment_versions.test_deployment_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_versions", "test_deployment_versions", acctest.Required, acctest.Create, GoldenGateDeploymentVersionDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "deployment_type", "DATABASE_ORACLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_version_collection.#"),
			),
		},
	})
}
