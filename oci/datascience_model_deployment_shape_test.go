// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	modelDeploymentShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	ModelDeploymentShapeResourceConfig = ""
)

// issue-routing-tag: datascience/default
func TestDatascienceModelDeploymentShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_model_deployment_shapes.test_model_deployment_shapes"

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployment_shapes", "test_model_deployment_shapes", Required, Create, modelDeploymentShapeDataSourceRepresentation) +
				compartmentIdVariableStr + ModelDeploymentShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.core_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.name"),
			),
		},
	})
}
