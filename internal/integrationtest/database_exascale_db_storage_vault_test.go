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
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseExascaleDbStorageVaultRequiredOnlyResource = DatabaseExascaleDbStorageVaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Required, acctest.Create, DatabaseExascaleDbStorageVaultRepresentation)

	DatabaseExascaleDbStorageVaultResourceConfig = DatabaseExascaleDbStorageVaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Optional, acctest.Update, DatabaseExascaleDbStorageVaultRepresentation)

	DatabaseExascaleDbStorageVaultSingularDataSourceRepresentation = map[string]interface{}{
		"exascale_db_storage_vault_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
	}

	DatabaseExascaleDbStorageVaultDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TFExascaleDbStorageVault`, Update: `TFExascaleDbStorageVaultUpdatedName`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExascaleDbStorageVaultDataSourceFilterRepresentation}}

	DatabaseExascaleDbStorageVaultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`}},
	}

	DatabaseExascaleDbStorageVaultRepresentation = map[string]interface{}{
		"availability_domain":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                      acctest.Representation{RepType: acctest.Required, Create: `TFExascaleDbStorageVault`, Update: `TFExascaleDbStorageVaultUpdatedName`},
		"high_capacity_database_storage":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExascaleDbStorageVaultHighCapacityDatabaseStorageRepresentation},
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `20`, Update: `25`},
		"description":                       acctest.Representation{RepType: acctest.Optional, Create: `ExaScale DB Storage Vault - description`, Update: `ExaScale DB Storage Vault - updated description`},
		"time_zone":                         acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"defined_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"example-tag-namespace-all.example-tag": "value"}, Update: map[string]string{"example-tag-namespace-all.example-tag": "updatedValue"}},
		"freeform_tags":                     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExascaleDbStorageIgnoreDefinedTagsRepresentation},
	}

	DatabaseExascaleDbStorageVaultHighCapacityDatabaseStorageRepresentation = map[string]interface{}{
		"total_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `800`, Update: `1600`},
	}

	DatabaseExascaleDbStorageIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatabaseExascaleDbStorageVaultResourceDependencies = AvailabilityDomainConfig
)

// issue-routing-tag: database/ExaCS
func TestDatabaseExascaleDbStorageVaultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExascaleDbStorageVaultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault"
	datasourceName := "data.oci_database_exascale_db_storage_vaults.test_exascale_db_storage_vaults"
	singularDatasourceName := "data.oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExascaleDbStorageVaultResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Optional, acctest.Create, DatabaseExascaleDbStorageVaultRepresentation), "database", "exascaleDbStorageVault", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExascaleDbStorageVaultDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExascaleDbStorageVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Required, acctest.Create, DatabaseExascaleDbStorageVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExascaleDbStorageVault"),
				resource.TestCheckResourceAttr(resourceName, "high_capacity_database_storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "high_capacity_database_storage.0.total_size_in_gbs", "800"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExascaleDbStorageVaultResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseExascaleDbStorageVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Optional, acctest.Create, DatabaseExascaleDbStorageVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExascaleDbStorageVault"),
				//resource.TestCheckResourceAttr(resourceName, "vm_cluster_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "high_capacity_database_storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "high_capacity_database_storage.0.total_size_in_gbs", "800"),
				resource.TestCheckResourceAttr(resourceName, "additional_flash_cache_in_percent", "20"),
				resource.TestCheckResourceAttr(resourceName, "description", "ExaScale DB Storage Vault - description"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseExascaleDbStorageVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseExascaleDbStorageVaultRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExascaleDbStorageVault"),
				//resource.TestCheckResourceAttr(resourceName, "vm_cluster_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "high_capacity_database_storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "high_capacity_database_storage.0.total_size_in_gbs", "800"),
				resource.TestCheckResourceAttr(resourceName, "additional_flash_cache_in_percent", "20"),
				resource.TestCheckResourceAttr(resourceName, "description", "ExaScale DB Storage Vault - description"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseExascaleDbStorageVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Optional, acctest.Update, DatabaseExascaleDbStorageVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFExascaleDbStorageVaultUpdatedName"),
				//resource.TestCheckResourceAttr(resourceName, "vm_cluster_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "high_capacity_database_storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "high_capacity_database_storage.0.total_size_in_gbs", "1600"),
				resource.TestCheckResourceAttr(resourceName, "additional_flash_cache_in_percent", "25"),
				resource.TestCheckResourceAttr(resourceName, "description", "ExaScale DB Storage Vault - updated description"),
				resource.TestCheckResourceAttr(resourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exascale_db_storage_vaults", "test_exascale_db_storage_vaults", acctest.Optional, acctest.Update, DatabaseExascaleDbStorageVaultDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExascaleDbStorageVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Optional, acctest.Update, DatabaseExascaleDbStorageVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TFExascaleDbStorageVaultUpdatedName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "exascale_db_storage_vaults.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exascale_db_storage_vaults.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "exascale_db_storage_vaults.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "exascale_db_storage_vaults.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.display_name", "TFExascaleDbStorageVaultUpdatedName"),
				//resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.vm_cluster_count", "0"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.high_capacity_database_storage.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "exascale_db_storage_vaults.0.high_capacity_database_storage.0.available_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.high_capacity_database_storage.0.total_size_in_gbs", "1600"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.additional_flash_cache_in_percent", "25"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.description", "ExaScale DB Storage Vault - updated description"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(datasourceName, "exascale_db_storage_vaults.0.system_tags.%", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Required, acctest.Create, DatabaseExascaleDbStorageVaultSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExascaleDbStorageVaultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exascale_db_storage_vault_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TFExascaleDbStorageVaultUpdatedName"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "vm_cluster_count", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "high_capacity_database_storage.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "high_capacity_database_storage.0.available_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "high_capacity_database_storage.0.total_size_in_gbs", "1600"),
				resource.TestCheckResourceAttr(singularDatasourceName, "additional_flash_cache_in_percent", "25"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_zone", "US/Pacific"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "ExaScale DB Storage Vault - updated description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "0"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseExascaleDbStorageVaultRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseExascaleDbStorageVaultDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_exascale_db_storage_vault" {
			noResourceFound = false
			request := oci_database.GetExascaleDbStorageVaultRequest{}

			tmp := rs.Primary.ID
			request.ExascaleDbStorageVaultId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetExascaleDbStorageVault(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExascaleDbStorageVaultLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseExascaleDbStorageVault") {
		resource.AddTestSweepers("DatabaseExascaleDbStorageVault", &resource.Sweeper{
			Name:         "DatabaseExascaleDbStorageVault",
			Dependencies: acctest.DependencyGraph["exascaleDbStorageVault"],
			F:            sweepDatabaseExascaleDbStorageVaultResource,
		})
	}
}

func sweepDatabaseExascaleDbStorageVaultResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	exascaleDbStorageVaultIds, err := getDatabaseExascaleDbStorageVaultIds(compartment)
	if err != nil {
		return err
	}
	for _, exascaleDbStorageVaultId := range exascaleDbStorageVaultIds {
		if ok := acctest.SweeperDefaultResourceId[exascaleDbStorageVaultId]; !ok {
			deleteExascaleDbStorageVaultRequest := oci_database.DeleteExascaleDbStorageVaultRequest{}

			deleteExascaleDbStorageVaultRequest.ExascaleDbStorageVaultId = &exascaleDbStorageVaultId

			deleteExascaleDbStorageVaultRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExascaleDbStorageVault(context.Background(), deleteExascaleDbStorageVaultRequest)
			if error != nil {
				fmt.Printf("Error deleting ExascaleDbStorageVault %s %s, It is possible that the resource is already deleted. Please verify manually \n", exascaleDbStorageVaultId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &exascaleDbStorageVaultId, DatabaseExascaleDbStorageVaultSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseExascaleDbStorageVaultSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseExascaleDbStorageVaultIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExascaleDbStorageVaultId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listExascaleDbStorageVaultsRequest := oci_database.ListExascaleDbStorageVaultsRequest{}
	listExascaleDbStorageVaultsRequest.CompartmentId = &compartmentId
	listExascaleDbStorageVaultsRequest.LifecycleState = oci_database.ExascaleDbStorageVaultLifecycleStateAvailable
	listExascaleDbStorageVaultsResponse, err := databaseClient.ListExascaleDbStorageVaults(context.Background(), listExascaleDbStorageVaultsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExascaleDbStorageVault list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, exascaleDbStorageVault := range listExascaleDbStorageVaultsResponse.Items {
		id := *exascaleDbStorageVault.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExascaleDbStorageVaultId", id)
	}
	return resourceIds, nil
}

func DatabaseExascaleDbStorageVaultSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if exascaleDbStorageVaultResponse, ok := response.Response.(oci_database.GetExascaleDbStorageVaultResponse); ok {
		return exascaleDbStorageVaultResponse.LifecycleState != oci_database.ExascaleDbStorageVaultLifecycleStateTerminated
	}
	return false
}

func DatabaseExascaleDbStorageVaultSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetExascaleDbStorageVault(context.Background(), oci_database.GetExascaleDbStorageVaultRequest{
		ExascaleDbStorageVaultId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
