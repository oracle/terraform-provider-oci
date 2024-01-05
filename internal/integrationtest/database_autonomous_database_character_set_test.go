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
	autonomousDatabaseCharacterSetDataSourceRepresentation = map[string]interface{}{}

	autonomousDatabaseCharacterSetDataSourceRepresentationDedicatedDatabase = map[string]interface{}{
		"character_set_type": acctest.Representation{RepType: acctest.Optional, Create: `DATABASE`},
		"is_dedicated":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_shared":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	autonomousDatabaseCharacterSetDataSourceRepresentationSharedDatabase = map[string]interface{}{
		"character_set_type": acctest.Representation{RepType: acctest.Optional, Create: `DATABASE`},
		"is_dedicated":       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_shared":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	autonomousDatabaseCharacterSetDataSourceRepresentationDedicatedNational = map[string]interface{}{
		"character_set_type": acctest.Representation{RepType: acctest.Optional, Create: `NATIONAL`},
		"is_shared":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	autonomousDatabaseCharacterSetDataSourceRepresentationSharedNational = map[string]interface{}{
		"character_set_type": acctest.Representation{RepType: acctest.Optional, Create: `NATIONAL`},
		"is_shared":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

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

func TestDatabaseAutonomousDatabaseCharacterSetResource_DedicatedDatabase(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseCharacterSetResource_DedicatedDatabase(")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_character_sets", "test_autonomous_database_character_sets", acctest.Required, acctest.Create, autonomousDatabaseCharacterSetDataSourceRepresentationDedicatedDatabase) +
				compartmentIdVariableStr + AutonomousDatabaseCharacterSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.0.name"),
			),
		},
	})
}

func TestDatabaseAutonomousDatabaseCharacterSetResource_DedicatedNational(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseCharacterSetResource_DedicatedNational")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_character_sets", "test_autonomous_database_character_sets", acctest.Required, acctest.Create, autonomousDatabaseCharacterSetDataSourceRepresentationDedicatedNational) +
				compartmentIdVariableStr + AutonomousDatabaseCharacterSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.0.name"),
			),
		},
	})
}

func TestDatabaseAutonomousDatabaseCharacterSetResource_SharedDatabase(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseCharacterSetResource_SharedDatabase")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_character_sets", "test_autonomous_database_character_sets", acctest.Required, acctest.Create, autonomousDatabaseCharacterSetDataSourceRepresentationSharedDatabase) +
				compartmentIdVariableStr + AutonomousDatabaseCharacterSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.0.name"),
			),
		},
	})
}

func TestDatabaseAutonomousDatabaseCharacterSetResource_SharedNational(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseCharacterSetResource_SharedNational")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_database_character_sets.test_autonomous_database_character_sets"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + // change this line, refer to main test
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_character_sets", "test_autonomous_database_character_sets", acctest.Required, acctest.Create, autonomousDatabaseCharacterSetDataSourceRepresentationSharedNational) +
				compartmentIdVariableStr + AutonomousDatabaseCharacterSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_character_sets.0.name"),
			),
		},
	})
}
