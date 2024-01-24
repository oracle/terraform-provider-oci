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
	ContainerenginePodShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	ContainerenginePodShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: containerengine/default
func TestContainerenginePodShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerenginePodShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_pod_shapes.test_pod_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_pod_shapes", "test_pod_shapes", acctest.Required, acctest.Create, ContainerenginePodShapeDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerenginePodShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "pod_shapes.#"),
				resource.TestCheckResourceAttr(datasourceName, "pod_shapes.0.memory_options.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "pod_shapes.0.name"),
				resource.TestCheckResourceAttr(datasourceName, "pod_shapes.0.network_bandwidth_options.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "pod_shapes.0.ocpu_options.#", "0"),
			),
		},
	})
}
