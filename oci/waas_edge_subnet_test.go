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
	edgeSubnetDataSourceRepresentation = map[string]interface{}{}

	EdgeSubnetResourceConfig = ""
)

func TestWaasEdgeSubnetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasEdgeSubnetResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_waas_edge_subnets.test_edge_subnets"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_waas_edge_subnets", "test_edge_subnets", Required, Create, edgeSubnetDataSourceRepresentation) +
					compartmentIdVariableStr + EdgeSubnetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceName, "edge_subnets.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "edge_subnets.0.cidr"),
					resource.TestCheckResourceAttrSet(datasourceName, "edge_subnets.0.region"),
					resource.TestCheckResourceAttrSet(datasourceName, "edge_subnets.0.time_modified"),
				),
			},
		},
	})
}
