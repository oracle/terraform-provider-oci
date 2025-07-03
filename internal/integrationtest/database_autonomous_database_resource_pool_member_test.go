// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	autonomousDatabaseRepresentationRPNew = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"cpu_core_count", "db_tools_details"}), map[string]interface{}{
		"compute_count":         acctest.Representation{RepType: acctest.Required, Create: `4.0`, Update: `4.0`},
		"compute_model":         acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryRepresentation},
	})

	autonomousDatabaseRepresentationRPUpdateNew = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"cpu_core_count", "db_tools_details"}), map[string]interface{}{
		"compute_count":         acctest.Representation{RepType: acctest.Required, Create: `4.0`, Update: `4.0`},
		"compute_model":         acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryRepresentation},
	})

	DatabaseAutonomousDatabaseResourcePoolLeaderIdRepresentationNew = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"cpu_core_count", "db_tools_details"}), map[string]interface{}{
		"compute_count":           acctest.Representation{RepType: acctest.Required, Create: `10.0`, Update: `10.0`},
		"compute_model":           acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"resource_pool_leader_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database_leader.id}`, Update: ` `},
		"db_name":                 acctest.Representation{RepType: acctest.Required, Create: adbMemberName},
	})

	DatabaseAutonomousDatabaseResourcePoolLeaderIdUpdateRepresentationNew = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"cpu_core_count", "admin_password", "db_tools_details"}), map[string]interface{}{
		"compute_count":           acctest.Representation{RepType: acctest.Required, Create: `10.0`, Update: `12.0`},
		"resource_pool_leader_id": acctest.Representation{RepType: acctest.Required, Update: ` `},
		"db_name":                 acctest.Representation{RepType: acctest.Required, Create: adbMemberName},
	})

	DatabaseAutonomousDatabaseResourcePoolMemberDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database_leader.id}`},
	}
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseResourcePoolMemberResource_basic(t *testing.T) {
	shouldSkipADBStest := os.Getenv("TF_VAR_should_skip_adbs_test")

	if shouldSkipADBStest == "true" {
		t.Skip("Skipping TestDatabaseAutonomousDatabaseResourcePoolMemberResource_basic test.\n" + "Current TF_VAR_should_skip_adbs_test=" + shouldSkipADBStest)
	}

	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseResourcePoolMemberResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_leader"
	resourceMemberName := "oci_database_autonomous_database.test_autonomous_database_member"
	datasourceName := "data.oci_database_autonomous_database_resource_pool_members.test_autonomous_database_resource_pool_members"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Optional, acctest.Create, autonomousDatabaseRepresentationRPNew), "database", "autonomousDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. Verify leader create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPNew),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
				// verify computed field db_workload to be defaulted to OLTP
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_pool_summary.#"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_pool_summary.#"),
				resource.TestCheckResourceAttr(resourceName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_pool_summary.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "resource_pool_summary.0.pool_size", "128"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. Verify member create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationRPNew, map[string]interface{}{
						"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryRepresentation},
					})) + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseResourcePoolLeaderIdRepresentationNew),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceMemberName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceMemberName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceMemberName, "cpu_core_count", "0"),
				resource.TestCheckResourceAttr(resourceMemberName, "compute_count", "10"),
				resource.TestCheckResourceAttr(resourceMemberName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceMemberName, "db_name", adbMemberName),
				// verify computed field db_workload to be defaulted to OLTP
				resource.TestCheckResourceAttr(resourceMemberName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttrSet(resourceMemberName, "resource_pool_leader_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceMemberName, "id")
					return err
				},
			),
		},
		//2. Verify list resource pool member
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationRPNew, map[string]interface{}{
						"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryUpdateRepresentation},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseResourcePoolLeaderIdRepresentationNew) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_resource_pool_members", "test_autonomous_database_resource_pool_members", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseResourcePoolMemberDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "resource_pool_member_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "resource_pool_member_collection.0.items.#", "1"),
			),
		},
		//3. Verify member leaving the resource pool leader
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create,
				acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationRPNew, map[string]interface{}{
					"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryRepresentation},
				})) + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Update, DatabaseAutonomousDatabaseResourcePoolLeaderIdUpdateRepresentationNew),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceMemberName, "resource_pool_leader_id"),
				resource.TestCheckResourceAttr(resourceMemberName, "resource_pool_leader_id", " "),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceMemberName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//4. Verify disable resource pool leader
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationRPUpdateNew, map[string]interface{}{
						"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPDisableSummaryRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceName, "total_backup_storage_size_in_gbs", "1000"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_pool_summary.#"),
				resource.TestCheckResourceAttr(resourceName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_pool_summary.0.is_disabled", "true"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}
