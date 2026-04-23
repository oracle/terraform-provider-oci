// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceComputeTargetShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatascienceComputeTargetShapeResourceConfig = ""
)

// issue-routing-tag: datascience/default
func TestDatascienceComputeTargetShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceComputeTargetShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_compute_target_shapes.test_compute_target_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_compute_target_shapes", "test_compute_target_shapes", acctest.Required, acctest.Create, DatascienceComputeTargetShapeDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceComputeTargetShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "compute_target_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_target_shapes.0.core_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_target_shapes.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_target_shapes.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_target_shapes.0.shape_series"),
			),
		},
	})
}
