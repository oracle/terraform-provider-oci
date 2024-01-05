// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseCancelExacsBackupDatabaseRepresentation = map[string]interface{}{
		"backup_id":             acctest.Representation{RepType: acctest.Required, Create: `unknown`},
		"cancel_backup_trigger": acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	GetBackupIdDataSourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.id}`},
	}

	DatabaseCancelExacsBackupDatabaseResourceDependencies = DatabaseDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, DatabaseDatabaseRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseBackupCancelManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseBackupCancelManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	resourceName := "oci_database_backup_cancel_management.test_cancel_backup"
	databaseName := "oci_database_database.test_database"
	listBackupName := "data.oci_database_backups.test_backups"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseCancelExacsBackupDatabaseResourceDependencies+kmsKeyIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_cancel_management", "test_cancel_backup", acctest.Optional, acctest.Create, DatabaseCancelExacsBackupDatabaseRepresentation), "database", "backup", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{

			// Creating dependencies and most importantly the database with autoBackups enabled
			{
				Config: config + compartmentIdVariableStr + DatabaseCancelExacsBackupDatabaseResourceDependencies + kmsKeyIdVariableStr,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(databaseName, "db_backup_config.0.auto_backup_enabled", "true"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, databaseName, "id")
						return err
					},
				),
			},
			// Poll the database until it reaches a backup_in_progress state and then call list backup on the databaseId to get the backup_id
			{
				PreConfig: acctest.WaitTillCondition(acctest.TestAccProvider, &resId, databaseWaitTillBackupInProgressConditionExa, time.Duration(1*time.Minute),
					databaseResponseFetchOperationExa, "database", true),
				Config: config + compartmentIdVariableStr + DatabaseCancelExacsBackupDatabaseResourceDependencies + kmsKeyIdVariableStr + acctest.GenerateDataSourceFromRepresentationMap("oci_database_backups", "test_backups", acctest.Optional, acctest.Update, GetBackupIdDataSourceRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(listBackupName, "backups.#", "1"),
				),
			},

			// populate the cancel_backup request with the backup id
			{
				Config: config + compartmentIdVariableStr + DatabaseCancelExacsBackupDatabaseResourceDependencies + kmsKeyIdVariableStr +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_backups", "test_backups", acctest.Optional, acctest.Update, GetBackupIdDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_backup_cancel_management", "test_cancel_backup", acctest.Required, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseCancelExacsBackupDatabaseRepresentation, []string{"backup_id"}), map[string]interface{}{
							"backup_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_backups.test_backups.backups.0.id}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "backup_id"),
				),
			},

			// test if retries of cancel_backup on a backup that has already completed or been canceled will fail
			{
				Config: config + compartmentIdVariableStr + DatabaseCancelExacsBackupDatabaseResourceDependencies + kmsKeyIdVariableStr +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_backups", "test_backups", acctest.Optional, acctest.Update, GetBackupIdDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_backup_cancel_management", "test_cancel_backup", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseCancelExacsBackupDatabaseRepresentation, []string{"backup_id"}), map[string]interface{}{
							"backup_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_backups.test_backups.backups.0.id}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "backup_id"),
					resource.TestCheckResourceAttrSet(resourceName, "cancel_backup_trigger"),
				),
			},
		},
	})
}

func databaseWaitTillBackupInProgressConditionExa(response common.OCIOperationResponse) bool {
	if databaseResponse, ok := response.Response.(oci_database.GetDatabaseResponse); ok {
		fmt.Print("Checking whether the state of resource is BACKUP_IN_PROGRESS: ", databaseResponse.LifecycleState == oci_database.DatabaseLifecycleStateBackupInProgress, "\n")
		return (databaseResponse.LifecycleState != oci_database.DatabaseLifecycleStateBackupInProgress)
	}
	return false
}

func databaseResponseFetchOperationExa(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetDatabase(context.Background(), oci_database.GetDatabaseRequest{
		DatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
