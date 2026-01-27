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
	ManagedKafkaNodeShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.A1.Flex`},
	}
	ManagedKafkaNodeShapeResourceConfig = ""
)

// issue-routing-tag: managed_kafka/default
func TestManagedKafkaNodeShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagedKafkaNodeShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_managed_kafka_node_shapes.test_node_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_managed_kafka_node_shapes", "test_node_shapes", acctest.Optional, acctest.Create, ManagedKafkaNodeShapeDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedKafkaNodeShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "VM.Standard.A1.Flex"),

				resource.TestCheckResourceAttrSet(datasourceName, "node_shape_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "node_shape_collection.0.items.#", "2"),
			),
		},
	})
}
