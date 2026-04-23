// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

var (
	DataccInfrastructureScaleOptionSingularDataSourceRepresentation = map[string]interface{}{
		"infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${var.infrastructure_scale_option_infrastructure_id}`},
	}

	// DataccInfrastructureScaleOptionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Required, acctest.Create, DataccInfrastructureRepresentation)
)

// issue-routing-tag: datacc/default
func TestDataccInfrastructureScaleOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataccInfrastructureScaleOptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// override terraform-federation-test profile with our own user profile
	if overrideProfile := os.Getenv("datacc_custom_config_file_profile_override"); overrideProfile != "" {
		t.Setenv(globalvar.TfEnvPrefix+globalvar.ConfigFileProfileAttrName, overrideProfile)
		t.Setenv(globalvar.TfEnvPrefix+globalvar.AuthAttrName, "")
		t.Setenv(globalvar.AuthAttrName, globalvar.AuthSecurityToken)
	}

	const testResourceType = "infrastructure_scale_option"
	tfVariableStr := GenerateTFVariableStrings(testResourceType)
	getTFVar := func(variableName string) string {
		return os.Getenv(globalvar.TfEnvPrefix + testResourceType + "_" + variableName)
	}

	infrastructureId := getTFVar("infrastructure_id")

	singularDatasourceName := "data.oci_datacc_infrastructure_scale_option.test_infrastructure_scale_option"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + tfVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_infrastructure_scale_option", "test_infrastructure_scale_option", acctest.Required, acctest.Create, DataccInfrastructureScaleOptionSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "possible_ssd_configurations.#"),
			),
		},
	})
}
