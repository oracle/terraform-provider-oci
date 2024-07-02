// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	DatabaseGiVersionMinorVersionDataSourceRepresentation = map[string]interface{}{
		"version":                        acctest.Representation{RepType: acctest.Required, Create: `23.0.0.0`},
		"availability_domain":            acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"is_gi_version_for_provisioning": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"shape":                          acctest.Representation{RepType: acctest.Optional, Create: `ExaDbXS`},
		"shape_family":                   acctest.Representation{RepType: acctest.Optional, Create: `EXADB_XS`},
	}

	DatabaseGiVersionMinorVersionResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: database/default
func TestDatabaseGiVersionMinorVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseGiVersionMinorVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_gi_version_minor_versions.test_gi_version_minor_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_gi_version_minor_versions", "test_gi_version_minor_versions", acctest.Optional, acctest.Create, DatabaseGiVersionMinorVersionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseGiVersionMinorVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "is_gi_version_for_provisioning", "false"),
				resource.TestCheckResourceAttr(datasourceName, "shape", "ExaDbXS"),
				resource.TestCheckResourceAttr(datasourceName, "shape_family", "EXADB_XS"),
				resource.TestCheckResourceAttr(datasourceName, "version", "23.0.0.0"),

				resource.TestCheckResourceAttrSet(datasourceName, "gi_minor_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "gi_minor_versions.0.grid_image_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "gi_minor_versions.0.version"),
			),
		},
	})
}
