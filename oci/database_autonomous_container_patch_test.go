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
	autonomousContainerPatchDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   Representation{repType: Required, create: `${var.compartment_id}`},
	}

	AutonomousContainerPatchResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Required, Create, autonomousContainerDatabaseRepresentation) +
		AutonomousExadataInfrastructureResourceConfig
)

func TestDatabaseAutonomousContainerPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerPatchResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_patches.test_autonomous_container_patches"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_patches", "test_autonomous_container_patches", Required, Create, autonomousContainerPatchDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerPatchResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.patch_model"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.quarter"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.type"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.version"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.year"),
				),
			},
		},
	})
}
