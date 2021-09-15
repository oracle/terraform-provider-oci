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
	jobShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	JobShapeResourceConfig = ""
)

// issue-routing-tag: datascience/default
func TestDatascienceJobShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceJobShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_job_shapes.test_job_shapes"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_datascience_job_shapes", "test_job_shapes", Required, Create, jobShapeDataSourceRepresentation) +
					compartmentIdVariableStr + JobShapeResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
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
