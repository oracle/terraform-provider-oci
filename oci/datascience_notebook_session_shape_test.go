// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	notebookSessionShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	NotebookSessionShapeResourceConfig = ""
)

func TestDatascienceNotebookSessionShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_notebook_session_shapes.test_notebook_session_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_datascience_notebook_session_shapes", "test_notebook_session_shapes", Required, Create, notebookSessionShapeDataSourceRepresentation) +
					compartmentIdVariableStr + NotebookSessionShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.0.core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_session_shapes.0.name"),
				),
			},
		},
	})
}
