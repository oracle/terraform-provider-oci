// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatascienceDatascienceJobShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatascienceJobShapeResourceConfig = ""
)

// issue-routing-tag: datascience/default
func TestDatascienceJobShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceJobShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_job_shapes.test_job_shapes"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_job_shapes", "test_job_shapes", acctest.Required, acctest.Create, DatascienceDatascienceJobShapeDataSourceRepresentation) +
					compartmentIdVariableStr + DatascienceJobShapeResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "job_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_shapes.0.core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_shapes.0.memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_shapes.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_shapes.0.shape_series"),
				),
			},
		},
	})
}
