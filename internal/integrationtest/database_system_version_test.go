// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseSystemVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"gi_version":     acctest.Representation{RepType: acctest.Required, Create: `18.0.0.0`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `Exadata.X9M`},
	}

	DatabaseSystemVersionResourceConfig = ""
)

// issue-routing-tag: database/default
func TestDatabaseSystemVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseSystemVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_system_versions.test_system_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_system_versions", "test_system_versions", acctest.Required, acctest.Create, DatabaseSystemVersionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSystemVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "gi_version", "18.0.0.0"),
				resource.TestCheckResourceAttr(datasourceName, "shape", "Exadata.X9M"),

				resource.TestCheckResourceAttrSet(datasourceName, "system_version_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "system_version_collection.0.items.#", "1"),
			),
		},
	})
}
