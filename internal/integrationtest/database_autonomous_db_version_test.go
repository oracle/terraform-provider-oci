// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousDbVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_workload":    acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
	}

	AutonomousDbVersionResourceConfig = ""
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousDbVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDbVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_db_versions.test_autonomous_db_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Required, acctest.Create, autonomousDbVersionDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDbVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.db_workload"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.details"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.is_dedicated"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.version"),
			),
		},

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Optional, acctest.Create, autonomousDbVersionDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDbVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_workload", "OLTP"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.db_workload"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.details"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.is_dedicated"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.is_default_for_free"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.is_default_for_paid"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.is_free_tier_enabled"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.is_paid_enabled"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.version"),
			),
		},
	})
}
