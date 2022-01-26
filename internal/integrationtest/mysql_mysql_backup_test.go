// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	MysqlBackupRequiredOnlyResource = MysqlBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Required, acctest.Create, mysqlBackupRepresentation)

	MysqlBackupResourceConfig = MysqlBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Optional, acctest.Update, mysqlBackupRepresentation)

	mysqlBackupSingularDataSourceRepresentation = map[string]interface{}{
		"backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_backup.test_mysql_backup.id}`},
	}

	mysqlBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"backup_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_backup.test_mysql_backup.id}`},
		"creation_type":  acctest.Representation{RepType: acctest.Optional, Create: `MANUAL`},
		"db_system_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_backup_db_system.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: mysqlBackupDataSourceFilterRepresentation}}
	mysqlBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_mysql_mysql_backup.test_mysql_backup.id}`}},
	}

	mysqlBackupRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_system_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_backup_db_system.id}`},
		"backup_type":       acctest.Representation{RepType: acctest.Optional, Create: `INCREMENTAL`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"retention_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	MysqlBackupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		MysqlConfigurationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_backup_db_system", acctest.Required, acctest.Create, mysqlDbSystemRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_mysql_mysql_backup.test_mysql_backup"
	datasourceName := "data.oci_mysql_mysql_backups.test_mysql_backups"
	singularDatasourceName := "data.oci_mysql_mysql_backup.test_mysql_backup"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MysqlBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Optional, acctest.Create, mysqlBackupRepresentation), "mysql", "mysqlBackup", t)

	acctest.ResourceTest(t, testAccCheckMysqlMysqlBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MysqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Required, acctest.Create, mysqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MysqlBackupResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MysqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Optional, acctest.Create, mysqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "creation_type"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MysqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(mysqlBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "creation_type"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + MysqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Optional, acctest.Update, mysqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "creation_type"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_in_days", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_backups", "test_mysql_backups", acctest.Optional, acctest.Update, mysqlBackupDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Optional, acctest.Update, mysqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "backup_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "creation_type", "MANUAL"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "backups.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.backup_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.creation_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.data_storage_size_in_gb"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.mysql_version"),
				resource.TestCheckResourceAttr(datasourceName, "backups.0.retention_in_days", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.shape_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "backups.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup", acctest.Required, acctest.Create, mysqlBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "creation_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_storage_size_in_gb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_system_snapshot.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mysql_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retention_in_days", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + MysqlBackupResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMysqlMysqlBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbBackupsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_mysql_backup" {
			noResourceFound = false
			request := oci_mysql.GetBackupRequest{}

			tmp := rs.Primary.ID
			request.BackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")

			response, err := client.GetBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.BackupLifecycleStateDeleted): true,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MysqlMysqlBackup") {
		resource.AddTestSweepers("MysqlMysqlBackup", &resource.Sweeper{
			Name:         "MysqlMysqlBackup",
			Dependencies: acctest.DependencyGraph["mysqlBackup"],
			F:            sweepMysqlMysqlBackupResource,
		})
	}
}

func sweepMysqlMysqlBackupResource(compartment string) error {
	dbBackupsClient := acctest.GetTestClients(&schema.ResourceData{}).DbBackupsClient()
	mysqlBackupIds, err := getMysqlBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, mysqlBackupId := range mysqlBackupIds {
		if ok := acctest.SweeperDefaultResourceId[mysqlBackupId]; !ok {
			deleteBackupRequest := oci_mysql.DeleteBackupRequest{}
			deleteBackupRequest.BackupId = &mysqlBackupId

			deleteBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")
			_, error := dbBackupsClient.DeleteBackup(context.Background(), deleteBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting MysqlBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", mysqlBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mysqlBackupId, mysqlBackupSweepWaitCondition, time.Duration(3*time.Minute),
				mysqlBackupSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getMysqlBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MysqlBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbBackupsClient := acctest.GetTestClients(&schema.ResourceData{}).DbBackupsClient()

	listBackupsRequest := oci_mysql.ListBackupsRequest{}
	listBackupsRequest.CompartmentId = &compartmentId
	listBackupsRequest.LifecycleState = oci_mysql.BackupLifecycleStateActive
	listBackupsResponse, err := dbBackupsClient.ListBackups(context.Background(), listBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MysqlBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mysqlBackup := range listBackupsResponse.Items {
		id := *mysqlBackup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MysqlBackupId", id)
	}
	return resourceIds, nil
}

func mysqlBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mysqlBackupResponse, ok := response.Response.(oci_mysql.GetBackupResponse); ok {
		return mysqlBackupResponse.LifecycleState != oci_mysql.BackupLifecycleStateDeleted
	}
	return false
}

func mysqlBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbBackupsClient().GetBackup(context.Background(), oci_mysql.GetBackupRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
