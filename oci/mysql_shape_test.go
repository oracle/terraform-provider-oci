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
	mysqlShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"is_supported_for":    Representation{repType: Optional, create: []string{`DBSYSTEM`}},
		"name":                Representation{repType: Optional, create: `name`},
	}

	MySQLShapeResourceConfig = AvailabilityDomainConfig
)

func TestMysqlShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_mysql_shapes.test_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", Required, Create, mysqlShapeDataSourceRepresentation) +
					compartmentIdVariableStr + MySQLShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.cpu_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.memory_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				),
			},
			// verify datasource with optionals
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", Optional, Create, mysqlShapeDataSourceRepresentation) +
					compartmentIdVariableStr + MySQLShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "is_supported_for.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "name", "name"),

					resource.TestCheckResourceAttrSet(datasourceName, "shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.cpu_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.is_supported_for.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.memory_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				),
			},
		},
	})
}
