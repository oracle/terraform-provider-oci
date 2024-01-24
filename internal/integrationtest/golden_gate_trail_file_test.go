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
	GoldenGateGoldenGateTrailFileSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"trail_file_id": acctest.Representation{RepType: acctest.Required, Create: `T1`},
	}

	GoldenGateGoldenGateTrailFileDataSourceRepresentation = map[string]interface{}{
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"trail_file_id": acctest.Representation{RepType: acctest.Required, Create: `T1`},
	}

	GoldenGateTrailFileResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateTrailFileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateTrailFileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"test_subnet_id\" { default = \"%s\" }\n", subnetId)

	datasourceName := "data.oci_golden_gate_trail_files.test_trail_files"
	singularDatasourceName := "data.oci_golden_gate_trail_file.test_trail_file"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_trail_files", "test_trail_files", acctest.Required, acctest.Create, GoldenGateGoldenGateTrailFileDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + GoldenGateTrailFileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.consumers"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.max_sequence_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.min_sequence_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.number_of_sequences"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.producer"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.size_in_bytes"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.time_last_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.items.0.trail_file_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_collection.0.time_last_fetched"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_trail_file", "test_trail_file", acctest.Required, acctest.Create, GoldenGateGoldenGateTrailFileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + GoldenGateTrailFileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trail_file_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.0", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "items.0.consumers"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.max_sequence_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.min_sequence_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.number_of_sequences"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.producer"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.size_in_bytes"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.time_last_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.trail_file_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_fetched"),
			),
		},
	})
}
