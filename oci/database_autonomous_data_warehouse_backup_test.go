// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousDataWarehouseBackupResourceConfig = AutonomousDataWarehouseBackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse_backup", "test_autonomous_data_warehouse_backup", Optional, Update, autonomousDataWarehouseBackupRepresentation)

	autonomousDataWarehouseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_data_warehouse_backup_id": Representation{repType: Required, create: `${oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup.id}`},
	}

	autonomousDataWarehouseBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_data_warehouse_id": Representation{repType: Optional, create: `${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}`},
		"display_name":                 Representation{repType: Optional, create: `Monthly Backup`},
		"state":                        Representation{repType: Optional, create: `ACTIVE`},
		"filter":                       RepresentationGroup{Required, autonomousDataWarehouseBackupDataSourceFilterRepresentation}}
	autonomousDataWarehouseBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup.id}`}},
	}

	autonomousDataWarehouseBackupRepresentation = map[string]interface{}{
		"autonomous_data_warehouse_id": Representation{repType: Required, create: `${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}`},
		"display_name":                 Representation{repType: Required, create: `Monthly Backup`},
	}

	AutonomousDataWarehouseBackupResourceDependencies = AutonomousDataWarehouseResourceConfig
)

func TestDatabaseAutonomousDataWarehouseBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDataWarehouseBackupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup"
	datasourceName := "data.oci_database_autonomous_data_warehouse_backups.test_autonomous_data_warehouse_backups"
	singularDatasourceName := "data.oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup"

	testResourceName := randomString(14, charset)
	setEnvSetting("TF_VAR_autonomous_data_warehouse_db_name", testResourceName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDataWarehouseBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse_backup", "test_autonomous_data_warehouse_backup", Required, Create, autonomousDataWarehouseBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_data_warehouse_backups", "test_autonomous_data_warehouse_backups", Optional, Update, autonomousDataWarehouseBackupDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDataWarehouseBackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse_backup", "test_autonomous_data_warehouse_backup", Optional, Update, autonomousDataWarehouseBackupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "Monthly Backup"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouse_backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_data_warehouse_backups.0.display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.is_automatic"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_data_warehouse_backups.0.type"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_data_warehouse_backup", "test_autonomous_data_warehouse_backup", Required, Create, autonomousDataWarehouseBackupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDataWarehouseBackupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_warehouse_backup_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_warehouse_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_automatic"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AutonomousDataWarehouseBackupResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
