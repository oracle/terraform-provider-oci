// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RecoveryLongTermBackupRequiredOnlyResource = RecoveryLongTermBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Required, acctest.Create, RecoveryLongTermBackupRepresentation)

	RecoveryLongTermBackupResourceConfig = RecoveryLongTermBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Optional, acctest.Update, RecoveryLongTermBackupRepresentation)

	RecoveryLongTermBackupSingularDataSourceRepresentation = map[string]interface{}{
		"long_term_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_recovery_long_term_backup.test_long_term_backup.id}`},
	}

	RecoveryLongTermBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_recovery_long_term_backup.test_long_term_backup.id}`},
		"protected_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_recovery_protected_database.test_protected_database.id}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: RecoveryLongTermBackupDataSourceFilterRepresentation}}
	RecoveryLongTermBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_recovery_long_term_backup.test_long_term_backup.id}`}},
	}

	RecoveryLongTermBackupRepresentation = map[string]interface{}{
		"protected_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_recovery_protected_database.test_protected_database.id}`},
		"retention_period":        acctest.RepresentationGroup{RepType: acctest.Required, Group: RecoveryLongTermBackupRetentionPeriodRepresentation},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"retention_point_in_time": acctest.Representation{RepType: acctest.Optional, Create: recoveryLongTermBackupRetentionTime},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: recoveryIgnoreDefinedTagsRepresentation},
	}
	RecoveryLongTermBackupRetentionPeriodRepresentation = map[string]interface{}{
		"retention_count":       acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `90`},
		"retention_period_type": acctest.Representation{RepType: acctest.Required, Create: `YEAR`, Update: `DAY`},
	}
	recoveryLongTermBackupRetentionTime = time.Now().UTC().Truncate(time.Second).Add(-(time.Duration(60) * time.Hour)).Add(time.Millisecond).Format(time.RFC3339Nano)

	RecoveryLongTermBackupResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_protected_database", "test_protected_database", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("protected_database_id", acctest.Representation{RepType: acctest.Required, Create: `${var.protected_database_id}`}, RecoveryProtectedDatabaseSingularDataSourceRepresentation)) +
		protectedDatabaseIdVariableStr +
		DefinedTagsDependencies
	protectedDatabaseId            = utils.GetEnvSettingWithBlankDefault("protected_database_id")
	protectedDatabaseIdVariableStr = fmt.Sprintf("variable \"protected_database_id\" { default = \"%s\" }\n", protectedDatabaseId)
)

// issue-routing-tag: recovery/default
func TestRecoveryLongTermBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRecoveryLongTermBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_recovery_long_term_backup.test_long_term_backup"
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RecoveryLongTermBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Optional, acctest.Create, RecoveryLongTermBackupRepresentation), "recovery", "longTermBackup", t)

	acctest.ResourceTest(t, testAccCheckRecoveryLongTermBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RecoveryLongTermBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Required, acctest.Create, RecoveryLongTermBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "protected_database_id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_period", map[string]string{
					"retention_count":       "1",
					"retention_period_type": "YEAR",
				},
					[]string{}),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RecoveryLongTermBackupResourceDependencies,
		},
	})
}
func TestRecoveryLongTermBackupResource_optionalAndDataSource(t *testing.T) {
	httpreplay.SetScenario("TestRecoveryLongTermBackupResource_optionalAndDataSource")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_recovery_long_term_backup.test_long_term_backup"
	datasourceName := "data.oci_recovery_long_term_backups.test_long_term_backups"
	singularDatasourceName := "data.oci_recovery_long_term_backup.test_long_term_backup"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RecoveryLongTermBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Optional, acctest.Create, RecoveryLongTermBackupRepresentation), "recovery", "longTermBackup", t)

	acctest.ResourceTest(t, testAccCheckRecoveryLongTermBackupDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RecoveryLongTermBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Optional, acctest.Create, RecoveryLongTermBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "protected_database_id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_period", map[string]string{
					"retention_count":       "1",
					"retention_period_type": "YEAR",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "retention_point_in_time", recoveryLongTermBackupRetentionTime),
				resource.TestCheckResourceAttrSet(resourceName, "retention_until_date_time"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + RecoveryLongTermBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Optional, acctest.Update, RecoveryLongTermBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "protected_database_id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_period", map[string]string{
					"retention_count":       "90",
					"retention_period_type": "DAY",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "retention_point_in_time", recoveryLongTermBackupRetentionTime),
				resource.TestCheckResourceAttrSet(resourceName, "retention_until_date_time"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_long_term_backups", "test_long_term_backups", acctest.Optional, acctest.Update, RecoveryLongTermBackupDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryLongTermBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Optional, acctest.Update, RecoveryLongTermBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "protected_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "long_term_backup_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "long_term_backup_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_long_term_backup", "test_long_term_backup", acctest.Required, acctest.Create, RecoveryLongTermBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryLongTermBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "long_term_backup_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_identifier"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retention_period.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "retention_period", map[string]string{
					"retention_count":       "90",
					"retention_period_type": "DAY",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_point_in_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_until_date_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "rman_tag"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_backup_completed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_backup_initiated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + RecoveryLongTermBackupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckRecoveryLongTermBackupDestroy(s *terraform.State) error {
	noResourceFound := false
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseRecoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_recovery_long_term_backup" {
			noResourceFound = false
			request := oci_recovery.GetLongTermBackupRequest{}

			tmp := rs.Primary.ID
			request.LongTermBackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "recovery")

			response, err := client.GetLongTermBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_recovery.LongTermBackupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("RecoveryLongTermBackup") {
		resource.AddTestSweepers("RecoveryLongTermBackup", &resource.Sweeper{
			Name:         "RecoveryLongTermBackup",
			Dependencies: acctest.DependencyGraph["longTermBackup"],
			F:            sweepRecoveryLongTermBackupResource,
		})
	}
}

func sweepRecoveryLongTermBackupResource(compartment string) error {
	databaseRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseRecoveryClient()
	longTermBackupIds, err := getRecoveryLongTermBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, longTermBackupId := range longTermBackupIds {
		if ok := acctest.SweeperDefaultResourceId[longTermBackupId]; !ok {
			deleteLongTermBackupRequest := oci_recovery.DeleteLongTermBackupRequest{}

			deleteLongTermBackupRequest.LongTermBackupId = &longTermBackupId

			deleteLongTermBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "recovery")
			_, error := databaseRecoveryClient.DeleteLongTermBackup(context.Background(), deleteLongTermBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting LongTermBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", longTermBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &longTermBackupId, RecoveryLongTermBackupSweepWaitCondition, time.Duration(3*time.Minute),
				RecoveryLongTermBackupSweepResponseFetchOperation, "recovery", true)
		}
	}
	return nil
}

func getRecoveryLongTermBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LongTermBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseRecoveryClient()

	listLongTermBackupsRequest := oci_recovery.ListLongTermBackupsRequest{}
	listLongTermBackupsRequest.CompartmentId = &compartmentId
	listLongTermBackupsRequest.LifecycleState = oci_recovery.ListLongTermBackupsLifecycleStateActive
	listLongTermBackupsResponse, err := databaseRecoveryClient.ListLongTermBackups(context.Background(), listLongTermBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LongTermBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, longTermBackup := range listLongTermBackupsResponse.Items {
		id := *longTermBackup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LongTermBackupId", id)
	}
	return resourceIds, nil
}

func RecoveryLongTermBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if longTermBackupResponse, ok := response.Response.(oci_recovery.GetLongTermBackupResponse); ok {
		return longTermBackupResponse.LifecycleState != oci_recovery.LongTermBackupLifecycleStateDeleted
	}
	return false
}

func RecoveryLongTermBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseRecoveryClient().GetLongTermBackup(context.Background(), oci_recovery.GetLongTermBackupRequest{
		LongTermBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
