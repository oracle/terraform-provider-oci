// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

//import (
//	"fmt"
//	"testing"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//
//	"github.com/oracle/terraform-provider-oci/httpreplay"
//	"github.com/oracle/terraform-provider-oci/internal/acctest"
//
//	"github.com/oracle/terraform-provider-oci/internal/utils"
//)
//
//var (
//	DatabaseDatabaseAutonomousDatabaseRefreshableCloneSingularDataSourceRepresentation = map[string]interface{}{
//		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
//	}
//
//	DatabaseDatabaseAutonomousDatabaseRefreshableCloneDataSourceRepresentation = map[string]interface{}{
//		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
//	}
//
//	DatabaseAutonomousDatabaseRefreshableCloneResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Required, acctest.Create, autonomousDatabaseBackupRepresentation) +
//		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation)
//)
//
//// issue-routing-tag: database/dbaas-adb
//func TestDatabaseAutonomousDatabaseRefreshableCloneResource_basic(t *testing.T) {
//	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseRefreshableCloneResource_basic")
//	defer httpreplay.SaveScenario()
//
//	config := acctest.ProviderTestConfig()
//
//	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
//	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
//
//	datasourceName := "data.oci_database_autonomous_database_refreshable_clones.test_autonomous_database_refreshable_clones"
//	singularDatasourceName := "data.oci_database_autonomous_database_refreshable_clone.test_autonomous_database_refreshable_clone"
//
//	acctest.SaveConfigContent("", "", "", t)
//
//	acctest.ResourceTest(t, nil, []resource.TestStep{
//		// verify datasource
//		{
//			Config: config +
//				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_refreshable_clones", "test_autonomous_database_refreshable_clones", acctest.Required, acctest.Create, DatabaseautonomousDatabaseRefreshableCloneDataSourceRepresentation) +
//				compartmentIdVariableStr + DatabaseAutonomousDatabaseRefreshableCloneResourceConfig,
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_id"),
//
//				resource.TestCheckResourceAttrSet(datasourceName, "refreshable_clone_collection.#"),
//				resource.TestCheckResourceAttr(datasourceName, "refreshable_clone_collection.0.items.#", "1"),
//			),
//		},
//		// verify singular datasource
//		{
//			Config: config +
//				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_refreshable_clone", "test_autonomous_database_refreshable_clone", acctest.Required, acctest.Create, DatabaseautonomousDatabaseRefreshableCloneSingularDataSourceRepresentation) +
//				compartmentIdVariableStr + DatabaseAutonomousDatabaseRefreshableCloneResourceConfig,
//			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
//				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
//
//				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
//			),
//		},
//	})
//}
