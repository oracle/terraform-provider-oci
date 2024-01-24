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
	DatascienceDatascienceNotebookSessionShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatascienceNotebookSessionShapeResourceConfig = ""
)

// issue-routing-tag: datascience/default
func TestDatascienceNotebookSessionShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_notebook_session_shapes.test_notebook_session_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_session_shapes", "test_notebook_session_shapes", acctest.Required, acctest.Create, DatascienceDatascienceNotebookSessionShapeDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceNotebookSessionShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.0.core_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.0.shape_series"),
			),
		},
	})
}
