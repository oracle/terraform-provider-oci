// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementCloudDbSystemRequiredOnlyResource = DatabaseManagementCloudDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemRepresentation)

	DatabaseManagementCloudDbSystemResourceConfig = DatabaseManagementCloudDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemRepresentation)

	DatabaseManagementCloudDbSystemSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_db_system.test_cloud_db_system.id}`},
	}

	DatabaseManagementCloudDbSystemDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.disc_compartment_id}`},
		"dbaas_parent_infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.disc_dbaas_dbsystem_id}`},
		"deployment_type":                acctest.Representation{RepType: acctest.Optional, Create: `VM`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.disc_dbaas_dbsystem_name}`, Update: `displayName2`},
		"state":                          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementCloudDbSystemDataSourceFilterRepresentation}}
	DatabaseManagementCloudDbSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_db_system.test_cloud_db_system.id}`}},
	}

	DatabaseManagementCloudDbSystemDiscoveryCdbPatchOperationsRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseManagementCloudDbSystemDiscoveryRepresentation,
		map[string]interface{}{
			"patch_operations": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDiscoveryCdbPatchRepresentation}},
		},
	)
	DatabaseManagementCloudDbSystemDiscoveryPdbPatchOperationsRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseManagementCloudDbSystemDiscoveryRepresentation,
		map[string]interface{}{
			"patch_operations": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDiscoveryPdbPatchRepresentation}},
		},
	)
	DatabaseManagementCloudDbSystemManagementRepresentation = map[string]interface{}{
		"cloud_db_system_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_db_system.test_cloud_db_system.id}`},
		"enable_cloud_database_management": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"is_enabled":                       acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"metadata":                         acctest.Representation{RepType: acctest.Optional, Create: `{}`},
	}

	ignoreFeatureChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`database_management_config`, `stack_monitoring_config`, `defined_tags`}},
	}
	DatabaseManagementCloudDbSystemRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.disc_compartment_id}`},
		"db_system_discovery_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_db_system_discovery.test_cloud_db_system_discovery.id}`},
		"database_management_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementCloudDbSystemDatabaseManagementConfigRepresentation},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `${var.disc_dbaas_dbsystem_name}`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		// Stack Monitoring for cloud db system is not yet available.
		//"stack_monitoring_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseManagementCloudDbSystemStackMonitoringConfigRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreFeatureChangesRepresentation},
	}
	DatabaseManagementCloudDbSystemDatabaseManagementConfigRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}
	DatabaseManagementCloudDbSystemStackMonitoringConfigRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	DatabaseManagementCloudDbSystemResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudDbSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudDbSystemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("disc_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"disc_compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	dbaasDbsystemId := utils.GetEnvSettingWithBlankDefault("disc_dbaas_dbsystem_id")
	dbaasDbsystemIdVariableStr := fmt.Sprintf("variable \"disc_dbaas_dbsystem_id\" { default = \"%s\" }\n", dbaasDbsystemId)

	dbaasDbsystemName := utils.GetEnvSettingWithBlankDefault("disc_dbaas_dbsystem_name")
	dbaasDbsystemNameVariableStr := fmt.Sprintf("variable \"disc_dbaas_dbsystem_name\" { default = \"%s\" }\n", dbaasDbsystemName)

	cloudAgentId := utils.GetEnvSettingWithBlankDefault("disc_cloud_agent_id")
	cloudAgentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", cloudAgentId)

	dbHostName := utils.GetEnvSettingWithBlankDefault("db_host_name")
	dbHostNameVariableStr := fmt.Sprintf("variable \"db_host_name\" { default = \"%s\" }\n", dbHostName)

	dbPort := utils.GetEnvSettingWithBlankDefault("db_port")
	dbPortVariableStr := fmt.Sprintf("variable \"db_port\" { default = \"%s\" }\n", dbPort)

	dbServiceName := utils.GetEnvSettingWithBlankDefault("db_service_name")
	dbServiceNameVariableStr := fmt.Sprintf("variable \"db_service_name\" { default = \"%s\" }\n", dbServiceName)

	dbCredentialName := utils.GetEnvSettingWithBlankDefault("db_credential_name")
	dbCredentialNameVariableStr := fmt.Sprintf("variable \"db_credential_name\" { default = \"%s\" }\n", dbCredentialName)

	dbUserName := utils.GetEnvSettingWithBlankDefault("db_user_name")
	dbUserNameVariableStr := fmt.Sprintf("variable \"db_user_name\" { default = \"%s\" }\n", dbUserName)

	dbPasswordSecretId := utils.GetEnvSettingWithBlankDefault("db_password_secret_id")
	dbPasswordSecretIdVariableStr := fmt.Sprintf("variable \"db_password_secret_id\" { default = \"%s\" }\n", dbPasswordSecretId)

	resourceName := "oci_database_management_cloud_db_system.test_cloud_db_system"
	datasourceName := "data.oci_database_management_cloud_db_systems.test_cloud_db_systems"
	singularDatasourceName := "data.oci_database_management_cloud_db_system.test_cloud_db_system"
	discoveryResourceName := "oci_database_management_cloud_db_system_discovery.test_cloud_db_system_discovery"

	var variableStr = compartmentIdVariableStr + compartmentIdUVariableStr + dbaasDbsystemIdVariableStr +
		cloudAgentIdVariableStr + dbaasDbsystemNameVariableStr + dbHostNameVariableStr + dbPortVariableStr +
		dbServiceNameVariableStr + dbCredentialNameVariableStr + dbUserNameVariableStr + dbPasswordSecretIdVariableStr

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+DatabaseManagementCloudDbSystemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system",
			acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemRepresentation), "databasemanagement", "cloudDbSystem", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementCloudDbSystemDestroy, []resource.TestStep{
		// Patch discovery and add connector to CDB
		{
			Config: config + variableStr + DatabaseManagementCloudDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryCdbPatchOperationsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(discoveryResourceName, "agent_id"),
				resource.TestCheckResourceAttr(discoveryResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(discoveryResourceName, "dbaas_parent_infrastructure_id"),
				resource.TestCheckResourceAttr(discoveryResourceName, "deployment_type", "VM"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, discoveryResourceName, "id")
					return err
				},
			),
		},
		// verify Create
		{
			Config: config + variableStr + DatabaseManagementCloudDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_discovery_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + variableStr + DatabaseManagementCloudDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system",
					acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_management_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_management_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_discovery_id"),
				resource.TestCheckResourceAttrSet(resourceName, "dbaas_parent_infrastructure_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_type"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "stack_monitoring_config.#", "1"),
				// Stack Monitoring for cloud db system is not yet available.
				//resource.TestCheckResourceAttr(resourceName, "stack_monitoring_config.0.is_enabled", "false"),
				//resource.TestCheckResourceAttr(resourceName, "stack_monitoring_config.0.metadata", "metadata"),
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
			Config: config + variableStr + DatabaseManagementCloudDbSystemResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_systems", "test_cloud_db_systems",
					acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system",
					acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "dbaas_parent_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_type", "VM"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_db_system_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_db_system_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + variableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				DatabaseManagementCloudDbSystemResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovery_agent_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cluster"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stack_monitoring_config.#", "1"),
				// Stack Monitoring for cloud db system is not yet available.
				//resource.TestCheckResourceAttr(singularDatasourceName, "stack_monitoring_config.0.is_enabled", "false"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "stack_monitoring_config.0.metadata", "metadata"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// disable DB Management
		{
			Config: config + variableStr + DatabaseManagementCloudDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryPdbPatchOperationsRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system",
					acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_database_managements_management",
					"test_cloud_db_system_cloud_database_managements_management",
					acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_discovery_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudDbSystemRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				// DBM was disabled in the previous step
				"database_management_config",
				"lifecycle_details",
				"time_updated",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementCloudDbSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_cloud_db_system" {
			noResourceFound = false
			request := oci_database_management.GetCloudDbSystemRequest{}

			tmp := rs.Primary.ID
			request.CloudDbSystemId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetCloudDbSystem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.CloudDbSystemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementCloudDbSystem") {
		resource.AddTestSweepers("DatabaseManagementCloudDbSystem", &resource.Sweeper{
			Name:         "DatabaseManagementCloudDbSystem",
			Dependencies: acctest.DependencyGraph["cloudDbSystem"],
			F:            sweepDatabaseManagementCloudDbSystemResource,
		})
	}
}

func sweepDatabaseManagementCloudDbSystemResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	cloudDbSystemIds, err := getDatabaseManagementCloudDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudDbSystemId := range cloudDbSystemIds {
		if ok := acctest.SweeperDefaultResourceId[cloudDbSystemId]; !ok {
			deleteCloudDbSystemRequest := oci_database_management.DeleteCloudDbSystemRequest{}

			deleteCloudDbSystemRequest.CloudDbSystemId = &cloudDbSystemId

			deleteCloudDbSystemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteCloudDbSystem(context.Background(), deleteCloudDbSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudDbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudDbSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cloudDbSystemId, DatabaseManagementCloudDbSystemSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementCloudDbSystemSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementCloudDbSystemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudDbSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listCloudDbSystemsRequest := oci_database_management.ListCloudDbSystemsRequest{}
	listCloudDbSystemsRequest.CompartmentId = &compartmentId
	listCloudDbSystemsRequest.LifecycleState = oci_database_management.ListCloudDbSystemsLifecycleStateActive
	listCloudDbSystemsResponse, err := dbManagementClient.ListCloudDbSystems(context.Background(), listCloudDbSystemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudDbSystem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudDbSystem := range listCloudDbSystemsResponse.Items {
		id := *cloudDbSystem.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudDbSystemId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementCloudDbSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudDbSystemResponse, ok := response.Response.(oci_database_management.GetCloudDbSystemResponse); ok {
		return cloudDbSystemResponse.LifecycleState != oci_database_management.CloudDbSystemLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementCloudDbSystemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetCloudDbSystem(context.Background(), oci_database_management.GetCloudDbSystemRequest{
		CloudDbSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
