// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	mysqlShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"is_supported_for":    Representation{RepType: Optional, Create: []string{`DBSYSTEM`}},
		"name":                Representation{RepType: Optional, Create: `name`},
	}

	MySQLShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: mysql/default
func TestMysqlShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_mysql_shapes.test_shapes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", Required, Create, mysqlShapeDataSourceRepresentation) +
				compartmentIdVariableStr + MySQLShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", Optional, Create, mysqlShapeDataSourceRepresentation) +
				compartmentIdVariableStr + MySQLShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	})
}
