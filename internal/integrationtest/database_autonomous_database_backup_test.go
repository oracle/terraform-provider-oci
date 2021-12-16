// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	adbBackupDbName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	AutonomousDatabaseBackupResourceConfig = AutonomousDatabaseBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Optional, acctest.Update, autonomousDatabaseBackupRepresentation)

	autonomousDatabaseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`},
	}

	autonomousDatabaseBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `Monthly Backup`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseBackupDataSourceFilterRepresentation}}
	autonomousDatabaseBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`}},
	}

	autonomousDatabaseBackupRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `Monthly Backup`},
	}

	AutonomousDatabaseBackupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create,
		acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbBackupDbName}, autonomousDatabaseRepresentation))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_backup.test_autonomous_database_backup"
	datasourceName := "data.oci_database_autonomous_database_backups.test_autonomous_database_backups"
	singularDatasourceName := "data.oci_database_autonomous_database_backup.test_autonomous_database_backup"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AutonomousDatabaseBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Required, acctest.Create, autonomousDatabaseBackupRepresentation), "database", "autonomousDatabaseBackup", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Required, acctest.Create, autonomousDatabaseBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_backups", "test_autonomous_database_backups", acctest.Optional, acctest.Update, autonomousDatabaseBackupDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Optional, acctest.Update, autonomousDatabaseBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Monthly Backup"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_database_backups.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.autonomous_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_database_backups.0.display_name", "Monthly Backup"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.is_automatic"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.is_restorable"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.time_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_backups.0.type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Required, acctest.Create, autonomousDatabaseBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_backup_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Monthly Backup"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_automatic"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseBackupResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
			ExpectNonEmptyPlan:      true,
		},
	})
}
