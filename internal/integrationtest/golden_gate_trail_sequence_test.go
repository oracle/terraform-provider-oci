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
	GoldenGateGoldenGateTrailSequenceSingularDataSourceRepresentation = map[string]interface{}{
		"deployment_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
		"trail_file_id":     acctest.Representation{RepType: acctest.Required, Create: `TI`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"trail_sequence_id": acctest.Representation{RepType: acctest.Required, Create: `2`},
	}

	GoldenGateGoldenGateTrailSequenceDataSourceRepresentation = map[string]interface{}{
		"deployment_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_deployment.id}`},
		"trail_file_id":     acctest.Representation{RepType: acctest.Required, Create: `TI`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"trail_sequence_id": acctest.Representation{RepType: acctest.Required, Create: `2`},
	}

	GoldenGateTrailSequenceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_trail_files", "test_trail_files", acctest.Required, acctest.Create, GoldenGateGoldenGateTrailFileDataSourceRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateTrailSequenceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateTrailSequenceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("TF_VAR_subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"test_subnet_id\" { default = \"%s\" }\n", subnetId)

	datasourceName := "data.oci_golden_gate_trail_sequences.test_trail_sequences"
	singularDatasourceName := "data.oci_golden_gate_trail_sequence.test_trail_sequence"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_trail_sequences", "test_trail_sequences", acctest.Required, acctest.Create, GoldenGateGoldenGateTrailSequenceDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + GoldenGateTrailSequenceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_file_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_sequence_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "trail_sequence_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "trail_sequence_collection.0", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_sequence_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_sequence_collection.0.items.0.sequence_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_sequence_collection.0.items.0.size_in_bytes"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_sequence_collection.0.items.0.time_last_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "trail_sequence_collection.0.time_last_fetched"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_trail_sequence", "test_trail_sequence", acctest.Required, acctest.Create, GoldenGateGoldenGateTrailSequenceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + GoldenGateTrailSequenceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trail_file_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trail_sequence_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_fetched"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.sequence_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.time_last_updated"),
			),
		},
	})
}
