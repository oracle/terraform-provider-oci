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
	PsqlShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `PostgreSQL.VM.Standard.E4.Flex.2.32GB`},
	}
)

// issue-routing-tag: psql/default
func TestPsqlShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_psql_shapes.test_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_shapes", "test_shapes", acctest.Optional, acctest.Create, PsqlShapeDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "shape_collection.0.items.0.id", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttrSet(datasourceName, "shape_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "shape_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "shape_collection.0.items.0.shape_memory_options.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "shape_collection.0.items.0.shape_ocpu_options.#"),
			),
		},
	})
}
