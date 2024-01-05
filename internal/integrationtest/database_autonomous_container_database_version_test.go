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
	DatabaseautonomousContainerDatabaseVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_component": acctest.Representation{RepType: acctest.Required, Create: `ADBD`},
	}

	DatabaseExaccAutonomousContainerDatabaseVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_component": acctest.Representation{RepType: acctest.Required, Create: `EXACC`},
	}

	DatabaseAutonomousContainerDatabaseVersionResourceConfig = ""
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousContainerDatabaseVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_versions.test_autonomous_container_database_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_versions", "test_autonomous_container_database_versions", acctest.Required, acctest.Create, DatabaseautonomousContainerDatabaseVersionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "service_component", "ADBD"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.#"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.details"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_versions.0.supported_apps.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.supported_apps.0.release_date"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.supported_apps.0.end_of_support"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.supported_apps.0.is_certified"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.supported_apps.0.supported_app_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.version"),
			),
		},
	})
}

// issue-routing-tag: database/default
func TestDatabaseExaccAutonomousContainerDatabaseVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_versions.test_autonomous_container_database_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_versions", "test_autonomous_container_database_versions", acctest.Required, acctest.Create, DatabaseExaccAutonomousContainerDatabaseVersionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "service_component", "EXACC"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.#"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_versions.0.supported_apps.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.supported_apps.0.release_date"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.supported_apps.0.end_of_support"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.supported_apps.0.is_certified"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.supported_apps.0.supported_app_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_versions.0.version"),
			),
		},
	})
}
