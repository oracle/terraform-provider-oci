// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	autonomousDatabaseCharacterSetDataSourceRepresentation = map[string]interface{}{}

	AutonomousDatabaseCharacterSetResourceConfig = ""
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousDatabaseCharacterSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseCharacterSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_database_character_sets.test_autonomous_database_character_sets"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_character_sets", "test_autonomous_database_character_sets", acctest.Required, acctest.Create, autonomousDatabaseCharacterSetDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseCharacterSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.0.name"),
			),
		},
	})
}
