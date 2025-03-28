// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MysqlMysqlMysqlVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	MysqlMysqlVersionResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_versions", "test_mysql_versions", acctest.Required, acctest.Create, MysqlMysqlMysqlVersionDataSourceRepresentation)
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_mysql_mysql_versions.test_mysql_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "versions.0.version_family"),
			),
		},
	})
}
