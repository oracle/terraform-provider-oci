// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowDayOfWeekRepresentationSaturday = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `SATURDAY`},
	}

	DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowDayOfWeekRepresentationSunday = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `SUNDAY`},
	}

	DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowRepresentation = map[string]interface{}{
		"day_of_week":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowDayOfWeekRepresentationSaturday},
		"maintenance_start_time": acctest.Representation{RepType: acctest.Required, Create: `09:00`},
		"maintenance_end_time":   acctest.Representation{RepType: acctest.Required, Create: `11:00`},
	}

	DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowRepresentationUpdated = map[string]interface{}{
		"day_of_week":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowDayOfWeekRepresentationSunday},
		"maintenance_start_time": acctest.Representation{RepType: acctest.Required, Create: `11:00`},
		"maintenance_end_time":   acctest.Representation{RepType: acctest.Required, Create: `13:00`},
	}

	DatabaseAutonomousDatabaseRPSummaryRepresentation256 = map[string]interface{}{
		"is_disabled":              acctest.Representation{RepType: acctest.Required, Create: `false`},
		"pool_size":                acctest.Representation{RepType: acctest.Required, Create: `256`},
		"pool_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Create: `256`},
	}

	DatabaseAutonomousDatabaseRPSummaryRepresentation512 = map[string]interface{}{
		"is_disabled":              acctest.Representation{RepType: acctest.Required, Create: `false`},
		"pool_size":                acctest.Representation{RepType: acctest.Required, Create: `512`},
		"pool_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Create: `512`},
	}

	timeOfMaintenancePause = time.Now().UTC().AddDate(0, 0, 20).Truncate(time.Millisecond).Format(time.RFC3339)

	autonomousDatabaseRepresentationRPLocalStandby = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"cpu_core_count", "db_tools_details", "display_name"}), map[string]interface{}{
		"db_name":                                acctest.Representation{RepType: acctest.Required, Create: longAdbName1},
		"compute_count":                          acctest.Representation{RepType: acctest.Required, Create: `4.0`},
		"compute_model":                          acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"data_storage_size_in_tbs":               acctest.Representation{RepType: acctest.Required, Create: `5`},
		"resource_pool_summary":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryRepresentation256},
		"autonomous_database_maintenance_window": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowRepresentation},
	})

	autonomousDatabaseRepresentationRPInitial = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"cpu_core_count", "db_tools_details", "display_name"}), map[string]interface{}{
		"compute_count":                          acctest.Representation{RepType: acctest.Required, Create: `4.0`},
		"compute_model":                          acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"data_storage_size_in_tbs":               acctest.Representation{RepType: acctest.Required, Create: `5`},
		"resource_pool_summary":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryRepresentation256},
		"autonomous_database_maintenance_window": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowRepresentation},
	})

	autonomousDatabaseRepresentationRPUpdated1 = acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationRPInitial, map[string]interface{}{
		"autonomous_database_maintenance_window": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseAutonomousDatabaseMaintenanceWindowRepresentationUpdated},
		"time_maintenance_pause_until":           acctest.Representation{RepType: acctest.Required, Create: timeOfMaintenancePause},
	})

	autonomousDatabaseRepresentationRPUpdated2 = acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationRPUpdated1, map[string]interface{}{
		"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryRepresentation512},
	})

	autonomousDatabaseResourcePoolMemberJoined = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseRepresentation, []string{"cpu_core_count", "db_tools_details", "display_name"}), map[string]interface{}{
		"compute_count":           acctest.Representation{RepType: acctest.Required, Create: `10.0`},
		"compute_model":           acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"resource_pool_leader_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database_leader.id}`},
		"db_name":                 acctest.Representation{RepType: acctest.Required, Create: adbMemberName},
	})

	autonomousDatabaseResourcePoolMemberLocalADGEnabled = acctest.RepresentationCopyWithNewProperties(autonomousDatabaseResourcePoolMemberJoined, map[string]interface{}{
		"is_local_data_guard_enabled":       acctest.Representation{RepType: acctest.Required, Create: `true`},
		"local_adg_resource_pool_leader_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database_local_standby_leader.id}`},
	})

	autonomousDatabaseResourcePoolMemberLocalADGDisabled = acctest.RepresentationCopyWithNewProperties(autonomousDatabaseResourcePoolMemberLocalADGEnabled, map[string]interface{}{
		"is_local_data_guard_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
	})

	autonomousDatabaseResourcePoolMemberLeft = acctest.RepresentationCopyWithNewProperties(autonomousDatabaseResourcePoolMemberLocalADGDisabled, map[string]interface{}{
		"resource_pool_leader_id": acctest.Representation{RepType: acctest.Required, Create: ` `},
	})

	autonomousDatabaseResourcePoolMemberDataSourceRepresentation = map[string]interface{}{
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

	resourceLeaderName := "oci_database_autonomous_database.test_autonomous_database_leader"
	resourceLocalStandbyLeaderName := "oci_database_autonomous_database.test_autonomous_database_local_standby_leader"
	resourceMemberName := "oci_database_autonomous_database.test_autonomous_database_member"
	datasourceName := "data.oci_database_autonomous_database_resource_pool_members.test_autonomous_database_resource_pool_members"

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. Verify leader with dedicated storage create with maintenance window
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPInitial),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceLeaderName, "data_storage_size_in_tbs", "5"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.pool_size", "256"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.pool_storage_size_in_tbs", "256"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.available_storage_capacity_in_tbs", "251"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.day_of_week.0.name", "SATURDAY"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.maintenance_start_time", "09:00"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.maintenance_end_time", "11:00"),
			),
		},
		//1. Update leader's paused Maintenance and maintenance window
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPUpdated1),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceLeaderName, "data_storage_size_in_tbs", "5"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.pool_size", "256"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.pool_storage_size_in_tbs", "256"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.available_storage_capacity_in_tbs", "251"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.day_of_week.0.name", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.maintenance_start_time", "11:00"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.maintenance_end_time", "13:00"),
				resource.TestCheckResourceAttrSet(resourceLeaderName, "time_maintenance_pause_until"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLeaderName, "id")
					return err
				},
			),
		},
		//2. Update leader's pool storage
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPUpdated2),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceLeaderName, "data_storage_size_in_tbs", "5"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.pool_size", "512"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.pool_storage_size_in_tbs", "512"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.available_storage_capacity_in_tbs", "507"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.day_of_week.0.name", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.maintenance_start_time", "11:00"),
				resource.TestCheckResourceAttr(resourceLeaderName, "autonomous_database_maintenance_window.0.maintenance_end_time", "13:00"),
				resource.TestCheckResourceAttrSet(resourceLeaderName, "time_maintenance_pause_until"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLeaderName, "id")
					return err
				},
			),
		},
		//3. Create member under the primary ERP leader
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPUpdated2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Create, autonomousDatabaseResourcePoolMemberJoined),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceMemberName, "resource_pool_leader_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLeaderName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceMemberName, "id")
					return err
				},
			),
		},
		//4. List resource pool member
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPUpdated2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Create, autonomousDatabaseResourcePoolMemberJoined) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_resource_pool_members", "test_autonomous_database_resource_pool_members", acctest.Required, acctest.Create, autonomousDatabaseResourcePoolMemberDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "resource_pool_member_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "resource_pool_member_collection.0.items.#", "1"),
			),
		},
		//5. Create another dedicated leader for member's local ADG
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPUpdated2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_local_standby_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPLocalStandby) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Create, autonomousDatabaseResourcePoolMemberJoined),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceLocalStandbyLeaderName, "data_storage_size_in_tbs", "5"),
				resource.TestCheckResourceAttr(resourceLocalStandbyLeaderName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceLocalStandbyLeaderName, "resource_pool_summary.0.is_disabled", "false"),
				resource.TestCheckResourceAttr(resourceLocalStandbyLeaderName, "resource_pool_summary.0.pool_size", "256"),
				resource.TestCheckResourceAttr(resourceLocalStandbyLeaderName, "resource_pool_summary.0.pool_storage_size_in_tbs", "256"),
				resource.TestCheckResourceAttr(resourceLocalStandbyLeaderName, "resource_pool_summary.0.available_storage_capacity_in_tbs", "251"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLeaderName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceMemberName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLocalStandbyLeaderName, "id")
					return err
				},
			),
		},
		//6. Enable local ADG on a member of dedicated ERP
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPUpdated2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_local_standby_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPLocalStandby) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Create, autonomousDatabaseResourcePoolMemberLocalADGEnabled),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceMemberName, "resource_pool_leader_id"),
				resource.TestCheckResourceAttr(resourceMemberName, "is_local_data_guard_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceMemberName, "local_adg_resource_pool_leader_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLeaderName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceMemberName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLocalStandbyLeaderName, "id")
					return err
				},
			),
		},
		//7. Disable local ADG on a member of dedicated ERP
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPUpdated2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_local_standby_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPLocalStandby) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Create, autonomousDatabaseResourcePoolMemberLocalADGDisabled),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceMemberName, "resource_pool_leader_id"),
				resource.TestCheckResourceAttr(resourceMemberName, "is_local_data_guard_enabled", "false"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLeaderName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceMemberName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLocalStandbyLeaderName, "id")
					return err
				},
			),
		},
		//8. Member leaving the resource pool leader
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPUpdated2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_local_standby_leader", acctest.Required, acctest.Create, autonomousDatabaseRepresentationRPLocalStandby) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_member", acctest.Required, acctest.Create, autonomousDatabaseResourcePoolMemberLeft),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceMemberName, "resource_pool_leader_id", " "),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLeaderName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceMemberName, "id")
					return err
				},
				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLocalStandbyLeaderName, "id")
					return err
				},
			),
		},
		//9. Disable resource pool leaders
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationRPUpdated2, map[string]interface{}{
						"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPDisableSummaryRepresentation},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_local_standby_leader", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autonomousDatabaseRepresentationRPLocalStandby, map[string]interface{}{
						"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPDisableSummaryRepresentation},
					})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceLeaderName, "resource_pool_summary.0.is_disabled", "true"),

				resource.TestCheckResourceAttr(resourceLocalStandbyLeaderName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceLocalStandbyLeaderName, "resource_pool_summary.0.is_disabled", "true"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLeaderName, "id")
					return err
				},

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceLocalStandbyLeaderName, "id")
					return err
				},
			),
		},
	})
}
