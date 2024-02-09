// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	resultId                       string
	DatabaseAutonomousDatabaseRepr = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_model":                        acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"compute_count":                        acctest.Representation{RepType: acctest.Required, Create: `2.0`},
		"data_storage_size_in_tbs":             acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_name":                              acctest.Representation{RepType: acctest.Required, Create: adbName},
		"admin_password":                       acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"db_workload":                          acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"character_set":                        acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"display_name":                         acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`, Update: `displayName2`},
		"is_auto_scaling_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_auto_scaling_for_storage_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_dedicated":                         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_mtls_connection_required":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"autonomous_maintenance_schedule_type": acctest.Representation{RepType: acctest.Optional, Create: `REGULAR`},
		"is_preview_version_with_service_terms_accepted": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":              acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"whitelisted_ips":            acctest.Representation{RepType: acctest.Optional, Create: []string{`1.1.1.1/28`}},
		"operations_insights_status": acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`, Update: `ENABLED`},
		"ncharacter_set":             acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	source_id_undelete = utils.GetEnvSettingWithBlankDefault("source_id_undelete")
	db_name_undelete   = utils.GetEnvSettingWithBlankDefault("db_name_undelete")

	AdbDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_workload":                  acctest.Representation{RepType: acctest.Optional, Create: `DW`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"lifecycle_state_not_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `UNAVAILABLE`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: AdbDataSourceFilterRepresentation}}
	AdbDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_database.test_autonomous_database.id}`}},
	}

	DbVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_workload":    acctest.Representation{RepType: acctest.Optional, Create: `DW`},
	}
	AdbResourceDependenciesUndelete = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Required, acctest.Create, DbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))
	AdbRepresentationUndeleteSource = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_model":            acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"compute_count":            acctest.Representation{RepType: acctest.Required, Create: `2.0`},
		"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_workload":              acctest.Representation{RepType: acctest.Optional, Create: `DW`},
		"is_auto_scaling_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_dedicated":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"db_name":                  acctest.Representation{RepType: acctest.Optional, Create: db_name_undelete},
		"is_preview_version_with_service_terms_accepted": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                 acctest.Representation{RepType: acctest.Optional, Create: ``},
		"license_model":              acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"vault_id":                   acctest.Representation{RepType: acctest.Optional, Create: ``},
		"operations_insights_status": acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`, Update: `ENABLED`},
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `UNDELETE_ADB`},
	}
	AdbRepresentationUndelete = map[string]interface{}{}
)

func TestDatabaseAutonomousDatabaseResource_undelete(t *testing.T) {
	datasourceName := "data.oci_database_autonomous_databases.test_autonomous_databases"

	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseResource_undelete")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	AdbRepresentationUndelete = acctest.RepresentationCopyWithNewProperties(
		AdbRepresentationUndeleteSource,
		map[string]interface{}{
			"source_id": acctest.Representation{RepType: acctest.Optional, Create: source_id_undelete},
		})
	UndeleteAutonomousDatabaseConfig := acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, AdbRepresentationUndelete)
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"
	resourceNameUndelete := "oci_database_autonomous_database.test_autonomous_database_undelete"
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AdbResourceDependenciesUndelete+UndeleteAutonomousDatabaseConfig, "database", "autonomousDatabase", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			//1. Verify datasource
			{
				Config: config + compartmentIdVariableStr + AdbResourceDependenciesUndelete +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_databases", "test_autonomous_databases", acctest.Optional, acctest.Update, AdbDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseRepr),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "db_name", adbName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
					resource.TestCheckResourceAttr(resourceName, "compute_count", "2"),
				),
			},
			//2. Delete before next Create
			{
				Config: config + compartmentIdVariableStr + AdbResourceDependenciesUndelete,
			},
			//3. undelete adb
			{
				Config: config + compartmentIdVariableStr + AdbResourceDependenciesUndelete +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_undelete", acctest.Optional, acctest.Create,
						AdbRepresentationUndelete),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameUndelete, "id"),
					resource.TestCheckResourceAttrSet(resourceNameUndelete, "time_undeleted"),
					resource.TestCheckResourceAttr(resourceNameUndelete, "db_name", db_name_undelete),
				),
			},
			//4.delete resource
			{
				Config: config + compartmentIdVariableStr + AdbResourceDependenciesUndelete,
			},
		},
	})
}
