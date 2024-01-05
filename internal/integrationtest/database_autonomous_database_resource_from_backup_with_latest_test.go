// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	autonomousDatabaseRepresentationForSourceFromLatestBackup = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbBackupTimestampName}, autonomousDatabaseRepresentationForSourceFromBackupTimestamp), []string{"timestamp"}),
		map[string]interface{}{
			"use_latest_available_backup_time_stamp": acctest.Representation{RepType: acctest.Required, Create: `true`},
		})
)

// issue-routing-tag: database/dbaas-adb
func TestResourceDatabaseAutonomousDatabaseResource_FromBackupWithLatest(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseAutonomousDatabaseResource_FromBackupWithLatest")
	defer httpreplay.SaveScenario()
	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database_from_backuptimestamp"

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//1. Create dependencies
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
		},
		//2. verify create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backuptimestamp", acctest.Required, acctest.Create, autonomousDatabaseRepresentationForSourceFromLatestBackup),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),

				func(s *terraform.State) (err error) {
					resId, err := acctest.FromInstanceState(s, resourceName, "id")
					sourceresId, err := acctest.FromInstanceState(s, "oci_database_autonomous_database.test_autonomous_database", "id")
					if resId == sourceresId {
						return fmt.Errorf("resource not created when it was supposed to be created")
					}
					return err
				},
			),
		},
		//3. delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies,
		},
		//4. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseFromBackupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database_from_backuptimestamp", acctest.Optional, acctest.Create, autonomousDatabaseRepresentationForSourceFromLatestBackup),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_name"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dedicated", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
	})
}
