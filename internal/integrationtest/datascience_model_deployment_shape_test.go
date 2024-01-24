// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatascienceDatascienceModelDeploymentShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatascienceModelDeploymentShapeResourceConfig = ""
)

// issue-routing-tag: datascience/default
func TestDatascienceModelDeploymentShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_model_deployment_shapes.test_model_deployment_shapes"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployment_shapes", "test_model_deployment_shapes", acctest.Required, acctest.Create, DatascienceDatascienceModelDeploymentShapeDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelDeploymentShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.core_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.shape_series"),
			),
		},
	})
}
