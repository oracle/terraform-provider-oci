// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseAutonomousDatabaseDisplayCostsRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_count":            acctest.Representation{RepType: acctest.Required, Create: `4.0`, Update: `4.0`},
		"compute_model":            acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_name":                  acctest.Representation{RepType: acctest.Required, Create: adbName},
		"admin_password":           acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#11`},
		"db_version":               acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}`},
		"db_workload":              acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `displayName2`, Update: `displayName2`},
		"is_dedicated":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":            acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DatabaseAutonomousDatabasesEstimateCostSavingDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database_leader.id}`},
		"is_cpu_autoscale":       acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	DatabaseAutonomousDatabasesEstimateCostSavingResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create, DatabaseAutonomousRepresentationDisplayCostsNewResource)

	DatabaseAutonomousRepresentationDisplayCostsNewResource = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseDisplayCostsRepresentation, []string{"cpu_core_count", "db_tools_details"}), map[string]interface{}{
		"compute_count":         acctest.Representation{RepType: acctest.Required, Create: `4.0`, Update: `4.0`},
		"compute_model":         acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseRPSummaryRepresentation},
	})

	DatabaseAutonomousDisplayCostsRPSummaryRepresentationResource = map[string]interface{}{
		"is_disabled": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
	}

	DatabaseAutonomousDisplayCostsRepresentationRPUpdateNewResource = acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousDatabaseDisplayCostsRepresentation, []string{"cpu_core_count", "db_tools_details"}), map[string]interface{}{
		"compute_count":         acctest.Representation{RepType: acctest.Required, Create: `4.0`, Update: `4.0`},
		"compute_model":         acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDisplayCostsRPSummaryRepresentationResource},
	})

	DatabaseAutonomousDisplayCostsRPDisableSummaryRepresentation = map[string]interface{}{
		"is_disabled": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
	}
	resId string
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabasesEstimateCostSavingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabasesEstimateCostSavingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	resourceName := "oci_database_autonomous_database.test_autonomous_database_leader"
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_databases_estimate_cost_savings.test_autonomous_databases_estimate_cost_savings"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Optional, acctest.Create, DatabaseAutonomousRepresentationDisplayCostsNewResource), "database", "autonomousDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases_estimate_cost_savings", "test_autonomous_databases_estimate_cost_savings", acctest.Required, acctest.Create, DatabaseAutonomousDatabasesEstimateCostSavingDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabasesEstimateCostSavingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "is_cpu_autoscale", "false"),

				resource.TestCheckResourceAttrSet(datasourceName, "estimate_cost_savings_summary_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "estimate_cost_savings_summary_collection.0.items.#", "0"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, datasourceName, "id")
					fmt.Printf("Error running elastic resource %s with resId %s \n", err, resId)
					return err
				},
			),
		},
		// verify disable resource pool leader
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_leader", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDisplayCostsRepresentationRPUpdateNewResource, map[string]interface{}{
						"resource_pool_summary": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDisplayCostsRPDisableSummaryRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceName, "total_backup_storage_size_in_gbs", "1000"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_pool_summary.#"),
				resource.TestCheckResourceAttr(resourceName, "resource_pool_summary.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_pool_summary.0.is_disabled", "true"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Printf("Error running elastic resource %s with resId %s \n", err, resId)
					return err
				},
			),
		},
	})
}
