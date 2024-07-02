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
	ExaDbVmClusterLocalClonedPDBRequiredOnlyResource = ExaDbVmClusterLocalClonePdbResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, ExaDbVmClusterLocalClonedPDBRepresentation)

	ExaDbVmClusterLocalClonedPDBSingularDataSourceRepresentation = map[string]interface{}{
		"pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
	}

	ExaDbVmClusterLocalClonedPDBDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"pdb_name":       acctest.Representation{RepType: acctest.Optional, Create: `LocalThinClonedPdb`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ExaDbVmClusterLocalClonedPDBDataSourceFilterRepresentation}}

	ExaDbVmClusterLocalClonedPDBDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_pluggable_database.test_pluggable_database.id}`}},
	}

	ExaDbVmClusterLocalClonedPDBRepresentation = map[string]interface{}{
		"container_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_pluggable_database.source_pdb.container_database_id}`},
		"pdb_name":                  acctest.Representation{RepType: acctest.Required, Create: `LocalThinClonedPdb`},
		"pdb_admin_password":        acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"tde_wallet_password":       acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pdb_creation_type_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: ExaDbVmClusterLocalClonePdbCreationTypeDetailsRepresentation},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ExaDbVmClusterLocalClonePdbIgnoreDefinedTagsRepresentation},
	}

	ExaDbVmClusterLocalClonePdbCreationTypeDetailsRepresentation = map[string]interface{}{
		"creation_type":                acctest.Representation{RepType: acctest.Required, Create: `LOCAL_CLONE_PDB`},
		"source_pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_pluggable_database_id}`},
		"is_thin_clone":                acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	ExaDbVmClusterLocalClonePdbIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	// Note: set env variable TF_VAR_source_pluggable_database_id before running this test
	ExaDbVmClusterLocalCloneSourcePDBSingularDataSourceRepresentation = map[string]interface{}{
		"pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_pluggable_database_id}`},
	}
	ExaDbVmClusterLocalClonePdbResourceDependencies = `variable "source_pluggable_database_id" {}` +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_pluggable_database", "source_pdb", acctest.Optional, acctest.Create, ExaDbVmClusterLocalCloneSourcePDBSingularDataSourceRepresentation)
)

// issue-routing-tag: database/ExaCS
func TestDatabaseExaDbVmClusterPluggableDatabaseResource_localThinClone(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaDbVmClusterPluggableDatabaseResource_localThinClone")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_pluggable_database.test_pluggable_database"
	datasourceName := "data.oci_database_pluggable_databases.test_pluggable_databases"
	singularDatasourceName := "data.oci_database_pluggable_database.test_pluggable_database"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExaDbVmClusterLocalClonePdbResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, ExaDbVmClusterLocalClonedPDBRepresentation), "database", "pluggableDatabase", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExaDbVmClusterLocalClonePdbResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, ExaDbVmClusterLocalClonedPDBRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "LocalThinClonedPdb"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "pdb_creation_type_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "pdb_creation_type_details.0.creation_type", "LOCAL_CLONE_PDB"),
				resource.TestCheckResourceAttr(resourceName, "pdb_creation_type_details.0.is_thin_clone", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "pdb_creation_type_details.0.source_pluggable_database_id"),
			),
		},
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + ExaDbVmClusterLocalClonePdbResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, ExaDbVmClusterLocalClonedPDBRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_pluggable_databases", "test_pluggable_databases", acctest.Optional, acctest.Create, ExaDbVmClusterLocalClonedPDBDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "state"),

				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.pdb_name", "LocalThinClonedPdb"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_databases.0.connection_strings.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_databases.0.container_database_id"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + ExaDbVmClusterLocalClonePdbResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, ExaDbVmClusterLocalClonedPDBRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, ExaDbVmClusterLocalClonedPDBSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "pdb_name", "LocalThinClonedPdb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "container_database_id"),
			),
		},
		// verify resource import
		{
			Config:            config + ExaDbVmClusterLocalClonedPDBRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"pdb_admin_password",
				"pdb_creation_type_details",
				"tde_wallet_password",
			},
			ResourceName: resourceName,
		},
		{
			Config: config + compartmentIdVariableStr + ExaDbVmClusterLocalClonePdbResourceDependencies,
		},
	})
}
