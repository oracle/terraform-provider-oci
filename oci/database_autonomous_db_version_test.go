// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	autonomousDbVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"db_workload":    Representation{repType: Optional, create: `OLTP`},
	}

	AutonomousDbVersionResourceConfig = ""
)

func TestDatabaseAutonomousDbVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDbVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_db_versions.test_autonomous_db_versions"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", Required, Create, autonomousDbVersionDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDbVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", Optional, Create, autonomousDbVersionDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDbVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "db_workload", "OLTP"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.db_workload"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.details"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.is_dedicated"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.is_free_tier_enabled"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_db_versions.0.version"),
				),
			},
		},
	})
}
