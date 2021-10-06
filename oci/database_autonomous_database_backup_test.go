// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	adbBackupDbName = RandomString(1, charsetWithoutDigits) + RandomString(13, charset)

	AutonomousDatabaseBackupResourceConfig = AutonomousDatabaseBackupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Optional, Update, autonomousDatabaseBackupRepresentation)

	autonomousDatabaseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_backup_id": Representation{RepType: Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`},
	}

	autonomousDatabaseBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{RepType: Optional, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":           Representation{RepType: Optional, Create: `Monthly Backup`},
		"state":                  Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":                 RepresentationGroup{Required, autonomousDatabaseBackupDataSourceFilterRepresentation}}
	autonomousDatabaseBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`}},
	}

	autonomousDatabaseBackupRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{RepType: Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":           Representation{RepType: Required, Create: `Monthly Backup`},
	}

	AutonomousDatabaseBackupResourceDependencies = GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
		GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbBackupDbName}, autonomousDatabaseRepresentation))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_backup.test_autonomous_database_backup"
	datasourceName := "data.oci_database_autonomous_database_backups.test_autonomous_database_backups"
	singularDatasourceName := "data.oci_database_autonomous_database_backup.test_autonomous_database_backup"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+AutonomousDatabaseBackupResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Required, Create, autonomousDatabaseBackupRepresentation), "database", "autonomousDatabaseBackup", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseBackupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Required, Create, autonomousDatabaseBackupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_backups", "test_autonomous_database_backups", Optional, Update, autonomousDatabaseBackupDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseBackupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Optional, Update, autonomousDatabaseBackupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", Required, Create, autonomousDatabaseBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseBackupResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
