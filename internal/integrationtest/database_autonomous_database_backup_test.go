// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var (
	AutonomousDatabaseBackupRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseBackupRepresentation)

	adbBackupDbName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	DatabaseAutonomousDatabaseBackupResourceConfigForLongTermBackup = DatabaseAutonomousDatabaseBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseBackupRepresentationForLongTermBackup)

	DatabaseAutonomousDatabaseBackupResourceConfig = DatabaseAutonomousDatabaseBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseBackupRepresentation)

	DatabaseDatabaseAutonomousDatabaseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`},
	}

	DatabaseDatabaseAutonomousDatabaseBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `Monthly Backup`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousDatabaseBackupDataSourceFilterRepresentation}}
	DatabaseAutonomousDatabaseBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}`}},
	}

	DatabaseAutonomousDatabaseBackupRepresentation = map[string]interface{}{
		"autonomous_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `Monthly Backup`},
		"is_long_term_backup":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"retention_period_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	DatabaseAutonomousDatabaseBackupRepresentationForLongTermBackup = map[string]interface{}{
		"autonomous_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `LongTerm Backup`},
		"is_long_term_backup":      acctest.Representation{RepType: acctest.Required, Create: `true`},
		"retention_period_in_days": acctest.Representation{RepType: acctest.Required, Create: `90`, Update: `91`},
	}

	DatabaseAutonomousDatabaseBackupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create,
		acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbBackupDbName}, DatabaseAutonomousDatabaseRepresentation))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_backup.test_autonomous_database_backup"
	singularDatasourceName := "data.oci_database_autonomous_database_backup.test_autonomous_database_backup"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousDatabaseBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseBackupRepresentationForLongTermBackup), "database", "autonomousDatabaseBackup", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseBackupRepresentationForLongTermBackup),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseBackupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseBackupRepresentationForLongTermBackup),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "LongTerm Backup"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_automatic"),
				resource.TestCheckResourceAttr(resourceName, "is_long_term_backup", "true"),
				resource.TestCheckResourceAttr(resourceName, "retention_period_in_days", "90"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseBackupRepresentationForLongTermBackup),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "LongTerm Backup"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_automatic"),
				resource.TestCheckResourceAttr(resourceName, "is_long_term_backup", "true"),
				resource.TestCheckResourceAttr(resourceName, "retention_period_in_days", "91"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_backup", "test_autonomous_database_backup", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseBackupResourceConfigForLongTermBackup,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_backup_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "LongTerm Backup"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_automatic"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retention_period_in_days", "91"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_available_till"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// verify resource import
		{
			Config:            config + AutonomousDatabaseBackupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_long_term_backup",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseAutonomousDatabaseBackupDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_database_backup" {
			request := oci_database.GetAutonomousDatabaseBackupRequest{}

			tmp := rs.Primary.ID
			request.AutonomousDatabaseBackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			continue
		}
	}
	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseAutonomousDatabaseBackup") {
		resource.AddTestSweepers("DatabaseAutonomousDatabaseBackup", &resource.Sweeper{
			Name:         "DatabaseAutonomousDatabaseBackup",
			Dependencies: acctest.DependencyGraph["autonomousDatabaseBackup"],
			F:            sweepDatabaseAutonomousDatabaseBackupResource,
		})
	}
}

func sweepDatabaseAutonomousDatabaseBackupResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	autonomousDatabaseBackupIds, err := getDatabaseAutonomousDatabaseBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousDatabaseBackupId := range autonomousDatabaseBackupIds {
		if ok := acctest.SweeperDefaultResourceId[autonomousDatabaseBackupId]; !ok {
			deleteAutonomousDatabaseBackupRequest := oci_database.DeleteAutonomousDatabaseBackupRequest{}

			deleteAutonomousDatabaseBackupRequest.AutonomousDatabaseBackupId = &autonomousDatabaseBackupId

			deleteAutonomousDatabaseBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteAutonomousDatabaseBackup(context.Background(), deleteAutonomousDatabaseBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousDatabaseBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousDatabaseBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &autonomousDatabaseBackupId, DatabaseAutonomousDatabaseBackupSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseAutonomousDatabaseBackupSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseAutonomousDatabaseBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AutonomousDatabaseBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listAutonomousDatabaseBackupsRequest := oci_database.ListAutonomousDatabaseBackupsRequest{}
	listAutonomousDatabaseBackupsRequest.CompartmentId = &compartmentId
	listAutonomousDatabaseBackupsRequest.LifecycleState = oci_database.AutonomousDatabaseBackupSummaryLifecycleStateActive
	listAutonomousDatabaseBackupsResponse, err := databaseClient.ListAutonomousDatabaseBackups(context.Background(), listAutonomousDatabaseBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousDatabaseBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousDatabaseBackup := range listAutonomousDatabaseBackupsResponse.Items {
		id := *autonomousDatabaseBackup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousDatabaseBackupId", id)
	}
	return resourceIds, nil
}

func DatabaseAutonomousDatabaseBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDatabaseBackupResponse, ok := response.Response.(oci_database.GetAutonomousDatabaseBackupResponse); ok {
		return autonomousDatabaseBackupResponse.LifecycleState != oci_database.AutonomousDatabaseBackupLifecycleStateDeleted
	}
	return false
}

func DatabaseAutonomousDatabaseBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetAutonomousDatabaseBackup(context.Background(), oci_database.GetAutonomousDatabaseBackupRequest{
		AutonomousDatabaseBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
