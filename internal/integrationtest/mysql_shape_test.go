// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	mysqlShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"is_supported_for":    acctest.Representation{RepType: acctest.Optional, Create: []string{`DBSYSTEM`}},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	MySQLShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: mysql/default
func TestMysqlShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_mysql_shapes.test_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", acctest.Required, acctest.Create, mysqlShapeDataSourceRepresentation) +
				compartmentIdVariableStr + MySQLShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", acctest.Optional, acctest.Create, mysqlShapeDataSourceRepresentation) +
				compartmentIdVariableStr + MySQLShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
