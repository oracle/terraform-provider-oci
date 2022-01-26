// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v56/datasafe"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	TargetDatabaseRequiredOnlyResource = TargetDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseRepresentation)

	TargetDatabaseResourceConfig = TargetDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Optional, acctest.Update, targetDatabaseRepresentation)

	targetDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"target_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database.test_target_database.id}`},
	}

	targetDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"target_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_target_database.test_target_database.id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: targetDatabaseDataSourceFilterRepresentation}}
	targetDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_target_database.test_target_database.id}`}},
	}

	targetDatabaseRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_details":  acctest.RepresentationGroup{RepType: acctest.Required, Group: targetDatabaseDatabaseDetailsRepresentation},
		"connection_option": acctest.RepresentationGroup{RepType: acctest.Optional, Group: targetDatabaseConnectionOptionRepresentation},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	targetDatabaseDatabaseDetailsRepresentation = map[string]interface{}{
		"database_type":          acctest.Representation{RepType: acctest.Required, Create: `AUTONOMOUS_DATABASE`, Update: `AUTONOMOUS_DATABASE`},
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"infrastructure_type":    acctest.Representation{RepType: acctest.Required, Create: `ORACLE_CLOUD`, Update: `ORACLE_CLOUD`},
	}
	targetDatabaseConnectionOptionRepresentation = map[string]interface{}{
		"connection_type":              acctest.Representation{RepType: acctest.Required, Create: `PRIVATE_ENDPOINT`, Update: `PRIVATE_ENDPOINT`},
		"datasafe_private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint.id}`},
	}
	targetDatabaseCredentialsRepresentation = map[string]interface{}{
		"password":  acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"user_name": acctest.Representation{RepType: acctest.Required, Create: `ADMIN`},
	}
	targetDatabaseTlsConfigRepresentation = map[string]interface{}{
		"status":                 acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `DISABLED`},
		"certificate_store_type": acctest.Representation{RepType: acctest.Optional, Create: `JKS`},
		"key_store_content":      acctest.Representation{RepType: acctest.Optional, Create: `keyStoreContent`, Update: `keyStoreContent2`},
		"store_password":         acctest.Representation{RepType: acctest.Optional, Create: `storePassword`, Update: `storePassword2`},
		"trust_store_content":    acctest.Representation{RepType: acctest.Optional, Create: `trustStoreContent`, Update: `trustStoreContent2`},
	}

	TargetDatabaseResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", acctest.Required, acctest.Create, dataSafePrivateEndpointRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeTargetDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_target_database.test_target_database"
	datasourceName := "data.oci_data_safe_target_databases.test_target_databases"
	singularDatasourceName := "data.oci_data_safe_target_database.test_target_database"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+TargetDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Optional, acctest.Create, targetDatabaseRepresentation), "datasafe", "targetDatabase", t)

	acctest.ResourceTest(t, testAccCheckDataSafeTargetDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + TargetDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.database_type", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.infrastructure_type", "ORACLE_CLOUD"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + TargetDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + TargetDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Optional, acctest.Create, targetDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_option.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_option.0.connection_type", "PRIVATE_ENDPOINT"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_option.0.datasafe_private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "database_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.database_type", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "database_details.0.autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.infrastructure_type", "ORACLE_CLOUD"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + TargetDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(targetDatabaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "connection_option.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_option.0.connection_type", "PRIVATE_ENDPOINT"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_option.0.datasafe_private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "database_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.database_type", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "database_details.0.autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.infrastructure_type", "ORACLE_CLOUD"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + TargetDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Optional, acctest.Update, targetDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_option.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_option.0.connection_type", "PRIVATE_ENDPOINT"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_option.0.datasafe_private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "database_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.database_type", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "database_details.0.autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.infrastructure_type", "ORACLE_CLOUD"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_databases", "test_target_databases", acctest.Optional, acctest.Update, targetDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + TargetDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Optional, acctest.Update, targetDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_database_id"),

				resource.TestCheckResourceAttr(datasourceName, "target_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "target_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "target_databases.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "target_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "target_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + TargetDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_option.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_option.0.connection_type", "PRIVATE_ENDPOINT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_details.0.database_type", "AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_details.0.infrastructure_type", "ORACLE_CLOUD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + TargetDatabaseResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"credentials",
				"tls_config",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataSafeTargetDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_target_database" {
			noResourceFound = false
			request := oci_data_safe.GetTargetDatabaseRequest{}

			tmp := rs.Primary.ID
			request.TargetDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetTargetDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeTargetDatabase") {
		resource.AddTestSweepers("DataSafeTargetDatabase", &resource.Sweeper{
			Name:         "DataSafeTargetDatabase",
			Dependencies: acctest.DependencyGraph["targetDatabase"],
			F:            sweepDataSafeTargetDatabaseResource,
		})
	}
}

func sweepDataSafeTargetDatabaseResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	targetDatabaseIds, err := getTargetDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, targetDatabaseId := range targetDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[targetDatabaseId]; !ok {
			deleteTargetDatabaseRequest := oci_data_safe.DeleteTargetDatabaseRequest{}

			deleteTargetDatabaseRequest.TargetDatabaseId = &targetDatabaseId

			deleteTargetDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteTargetDatabase(context.Background(), deleteTargetDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting TargetDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", targetDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &targetDatabaseId, targetDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				targetDatabaseSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getTargetDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TargetDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listTargetDatabasesRequest := oci_data_safe.ListTargetDatabasesRequest{}
	listTargetDatabasesRequest.CompartmentId = &compartmentId
	listTargetDatabasesRequest.LifecycleState = oci_data_safe.ListTargetDatabasesLifecycleStateActive
	listTargetDatabasesResponse, err := dataSafeClient.ListTargetDatabases(context.Background(), listTargetDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TargetDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, targetDatabase := range listTargetDatabasesResponse.Items {
		id := *targetDatabase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TargetDatabaseId", id)
	}
	return resourceIds, nil
}

func targetDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if targetDatabaseResponse, ok := response.Response.(oci_data_safe.GetTargetDatabaseResponse); ok {
		return targetDatabaseResponse.LifecycleState != oci_data_safe.LifecycleStateDeleted
	}
	return false
}

func targetDatabaseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetTargetDatabase(context.Background(), oci_data_safe.GetTargetDatabaseRequest{
		TargetDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
