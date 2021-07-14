// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	modelDeploymentShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	ModelDeploymentShapeResourceConfig = ""
)

func TestDatascienceModelDeploymentShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_model_deployment_shapes.test_model_deployment_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_datascience_model_deployment_shapes", "test_model_deployment_shapes", Required, Create, modelDeploymentShapeDataSourceRepresentation) +
					compartmentIdVariableStr + ModelDeploymentShapeResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "model_deployment_shapes.0.name"),
				),
			},
		},
	})
}
