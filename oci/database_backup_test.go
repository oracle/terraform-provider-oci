// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v29/common"
	oci_database "github.com/oracle/oci-go-sdk/v29/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	backupDataSourceRepresentation = map[string]interface{}{
		"database_id": Representation{repType: Optional, create: `${data.oci_database_databases.db.databases.0.id}`},
		"filter":      RepresentationGroup{Required, backupDataSourceFilterRepresentation}}
	backupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_backup.test_backup.id}`}},
	}

	backupRepresentation = map[string]interface{}{
		"database_id":  Representation{repType: Required, create: `${data.oci_database_databases.db.databases.0.id}`},
		"display_name": Representation{repType: Required, create: `Monthly Backup`},
	}

	BackupResourceDependencies = DbSystemResourceConfig + `
data "oci_database_databases" "db" {
       compartment_id = "${var.compartment_id}"
       db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}`
)

func TestDatabaseBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_backup.test_backup"
	datasourceName := "data.oci_database_backups.test_backups"

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseBackupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_backup", "test_backup", Required, Create, backupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "database_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Monthly Backup"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateDataSourceFromRepresentationMap("oci_database_backups", "test_backups", Optional, Update, backupDataSourceRepresentation) +
					compartmentIdVariableStr + BackupResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_backup", "test_backup", Optional, Update, backupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_id"),

					resource.TestCheckResourceAttr(datasourceName, "backups.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_edition"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.database_size_in_gbs"),
					resource.TestCheckResourceAttr(datasourceName, "backups.0.display_name", "Monthly Backup"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.shape"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.time_ended"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.time_started"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.type"),
					resource.TestCheckResourceAttrSet(datasourceName, "backups.0.version"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// Need this workaround due to import behavior change introduced by https://github.com/hashicorp/terraform/issues/20985
					"database_size_in_gbs",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckDatabaseBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_backup" {
			noResourceFound = false
			request := oci_database.GetBackupRequest{}

			tmp := rs.Primary.ID
			request.BackupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

			response, err := client.GetBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.BackupLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseBackup") {
		resource.AddTestSweepers("DatabaseBackup", &resource.Sweeper{
			Name:         "DatabaseBackup",
			Dependencies: DependencyGraph["backup"],
			F:            sweepDatabaseBackupResource,
		})
	}
}

func sweepDatabaseBackupResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	backupIds, err := getBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, backupId := range backupIds {
		if ok := SweeperDefaultResourceId[backupId]; !ok {
			deleteBackupRequest := oci_database.DeleteBackupRequest{}

			deleteBackupRequest.BackupId = &backupId

			deleteBackupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteBackup(context.Background(), deleteBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting Backup %s %s, It is possible that the resource is already deleted. Please verify manually \n", backupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &backupId, backupSweepWaitCondition, time.Duration(3*time.Minute),
				backupSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getBackupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "BackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

	listBackupsRequest := oci_database.ListBackupsRequest{}
	listBackupsRequest.CompartmentId = &compartmentId
	listBackupsResponse, err := databaseClient.ListBackups(context.Background(), listBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Backup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, backup := range listBackupsResponse.Items {
		id := *backup.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "BackupId", id)
	}
	return resourceIds, nil
}

func backupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if backupResponse, ok := response.Response.(oci_database.GetBackupResponse); ok {
		return backupResponse.LifecycleState != oci_database.BackupLifecycleStateDeleted
	}
	return false
}

func backupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetBackup(context.Background(), oci_database.GetBackupRequest{
		BackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
