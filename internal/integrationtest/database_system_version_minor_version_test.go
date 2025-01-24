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
	DatabaseSystemVersionMinorVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"gi_version":     acctest.Representation{RepType: acctest.Required, Create: `18.0.0.0`},
		"major_version":  acctest.Representation{RepType: acctest.Required, Create: `23`},
		"is_latest":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"resource_id":    acctest.Representation{RepType: acctest.Optional, Create: `resource_id`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `Exadata.X9M`},
	}

	DatabaseSystemVersionMinorVersionResourceConfig = ""
)

// issue-routing-tag: database/default
func TestDatabaseSystemVersionMinorVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseSystemVersionMinorVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_system_version_minor_versions.test_system_version_minor_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_system_version_minor_versions", "test_system_version_minor_versions", acctest.Required, acctest.Create, DatabaseSystemVersionMinorVersionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSystemVersionMinorVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "system_version_minor_version_collection.#"),
			),
		},
	})
}
