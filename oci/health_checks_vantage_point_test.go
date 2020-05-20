// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vantagePointDataSourceRepresentation = map[string]interface{}{
		"display_name": Representation{repType: Optional, create: `AWS Asia Pacific South 1`},
		"name":         Representation{repType: Optional, create: `aws-bom`},
	}

	VantagePointResourceConfig = ""
)

func TestHealthChecksVantagePointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksVantagePointResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_health_checks_vantage_points.test_vantage_points"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_health_checks_vantage_points", "test_vantage_points", Optional, Create, vantagePointDataSourceRepresentation) +
					compartmentIdVariableStr + VantagePointResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "display_name", "AWS Asia Pacific South 1"),
					resource.TestCheckResourceAttr(datasourceName, "name", "aws-bom"),

					resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.0.display_name"),
					resource.TestCheckResourceAttr(datasourceName, "health_checks_vantage_points.0.geo.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.0.provider_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "health_checks_vantage_points.0.routing.#"),
				),
			},
		},
	})
}
